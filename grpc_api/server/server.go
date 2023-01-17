package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"grpc/config"
	"grpc/interceptors"
	pb "grpc/proto"
	"log"
	"net"

	"gorm.io/gorm"
)

var err error

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "root@123",
		DBName:   "User",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

type User struct {
	Id       string
	Name     string
	MobileNo string
	Address  string
	Username string
	Password string
}

func init() {

	DB, err = gorm.Open(mysql.Open(DbURL(BuildDBConfig())), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connected...")

	config.ApiLogger()
}

func (u *User) TableName() string {
	return "User"
}

type Server struct {
	pb.UnimplementedLoginServer
}

func NewServer() *Server {
	return &Server{}
}

func (*Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	var user User
	res := DB.Find(&user, "username = ? AND password = ?", username, password)

	if res.RowsAffected == 0 {
		return &pb.LoginResponse{
			Result: "Failed",
		}, nil
	}
	return &pb.LoginResponse{
		Result: "Login",
	}, nil

}

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.ServerInterceptor))
	pb.RegisterLoginServer(s, NewServer())

	log.Printf("Server listening at %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
