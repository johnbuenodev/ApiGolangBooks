package database

import (
	"apiGolang/database/migrations"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro: ", err)
	}
	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	//Rodando migrations para ter mapeado o banco dentro do container docker
	migrations.RunMigrations(db)
}

//Retornar o banco que foi instancia, senão teria varias instancias do banco
func GetDatabase() *gorm.DB {
	return db
}

/* JSON Base para CRUD

{
	"name": "Livro1",
	"description": "Um livro",
	"medium_price": 99.85,
	"author": "um author",
	"img_url": "https://cdn.shopify.com/s/files/1/2022/6883/products/IMG_2002_1024x1024@2x.JPG?v=1538235544"
}

*/
