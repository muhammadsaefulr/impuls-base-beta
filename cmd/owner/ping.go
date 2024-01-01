package cmd

import (
	"fmt"
	x "mywabot/system"
	"time"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "ping",
		Cmd:     []string{"ping"},
		Tags:    "owner",
		Desc:    "Test Ping To Host",
		Prefix:  true,
		IsOwner: true,
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

			now := time.Now()
			AtDate := time.Unix(m.Timestamp.Unix(), 0)
			TimeDuration := now.Sub(AtDate)
			setToSeconds := TimeDuration.Seconds()
			res := fmt.Sprintf("%.3f seconds", setToSeconds)
			m.Reply(res)
		},
	})
}
