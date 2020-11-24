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

func (r *mutationResolver) SignUp(ctx context.Context, email string, password string, name string, bio *string, profileImage string, birthday *string) (string, error) {
		
	db, err := database.GetDatabase()

	if err != nil {
		log.Println("Unable to connect to database", err)
		panic(fmt.Errorf("데이터 베이스 연결 에러"))
	}
	defer db.Close()
	
	//* 비밀번호 bcrypt 해쉬 하기
	hash,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	
	if err !=nil{
		panic(fmt.Errorf("비밀번호 해슁 에러"))
	}
	now := time.Now().String()
	
	user := model.User{}
	user.Email = email
	user.Name = name
	user.Password=string(hash)
	user.Bio=bio
	user.ProfileImage="https://api.miniintern.com/images/profile/profile_image_default.svg"
	birthday= &now
	
	db.Create(&user)
	return "access_token", nil
}
