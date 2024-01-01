package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	x "mywabot/system"
	"net/url"
	"os"
	"os/exec"
)

type DataDetail struct {
	Username     string `json:"username"`
	PostURL      string `json:"postUrl"`
	ThumbnailURL string `json:"thumbnailUrl"`
	PostDesc     string `json:"postDesc"`
}

// JSONData represents the overall structure of the JSON file
type JSONData struct {
	DataDetail []DataDetail `json:"dataDetail"`
}

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "igreels",
		Cmd:     []string{"igreels <query_link>"},
		Tags:    "downloader",
		Desc:    "Example downloader using py",
		Prefix:  true,
		IsOwner: true,
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

			cmdActivate := exec.Command("/bin/sh", "-c", "source "+"lib/env/bin/activate")
			cmdActivate.Stdout = os.Stdout
			cmdActivate.Stderr = os.Stderr
			err := cmdActivate.Run()
			if err != nil {
				fmt.Println("Error activating virtual environment:", err)
				return
			}

			// m.Reply("Example Res : " + url.QueryEscape(m.Query))
			cmdPython := exec.Command("python", "lib/downloader.py", url.QueryEscape(m.Query))
			cmdPython.Stdout = os.Stdout
			cmdPython.Stderr = os.Stderr

			// Pass the environment variables to the Python script
			cmdPython.Env = os.Environ()

			err = cmdPython.Run()
			if err != nil {
				fmt.Println("Error running Python script:", err)
				return
			}

			filePath := "result_instagram_download.json"

			// Read JSON data from the file
			data, err := readJSONFromFile(filePath)
			if err != nil {
				fmt.Println("Error reading JSON file:", err)
				return
			}

			// Print the read data
			for _, detail := range data.DataDetail {
				// fmt.Printf("Username: %s\n", detail.Username)
				fmt.Printf("Post URL: %s\n", detail.PostURL)
				// fmt.Printf("Thumbnail URL: %s\n", detail.ThumbnailURL)
				// fmt.Printf("Post Description: %s\n", detail.PostDesc)
				sock.SendVideo(m.From, detail.PostURL, "\n\ndescription\n\n :"+detail.PostDesc, *m)
			}

			err = deleteJSONFile(filePath)
			if err != nil {
				fmt.Println("Error deleting JSON file:", err)
				return
			}

			// fmt.Println("Python script executed successfully")
		},
	})
}

func readJSONFromFile(filePath string) (JSONData, error) {
	// Read the entire file
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return JSONData{}, err
	}

	// Create a variable to store the decoded JSON data
	var jsonData JSONData

	// Unmarshal the JSON data into the JSONData struct
	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		return JSONData{}, err
	}

	return jsonData, nil
}

func deleteJSONFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
