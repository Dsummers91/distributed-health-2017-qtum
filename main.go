package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Response struct {
	Version int `json:version"`
}

func main() {
	var response Response
	out, err := exec.Command("qtum-cli", "getnetworkinfo").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	outAsString := string(out)
	text := strings.Replace(outAsString, "\n", "", -1)
	text = strings.Replace(text, " ", "", -1)

	fmt.Println(text)
	_ = json.Unmarshal([]byte(text), &response)
	fmt.Println(response.Version)
}
