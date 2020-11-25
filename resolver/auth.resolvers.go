package resolver

import (
	"context"
	"fmt"
	"log"
	"time"

	database "github.com/jerrynim/gql-leave/db"
	"github.com/jerrynim/gql-leave/graph/model"
	"golang.org/x/crypto/bcrypt"
)

//? Query

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}


//? Mutation



func (r *mutationResolver) SignUp(ctx context.Context, email string, password string, name string, bio *string, position string, profileImage string, birthday string, remainLeaves int) (string, error) {

	db, err := database.GetDatabase()

	if err != nil {
		log.Println("Unable to connect to database", err)
		panic(fmt.Errorf("데이터 베이스 연결 에러"))
	}
	defer db.Close()
	
	//* 비밀번호 bcrypt 해쉬 하기
	hash,hashErr := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	
	if hashErr !=nil{
		panic(fmt.Errorf("비밀번호 해싱 에러"))
	}
	
	birthdayDate,parseErr :=time.Parse(time.RFC3339, birthday)
	if parseErr !=nil{
		panic(fmt.Errorf("생년월일 시간 파싱 에러"))
	}
	user := model.User{
		Email : email,
		Name : name,
		Password:string(hash),
		Position: position,
		Bio:bio,
		Birthday: birthdayDate,
		RemainLeaves: 15,
	}
	
	
	db.Create(&user)
	return "access_token", nil
}
