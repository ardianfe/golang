package productcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"Products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	product := productmodel.Detail(id)
	data := map[string]any{
		"Product": product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"Categories": categories,
		}

		temp.Execute(w, data)
		return
	}

	if r.Method == "POST" {
		var product entities.Product
		var validationErrors []string

		stockStr := r.FormValue("stock")
		CidStr := r.FormValue("category_id")

		if stockStr == "" {
			validationErrors = append(validationErrors, "Stock tidak boleh kosong")
		}

		if CidStr == "" {
			validationErrors = append(validationErrors, "Category ID tidak boleh kosong")
		}

		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			validationErrors = append(validationErrors, "Stock harus berupa angka yang valid")
		}

		Cid, err := strconv.Atoi(CidStr)
		if err != nil {
			validationErrors = append(validationErrors, "Category ID harus berupa angka yang valid")
		}

		product.Name = r.FormValue("name")
		product.Description = r.FormValue("description")

		if product.Name == "" {
			validationErrors = append(validationErrors, "Nama tidak boleh kosong")
		}
		if product.Description == "" {
			validationErrors = append(validationErrors, "Deskripsi tidak boleh kosong")
		}

		if len(validationErrors) > 0 {
			temp, _ := template.ParseFiles("views/product/create.html")
			categories := categorymodel.GetAll()
			data := struct {
				Errors     []string
				Product    entities.Product
				Categories []entities.Category
			}{
				Errors:     validationErrors,
				Product:    product,
				Categories: categories,
			}
			temp.Execute(w, data)
			return
		}

		product.Stock = int64(stock)
		product.Category.Id = uint(Cid)
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productmodel.Create(product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()
		data := map[string]any{
			"Product":  product,
			"Category": categories,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product
		var validationErrors []string

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			validationErrors = append(validationErrors, "Stock tidak boleh kosong")
		}

		Cid, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		product.Id = uint(id)
		product.Name = r.FormValue("name")
		product.Category.Id = uint(Cid)
		product.Stock = int64(stock)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		if product.Name == "" {
			validationErrors = append(validationErrors, "Nama tidak boleh kosong")
		}
		if product.Description == "" {
			validationErrors = append(validationErrors, "Deskripsi tidak boleh kosong")
		}

		if len(validationErrors) > 0 {
			temp, _ := template.ParseFiles("views/product/edit.html")
			categories := categorymodel.GetAll()
			data := struct {
				Errors   []string
				Product  entities.Product
				Category []entities.Category
			}{
				Errors:   validationErrors,
				Product:  product,
				Category: categories,
			}
			temp.Execute(w, data)
			return
		}

		if ok := productmodel.Update(id, product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := productmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
