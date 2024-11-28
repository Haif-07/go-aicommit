package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ysicing/openai/openai"
)

func main() {
	locale := "zh"
	maxLength := 500
	commitType := CommitTypeConventional
	prompt := generatePrompt(locale, maxLength, commitType)

	diff, err := GetRepoDiff("./")
	if err != nil {
		log.Fatalf("Error getting repo diff: %v", err)
	}
	if len(diff) == 0 {
		// 这里可以设置一个默认的提交消息
		fmt.Println("first commit😊")
		return
	}

	client, err := openai.New(
		openai.WithToken(os.Getenv("OPENAI_API")),
		openai.WithProvider(openai.DEEPSEEK),
		openai.WithModel(openai.DeepseekChat),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Completion(context.Background(), prompt, diff)
	if err != nil {
		if strings.Contains(resp.Content, "yes") {
			log.Printf("rejected")
		} else {
			log.Printf("approved")
		}
	}
	// log.Printf("content:%s, prompt:%d,completion:%d,total:%d", resp.Content, resp.Usage.PromptTokens, resp.Usage.CompletionTokens, resp.Usage.TotalTokens)
	fmt.Printf("%s", resp.Content)

}
