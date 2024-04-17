package cmd

import (
	"fmt"
	x "mywabot/system"
	"net/url"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "gpt4",
		Cmd:     []string{"gpt4"},
		Tags:    "ai",
		IsQuery: true,
		Prefix:  true,
		ValueQ:  ".gpt4 siapa kamu?",
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")
			var res struct {
				Result string `json:"result"`
			}
			err := x.GetResult("https://aemt.me/gpt4?text="+url.QueryEscape(m.Query), &res)
			if err != nil {
				m.Reply(fmt.Sprint(err))
				return
			}
			m.Reply(res.Result)
			m.React("✅")
		},
	})
}
