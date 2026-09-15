package user

import (
	"messenger/internal/domain"
	"github.com/google/uuid"
	"time"
)

type Usecase struct {
	Users []domain.User
}

func (u *Usecase)AddUser(name string) domain.User {
	user := domain.User{ID: uuid.New(), Username: name, CreatedAt: time.Now()}
	u.Users = append(u.Users, user)
	return user
}

func (u *Usecase)IsExist(id uuid.UUID) bool {
	for i := 0; i < len(u.Users); i++ {
		if id == u.Users[i].ID {
			return true
		}
	}
	return false
}

func (u *Usecase)FindUserFromID(id uuid.UUID) *domain.User {
	for i := 0; i < len(u.Users); i++ {
		if id == u.Users[i].ID {
			return &u.Users[i]
		}
	}
	return nil
}