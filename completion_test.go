package gollamafileclient

import (
	"context"
	"fmt"
	"strings"
	"testing"

	usefulgo "github.com/pthomison/go-useful"
)

var (
	url    = "http://127.0.0.1:8080/completion"
	binary = "./mistral-7b-instruct-v0.1-Q4_K_M-server.llamafile"
	flags  = []string{
		"--nobrowser",
	}

	systemPrompt = "I want you to become my Prompt engineer. Your goal is to help me craft the best possible prompt for my needs. The prompt will be used by you, ChatGPT. You will follow the following process:1. Your first response will be to ask me what the prompt should be about. I will provide my answer, but we will need to improve it through continual iterations by going through the next steps.2. Based on my input, you will generate 2 sections, a) Revised prompt (provide your rewritten prompt, it should be clear, concise, and easily understood by you), b) Questions (ask any relevant questions pertaining to what additional information is needed from me to improve the prompt).3. We will continue this iterative process with me providing additional information to you and you updating the prompt in the Revised prompt section until I say we are done."

	prompts = []string{
		"limit your responses to 10 words or less. finish this sentence: The best thing to do during the day is",
		"limit your responses to 100 words or less. finish this sentence: The best thing to do during the day is",
		"limit your responses to 1000 words or less. finish this sentence: The best thing to do during the day is",
		"limit your responses to 10 words or less. What is the best thing to do during the day? Provide a concise answer that includes an actionable suggestion.",
		"limit your responses to 100 words or less. What is the best thing to do during the day? Provide a concise answer that includes an actionable suggestion.",
		"limit your responses to 1000 words or less. What is the best thing to do during the day? Provide a concise answer that includes an actionable suggestion.",
	}
	promptTryCount = 10
)

func TestMain(t *testing.T) {

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	RunMistralLlamafileServer(ctx, "./mistral-7b-server.llamafile", []string{"--nobrowser"})

	printPromptResponsesFn := func(responses []CompetionResponse) string {
		out := "Prompt: "

		out = fmt.Sprintf("%v%v\n", out, responses[0].Prompt)

		for i, resp := range responses {
			out = fmt.Sprintf("%v%v: %v\n", out, i, resp.Content)
		}
		return out
	}

	for _, prompt := range prompts {
		responses := []CompetionResponse{}
		for i := 0; i < promptTryCount; i++ {
			response, err := SendCompletionRequest(url, DefaultCompetionRequestWithPrompt(fmt.Sprintf("%v %v", systemPrompt, prompt)))
			usefulgo.CheckTest(err, t)
			response.Content = strings.Join(strings.Split(response.Content, "\n"), "\\n")
			responses = append(responses, response)
		}
		fmt.Println(printPromptResponsesFn(responses))
	}

}
