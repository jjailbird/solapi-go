package main

import (
	"fmt"

	"github.com/jjailbird/solapi-go"
)

func main() {
	client := solapi.NewClient()

	// 발송요청을 할 그룹 아이디
	groupId := "G4V20200824233846LAI57B7SKJACJ40"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.SendGroup(groupId)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
