package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	//file paths
	inputDir := "content"
	inputLanguage := "en"

	//configs to authenticate
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithSharedConfigProfile("edsoncelio"),
	)

	if err != nil {
		log.Fatal(err)
	}

	//reading the content dir
	entries, err := os.ReadDir(fmt.Sprintf("%s/%s/", inputDir, inputLanguage))
	if err != nil {
		log.Fatal(err)
	}

	//check if exists or create the output dir
	if _, err := os.Stat("translated/pt-br/"); os.IsNotExist(err) {
		err := os.MkdirAll("translated/pt-br/", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	//initialize client
	client := translate.NewFromConfig(cfg)

	for _, e := range entries {
		fmt.Printf("%s\n", e.Name())
		fileInfo, err := os.Stat(fmt.Sprintf("%s/%s/%s", inputDir, inputLanguage, e.Name()))
		if err != nil {
			log.Fatal(err)
		}

		// iterate only in files
		if !fileInfo.IsDir() {
			text, err := os.ReadFile(fmt.Sprintf("%s/%s/%s", inputDir, inputLanguage, e.Name()))
			if err != nil {
				fmt.Print(err)
			}

			//translate file
			demoText, err := client.TranslateText(context.TODO(), &translate.TranslateTextInput{
				SourceLanguageCode: aws.String("en"),
				TargetLanguageCode: aws.String("pt"),
				Text:               aws.String(string(text)),
			})

			outErr := os.WriteFile(fmt.Sprintf("translated/pt-br/%s", filepath.Base(e.Name())), []byte(*demoText.TranslatedText), 0644)
			if outErr != nil {
				log.Fatal(outErr)
			}
			fmt.Println("Done!")
		}

	}

}
