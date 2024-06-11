package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {
	// Mengambil data kategori dari database
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data, hentikan program dan cetak pesan kesalaha
		panic(err)
	}

	// Pastikan untuk menutup koneksi ke database setelah selesai digunakan
	defer rows.Close()

	// Membuat variabel untuk menyimpan daftar kategori yang akan dikembalikan
	var categories []entities.Category

	// Melakukan pengulangan melalui setiap baris hasil query
	for rows.Next() {
		// Membuat variabel untuk menyimpan satu kategori
		var category entities.Category
		// Memindai nilai dari setiap kolom dalam baris hasil query ke variabel category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}
		// Menambahkan kategori yang telah dipindai ke dalam slice categories
		categories = append(categories, category)
	}

	// Mengembalikan slice categories yang berisi semua kategori yang telah diambil dari database
	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec("INSERT INTO categories (name, created_at, updated_at)value (?, ?, ?)",
		category.Name, category.CreatedAt, category.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	Lastinsertidd, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return Lastinsertidd > 0

}

func Detail(id int) entities.Category {
	row := config.DB.QueryRow("SELECT id, name FROM Categories WHERE id = ?", id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec("UPDATE categories SET name = ?, updated_at = ? WHERE id = ?", category.Name, category.UpdatedAt, id)
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
	_, err := config.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}
