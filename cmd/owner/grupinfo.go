package cmd

import (
	"fmt"
	x "mywabot/system"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:   "grupinfo",
		Cmd:    []string{"grupinfo"},
		Tags:   "owner",
		Prefix: true,
		Exec: func(sock *x.Nc, m *x.IMsg) {

			resp, err := sock.WA.GetGroupInfo(m.From)
			if err != nil {
				fmt.Println(err)
			}

			m.Reply("Nama grup : " + resp.Name)
			// sock.SendText(m.From, resp.Name, opts. )
		},
	})
}
