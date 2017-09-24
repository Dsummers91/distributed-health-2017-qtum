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

type CreateContractResponse struct {
	TransactionID string `json:"txid"`
	Sender        string `json:"sender"`
	Hash160       string `json:"hash160"`
	Address       string `json:"address"`
}

var bin = "6060604052341561000f57600080fd5b60405162000d0938038062000d09833981016040528080519060200190919080518201919060200180518201919050506000804260058190555084156100b75760008054806001018281610063919061028c565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061011b565b600180548060010182816100cb919061028c565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b600091505b83518210156101ae576000805480600101828161013d919061028c565b91600052602060002090016000868581518110151561015857fe5b90602001906020020151909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050816001019150610120565b600090505b825181101561024157600180548060010182816101d0919061028c565b9160005260206000209001600085848151811015156101eb57fe5b90602001906020020151909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508060010190506101b3565b33600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050506102dd565b8154818355818115116102b3578183600052602060002091820191016102b291906102b8565b5b505050565b6102da91905b808211156102d65760008160009055506001016102be565b5090565b90565b610a1c80620002ed6000396000f300606060405236156100ce576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630b97bc86146100d357806311b5e6f0146100fc5780633e104a9514610125578063629e4ed01461017a57806366542aa7146101e457806366eae86e1461024757806393066b4c146102895780639ce110d7146102ec5780639f272ee214610341578063b116619d1461036a578063b60d42881461038d578063c24a0f8b146103af578063cedece74146103d8578063e89e4ed614610401575b600080fd5b34156100de57600080fd5b6100e6610467565b6040518082815260200191505060405180910390f35b341561010757600080fd5b61010f61046d565b6040518082815260200191505060405180910390f35b341561013057600080fd5b6101386104a1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561018557600080fd5b61018d6104cb565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101d05780820151818401526020810190506101b5565b505050509050019250505060405180910390f35b34156101ef57600080fd5b610205600480803590602001909190505061056f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561025257600080fd5b6102876004808035600019169060200190919080359060200190919080359060200190919080359060200190919050506105ae565b005b341561029457600080fd5b6102aa600480803590602001909190505061066f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156102f757600080fd5b6102ff6106ae565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561034c57600080fd5b6103546106d4565b6040518082815260200191505060405180910390f35b341561037557600080fd5b61038b600480803590602001909190505061077f565b005b6103956108e2565b604051808215151515815260200191505060405180910390f35b34156103ba57600080fd5b6103c26108eb565b6040518082815260200191505060405180910390f35b34156103e357600080fd5b6103eb6108f1565b6040518082815260200191505060405180910390f35b341561040c57600080fd5b61042260048080359060200190919050506108f7565b60405180876000191660001916815260200186815260200185815260200184815260200183815260200182151515158152602001965050505050505060405180910390f35b60055481565b6000806018603c80600554420381151561048357fe5b0481151561048d57fe5b0481151561049757fe5b0490508091505090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6104d361094f565b6104db61094f565b60006003805490506040518059106104f05750595b90808252806020026020018201604052509150600090505b6003805490508110156105675760038181548110151561052457fe5b906000526020600020906006020160000154828281518110151561054457fe5b906020019060200201906000191690816000191681525050806001019050610508565b819250505090565b60018181548110151561057e57fe5b90600052602060002090016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600380548060010182816105c29190610963565b9160005260206000209060060201600060c060405190810160405280886000191681526020016005548152602001878152602001868152602001858152602001600015158152509091909150600082015181600001906000191690556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050160006101000a81548160ff02191690831515021790555050505050505050565b60008181548110151561067e57fe5b90600052602060002090016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060006106e161046d565b9150600090505b6003805490508110156107795760016003805490500381141561070d5780925061077a565b8160038281548110151561071d57fe5b90600052602060002090600602016002015410801561075e57508160036001830181548110151561074a57fe5b906000526020600020906006020160020154115b1561076e5760018101925061077a565b8060010190506106e8565b5b505090565b600080600061078c61046d565b92506107966106d4565b91506003828154811015156107a757fe5b90600052602060002090600602016004015484101580156107ea57506000836003848154811015156107d557fe5b90600052602060002090600602016002015403115b156108dc57600090505b6001805490508110156108db5760018181548110151561081057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6001805490503073ffffffffffffffffffffffffffffffffffffffff1631606460038681548110151561088457fe5b90600052602060002090600602016003015481151561089f57fe5b04028115156108aa57fe5b049081150290604051600060405180830381858888f1935050505015156108d057600080fd5b8060010190506107f4565b5b50505050565b60006001905090565b60065481565b60045481565b60038181548110151561090657fe5b90600052602060002090600602016000915090508060000154908060010154908060020154908060030154908060040154908060050160009054906101000a900460ff16905086565b602060405190810160405280600081525090565b8154818355818115116109905760060281600602836000526020600020918201910161098f9190610995565b5b505050565b6109ed91905b808211156109e95760008082016000905560018201600090556002820160009055600382016000905560048201600090556005820160006101000a81549060ff02191690555060060161099b565b5090565b905600a165627a7a7230582007693cd5813cb26eb92aad0508ea188849e62536cd3b897596789f5a41c3e63c0029"

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

func DeployContract() []byte {
	entity := getEntity(1)
	fmt.Printf("%s %s %s %s %s %s %s", "qtum-cli", "--regtest", entity.datadir, entity.rpcPort, "createcontract", bin, "50000")
	out, err := exec.Command("qtum-cli", "--regtest", entity.datadir, entity.rpcPort, "createcontract", bin, "50000").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	outAsString := string(out)
	text := strings.Replace(outAsString, "\n", "", -1)
	text = strings.Replace(text, " ", "", -1)
	fmt.Println(text)
	return []byte(text)
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
