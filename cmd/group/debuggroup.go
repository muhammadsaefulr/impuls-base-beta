package cmd

import (
	"fmt"
	x "mywabot/system"
	"strings"

	waProto "github.com/amiruldev20/waSocket/binary/proto"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:   "debuggroup",
		Cmd:    []string{"debuggroup"},
		Tags:   "group",
		Prefix: true,
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")
			// sock.FetchGroupAdmin()

			var participantNames []string

			resp, err := sock.WA.GetGroupInfo(m.From)

			if err != nil {
				fmt.Sprint(err)
			}

			for _, participant := range resp.Participants {
				participantNames = append(participantNames, participant.JID.User)
			}

			participantStr := fmt.Sprintf("Participants:\n%s", strings.Join(participantNames, "\n"))

			sock.SendText(m.From, participantStr, &waProto.ContextInfo{
				MentionedJid: participantNames,
			})

		},
	})
}
