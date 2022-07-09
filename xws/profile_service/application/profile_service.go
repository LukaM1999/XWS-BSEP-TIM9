package application

import (
	auth "dislinkt/common/domain"
	"dislinkt/common/loggers"
	"dislinkt/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
)

var log = loggers.NewProfileLogger()

type ProfileService struct {
	store        domain.ProfileStore
	orchestrator *UpdateProfileOrchestrator
}

func NewProfileService(store domain.ProfileStore, orchestrator *UpdateProfileOrchestrator) *ProfileService {
	return &ProfileService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *ProfileService) Get(profileId string) (*domain.Profile, error) {
	return service.store.Get(profileId)
}

func (service *ProfileService) GetAll(search string) ([]*domain.Profile, error) {
	return service.store.GetAll(search)
}

func (service *ProfileService) Create(profile *domain.Profile) error {
	return service.store.Create(profile)
}

func (service *ProfileService) RollbackUpdate(profile *domain.Profile) error {
	return service.store.Update(profile.Id.Hex(), profile)
}

func (service *ProfileService) Update(profileId string, profile *domain.Profile) error {
	oldProfile, err := service.Get(profileId)
	if err != nil {
		log.WithField("profileId", profileId).Errorf("Cannot get profile: %v", err)
		return err
	}
	err = service.store.Update(profileId, profile)
	if err != nil {
		log.WithField("profileId", profileId).Errorf("Cannot update profile: %v", err)
		return err
	}
	newProfile := &auth.Profile{
		Id:             profile.Id,
		Username:       profile.Username,
		FirstName:      profile.FirstName,
		LastName:       profile.LastName,
		FullName:       profile.FirstName + profile.LastName,
		DateOfBirth:    profile.DateOfBirth,
		PhoneNumber:    profile.PhoneNumber,
		Email:          profile.Email,
		Gender:         profile.Gender,
		Biography:      profile.Biography,
		Education:      make([]auth.Education, 0),
		WorkExperience: make([]auth.WorkExperience, 0),
		Skills:         make([]string, 0),
		Interests:      make([]string, 0),
		IsPrivate:      profile.IsPrivate,
		AgentToken:     profile.AgentToken,
	}
	for _, education := range profile.Education {
		education := &domain.Education{
			School:       education.School,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			StartDate:    education.StartDate,
			EndDate:      education.EndDate,
			Grade:        education.Grade,
			Description:  education.Description,
		}
		profile.Education = append(profile.Education, *education)
	}

	for _, workExperience := range profile.WorkExperience {
		workExperience := &domain.WorkExperience{
			Title:          workExperience.Title,
			Company:        workExperience.Company,
			EmploymentType: workExperience.EmploymentType,
			Location:       workExperience.Location,
			StartDate:      workExperience.StartDate,
			EndDate:        workExperience.StartDate,
		}
		profile.WorkExperience = append(profile.WorkExperience, *workExperience)
	}

	for _, skill := range profile.Skills {
		profile.Skills = append(profile.Skills, skill)
	}

	for _, interest := range profile.Interests {
		profile.Interests = append(profile.Interests, interest)
	}
	err = service.orchestrator.Start(newProfile, oldProfile.Username, oldProfile.FirstName, oldProfile.LastName, oldProfile.IsPrivate)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProfileService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *ProfileService) GetByToken(token string) (*domain.Profile, error) {
	return service.store.GetByToken(token)
}

func (service *ProfileService) GenerateToken(profileId string) (string, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return "", err
	}

	return service.store.GenerateToken(id)
}

func (service *ProfileService) GetLogs() ([]string, error) {
	logPathPrefix := "../../logs/"
	if os.Getenv("OS_ENV") == "docker" {
		logPathPrefix = "./logs/"
	}
	content, err := os.ReadFile(logPathPrefix + "profile_service/profile.log")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
