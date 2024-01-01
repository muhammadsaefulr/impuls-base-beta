package cmd

import (
	"fmt"
	x "mywabot/system"
	"net/url"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "gpt",
		Cmd:     []string{"gpt"},
		Tags:    "ai",
		IsQuery: true,
		Prefix:  true,
		ValueQ:  ".gpt siapa kamu?",
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")
			var res struct {
				Result string `json:"result"`
			}
			err := x.GetResult("https://skizo.tech/api/openai?apikey=zpfnzf10zkrpvb&text="+url.QueryEscape(m.Query), &res)
			if err != nil {
				m.Reply(fmt.Sprint(err))
				return
			}
			m.Reply(res.Result)
			m.React("✅")
		},
	})
}
