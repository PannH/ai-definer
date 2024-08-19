package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/definition/:lang/:term", getTermDefinition)

	router.Run("localhost:8080")
}

func getTermDefinition(c *gin.Context) {
	lang := c.Param("lang")
	term := c.Param("term")

	client := openai.NewClient(os.Getenv("OPENAI_KEY"))

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: `
					You are an assistant that provides term and phrase definitions in different languages.

					You will be given a JSON object matching the following schema :
					{
						"term": "string" // The term or phrase to define
						"lang": "string" // The language of the term or phrase (code ISO 639-1)
					}

					Your goal is to provide data about the term or phrase in the requested language. Your response must be a JSON object matching the following schema :
					{
						"term": "string" // The term or phrase
						"lang": "string" // The language of the term or phrase (code ISO 639-1)
						"definition": "string" // The definition of the term or phrase
						"type": "string" // The type of word (noun, verb, adjective, etc.)
						"pronunciation": "string" // The pronunciation of the term or phrase using the International Phonetic Alphabet
					}

					If the term or phrase seems to be wrong, like a set of random characters, simply return the following error :
					{
						"error": "Unknown term"
					}

					If the language does not match the ISO 639-1 standard, return the following error :
					{
						"error": "Unknown language, please use ISO 639-1 language codes"
					}
				`,
			},
			{
				Role: openai.ChatMessageRoleUser,
				Content: fmt.Sprintf(`{"term": "%s", "lang": "%s"}`, term, lang),
			},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if _, ok := data["error"]; ok {
		c.JSON(http.StatusBadRequest, data)
	} else {
		c.JSON(http.StatusOK, data)
	}

}