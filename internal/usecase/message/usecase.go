package message

import (
	"errors"
	"messenger/internal/domain"
	"messenger/internal/usecase/user"
	"time"

	"github.com/google/uuid"
)

type Usecase struct {
	Messages    []domain.Message
	UserUsecase user.Usecase
}

func (u *Usecase) Create(text string, authorID uuid.UUID) error {
	exist := u.UserUsecase.IsExist(authorID)
	if !exist {
		return errors.New("user not found")
	}
	msg := domain.Message{Text: text, CreatedAt: time.Now(), AuthorID: authorID}
	u.Messages = append(u.Messages, msg)
	return nil
}

func (u *Usecase) ReadMesseng() []domain.Message {
	for i := 0; i < len(u.Messages); i++ {
		user := u.UserUsecase.FindUserFromID(u.Messages[i].AuthorID)
		if user == nil {
			continue
		}
		u.Messages[i].Meta.AuthorName = user.Username
	}
	return u.Messages

}
