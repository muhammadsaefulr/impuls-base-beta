package cmd

import (
	"fmt"
	x "mywabot/system"
	"os"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "(up|uploader)",
		Cmd:     []string{"uploader"},
		Tags:    "tools",
		IsMedia: true,
		Prefix:  true,
		Exec: func(sock *x.Nc, m *x.IMsg) {
			if !m.IsQuotedImage && !m.IsImage {
				m.Reply("Cannot Request This Metadata")
				return
			}
			m.React("⏱️")

			// quoted image
			if m.IsQuotedImage {
				conjp := "./tmp/" + m.ID + ".jpg"
				conwp := "./tmp/" + m.ID + ".webp"
				byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.ImageMessage)
				err := os.WriteFile(conjp, byte, 0644)
				if err != nil {
					fmt.Println("Failed saved jpg")
					return
				}
				x.ImgToWebp(conjp, conwp)
				sock.StickerPath(m.From, conwp, *m)
				os.Remove(conwp)
				os.Remove(conjp)
				m.React("✅")
			}

			// from image
			if m.IsImage {
				conjp := "./tmp/" + m.ID + ".jpg"
				conwp := "./tmp/" + m.ID + ".webp"
				byte, _ := sock.WA.Download(m.Media)
				err := os.WriteFile(conjp, byte, 0644)
				if err != nil {
					fmt.Println("Failed saved jpg")
					return
				}
				x.ImgToWebp(conjp, conwp)
				sock.StickerPath(m.From, conwp, *m)
				os.Remove(conwp)
				os.Remove(conjp)
				m.React("✅")
			}

		},
	})
}
