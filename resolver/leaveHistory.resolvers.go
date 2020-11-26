package resolver

import (
	"context"
	"fmt"
	"log"
	"time"

	database "github.com/jerrynim/gql-leave/db"
	"github.com/jerrynim/gql-leave/graph/model"
)
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