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

			var participantNumber []string
			var participantServer []string

			resp, err := sock.WA.GetGroupInfo(m.From)

			if err != nil {
				fmt.Sprint(err)
			}

			for _, participant := range resp.Participants {
				participantNumber = append(participantNumber, participant.JID.User)
				participantServer = append(participantServer, participant.JID.Server)
			}

			// var userParticipant = strings.Join(participantNumber, "\n")
			var participantStr strings.Builder
			var userParticipant strings.Builder

			for i := 0; i < len(participantNumber); i++ {
				participantStr.WriteString(participantNumber[i] + "@" + participantServer[i] + "\n")
			}

			for i := 0; i < len(participantNumber); i++ {
				userParticipant.WriteString("@" + participantNumber[i] + "\n")
			}

			m.Reply(participantStr.String())

			sock.SendText(m.From, userParticipant.String(), &waProto.ContextInfo{
				MentionedJid: []string{participantStr.String()},
			})

		},
	})
}
