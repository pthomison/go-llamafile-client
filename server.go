package gollamafileclient

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/commander-cli/cmd"
	"github.com/sirupsen/logrus"
)

func RunMistralLlamafileServer(ctx context.Context, binary string, flags []string) {
	go func() {
		c := cmd.NewCommand(fmt.Sprintf("%v %v", binary, strings.Join(flags, " ")))

		err := c.ExecuteContext(ctx)
		if err != nil {
			panic(err.Error())
		}

		logrus.Info(c.Stdout(), c.Stderr())
	}()

	for {
		waitForRequest := DefaultCompetionRequest()
		waitForRequest.NPredict = 10
		_, err := SendCompletionRequest("http://127.0.0.1:8080/completion", waitForRequest)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
