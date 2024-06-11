package productmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Product {
	// Mengambil data produk dari database
	rows, err := config.DB.Query("SELECT products.id, products.name, categories.name as category_name, products.stock, products.description, products.created_at, products.updated_at FROM products JOIN categories ON products.category_id = categories.id")
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data, hentikan program dan cetak pesan kesalaha
		panic(err)
	}

	// Pastikan untuk menutup koneksi ke database setelah selesai digunakan
	defer rows.Close()

	// Membuat variabel untuk menyimpan daftar produk yang akan dikembalikan
	var products []entities.Product

	// Melakukan pengulangan melalui setiap baris hasil query
	for rows.Next() {
		// Membuat variabel untuk menyimpan satu kategori
		var product entities.Product
		// Memindai nilai dari setiap kolom dalam baris hasil query ke variabel category
		if err := rows.Scan(&product.Id, &product.Name, &product.Category.Name, &product.Stock, &product.Description, &product.CreatedAt, &product.UpdatedAt); err != nil {
			panic(err)
		}
		// Menambahkan kategori yang telah dipindai ke dalam slice categories
		products = append(products, product)
	}

	// Mengembalikan slice categories yang berisi semua kategori yang telah diambil dari database
	return products
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec("INSERT INTO products (name, category_id, stock, description, created_at, updated_at) values (?, ?, ?, ?, ?, ?)",
		product.Name, product.Category.Id, product.Stock, product.Description, product.CreatedAt, product.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow("SELECT products.id, products.name, categories.name as category_name, products.stock, products.description, products.created_at, products.updated_at FROM products JOIN categories ON products.category_id = categories.id WHERE products.id =?", id)

	var product entities.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Category.Name, &product.Stock, &product.Description, &product.CreatedAt, &product.UpdatedAt); err != nil {
		panic(err.Error())
	}

	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec("UPDATE products SET name = ?, category_id = ?, stock = ?, description = ?, updated_at = ? WHERE id = ?",
		product.Name, product.Category.Id, product.Stock, product.Description, product.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
