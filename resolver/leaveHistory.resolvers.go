package resolver

import (
	"context"
	"fmt"
	"log"
	"time"

	database "github.com/jerrynim/gql-leave/db"
	"github.com/jerrynim/gql-leave/graph/model"
)

//? Query
func (r *queryResolver) GetMyLeaves(ctx context.Context) ([]*model.LeaveHistory, error) {
	userContext := ctx.Value("user")

	if userContext ==nil{
		panic(fmt.Errorf("잘못된 토큰입니다."))
	}
	loggedUser:= userContext.(model.User)

	db, err := database.GetDatabase()

	if err != nil {
		log.Println("Unable to connect to database", err)
		panic(fmt.Errorf("데이터 베이스 연결 에러"))
	}
	defer db.Close()
	
	var leaves []*model.LeaveHistory
	

	db.Limit(10).Find(&leaves, &model.LeaveHistory{UserID: loggedUser.ID})

	var result []*model.LeaveHistory

	for _,leave := range leaves{
		var user model.User
		var approver model.User
		fmt.Print(&leave,"마맘마")
		userResult:= db.First(&user, &model.User{ID:leave.UserID})
		if userResult.Error != nil{
			fmt.Print("유저 찾기 에러")
		}

		approverResult:= db.First(&approver, &model.User{ID:loggedUser.ID})
		if approverResult.Error != nil{
			fmt.Print("승인 유저 찾기 에러")
		}
		
		leave.User=&user
		leave.Approver=&approver
		
		result = append(result, leave)
	}

	return result,nil
}

func (r *queryResolver) GetAppliedLeaves(ctx context.Context) ([]*model.LeaveHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

//? Mutation
func (r *mutationResolver) MakeLeaveHistory(ctx context.Context, date string, reason *string, typeArg model.LeaveType) (bool, error) {

	userContext := ctx.Value("user")

	if userContext ==nil{
		panic(fmt.Errorf("잘못된 토큰입니다."))
	}
	loggedUser:= userContext.(model.User)

	leaveDate,leaveDateParseErr:= time.Parse(time.RFC3339, date)
	if leaveDateParseErr !=nil{
		panic(fmt.Errorf("생년월일 시간 파싱 에러"))
	}

	db, err := database.GetDatabase()

	if err != nil {
		log.Println("Unable to connect to database", err)
		panic(fmt.Errorf("데이터 베이스 연결 에러"))
	}
	defer db.Close()


	var existLeave model.LeaveHistory

	db.Find(&existLeave, &model.LeaveHistory{Date: leaveDate})
	
	if existLeave.ID !=0{
		panic(fmt.Errorf("해당 일에 이미 휴가가 존재합니다."))
	}
	

	leave:= model.LeaveHistory{
		User: &loggedUser,
		Date:leaveDate,
		Reason: reason,
		Type: typeArg,
		Status: "applied",
	}
	createErr:= db.Create(&leave).Error;

	if createErr != nil {
		panic(fmt.Errorf("휴가 내역 생성 에러", createErr.Error()))
	}

	return true, nil
}


func (r *mutationResolver) ChangeLeaveStatus(ctx context.Context, leaveID int, status model.LeaveStatus) (bool, error) {


	fmt.Print(ctx)	
	userContext := ctx.Value("user")

	if userContext ==nil{
		panic(fmt.Errorf("잘못된 토큰입니다."))
	}
	loggedUser:= userContext.(model.User)

	if loggedUser.Role =="normal"{
		panic(fmt.Errorf("권한이 없습니다."))
	}

	db, err := database.GetDatabase()

	if err != nil {
		log.Println("Unable to connect to database", err)
		panic(fmt.Errorf("데이터 베이스 연결 에러"))
	}
	defer db.Close()

	var leave model.LeaveHistory

	result := db.Find(&leave, &model.LeaveHistory{ID: leaveID})
	if result.Error != nil{
		panic(fmt.Errorf("휴가자 조회 에러"))
	}
	
	leave.Status=status
	leave.ApproverID=&loggedUser.ID
	leave.UpdatedAt=time.Now()
	fmt.Print(status)
	db.Save(leave)

	return true,nil
}