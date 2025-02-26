package insights

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"moonlogs/internal/lib/serialize"
	"net/http"
	"net/url"
	"strings"
)

const openRouterUri = "https://openrouter.ai/api/v1/chat/completions"
const acceptKey = "Accept"
const authKey = "Authorization"
const bearerPrefix = "Bearer "

type OpenRouterInsightsAdapter struct {
	token string
	model string
}

func NewOpenRouterInsightsAdapter(token string, model string) *OpenRouterInsightsAdapter {
	return &OpenRouterInsightsAdapter{token: token, model: model}
}

type OpenRouterResponse struct {
	Choices []*OpenRouterChoice `json:"choices"`
}

type OpenRouterChoice struct {
	Message *OpenRouterContentMessage `json:"message"`
}

type OpenRouterContentMessage struct {
	Content string `json:"content"`
}

func (a *OpenRouterInsightsAdapter) GenerateContent(ctx context.Context, prompt string, httpClient *http.Client) (string, error) {
	req, err := a.buildRequest(ctx, openRouterUri, a.model, a.token, requestMethod, contentTypeKey, contentTypeValue, prompt)
	if err != nil {
		return "", fmt.Errorf("preparing request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}

	// OpenRouter returns strange responses with a lot of new lines at the beginning of response body.
	// Trim leading spaces.
	respStr := string(b)
	respStr = strings.TrimSpace(respStr)

	var openRouterResp OpenRouterResponse
	err = serialize.NewJSONDecoder(bytes.NewBufferString(respStr)).Decode(&openRouterResp)
	if err != nil {
		return "", fmt.Errorf("decoding open router response: %w", err)
	}

	if len(openRouterResp.Choices) == 0 {
		return "", errors.New("empty choices")
	}

	choice := openRouterResp.Choices[0]
	if choice.Message == nil {
		return "", errors.New("empty choice message")
	}

	content := choice.Message.Content
	if len(content) == 0 {
		return "", errors.New("empty choice message content")
	}

	return content, nil
}

type OpenRouterRequest struct {
	Model    string              `json:"model"`
	Messages []OpenRouterMessage `json:"messages"`
}

type OpenRouterMessage struct {
	Role    string              `json:"role"`
	Content []OpenRouterContent `json:"content"`
}

type OpenRouterContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (a *OpenRouterInsightsAdapter) buildRequest(ctx context.Context, baseURI, modelName, token, requestMethod, contentTypeKey, contentTypeValue, propmt string) (*http.Request, error) {
	_, err := url.Parse(baseURI)
	if err != nil {
		return nil, fmt.Errorf("parsing base uri: %w", err)
	}

	content := OpenRouterContent{Text: propmt, Type: "text"}
	message := OpenRouterMessage{Role: "user", Content: []OpenRouterContent{content}}
	payload := OpenRouterRequest{Model: modelName, Messages: []OpenRouterMessage{message}}

	jsonData, err := serialize.JSONMarshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshalling payload body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, requestMethod, baseURI, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("building request with timeout: %w", err)
	}

	req.Header.Set(contentTypeKey, contentTypeValue)
	req.Header.Set(acceptKey, contentTypeValue)
	req.Header.Set(authKey, bearerPrefix+token)

	return req, nil
}
