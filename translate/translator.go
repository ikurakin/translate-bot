package translate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type (
	Translator interface {
		Translate(srcLang, dstLang, text string) ([]string, error)
	}

	translator struct {
		apiURL string
		client *http.Client
	}

	translateResponse struct {
		Code int      `json:"code"`
		Lang string   `json:"lang"`
		Text []string `json:"text"`
	}
)

var translateResponseErrCodes = map[int]string{
	401: "Invalid API key",
	402: "Blocked API key",
	404: "Exceeded the daily limit on the amount of translated text",
	413: "Exceeded the maximum text size",
	422: "The text cannot be translated",
	501: "The specified translation direction is not supported",
}

func New(key string, u string) Translator {
	return translator{
		apiURL: u + "?key=" + key,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (tr translator) Translate(srcLang, dstLang, text string) ([]string, error) {
	translateURL := tr.apiURL + "&lang=" + srcLang + "-" + dstLang + "&text=" + url.QueryEscape(text)
	resp, err := tr.client.Get(translateURL)
	if err != nil {
		return []string{}, fmt.Errorf("failed to translate %s", translateURL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if status, ok := translateResponseErrCodes[resp.StatusCode]; ok {
			return []string{}, fmt.Errorf("translation api error: %s", status)
		}
		return []string{}, fmt.Errorf("translations api response code is not 200 OK: %s", resp.Status)
	}

	var result translateResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result.Text, err
}
