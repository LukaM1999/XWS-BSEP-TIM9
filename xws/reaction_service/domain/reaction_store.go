package domain

type ReactionStore interface {
	Get(postId string) ([]*Reaction, error)
	Reaction(reaction *Reaction) (*Reaction, error)
	Delete(id string) error
	DeleteAll() error
}
