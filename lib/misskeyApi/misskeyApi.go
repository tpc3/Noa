package misskeyapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/tpc3/Noa/lib/config"
)

var (
	getnotesEndpoint  = "https://" + config.Loadconfig.Misskey.Host + "/api/notes/timeline"
	sendnotesEndpoint = "https://" + config.Loadconfig.Misskey.Host + "/api/notes/create"
)

func MisskeyGetnotesRequest(token string) ([]string, error) {
	check := true
	requestBody := GetnotesRequest{
		Limit: 100,
		Token: token,
	}

	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		check = false
	}

	req, err := http.NewRequest("POST", getnotesEndpoint, bytes.NewBuffer(requestJson))
	if err != nil {
		check = false
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		check = false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		check = false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		check = false
	}

	response := make([]*NotesResponse, 0)
	err = json.Unmarshal(body, &response)
	if err != nil {
		check = false
	}

	var resarray []string
	for i := 0; i < 100; i++ {
		if response[i].Text == "" {
			if response[i].Renote.Text == "" {
				continue
			}
			resarray = append(resarray, response[i].Renote.Text)
		} else {
			resarray = append(resarray, response[i].Text)
		}
	}

	if check {
		return resarray, nil
	} else {
		return resarray, errors.New("API error")
	}

}

func MisskeySendnotesRequest(token string, text string) error {
	check := true
	requestBody := NotesRequest{
		Visibility: "home",
		Text:       text,
		Token:      token,
		LocalOnly:  true,
	}

	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		check = false
	}

	req, err := http.NewRequest("POST", sendnotesEndpoint, bytes.NewBuffer(requestJson))
	if err != nil {
		check = false
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		check = false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		check = false
	}

	if check {
		return nil
	} else {
		return errors.New("API error")
	}

}
