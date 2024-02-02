package cmd

import (
	x "mywabot/system"

	waProto "github.com/amiruldev20/waSocket/binary/proto"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "jid",
		Cmd:     []string{"jid"},
		Tags:    "owner",
		Prefix:  true,
		IsOwner: true,
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

			// jid := fmt.Sprint("User JID : ", m.From)
			// m.Reply(jid)

			// userNumber := fmt.Sprint("User Number : ", m.Sender.User)
			// m.Reply(userNumber)

			// mentioned := fmt.Sprint("@" + m.Sender.User + "@" + m.From.User + "g.us")

			// dataName := "@" + m.From.User
			// sock.WA.SendMessage(context.Background(), m.From, &waProto.Message{
			// 	ExtendedTextMessage: &waProto.ExtendedTextMessage{
			// 		Text: proto.String("Halo, " + m.PushName),
			// 	},
			// })

			sock.SendText(m.From, "Halo, "+"@"+m.Sender.User, &waProto.ContextInfo{
				MentionedJid: []string{m.Sender.User + "@" + m.Sender.Server},
			})

			// m.Reply("@" + m.Sender.User + "@" + m.Sender.Server)
			// tags := fmt.Sprint(m.From.String() + m.String())
		},
	})
}
