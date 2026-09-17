package router

import (
	"fmt"
	"messenger/internal/mapping"
	"messenger/internal/usecase/message"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	//"encoding/json"
)

func newM(base *mux.Router, u *message.Usecase) *mux.Router {
	routerMess := base.PathPrefix("/mess").Subrouter()

	routerMess.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		//authorIDstr := r.URL.Query().Get("authorID")
		text := r.URL.Query().Get("text")
		cookie, err := r.Cookie("user_id")
		if err != nil {
			// Куки не найдено или другая ошибка
			fmt.Fprintf(w, "Cookie not found")
			http.Error(w, "Cookie not found", http.StatusNotFound)
			return
		}
		authorIDstr := cookie.Value
		if authorIDstr == "" {
			fmt.Fprintf(w, "authorID empty")
			return
		}
		if text == "" {
			fmt.Fprintf(w, "text empty")
			return
		}
		authorID, err := uuid.Parse(authorIDstr)
		if err != nil {
			fmt.Fprintf(w, "wrong authorID")
			return
		}
		err = u.Create(text, authorID)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		fmt.Fprintf(w, "ok")

	})
	routerMess.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		messages := u.ReadMesseng()
		m := mapping.MessageSFromDomainToStringS(messages)
		for i := 0; i < len(messages); i++ {
			fmt.Fprintln(w, m[i])
		}
	})

	return routerMess
}
