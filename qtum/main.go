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

type Entity struct {
	datadir string
	rpcPort string
	address string
}

func main() {
	startAPI() // Start Server
	var response Response

	out, err := exec.Command("qtum-cli", "--regtest", "-datadir=/home/deon/.qtum3", "-rpcport=2236", "getinfo").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	outAsString := string(out)
	text := strings.Replace(outAsString, "\n", "", -1)
	text = strings.Replace(text, " ", "", -1)

	_ = json.Unmarshal([]byte(text), &response)
	entity := getEntity(1)
	fmt.Printf("TEST: %s %s %s", entity.address, entity.datadir, entity.rpcPort)
	entity = getEntity(2)
	fmt.Printf("TEST: %s %s %s", entity.address, entity.datadir, entity.rpcPort)
	entity = getEntity(3)
	fmt.Printf("TEST: %s %s %s", entity.address, entity.datadir, entity.rpcPort)
	fmt.Scanln()
}

func DeployContract(isSponser bool, _sponsors []string, _individuals []string, _milestoneNames []string, _milestonePayoutDays []uint, _milestonePayoutPercentages []uint, _milestoneSteps []uint) {
	out, err := exec.Command("qtum-cli", "--regtest", "-datadir=/home/deon/.qtum3", "-rpcport=2236", "getinfo").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	outAsString := string(out)
	text := strings.Replace(outAsString, "\n", "", -1)
	text = strings.Replace(text, " ", "", -1)
	return
}

func (entity *Entity) getHexAddress() {
	out, err := exec.Command("qtum-cli", "--regtest", entity.datadir, entity.rpcPort, "getaccountaddress", "x").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	outAsString := string(out)
	text := strings.Replace(outAsString, "\n", "", -1)
	text = strings.Replace(text, " ", "", -1)
	entity.address = text
	return
}

func getEntity(person int) Entity {
	var entity Entity

	if person == 1 {
		entity.rpcPort = "-datadir=/home/deon/.qtum3"
		entity.datadir = "-rpcport=2236"
	}
	if person == 2 {
		entity.rpcPort = "-datadir=/home/deon/.qtum1"
		entity.datadir = "-rpcport=2234"
	}
	if person == 3 {
		entity.rpcPort = "-datadir=/home/deon/.qtum4"
		entity.datadir = "-rpcport=2237"
	}
	entity.getHexAddress()
	return entity
}
