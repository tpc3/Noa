package misskeyapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/tpc3/Noa/lib/config"
)

var (
	getnotesEndpoint  = "https://" + config.Loadconfig.Misskey.Host + "/api/notes/timeline"
	sendnotesEndpoint = "https://" + config.Loadconfig.Misskey.Host + "/api/notes/create"
)

func MisskeyGetnotesRequest(token string) []string {
	requestBody := GetnotesRequest{
		Limit: 100,
		Token: token,
	}

	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal("Marshal json error: ", err)
	}

	req, err := http.NewRequest("POST", getnotesEndpoint, bytes.NewBuffer(requestJson))
	if err != nil {
		log.Fatal("Creating http request error: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Sending https request error: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("API error(getnotes): ", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Reading body error: ", err)
	}

	response := make([]*NotesResponse, 0)
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Unmarshal json error: ", err)
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

	return resarray
}

func MisskeySendnotesRequest(token string, text string) {
	requestBody := NotesRequest{
		Visibility: "home",
		Text:       text,
		Token:      token,
		LocalOnly:  true,
	}

	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal("Marshaling json error: ", err)
	}

	req, err := http.NewRequest("POST", sendnotesEndpoint, bytes.NewBuffer(requestJson))
	if err != nil {
		log.Fatal("Creating http request error: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Sending https request error: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("API error(sendnotes): ", err)
	}
}
