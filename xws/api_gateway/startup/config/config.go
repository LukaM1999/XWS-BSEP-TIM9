package config

type Config struct {
	Port         string
	SecurityHost string
	SecurityPort string
	ProfileHost  string
	ProfilePort  string
	CommentHost  string
	CommentPort  string
	ReactionHost string
	ReactionPort string
}

func NewConfig() *Config {
	return &Config{
		Port:         "8000",
		SecurityHost: "security_service",
		SecurityPort: "8001",
		ProfileHost:  "profile_service",
		ProfilePort:  "8001",
		CommentHost:  "comment_service",
		CommentPort:  "8001",
		ReactionHost: "reaction_service",
		ReactionPort: "8001",
	}
}
