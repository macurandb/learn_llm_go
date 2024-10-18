# Learn LLM Go

This project demonstrates how to build a simple API using [Gin](https://gin-gonic.com/) and OpenAI's language model to generate text completions. The project is written in Go and interacts with OpenAI's API to handle user prompts and generate responses.

## Features

- RESTful API built with Gin
- Handles POST requests to generate text completions using OpenAI's API
- Reads API keys from `.env` file for secure management

## Prerequisites

- Go (1.19 or newer)
- OpenAI API Key
- Git

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/learn_llm_go.git
cd learn_llm_go
```

## 2. Install dependencies

```bash
go mod tidy
``` 

### 3. Set up your environment variables

Create a `.env` file in the root directory of the project with the following content:

```bash
OPENAI_API_KEY=your_openai_api_key_here
``` 

Replace `your_openai_api_key_here` with your actual OpenAI API key, which you can get from [OpenAI](https://platform.openai.com/account/api-keys).

### 4. Build and run the project

Use the provided Makefile to build and run the project:

```bash
make run
``` 

The API will start on `http://localhost:8080`.

### 5. Example request

You can make a POST request to the API to generate text completions using the following endpoint:

- **URL:** `http://localhost:8080/api/v1/generate`
- **Method:** `POST`
- **Payload:**

```json
{
  "prompt": "Explain the theory of relativity"
}
```
You can use tools like `curl`, Postman, or any HTTP client to test the API:

```bash
curl -X POST http://localhost:8080/api/v1/generate \
  -H "Content-Type: application/json" \
  -d '{"prompt": "Explain the theory of relativity"}'
``` 

The response will look like:

```json
{
  "completion": "The theory of relativity, proposed by Albert Einstein, revolutionized the understanding of space, time, and gravity. It consists of two parts: special relativity and general relativity..."
}
```

## Contact

If you have any questions, feel free to reach out to [macurandb](mailto:,macurandb@gmail.com).



