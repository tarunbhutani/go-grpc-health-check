package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/grpc_server/client/pkg"
	pb "github.com/grpc_server/health_info"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port   = flag.Int("port", 8443, "The server port")
	addr   = flag.String("addr", "localhost:50051", "the address to connect to")
	client pb.HealthInfoServiceClient
)

func main() {
	envAddr := os.Getenv("GRPC_ADDR")
	fmt.Println("envAddr ", envAddr)
	if envAddr != "" {
		fmt.Println("Using address from environment variable:", envAddr)
		// Update the port value
		*addr = envAddr
	}

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client = pb.NewHealthInfoServiceClient(conn)

	r := mux.NewRouter()
	r.HandleFunc("/health-info", HandleHealthInfo).Methods(http.MethodGet)
	r.HandleFunc("/service-info", HandleServiceInfo).Methods(http.MethodGet)
	http.Handle("/", r)
	server := pkg.RestfulServer{
		Port:    *port,
		Handler: r,
	}
	deamon := pkg.NewDaemonServer()
	deamon.Add(&server)
	deamon.Start()
}

func HandleHealthInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Health Info")
	w.Header().Add("Content-Type", "application/json")
	response, err := client.CheckHealth(r.Context(), &pb.HealthCheckRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func HandleServiceInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Server Info")
	w.Header().Add("Content-Type", "applicaiton/json")
	response, err := client.GetServiceInfo(r.Context(), &pb.ServiceInfoRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)

}

func validateAddr(addr string) error {
	// You can add custom validation logic here based on your requirements
	// For a basic check, you can attempt to dial the address
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("unable to dial the address: %v", err)
	}
	defer conn.Close()

	return nil
}
