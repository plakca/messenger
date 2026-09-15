package router

import (
	"fmt"
	"messenger/internal/usecase/message"
	"net/http"
	"github.com/gorilla/mux"
	"time"
)


func newA(base *mux.Router) *mux.Router {
	routerAuth := base.PathPrefix("/auth").Subrouter()

	u := message.Usecase{}

	routerAuth.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		if username == "" {
			fmt.Fprintf(w, "username empty")
			return
		}
		user := u.UserUsecase.AddUser(username)
		cookie := &http.Cookie{
			Name:     "user_id",
			Value:    user.ID.String(),
			Path:     "/",
			SameSite: http.SameSiteStrictMode,   // Защита от CSRF
			MaxAge:   3600,                      // Срок действия 1 час
			Expires:  time.Now().Add(time.Hour), // Абсолютное время истечения
		}
		http.SetCookie(w, cookie)
		fmt.Fprint(w, user)
	})

	return routerAuth
}

