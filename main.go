package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	model *genai.GenerativeModel
)

func init() {
	godotenv.Load(".env")
	gclient, _ := genai.NewClient(context.Background(), option.WithAPIKey(os.Getenv("API_KEY")))
	model = gclient.GenerativeModel("gemini-pro")
}

func main() {
	for {
		fmt.Printf("Enter Your Prompt : ")
		prompt := prompt()
		if prompt == "" {
			fmt.Println("ERROR : Prompt is empty")
			continue
		}
		fmt.Println(prompt)
		result(prompt)
	}
}

func prompt() string {
	var prompt string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	prompt = scanner.Text()

	return prompt
}

func result(prompt string) {
	resp, _ := model.GenerateContent(context.Background(), genai.Text(prompt))

	//Decorate the response
	fmt.Println("***********************************************")
	fmt.Println("************* Gemini Response ****************")
	fmt.Println("***********************************************")
	fmt.Println()

	//Print the response
	// fmt.Println(resp)
	// fmt.Println("-------------------------------")
	if resp == nil {
		fmt.Println("Sorry It is not available")
	} else {
		fmt.Println(resp.Candidates[0].Content.Parts[0])
	}

	//Close the decoration
	fmt.Println()
	fmt.Println("***********************************************")
	fmt.Println("*************** End of Response ***************")
	fmt.Println("***********************************************")
}
