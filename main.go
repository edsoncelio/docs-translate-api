package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithSharedConfigProfile("edsoncelio"),
	)

	if err != nil {
		panic(err)
	}

	text, err := os.ReadFile("content/en/kubernetes.md") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	client := translate.NewFromConfig(cfg)

	demoText, err := client.TranslateText(context.TODO(), &translate.TranslateTextInput{
		SourceLanguageCode: aws.String("en"),
		TargetLanguageCode: aws.String("pt"),
		Text:               aws.String(string(text)),
	})

	outFile := os.WriteFile("content/pt/kubernetes.md", []byte(*demoText.TranslatedText), 0644)
	fmt.Printf("Done! - %s", outFile)
}
