package voicetext

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client interface {
	Tts(text string, speaker string, o *ttsOptions) ([]byte, error)
}

type client struct {
	apiKey string
	apiUrl string
}

func NewClient(apiKey string) Client {
	return &client{apiKey, "https://api.voicetext.jp/%s/%s"}
}

func (c *client) Tts(text string, speaker string, o *ttsOptions) ([]byte, error) {
	params := map[string]string{}
	params["text"] = text
	params["speaker"] = speaker
	if o != nil {
		o.addOption(params)
	}
	endPoint := fmt.Sprintf(c.apiUrl, "v1", "tts")
	content, err := c.post("POST", endPoint, params)
	if err != nil {
		return []byte{}, err
	}
	return content, nil
}

func (c *client) post(method string, endPoint string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, endPoint, strings.NewReader(createValues(params).Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.apiKey, "")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var result struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if err = json.NewDecoder(res.Body).Decode(&result); err != nil || len(result.Error.Message) <= 0 {
			return []byte{}, fmt.Errorf("Invalid status: %s", res.Status)
		}
		return []byte{}, fmt.Errorf("Invalid status: %s, %s", res.Status, result.Error.Message)
	}
	return ioutil.ReadAll(res.Body)
}

func createValues(params map[string]string) url.Values {
	values := map[string][]string{}
	for k, v := range params {
		values[k] = []string{v}
	}
	return values
}
