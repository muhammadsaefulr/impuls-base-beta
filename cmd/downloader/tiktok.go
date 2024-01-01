package cmd

import (
	"fmt"
	x "mywabot/system"
	"net/url"
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

			type ApiData struct {
				Title string `json:"title"`
				Cover string `json:"cover"`
				Play  string `json:"play"`
			}

			var Res struct {
				Creator       string  `json:"creator"`
				Code          int     `json:"code"`
				Msg           string  `json:"msg"`
				ProcessedTime float64 `json:"processed_time"`
				Data          ApiData `json:"data"`
			}

			// Struktur untuk representasi data bagian "data" dalam JSON

			err := x.GetResult("https://skizo.tech/api/tiktok?apikey=zpfnzf10zkrpvb&url="+url.QueryEscape(m.Query), &Res)
			if err != nil {
				m.Reply(fmt.Sprint(err))
				return
			}
			sock.SendVideo(m.From, Res.Data.Play, Res.Data.Title, *m)
			m.React("✅")
		},
	})
}
