package homecontroller

import (
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	// validasi jika url tidak dikenali makan akan memunculkan eror page not found
	if r.URL.Path != "/dashboard" {
		http.NotFound(w, r)
		return
	}
	// menampilkan index html di view/home/index.html
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
