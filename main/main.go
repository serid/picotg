package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"picotg/api"
	"picotg/model"
	"strconv"
)

func main() {
	err := main2()

	if err != nil {
		panic(err)
	}
}

// Function sends message to user <cli arg1>
func main2() error {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		text := "usage: sendtext <userid>"
		fmt.Println(text)
		//return errors.New()
	}

	userid, err := strconv.Atoi(argsWithoutProg[0])

	if err != nil {
		return err
	}

	text, err := ioutil.ReadFile("./message.txt")

	if err != nil {
		return err
	}

	token, err := ioutil.ReadFile("./token.txt")

	if err != nil {
		return err
	}

	theApi := api.NewApi(string(token))

	arg := model.SendMessage{
		ChatId: userid,
		Text:   string(text),
	}

	bytes, err := json.Marshal(arg)

	if err != nil {
		return err
	}

	fmt.Println(string(bytes))

	returnedBytes, err := theApi.Request("sendMessage", bytes)

	if err != nil {
		return err
	}

	fmt.Println(string(returnedBytes))
	/*
		var msg model.Message

		err = json.Unmarshal(returnedBytes[:len(returnedBytes)-30], &msg)

		if err != nil {
			return err
		}
	*/
	return nil
}
