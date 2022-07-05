package startup

import (
	"dislinkt/job_offer_service/domain"
	"time"
)

var jobOffers = []*domain.JobOffer{
	{
		Id:          1,
		ProfileId:   "",
		Company:     "Vega IT",
		Position:    "Software Engineer",
		Description: "We are looking for a software engineer to join our team.",
		Criteria:    "We are looking for a software engineer to join our team.",
		Skills:      []string{"Ruby", "PHP", "React", "Go"},
		CreatedAt:   time.Now(),
	},
	{
		Id:          2,
		ProfileId:   "",
		Company:     "Quantox",
		Position:    "System Administrator",
		Description: "We are looking for a software engineer to join our team.",
		Criteria:    "We are looking for a software engineer to join our team.",
		Skills:      []string{"Shell", "Linux", "Docker"},
		CreatedAt:   time.Now(),
	},
	{
		Id:          3,
		ProfileId:   "",
		Company:     "HTEC",
		Position:    "Software Engineer",
		Description: "We are looking for a software engineer to join our team.",
		Criteria:    "We are looking for a software engineer to join our team.",
		Skills:      []string{"JavaScript", "Golang", "PostgreSQL"},
		CreatedAt:   time.Now(),
	},
	{
		Id:          4,
		ProfileId:   "",
		Company:     "NovaLite",
		Position:    "Software Engineer",
		Description: "We are looking for a software engineer to join our team.",
		Criteria:    "We are looking for a software engineer to join our team.",
		Skills:      []string{"Node", "MongoDB", "AWS"},
		CreatedAt:   time.Now(),
	},
	{
		Id:          5,
		ProfileId:   "",
		Company:     "LagerSoft",
		Position:    "Software Developer",
		Description: "We are looking for a software engineer to join our team.",
		Criteria:    "We are looking for a software engineer to join our team.",
		Skills:      []string{"Java", "C#", "Go"},
		CreatedAt:   time.Now(),
	},
}
