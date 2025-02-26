package insights

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/lib/serialize"
	"net/http"
	"net/url"
)

// TODO: Consider if this should be customizable
const geminiModel = "gemini-1.5-flash"
const baseURI = "https://generativelanguage.googleapis.com/v1beta/models/"
const action = "generateContent"
const requestMethod = http.MethodPost
const contentTypeKey = "Content-Type"
const contentTypeValue = "application/json"

type GeminiInsightsAdapter struct {
	token string
}

func NewGeminiInsightsAdapter(token string) *GeminiInsightsAdapter {
	return &GeminiInsightsAdapter{token: token}
}

type GeminiResponse struct {
	Candidates []*GeminiCandidate `json:"candidates"`
	Error      *GeminiError       `json:"error"`
}

type GeminiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type GeminiCandidate struct {
	Content *GeminiContent `json:"content"`
}

func (g *GeminiInsightsAdapter) GenerateContent(ctx context.Context, prompt string, httpClient *http.Client) (string, error) {
	req, err := g.buildRequest(ctx, baseURI, geminiModel, g.token, action, requestMethod, contentTypeKey, contentTypeValue, prompt)
	if err != nil {
		return "", fmt.Errorf("preparing request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	var geminiResp GeminiResponse
	err = serialize.NewJSONDecoder(resp.Body).Decode(&geminiResp)
	if err != nil {
		return "", fmt.Errorf("decoding gemini response: %w", err)
	}

	if geminiResp.Error != nil {
		return "", fmt.Errorf("error from gemini: %s", geminiResp.Error.Message)
	}

	if len(geminiResp.Candidates) == 0 {
		return "", errors.New("empty candidates")
	}

	cand := geminiResp.Candidates[0]
	if cand.Content == nil {
		return "", errors.New("empty content candidates")
	}

	if len(cand.Content.Parts) == 0 {
		return "", errors.New("empty content parts")
	}
	part := cand.Content.Parts[0]

	return part.Text, nil
}

type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

type GeminiContent struct {
	Parts []*GeminiPart `json:"parts"`
}

type GeminiPart struct {
	Text string `json:"text"`
}

func (g *GeminiInsightsAdapter) buildRequest(ctx context.Context, baseURI, modelName, token, action, requestMethod, contentTypeKey, contentTypeValue, propmt string) (*http.Request, error) {
	uri := baseURI + modelName + ":" + action + "?key=" + token
	_, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("parsing base uri: %w", err)
	}

	part := GeminiPart{Text: propmt}
	content := GeminiContent{Parts: []*GeminiPart{&part}}
	payload := GeminiRequest{Contents: []GeminiContent{content}}

	jsonData, err := serialize.JSONMarshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshalling payload body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, requestMethod, uri, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("building request with timeout: %w", err)
	}

	req.Header.Set(contentTypeKey, contentTypeValue)

	return req, nil
}
