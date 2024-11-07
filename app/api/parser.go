package api

import (
    "encoding/json"
    "fmt"
)

type Response struct {
    Choices []Choice `json:"choices"`
}

type Choice struct {
    MessageContent MessageContent `json:"message"`
}

type MessageContent struct {
    Content string `json:"content"`
}

func ParseResponse(body []byte) (string, error) {
    var response Response
    if err := json.Unmarshal(body, &response); err != nil {
        return "", fmt.Errorf("error unmarshalling JSON: %v", err)
    }

    if len(response.Choices) > 0 {
        return response.Choices[0].MessageContent.Content, nil
    }
    return "", fmt.Errorf("no choices found in the response")
}
