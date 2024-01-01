package cmd

import (
	"fmt"
	x "mywabot/system"
	"os/exec"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "testing",
		Cmd:     []string{"testing"},
		Tags:    "owner",
		Desc:    "Testing Reply And UnitTest",
		Prefix:  false,
		IsOwner: true,
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

			conwp := "./tmp/" + m.ID + ".webp"
			x.CreateExif("mywabot.exif", "🤖 TEST WA BOT 2023 🤖\n\nLibrary: WASOCKET\n\n Language: GoLang \n\n", "")

			createExif := fmt.Sprintf("webpmux -set exif %s %s -o %s", "tmp/exif/mywabot.exif", conwp, conwp)
			cmd := exec.Command("bash", "-c", createExif)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Failed to set webp metadata", err)
			}

			m.Reply("Ini Reply data")
		},
	})
}
