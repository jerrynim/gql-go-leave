package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/AnjaneyuluBatta505/gin-graphql-postgres/graph/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
}


func getDatabaseUrl() string {
	host:= os.Getenv("DB_HOST")
	port,err:= strconv.Atoi(os.Getenv("DB_PORT"))
	user:= os.Getenv("DB_USER")
	name:= os.Getenv("DB_NAME")
	password:= os.Getenv("DB_PASSWORD")
	if err !=nil {
		log.Printf("DB_PORT Atoi 에러")
		return ""
	}
	config := dbConfig{host, port, user, name, password}
	
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.host, config.port, config.user, config.dbname, config.password)
}

func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", getDatabaseUrl())
	return db, err
}

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Question{})
	if !db.HasTable(&model.Question{}) {
		log.Printf("question 테이블이 없어요.")
		db.CreateTable(&model.Question{})
	}
	if !db.HasTable(&model.Choice{}) {
		db.CreateTable(&model.Choice{})
		db.Model(&model.Choice{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")
	}
}
