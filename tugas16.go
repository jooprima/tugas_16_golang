package main

import "fmt"
import "database/sql"
import _ "mysql-master"

type data struct {
	id      int
	nama    string
	umur    int
	jabatan string
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_tugas16golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func tampil_data() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from 	tbl_karyawan")

	// rows, err := db.Query("select * from barang")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []data

	for rows.Next() {
		var each = data{}

		var err = rows.Scan(&each.id, &each.nama, &each.umur, &each.jabatan)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range result {
		fmt.Println(each.nama, each.umur, each.jabatan)
	}
}

func main() {
	tampil_data()
}
