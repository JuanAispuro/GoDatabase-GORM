package main

import "github.com/JuanAispuro/GoDatabase-GORM/storage"

func main() {
	driver := storage.Postgres

	storage.New(driver)
}
