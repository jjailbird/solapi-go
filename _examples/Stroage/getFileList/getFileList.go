package main

import (
	"fmt"

	"github.com/jjailbird/solapi-go"
)

func main() {
	client := solapi.NewClient()

	// SetCustomConfig
	/*
		client.Storage.Config = map[string]string{
			"APIKey": "Another API KEY",
		}
	*/

	// 검색조건값
	params := make(map[string]string)
	params["limit"] = "1"
	params["type"] = "MMS"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Storage.GetFileList(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
