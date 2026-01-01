package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	_ = godotenv.Load()

	// define and parse command line files
	apiMode := flag.Bool("server_api_mode", false, "Run in API Mode")
	port := flag.String("server_port", "8080", "Port for the server")
	filePath := flag.String("file_path", "", "File path for files to summarize")
	summaryRate := flag.Float64("summary_rate", 0.3, "Summary rate for extractive summary")
	targetPercent := flag.Float64("summary_percent", 30.0, "Target summary for abstractive summarization")
	summaryType := flag.String("summary_type", "extractive", "Summary type (extractive, abstractive_openapi, abstractive_hugging, hybrid)")

	// openai config
	openAIKey := flag.String("openai_key", os.Getenv("OPENAI_API_KEY"), "OpenAI API Key")
	openAIModel := flag.String("openai_model", os.Getenv("OPENAI_MODEL"), "OpenAI model to use")
	openAIBaseURL := flag.String("openai_url", os.Getenv("OPENAI_URL"), "OpenAI base url for OpenAI requests")

	// Hugging Face config
	huggingFaceKey := flag.String("hf_key", os.Getenv("HUGGING_FACE_KEY"), "HuggingFace API Key")
	huggingFaceModel := flag.String("hf_model", "", "HuggingFace model to use")
	huggingFaceURL := flag.String("hf_url", "", "HuggingFace url for inference API")
	maxLength := flag.Int("max_length", 0, "Maximum length for hugging face summary (0 to auto)")
	minLength := flag.Int("min_length", 0, "Minimum length for hugging face summary (0 to auto)")

	flag.Parse()

	// Create a config from the parsed flags

	config := Config{
		FilePath:         *filePath,
		SummaryRate:      *summaryRate,
		TargetPercent:    *targetPercent,
		SummarizerType:   *summaryType,
		OpenAIKey:        *openAIKey,
		OpenAIModel:      *openAIModel,
		OpenAIBaseURL:    *openAIBaseURL,
		HuggingFaceKey:   *huggingFaceKey,
		HuggingFaceModel: *huggingFaceModel,
		HuggingFaceURL:   *huggingFaceURL,
		MaxLength:        *maxLength,
		MinLength:        *minLength,
	}

	if *apiMode {
		fmt.Printf("Starting API server on port %v\n", *port)
	} else if config.FilePath == "" {
		fmt.Printf("Error: please provide valid file name with `%s` flag", "-file_path")
		flag.Usage()
		os.Exit(1)
	}
}
