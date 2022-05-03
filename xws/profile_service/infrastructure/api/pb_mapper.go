package api

import (
	pb "dislinkt/common/proto/profile_service"
	"dislinkt/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//Function to return a pb.User from a domain.User
func mapProfileToPb(profile *domain.Profile) *pb.Profile {
	pbProfile := &pb.Profile{
		Id:             profile.Id.Hex(),
		Username:       profile.Username,
		FirstName:      profile.FirstName,
		LastName:       profile.LastName,
		DateOfBirth:    timestamppb.New(profile.DateOfBirth),
		PhoneNumber:    profile.PhoneNumber,
		Email:          profile.Email,
		Gender:         profile.Gender,
		Biography:      profile.Biography,
		Education:      make([]*pb.Education, 0),
		WorkExperience: make([]*pb.WorkExperience, 0),
		Skills:         make([]string, 0),
		Interests:      make([]string, 0),
	}

	for _, education := range profile.Education {
		educationPb := &pb.Education{
			School:       education.School,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			StartDate:    timestamppb.New(education.StartDate),
			EndDate:      timestamppb.New(education.EndDate),
			Grade:        education.Grade,
			Description:  education.Description,
		}
		pbProfile.Education = append(pbProfile.Education, educationPb)
	}

	for _, workExperience := range profile.WorkExperience {
		workExperiencePb := &pb.WorkExperience{
			Title:          workExperience.Title,
			Company:        workExperience.Company,
			EmploymentType: 0,
			Location:       workExperience.Location,
			StartDate:      timestamppb.New(workExperience.StartDate),
			EndDate:        timestamppb.New(workExperience.StartDate),
		}
		pbProfile.WorkExperience = append(pbProfile.WorkExperience, workExperiencePb)
	}

	for _, skill := range profile.Skills {
		pbProfile.Skills = append(pbProfile.Skills, skill)
	}

	for _, interest := range profile.Interests {
		pbProfile.Interests = append(pbProfile.Interests, interest)
	}

	return pbProfile
}

//Function to return a domain.User from a pb.User
func mapPbToProfile(pbProfile *pb.Profile) *domain.Profile {
	profile := &domain.Profile{
		Id:             getObjectId(pbProfile.Id),
		Username:       pbProfile.Username,
		FirstName:      pbProfile.FirstName,
		LastName:       pbProfile.LastName,
		DateOfBirth:    pbProfile.DateOfBirth.AsTime(),
		PhoneNumber:    pbProfile.PhoneNumber,
		Email:          pbProfile.Email,
		Gender:         pbProfile.Gender,
		Biography:      pbProfile.Biography,
		Education:      make([]domain.Education, 0),
		WorkExperience: make([]domain.WorkExperience, 0),
		Skills:         make([]string, 0),
		Interests:      make([]string, 0),
	}
	for _, education := range pbProfile.Education {
		education := &domain.Education{
			School:       education.School,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			StartDate:    education.StartDate.AsTime(),
			EndDate:      education.EndDate.AsTime(),
			Grade:        education.Grade,
			Description:  education.Description,
		}
		profile.Education = append(profile.Education, *education)
	}

	for _, workExperience := range pbProfile.WorkExperience {
		workExperience := &domain.WorkExperience{
			Title:          workExperience.Title,
			Company:        workExperience.Company,
			EmploymentType: workExperience.EmploymentType.String(),
			Location:       workExperience.Location,
			StartDate:      workExperience.StartDate.AsTime(),
			EndDate:        workExperience.StartDate.AsTime(),
		}
		profile.WorkExperience = append(profile.WorkExperience, *workExperience)
	}

	for _, skill := range pbProfile.Skills {
		profile.Skills = append(profile.Skills, skill)
	}

	for _, interest := range pbProfile.Interests {
		profile.Interests = append(profile.Interests, interest)
	}
	return profile
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
