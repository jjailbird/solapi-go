package main

import (
	"fmt"

	"github.com/jjailbird/solapi-go"
)

func main() {
	client := solapi.NewClient()

	// 조회를 원하는 그룹아이디
	groupId := "G4V20200826105257ADZ4FNAUGO9NL1D"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.GetGroup(groupId)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
