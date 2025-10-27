package controller

import (
	"context"
	pb "resumeadvisor/proto/proto"
	"resumeadvisor/service"
	"strings"
)

type ResumeController struct {
	service *service.ResumeService
	pb.UnimplementedResumeServiceServer
}

func NewResumeController(s *service.ResumeService) *ResumeController {
	return &ResumeController{service: s}
}

func (c *ResumeController) GetFeedback(ctx context.Context, req *pb.ResumeRequest) (*pb.ResumeFeedback, error) {
	suggestions := c.service.GenerateSuggestions(req.Resumetext)
	return &pb.ResumeFeedback{Suggestions: suggestions}, nil
}

func (c *ResumeController) GetDetailedFeedback(req *pb.ResumeRequest, stream pb.ResumeService_GetDetailedFeedbackServer) error {
	lines := strings.Split(req.Resumetext, "\n")
	for _, line := range lines {
		sugg := c.service.GenerateSuggestions(line)
		if err := stream.Send(&pb.ResumeFeedback{Suggestions: sugg}); err != nil {
			return err
		}
	}
	return nil
}
