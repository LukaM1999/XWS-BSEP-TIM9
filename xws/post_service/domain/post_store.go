package domain

type PostStore interface {
	Get(id string) (*Post, error)
	GetProfilePosts(profileId string) ([]*Post, error)
	GetConnectionPosts(profileId string) ([]*Post, error)
	Create(post *Post) error
	CreateConnection(connection *Connection) error
	Update(id string, post *Post) error
	Delete(id string) error
	DeleteAll() error
}
