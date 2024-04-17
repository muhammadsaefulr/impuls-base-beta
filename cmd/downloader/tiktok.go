package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	x "mywabot/system"
	"net/http"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "tiktok",
		Cmd:     []string{"tiktok"},
		Tags:    "downloader",
		IsQuery: true,
		Prefix:  true,
		ValueQ:  ".tiktok <link>",
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

			type ResponseData struct {
				Creator string `json:"creator"`
				VidUrl  string `json:"url"`
			}

			type RequestData struct {
				AFormat         string `json:"aFormat"`
				DubLang         bool   `json:"dubLang"`
				FilenamePattern string `json:"filenamePattern"`
				URL             string `json:"url"`
				VQuality        string `json:"vQuality"`
			}

			url := "https://co.wuk.sh/api/json"
			data := RequestData{
				AFormat:         "mp3",
				DubLang:         false,
				FilenamePattern: "classic",
				URL:             m.Query,
				VQuality:        "1080",
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error encoding JSON:", err)
				return
			}

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			fmt.Println("Response Status:", resp.Status)

			var resData ResponseData
			err = json.NewDecoder(resp.Body).Decode(&resData)
			if err != nil {
				fmt.Println("Error decoding response:", err)
				return
			}

			sock.SendVideo(m.From, resData.VidUrl, fmt.Sprint("berhasil !"), *m)
			m.React("✅")
		},
	})
}
