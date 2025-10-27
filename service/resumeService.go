package service

import "strings"

type ResumeService struct{}

func (s *ResumeService) GenerateSuggestions(text string) []string {
	suggestion := []string{}

	if len(text) < 50 {
		suggestion = append(suggestion, "Resume is too short. Add more details .")
	}

	if !strings.Contains(strings.ToLower(text), "experience") {
		suggestion = append(suggestion, "Include your  professional experience")
	}

	if !strings.Contains(strings.ToLower(text), "skills") {
		suggestion = append(suggestion, "List your relevent skills")
	}

	if len(suggestion) == 0 {
		suggestion = append(suggestion, "Resume looks good.")
	}

	return suggestion
}
