package cmd

import (
	"fmt"
	x "mywabot/system"
	"net/url"
)

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

			type Res struct {
				WM        string `json:"wm"`
				Thumbnail string `json:"thumbnail"`
				URL       string `json:"url"`
			}

			type Response struct {
				Status  bool   `json:"status"`
				Code    int    `json:"code"`
				Creator string `json:"creator"`
				Result  []Res  `json:"result"`
			}

			var response Response
			err := x.GetResult("https://aemt.me/download/igdl?url="+url.QueryEscape(m.Query), &response)
			if err != nil {
				m.Reply(fmt.Sprint(err))
				return
			}
			if len(response.Result) > 0 {
				sock.SendVideo(m.From, response.Result[0].URL, response.Result[0].WM, *m) // Access the first result
			}
			m.React("✅")
		},
	})
}
