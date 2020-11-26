package resolver

import (
	"context"
	"fmt"
	"log"
	"time"

	database "github.com/jerrynim/gql-leave/db"
	"github.com/jerrynim/gql-leave/graph/model"
	"github.com/jerrynim/gql-leave/jwt"

	"golang.org/x/crypto/bcrypt"
)

//? Query

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}


//? Mutation

func (r *mutationResolver) SignUp(ctx context.Context, email string, password string, name string, bio *string, department string, position string, workSpace string, contact string, birthday string, enteredDate string, remainLeaves int) (*model.AuthResponse, error) {

	// loggedUser := ctx.Value("user")
    // fmt.Print(loggedUser,"유저??")
	
	db, err := database.GetDatabase()

	if err != nil {
		log.Println("Unable to connect to database", err)
		panic(fmt.Errorf("데이터 베이스 연결 에러"))
	}
	defer db.Close()
	
	var existUser model.User
	db.Find(&existUser, &model.User{Email: email})

	if existUser.ID !=0{
		panic(fmt.Errorf("이미 존재하는 이메일 입니다."))
	}


	//* 비밀번호 bcrypt 해쉬 하기
	hash,hashErr := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	
	if hashErr !=nil{
		panic(fmt.Errorf("비밀번호 해싱 에러"))
	}
	
	birthdayDate,birthdayParseErr:= time.Parse(time.RFC3339, birthday)
	if birthdayParseErr !=nil{
		panic(fmt.Errorf("생년월일 시간 파싱 에러"))
	}

	parsedEnteredDate,enterDateParseErr:= time.Parse(time.RFC3339, enteredDate)
	if enterDateParseErr !=nil{
		panic(fmt.Errorf("생년월일 시간 파싱 에러"))
	}


	user := model.User{
		Email : email,
		Password:string(hash),
		Name : name,
		Bio:bio,
		Department: department,
		Position: position,
		WorkSpace: workSpace,
		Contact: contact,
		Birthday: birthdayDate,
		EnteredDate: parsedEnteredDate,
		RemainLeaves: remainLeaves,
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

	db, err := database.GetDatabase()

	if err != nil {
		log.Println("Unable to connect to database", err)
		panic(fmt.Errorf("데이터 베이스 연결 에러"))
	}
	defer db.Close()

	var user model.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		panic(fmt.Errorf("해당 이메일의 유저를 찾을 수 없습니다."))
	  }

	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passwordErr !=nil{
		panic(fmt.Errorf("비밀번호가 일치하지 않습니다."))
	}

	token,jwtErr:=jwt.GenerateToken(fmt.Sprint(user.ID))
	
	
	if jwtErr != nil {
		panic(fmt.Errorf("토큰 생성 에러"))
	}

	 result:= model.AuthResponse{
		 Token: token,
		 User: &user,
	 }
	 return &result,nil
}

func (r *mutationResolver) Me(ctx context.Context) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}