package main

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	pbBank "github.com/vladimirvasconcelos/labGrpc/_pb"
	"google.golang.org/grpc"
)

var Account001 = &pbBank.Account{
	Id: 001,
	Balance: &pbBank.AccountBalance{
		TotalBalance: 700.0,
	},
}

var Account002 = &pbBank.Account{
	Id: 002,
	Balance: &pbBank.AccountBalance{
		TotalBalance: 0.10,
	},
}

func main() {
	println("Starting client...")

	client, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ts := pbBank.NewTransactionServiceClient(client)
	println("Client Started")

	transaction := &pbBank.Transaction{
		Destination:Account001,
		Value: 300.00,
	}
	println("[BEFORE]",transaction.Destination.String())
	DepositValue(ts, transaction)
	println("\n[Press enter]\n")
	fmt.Scanln()
	transaction2 := &pbBank.Transaction{
		Origin:Account001,
		Destination:Account002,
		Value: 50.00,
	}

	println("[BEFORE]",transaction2.Origin.String())
	println("[BEFORE]",transaction2.Destination.String())
	TransferValue(ts, transaction2)

}

//DepositValue
func DepositValue(ts pbBank.TransactionServiceClient, transaction *pbBank.Transaction) {
	creditResponse, err := ts.Deposit(context.Background(), transaction)
	if err != nil {
		log.Fatal(err.Error())
	}
	println("[AFTER]", creditResponse.String())
}

func TransferValue(ts pbBank.TransactionServiceClient, transaction *pbBank.Transaction) {
	creditedAccount, err := ts.Transfer(context.Background(), transaction)
	if err != nil {
		log.Fatal(err.Error())
	}
	println("[AFTER]",transaction.Origin.String())
	println("[AFTER]",creditedAccount.String())
}

