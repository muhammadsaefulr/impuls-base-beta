package cmd

import (
	"fmt"
	x "mywabot/system"
	"net/url"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "gopay",
		Cmd:     []string{"gpt"},
		Tags:    "tools",
		IsQuery: true,
		Prefix:  true,
		ValueQ:  ".gopay <no-hp>",
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

			var res struct {
				Result string `json:"data_ewallet`
			}
			err := x.GetResult("https://skizo.tech/api/cek-ewallet?ewallet_code=gopay&apikey=zpfnzf10zkrpvb&number="+url.QueryEscape(m.Query), &res)
			if err != nil {
				m.Reply(fmt.Sprint(err))
				return
			}
			m.Reply(res.Result)
			m.React("✅")
		},
	})
}
