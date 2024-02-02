package cmd

import (
	"fmt"
	x "mywabot/system"
	"os/exec"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "exif",
		Cmd:     []string{"exif"},
		Tags:    "owner",
		Prefix:  true,
		IsOwner: true,
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

			conwp := "./tmp/" + m.ID + ".webp"
			x.CreateExif("mywabot.exif", "Impuls - Simple Whatsapp Bot", "")

			createExif := fmt.Sprintf("webpmux -set exif %s %s -o %s", "tmp/exif/mywabot.exif", conwp, conwp)
			cmd := exec.Command("bash", "-c", createExif)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Failed to set webp metadata", err)
			}

			m.Reply("Berhasil Mengubah Exif")
		},
	})
}
