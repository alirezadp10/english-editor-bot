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
    apiURL     = "https://api.deepinfra.com/v1/openai/chat/completions"
    systemRole = "Act as an English teacher to check and correct the sentences in your responses. Provide the revised version in formal and informal(casual) form, list corrected errors in bullet form. For example, your answer format should be as follows, without additional introductory text or section titles:\n\n<b>👔 اگه بخوای رسمی بگی:</b>\n[Corrected form]\n\n<b>🦦 اگه بخوای یه خودمونی بگی:</b>\n[Casual form]\n\nیعنی: [translation in persian]\n\n<b>🚧 مواردی که اصلاح کردم برات:</b>\n<blockquote>[Corrected errors in bullet list]</blockquote>"
)

func CreateRequestBody(userInput string) []byte {
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
