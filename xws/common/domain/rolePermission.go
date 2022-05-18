package domain

type RolePermission struct {
	Role        string   `bson:"role"`
	Permissions []string `bson:"permissions"`
}
