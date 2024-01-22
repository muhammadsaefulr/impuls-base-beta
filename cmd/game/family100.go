package cmd

import (
	x "mywabot/system"
)

type JSONData struct {
	Soal    string   `json:"soal"`
	Jawaban []string `json:"jawaban"`
}

func init() {
	x.NewCmd(&x.ICmd{
		Name: "family100",
	})
}
