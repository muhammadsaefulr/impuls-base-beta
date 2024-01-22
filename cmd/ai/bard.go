package cmd

import (
	x "mywabot/system"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "bard",
		Cmd:     []string{"bard"},
		Tags:    "ai",
		IsQuery: true,
		Prefix:  true,
		ValueQ:  ".bard siapa kamu?",
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")
			// var res struct {
			// 	Result string `json:"result"`
			// }
			// err := x.GetResult("https://skizo.tech/api/bard-ai?apikey=zpfnzf10zkrpvb&text="+url.QueryEscape(m.Query), &res)
			// if err != nil {
			// 	m.Reply(fmt.Sprint(err))
			// 	return
			// }
			m.Reply("Maaf, Sedang Terjadi Masalah !")
			m.React("✅")
		},
	})
}
