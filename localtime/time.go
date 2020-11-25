package localtime

import (
	"fmt"
	"time"
)

func GetTime() (string,error) {
    loc, err := time.LoadLocation("Asia/Seoul")
    if err != nil {
        fmt.Println(err)
        return "",err
    }
    time.Local = loc
    formatTime := time.Now().Format(time.RFC3339)
    return formatTime,err
}

func ParseTime(receiveTime string) (string,error){
   parsed,parseErr:= time.Parse(time.RFC3339, receiveTime)
   if parseErr !=nil{
       return "",parseErr
   }
   formatted:= parsed.Format(time.RFC3339)
   return formatted,nil
}