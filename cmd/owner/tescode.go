package cmd

import (
	"fmt"
	x "mywabot/system"
	"os"
	"os/exec"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name:    "tes",
		Cmd:     []string{"tes"},
		Tags:    "owner",
		IsOwner: true,
		Prefix:  true,
		Exec: func(sock *x.Nc, m *x.IMsg) {

			sock.SendImage(m.From, "https://i.pinimg.com/1200x/05/85/a9/0585a9547d40b1908e30ed3d352d9116.jpg", "Halo ini caption", *m)
			sock.SendDocument(m.From, "https://i.pinimg.com/1200x/05/85/a9/0585a9547d40b1908e30ed3d352d9116.jpg", "Halo", "captionnya", *m)
			sock.SendAudio(m.From, "https://sf16-ies-music-va.tiktokcdn.com/obj/musically-maliva-obj/6841176534603041542.mp3", false, *m)
			if m.IsQuotedSticker {
				m.React("💤")
				conjp := "./tmp/" + m.ID + ".webp"
				conwp := "./tmp/" + m.ID + ".webp"
				byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.StickerMessage)
				err := os.WriteFile(conjp, byte, 0644)
				if err != nil {
					fmt.Println("Failed saved webp")
					return
				}
				x.CreateExif("mywabot.exif", "Impuls • Beta Test Bot", "")

				createExif := fmt.Sprintf("webpmux -set exif %s %s -o %s", "tmp/exif/mywabot.exif", conwp, conwp)
				cmd := exec.Command("bash", "-c", createExif)
				err = cmd.Run()
				if err != nil {
					fmt.Println("Failed to set webp metadata", err)
				}
				sock.StickerPath(m.From, conwp, *m)
				os.Remove(conwp)
				os.Remove(conjp)
			}

			if m.IsQuotedImage {
				conjp := "./tmp/" + m.ID + ".jpg"
				conwp := "./tmp/" + m.ID + ".webp"
				byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.ImageMessage)
				err := os.WriteFile(conjp, byte, 0644)
				if err != nil {
					fmt.Println("Failed saved image")
					return
				}
				err = x.ImgToWebp(conjp, conwp)

				if err != nil {
					fmt.Println("Failed to convert image to webp!!")
				}
				sock.StickerPath(m.From, conwp, *m)
			}
		},
	})
}
