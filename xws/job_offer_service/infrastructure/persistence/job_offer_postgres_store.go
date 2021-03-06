package persistence

import (
	"context"
	"dislinkt/common/tracer"
	"dislinkt/job_offer_service/domain"
	"dislinkt/job_offer_service/ent"
	"dislinkt/job_offer_service/ent/joboffer"
	"dislinkt/job_offer_service/ent/skill"
	"fmt"
	"log"
	"sort"
)

type JobOfferPostgresStore struct {
	jobOfferString string
}

func NewJobOfferPostgresStore(host string, port string) domain.JobOfferStore {
	jobOfferString := fmt.Sprintf("host=%s port=%s user=postgres password=ftn dbname=dislinkt sslmode=disable", host, port)
	client, err := ent.Open("postgres", jobOfferString)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//err = entc.Generate("./ent/schema", &gen.Config{}, entc.Extensions(entviz.Extension{}))
	//if err != nil {
	//	log.Fatalf("running ent codegen: %v", err)
	//}
	return &JobOfferPostgresStore{
		jobOfferString: jobOfferString,
	}
}

func (store *JobOfferPostgresStore) GetJobs(ctx context.Context) ([]*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "GetJobs Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	client, err := ent.Open("postgres", store.jobOfferString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	all, err := client.JobOffer.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	var jobOffers = make([]*domain.JobOffer, 0)

	for _, jobOffer := range all {
		skills, err := client.JobOffer.QueryRequires(jobOffer).All(ctx)
		if err != nil {
			return nil, err
		}
		var skillNames []string
		for _, s := range skills {
			skillNames = append(skillNames, s.Name)
		}
		jobOffers = append(jobOffers, &domain.JobOffer{
			Id:          jobOffer.ID,
			ProfileId:   jobOffer.ProfileID,
			Company:     jobOffer.Company,
			Position:    jobOffer.Position,
			Description: jobOffer.Description,
			Criteria:    jobOffer.Criteria,
			Skills:      skillNames,
			CreatedAt:   jobOffer.CreatedAt,
		})
	}
	return jobOffers, nil
}

func (store *JobOfferPostgresStore) GetMyJobs(ctx context.Context, profileId string) ([]*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "GetMyJobs Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	client, err := ent.Open("postgres", store.jobOfferString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	all, err := client.JobOffer.Query().Where(joboffer.ProfileID(profileId)).All(ctx)

	if err != nil {
		return nil, err
	}

	var jobOffers = make([]*domain.JobOffer, 0)

	for _, jobOffer := range all {
		skills, err := client.JobOffer.QueryRequires(jobOffer).All(ctx)
		if err != nil {
			return nil, err
		}
		var skillNames []string
		for _, s := range skills {
			skillNames = append(skillNames, s.Name)
		}
		jobOffers = append(jobOffers, &domain.JobOffer{
			Id:          jobOffer.ID,
			ProfileId:   jobOffer.ProfileID,
			Company:     jobOffer.Company,
			Position:    jobOffer.Position,
			Description: jobOffer.Description,
			Criteria:    jobOffer.Criteria,
			Skills:      skillNames,
			CreatedAt:   jobOffer.CreatedAt,
		})
	}
	return jobOffers, nil
}

func (store *JobOfferPostgresStore) GetJob(ctx context.Context, id int) (*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "GetJob Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	client, err := ent.Open("postgres", store.jobOfferString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	jobOffer, err := client.JobOffer.Query().Where(joboffer.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	skills, err := client.JobOffer.QueryRequires(jobOffer).All(ctx)
	if err != nil {
		return nil, err
	}

	var skillNames []string
	for _, s := range skills {
		skillNames = append(skillNames, s.Name)
	}

	return &domain.JobOffer{
		Id:          jobOffer.ID,
		ProfileId:   jobOffer.ProfileID,
		Position:    jobOffer.Position,
		Company:     jobOffer.Company,
		Description: jobOffer.Description,
		Criteria:    jobOffer.Criteria,
		CreatedAt:   jobOffer.CreatedAt,
		Skills:      skillNames,
	}, nil
}

func (store *JobOfferPostgresStore) CreateJob(ctx context.Context, jobOffer *domain.JobOffer) (*domain.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateJob Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	client, err := ent.Open("postgres", store.jobOfferString)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	var skills []*ent.Skill
	for _, s := range jobOffer.Skills {
		newSkill, err := client.Skill.Create().SetName(s).Save(ctx)
		if err != nil {
			existingSkill, err := client.Skill.Query().Where(skill.NameEQ(s)).Only(ctx)
			if err != nil {
				return nil, err
			}
			skills = append(skills, existingSkill)
			continue
		}
		skills = append(skills, newSkill)
	}
	newJobOffer, err := client.JobOffer.Create().
		SetProfileID(jobOffer.ProfileId).
		SetPosition(jobOffer.Position).
		SetCompany(jobOffer.Company).
		SetDescription(jobOffer.Description).
		SetCriteria(jobOffer.Criteria).
		SetCreatedAt(jobOffer.CreatedAt).
		AddRequires(skills...).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.JobOffer{
		Id:          newJobOffer.ID,
		ProfileId:   newJobOffer.ProfileID,
		Position:    newJobOffer.Position,
		Company:     newJobOffer.Company,
		Description: newJobOffer.Description,
		Criteria:    newJobOffer.Criteria,
		CreatedAt:   newJobOffer.CreatedAt,
	}, nil

}

func (store *JobOfferPostgresStore) DeleteAll(ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteAll Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	client, err := ent.Open("postgres", store.jobOfferString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	_, err = client.JobOffer.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.Skill.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (store *JobOfferPostgresStore) Delete(ctx context.Context, id int) error {
	span := tracer.StartSpanFromContext(ctx, "Delete Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	client, err := ent.Open("postgres", store.jobOfferString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	err = client.JobOffer.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (store *JobOfferPostgresStore) DeleteSkill(ctx context.Context, skillName string) error {
	span := tracer.StartSpanFromContext(ctx, "DeleteSkill Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	client, err := ent.Open("postgres", store.jobOfferString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	_, err = client.Skill.Delete().Where(skill.NameEQ(skillName)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (store *JobOfferPostgresStore) GetRecommendations(ctx context.Context, profileId string, skills []string) ([]*domain.JobRecommendation, error) {
	span := tracer.StartSpanFromContext(ctx, "GetRecommendations Store")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	client, err := ent.Open("postgres", store.jobOfferString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	jobOffers, err := client.JobOffer.Query().
		Where(joboffer.ProfileIDNEQ(profileId)).
		QueryRequires().
		Where(skill.NameIn(skills...)).
		QueryRequired().
		All(ctx)

	if err != nil {
		return nil, err
	}

	recommendations := make(map[int64]int64)

	for _, jobOffer := range jobOffers {
		skillCount, err := client.JobOffer.Query().
			Where(joboffer.IDEQ(jobOffer.ID)).
			QueryRequires().
			Where(skill.NameIn(skills...)).
			Count(ctx)
		if err != nil {
			return nil, err
		}
		if skillCount > 0 {
			recommendations[int64(jobOffer.ID)] = int64(skillCount)
		}
	}

	var recommendationsSorted []*domain.JobRecommendation
	for k, v := range recommendations {
		recommendationsSorted = append(recommendationsSorted, &domain.JobRecommendation{JobId: k, SkillCount: v})
	}

	sort.Slice(recommendationsSorted, func(i, j int) bool {
		return recommendationsSorted[i].SkillCount > recommendationsSorted[j].SkillCount
	})

	return recommendationsSorted, nil
}
