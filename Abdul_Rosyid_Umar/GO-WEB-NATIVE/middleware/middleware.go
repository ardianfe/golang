package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func init() {
	store.MaxAge(1800) // Set MaxAge to 30 minutes
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   1800, // 30 minutes
		HttpOnly: true,
	}
}

func RequireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		userRole, ok := session.Values["role"].(string)
		log.Printf("RequireRole - UserRole: %s, ExpectedRole: %s, Ok: %v", userRole, role, ok)
		if !ok || userRole != role {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func RequireAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		userID, ok := session.Values["user_id"].(int)
		log.Printf("RequireAuthentication - UserID: %d, Ok: %v", userID, ok)
		if !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
