package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	pb "github.com/kajikentaro/gRPC-test/grpc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func getEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

func connectDB() *sqlx.DB {
	dataSource := fmt.Sprintf("%s:%s@(%s)/%s",
		getEnv("MYAPP_MYSQL_USER", ""),
		getEnv("MYSQL_ROOT_PASSWORD", ""),
		getEnv("MYAPP_MYSQL_HOSTNAME", ""),
		getEnv("MYAPP_MYSQL_DB_NAME", ""),
	)
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	db.Exec("CREATE TABLE user(id int AUTO INCREMENT PRIMARY KEY, name TEXT, email TEXT)")
	db.Exec("CREATE TABLE data(id int AUTO INCREMENT PRIMARY KEY, user_id int, value TEXT)")
	return db
}

func main() {
	db := connectDB()
	fmt.Println(db)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
