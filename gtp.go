package gtp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 这里 github上收集的一些key，经过测试都已经无法使用了
var Keys = strings.Split(`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
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
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
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
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
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
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
sk-uaVxaKNMvobxyRkramoIjtT3BlbkFJjEOjcT1gj3cG9C2CcQ5
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
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
sk-kpHABG8aOsxSch5pA7pSLosxBImxjbb5SC1dnTU0ntNl17Nz
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
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
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
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
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
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
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
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
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
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
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
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
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
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
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`, "\n")

const BASEURL = "https://api.openai.com/v1/"

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChoiceItem           `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
}

type ChoiceItem struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	Logprobs     int    `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}

// Completions gtp文本模型回复
//curl https://api.openai.com/v1/completions
//-H "Content-Type: application/json"
//-H "Authorization: Bearer your chatGPT key"
//-d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
func Completions(msg string, key string) (string, error) {
	requestBody := ChatGPTRequestBody{
		Model:            "text-davinci-003",
		Prompt:           msg,
		MaxTokens:        1024,
		Temperature:      0.7,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	//log.Printf("request gtp json string : %v", string(requestData))
	req, err := http.NewRequest("POST", BASEURL+"completions", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := key
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("gtp api status code not equals 200,code is %d", response.StatusCode))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	//log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	var reply string
	if len(gptResponseBody.Choices) > 0 {
		reply = gptResponseBody.Choices[0].Text
	}
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}
