package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"dasha_puteshestvinichu/accounts/models"
	"dasha_puteshestvinichu/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type bankServer struct {
	proto.UnimplementedBankServer
	db *sql.DB
}

func (bs *bankServer) fetchAccount(name string) (models.Account, error) {
	query := "SELECT name, amount FROM accounts WHERE name=$1"
	row := bs.db.QueryRow(query, name)
	account := models.Account{}
	err := row.Scan(&account.Name, &account.Amount)
	return account, err
}

func (bs *bankServer) addAccount(account models.Account) error {
	query := "INSERT INTO accounts(name, amount) VALUES($1, $2)"
	_, err := bs.db.Exec(query, account.Name, account.Amount)
	return err
}

func (bs *bankServer) removeAccount(name string) error {
	query := "DELETE FROM accounts WHERE name=$1"
	_, err := bs.db.Exec(query, name)
	return err
}

func (bs *bankServer) updateAccountAmount(account models.Account) error {
	query := "UPDATE accounts SET amount=$1 WHERE name=$2"
	_, err := bs.db.Exec(query, account.Amount, account.Name)
	return err
}

func (bs *bankServer) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.Empty, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "empty name")
	}

	if _, err := bs.fetchAccount(req.Name); err == nil {
		return nil, status.Error(codes.AlreadyExists, "account already exists")
	}

	account := models.Account{Name: req.Name, Amount: int(req.Amount)}
	if err := bs.addAccount(account); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (bs *bankServer) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	account, err := bs.fetchAccount(req.Name)
	if err != nil {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	response := &proto.GetAccountResponse{
		Name:   account.Name,
		Amount: int32(account.Amount),
	}
	return response, nil
}

func (bs *bankServer) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.Empty, error) {
	if _, err := bs.fetchAccount(req.Name); err != nil {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	if err := bs.removeAccount(req.Name); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (bs *bankServer) PatchAccount(ctx context.Context, req *proto.PatchAccountRequest) (*proto.Empty, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "empty name")
	}

	if _, err := bs.fetchAccount(req.Name); err != nil {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	account := models.Account{Name: req.Name, Amount: int(req.Amount)}
	if err := bs.updateAccountAmount(account); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func main() {
	const dsn = "host=0.0.0.0 port=1323 dbname=postgres user=postgres password=password"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1323))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterBankServer(s, &bankServer{db: db})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
