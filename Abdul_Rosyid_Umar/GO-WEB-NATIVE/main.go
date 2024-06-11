package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/logincontroller"
	"go-web-native/controllers/productcontroller"
	"go-web-native/middleware"

	///import homecontroller dari folder homecontroller project go-web-native agar bisa memanggil handler
	// yang ada di dalam homecontroller. jadi tinggall di panggil panggil saja

	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// http.HandleFunc("/login", controller

	//homepage, memanggil function dari homecontroller.welcome, ini adalah contoh pemanggilan handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			logincontroller.LoginPage(w, r)
		case "POST":
			logincontroller.Login(w, r)
		}
	})
	// http.HandleFunc("/dashboard", logincontroller.Dashboard)
	// http.HandleFunc("/user", logincontroller.Dashboard)

	// http.HandleFunc("/", logincontroller.LoginPage)

	// // Route for login submission
	// http.HandleFunc("/-submit", logincontroller.Login)

	// Routes with role-based access control
	http.Handle("/admin/dashboard", middleware.RequireRole("admin", middleware.RequireAuthentication(http.HandlerFunc(logincontroller.Dashboard))))
	http.Handle("/user/dashboard", middleware.RequireRole("user", middleware.RequireAuthentication(http.HandlerFunc(logincontroller.Dashboard))))

	//jika masuk ke dalam /categories maka server akan running fungsi index yang ada di categorycontroller
	http.HandleFunc("/categories", categorycontroller.Index)
	//jika masuk ke dalam /categories/add maka server akan running fungsi Add yang ada di categorycontroller
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)
	http.HandleFunc("/products/detail", productcontroller.Detail)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
