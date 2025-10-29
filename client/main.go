package main

import (
	"context"
	"log"
	"time"

	pb "resumeadvisor/proto/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewResumeServiceClient(conn)

	resumeText := `Sharooq
                  Software Engineer
      Looking for opportunities in backend development`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.GetFeedback(ctx, &pb.ResumeRequest{
		Username:   "Sharooq",
		Resumetext: resumeText,
	})
	if err != nil {
		log.Fatalf("Error calling GetFeedback: %v", err)
	}
	log.Println("Feedback suggestions:")
	for _, s := range resp.Suggestions {
		log.Println("-", s)
	}
}
