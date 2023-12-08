package gollamafileclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type CompetionRequest struct {
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
	TopK        float64 `json:"top_k"`
	TopP        float64 `json:"top_p"`
	MinP        float64 `json:"min_p"`
	NPredict    float64 `json:"n_predict"`
	NKeep       float64 `json:"n_keep"`
}

type CompetionResponse struct {
	Prompt          string `json:"prompt"`
	TokensEvaluated int    `json:"tokens_evaluated"`
	TokensPredicted int    `json:"tokens_predicted"`
	Content         string `json:"content"`
	Model           string `json:"model"`
}

func DefaultCompetionRequest() CompetionRequest {
	return CompetionRequest{
		Prompt:      "",
		Temperature: 1.5,
		TopK:        40,
		TopP:        0.5,
		MinP:        0.05,
		NPredict:    1000,
		NKeep:       0,
	}
}

func DefaultCompetionRequestWithPrompt(prompt string) CompetionRequest {
	return CompetionRequest{
		Prompt:      prompt,
		Temperature: 1.5,
		TopK:        40,
		TopP:        0.5,
		MinP:        0.05,
		NPredict:    1000,
		NKeep:       0,
	}
}

func SendCompletionRequest(server string, request CompetionRequest) (CompetionResponse, error) {
	cr, err := json.Marshal(request)
	if err != nil {
		return CompetionResponse{}, err
	}

	resp, err := http.Post(server, "application/json", bytes.NewBuffer(cr))
	if err != nil {
		return CompetionResponse{}, err
	}

	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return CompetionResponse{}, err
	}

	m := map[string]interface{}{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return CompetionResponse{}, err
	}

	r := CompetionResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return CompetionResponse{}, err
	}

	return r, nil
}
