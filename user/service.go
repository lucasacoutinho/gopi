package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lucasacoutinho/gopi/user/db"
)

type Service struct {
	r *db.Queries
}

func NewService(r *db.Queries) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) Get(ctx context.Context, id ID) (*User, error) {
	m, err := s.r.Get(ctx, sql.NullInt32{Int32: int32(id), Valid: true})
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	return &User{
		ID:       ID(m.ID.Int32),
		Username: m.Username.String,
		Email:    m.Email.String,
		Password: m.Password.String,
	}, nil
}

func (s *Service) List(ctx context.Context) ([]*User, error) {
	p, err := s.r.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	var users []*User
	for _, j := range p {
		users = append(users, &User{
			ID:       ID(j.ID.Int32),
			Username: j.Username.String,
			Email:    j.Email.String,
			Password: j.Password.String,
		})
	}
	return users, nil
}

func (s *Service) Create(ctx context.Context, Username, Email, Password string) (ID, error) {
	result, err := s.r.Create(ctx, db.CreateParams{
		Username: sql.NullString{
			String: Username,
			Valid:  true,
		},
		Email: sql.NullString{
			String: Email,
			Valid:  true,
		},
		Password: sql.NullString{
			String: Password,
			Valid:  true,
		},
	})
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	return ID(id), nil
}

func (s *Service) Update(ctx context.Context, e *User) error {
	err := s.r.Update(ctx, db.UpdateParams{
		ID: sql.NullInt32{
			Int32: int32(e.ID),
			Valid: true,
		},
		Username: sql.NullString{
			String: e.Username,
			Valid:  true,
		},
		Email: sql.NullString{
			String: e.Email,
			Valid:  true,
		},
		Password: sql.NullString{
			String: e.Password,
			Valid:  true,
		},
	})
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, id ID) error {
	err := s.r.Delete(ctx, sql.NullInt32{
		Int32: int32(id),
		Valid: true,
	})
	if err != nil {
		return fmt.Errorf("error removing user: %w", err)
	}
	return nil
}
