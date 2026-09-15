package mapping

import (
	"messenger/internal/domain"
	"fmt"
	"time"
)

func MessageFromDomainToString(m domain.Message) string{
	//Азамат: Привет! 21:10
	return fmt.Sprintf("%v: %s (%v)", m.Meta.AuthorName, m.Text, m.CreatedAt.Format(time.DateTime))
}

func MessageSFromDomainToStringS(messages []domain.Message) []string {
	mess := []string{}
	for i := 0; i < len(messages); i++ {
		mess = append(mess ,MessageFromDomainToString(messages[i]))
	}
	return mess
}