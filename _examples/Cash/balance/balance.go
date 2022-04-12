package main

import (
	"fmt"

	"github.com/jjailbird/solapi-go"
)

func main() {
	client := solapi.NewClient()

	// SetCustomConfig
	/*
		client.Cash.Config = map[string]string{
			"APIKey": "Another API KEY",
		}
	*/

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Cash.Balance()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
