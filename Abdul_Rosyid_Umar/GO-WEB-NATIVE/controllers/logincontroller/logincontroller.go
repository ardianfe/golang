package logincontroller

import (
	"go-web-native/models/usermodel"
	"log"
	"net/http"
	"text/template"

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

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/auth/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		user, err := usermodel.Authenticate(username, password)
		if err != nil {
			tmpl, _ := template.ParseFiles("views/auth/login.html")
			data := map[string]string{
				"Error": "Invalid username or password",
			}
			tmpl.Execute(w, data)
			return
		}

		session, _ := store.Get(r, "session")
		session.Values["user_id"] = user.ID
		session.Values["username"] = user.Username
		session.Values["role"] = user.Role
		session.Save(r, w)

		if user.Role == "admin" {
			http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		}
		if user.Role == "user" {
			http.Redirect(w, r, "/user/dashboard", http.StatusSeeOther)
		}
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	userID, ok := session.Values["user_id"].(int)
	username, usernameOk := session.Values["username"].(string)
	role, roleOk := session.Values["role"].(string)
	if !ok || userID == 0 || !usernameOk || !roleOk {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	log.Printf("Dashboard accessed - UserID: %d, Username: %s, Role: %s", userID, username, role)

	data := map[string]interface{}{
		"Username": username,
		"Role":     role,
	}

	var tmpl *template.Template
	var err error
	if role == "admin" {
		tmpl, err = template.ParseFiles("views/home/index.html")
	} else {
		tmpl, err = template.ParseFiles("views/home/user.html")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
