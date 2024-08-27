package olly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MyLLM struct{}

func (llm *MyLLM) GenerateResponse(question string) (string, error) {
	answer, err := sendRequest(question)
	if err != nil {
		return "", err
	}

	// Check if the answer is incomplete (you may need to adjust this based on your API's behavior)
	if strings.HasSuffix(answer, "...") || len(answer) >= 4000 { // Assuming 4000 is close to your API's limit
		answer += "\n[Note: The response was truncated due to length limitations.]"
	}

	return answer, nil
}

func sendRequest(question string) (string, error) {
	url := "https://ollybackend.ambitiousflower-4724f605.centralindia.azurecontainerapps.io/api/generate"
	payload := map[string]string{"question": question}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response map[string]string
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	answer, ok := response["result"]
	if !ok {
		return "", fmt.Errorf("no answer found in response")
	}
	return answer, nil
}
