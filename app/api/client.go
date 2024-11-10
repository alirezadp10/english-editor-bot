package api

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

const (
    apiURL = "https://api.deepinfra.com/v1/openai/chat/completions"
)

func CreateRequestBody(userInput, systemRole string) []byte {
    requestBody := map[string]interface{}{
        "model": "meta-llama/Meta-Llama-3.1-70B-Instruct",
        "messages": []map[string]string{
            {"role": "system", "content": systemRole},
            {"role": "user", "content": userInput},
        },
    }

    fmt.Println(requestBody)
    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        log.Fatalf("Error marshalling JSON: %v", err)
    }
    return jsonData
}

func SendRequest(jsonData []byte, apiKey string) ([]byte, error) {
    req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, fmt.Errorf("error creating request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error sending request: %v", err)
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}
