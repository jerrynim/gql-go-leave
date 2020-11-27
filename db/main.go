package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jerrynim/gql-leave/graph/model"
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
	db.Exec(`
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'leave_type') THEN
			CREATE TYPE leave_type AS ENUM ('day', 'morning', 'afternoon');
		END IF;

		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'leave_status') THEN
			CREATE TYPE leave_status AS ENUM ('applied', 'accepted', 'rejected');
		END IF;

		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
			CREATE TYPE user_role AS ENUM ('master', 'manager', 'normal');
		END IF;
		
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
			CREATE TYPE user_status AS ENUM ('inOffice', 'resign');
   		END IF;
	END$$;
	`)	
	db.AutoMigrate(&model.User{}, &model.LeaveHistory{})
}
