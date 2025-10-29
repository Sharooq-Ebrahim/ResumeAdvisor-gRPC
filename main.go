package main

import (
	"log"
	"net"
	"resumeadvisor/controller"
	"resumeadvisor/service"

	"google.golang.org/grpc"

	pb "resumeadvisor/proto/proto"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Println("Falied to listen", err)
	}

	grpcServer := grpc.NewServer()
	resumeService := &service.ResumeService{}
	ResumeController := controller.NewResumeController(resumeService)
	pb.RegisterResumeServiceServer(grpcServer, ResumeController)

	log.Println("ResumeAdvisor (gRPC) server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
