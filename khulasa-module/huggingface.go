package khulasa

import (
	"net/http"
	"sync"
)

type HuggingFaceConfig struct {
	ApiKey       string
	ModelName    string
	InferenceURL string
	MaxLength    int
	MinLength    int
	Client       *http.Client
	clientmu     sync.Mutex
}
