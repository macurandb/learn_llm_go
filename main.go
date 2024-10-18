package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms/openai"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("generate", generateCompletion)
	}

	err = router.Run(":8080")
	if err != nil {
		return
	}
}

func generateCompletion(c *gin.Context) {

	var requestData struct {
		Prompt string `json:"prompt"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		c.JSON(500, gin.H{"error": "Missing OpenAI API key"})
		return
	}

	llm, err := openai.New()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	completion, err := llm.Call(context.Background(), requestData.Prompt)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"completion": completion})
}
