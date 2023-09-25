package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"proyecto-golang/application/entities"
	"proyecto-golang/pkg/helpers/configloader"
)

var Instance *gorm.DB
var err error

func Connect(conf configloader.Config) {

	pgConnString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Conn.Host, conf.Conn.Username, conf.Conn.Password, conf.Conn.DatabaseName, conf.Conn.Port)

	Instance, err = gorm.Open(postgres.Open(pgConnString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Module{})
	insertInitialDate()
	log.Println("Database Migration Completed...")
}

func insertInitialDate() {

	var modules = []entities.Module{
		{Name: "Inspecciones", Url: "/dashboard/inspeccion", Icon: "mdi-book-check", Active: true},
		{Name: "Asistencia", Url: "/dashboard/asistencia", Icon: "mdi-map-marker-distance", Active: true},
		{Name: "Siniestros", Url: "/dashboard/siniestro", Icon: "mdi-wrench", Active: true},
		{Name: "NPS", Url: "/dashboard/nps", Icon: "mdi-star", Active: true}}

	Instance.Create(&modules)
}
