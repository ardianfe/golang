package categorycontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category
		var validationErrors []string

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if category.Name == "" {
			validationErrors = append(validationErrors, "Nama kategori tidak boleh kosong.")
		}

		if len(validationErrors) > 0 {
			temp, _ := template.ParseFiles("views/category/create.html")
			data := struct {
				Errors []string
			}{
				Errors: validationErrors,
			}
			temp.Execute(w, data)
			return
		}

		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(id)
		data := map[string]any{
			"Category": category,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.Category
		var erorValidation []string

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category.Id = uint(id)
		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if category.Name == "" {
			erorValidation = append(erorValidation, "Nama tidak boleh kosong.")
		}

		if len(erorValidation) > 0 {
			temp, _ := template.ParseFiles("views/category/edit.html")
			data := struct {
				Errors   []string
				Category entities.Category
			}{
				Errors:   erorValidation,
				Category: category,
			}
			temp.Execute(w, data)

			return
		}

		if ok := categorymodel.Update(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := categorymodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}