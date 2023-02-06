package user

import "context"

// ID person ID
type ID int

// Person defines a person
type User struct {
	ID       ID
	Username string
	Email    string
	Password string
}

// UseCase defines the domain use case
type UseCase interface {
	List(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id ID) (*User, error)
	Create(ctx context.Context, Username, Email, Password string) (ID, error)
	Update(ctx context.Context, e *User) error
	Delete(ctx context.Context, id ID) error
}
