package resolver

import (
	"context"
	"fmt"
	"log"

	database "github.com/jerrynim/gql-leave/db"
	"github.com/jerrynim/gql-leave/graph/model"
	"github.com/jerrynim/gql-leave/jwt"
	"github.com/jerrynim/gql-leave/localtime"

	"golang.org/x/crypto/bcrypt"
)

//? Query

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}


//? Mutation

func (r *mutationResolver) SignUp(ctx context.Context, email string, password string, name string, bio *string, position string, contact string, profileImage string, birthday string, remainLeaves int) (*model.AuthResponse, error) {

	// loggedUser := ctx.Value("user")
    // fmt.Print(loggedUser,"유저??")
	
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
	
	birthdayDate,parseErr :=localtime.ParseTime(birthday)
	if parseErr !=nil{
		panic(fmt.Errorf("생년월일 시간 파싱 에러"))
	}
	now,parseErr :=localtime.GetTime()


	user := model.User{
		Email : email,
		Name : name,
		Password:string(hash),
		Position: position,
		Contact: contact,
		Bio:bio,
		Birthday: birthdayDate,
		RemainLeaves: remainLeaves,
		CreatedAt: now,
		UpdatedAt: now,
	}
	
	
	createErr:= db.Create(&user).Error;

	if createErr != nil {
		panic(fmt.Errorf("유저 생성 에러", createErr.Error()))
	}
	fmt.Print("새로운 유저",user.Name,"id: ",user.ID,"\n")
	token,jwtErr:=jwt.GenerateToken(fmt.Sprint(user.ID))
	
	
	if jwtErr != nil {
		panic(fmt.Errorf("토큰 생성 에러"))
	}

	response:= model.AuthResponse{Token: token,User:&user}
	return &response, nil
}


func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Me(ctx context.Context) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}