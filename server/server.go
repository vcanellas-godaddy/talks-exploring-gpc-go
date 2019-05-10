package main

import (
	"context"
	"net"
	"time"

	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
	pbBank "github.com/vladimirvasconcelos/labGrpc/_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Withdrawal(ctx context.Context, req *pbBank.Transaction) (*pbBank.Account, error) {
	// Receive transaction values
	debitAccount := req.GetOrigin()
	value := req.GetValue()

	// Operation
	debitAccount.Balance.TotalBalance -= value
	return debitAccount, nil

}

func (s *server) Deposit(ctx context.Context, req *pbBank.Transaction) (*pbBank.Account, error) {

	// Receive transaction values
	creditAccount := req.GetDestination()
	value := req.GetValue()
	// Operation
	creditAccount.Balance.TotalBalance += value

	return creditAccount, nil
}

func (s *server) Transfer(ctx context.Context, req *pbBank.Transaction) (creditedAcc *pbBank.Account, err error) {
	println("Transfer begin on server")
	// Receive transaction values
	debitAccount := req.GetOrigin()
	creditAccount := req.GetDestination()
	value := req.GetValue()
	// Operation
	debitAccount.Balance.TotalBalance -= value
	creditAccount.Balance.TotalBalance += value
	color.Printf("Transfer ends on server %v", time.Now())
	return creditAccount, nil
}

func main() {
	println("Starting server...")
	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pbBank.RegisterTransactionServiceServer(s, &server{})
	println("🔅 Server started")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("💀Server is dead: %v", err)
	}
}
