package cmd

import (
	"fmt"
	x "mywabot/system"
	"sort"
	"strings"

	waProto "github.com/amiruldev20/waSocket/binary/proto"
	"google.golang.org/protobuf/proto"
)

type item struct {
	Name   []string
	Prefix bool
}

type tagSlice []string

func (t tagSlice) Len() int {
	return len(t)
}

func (t tagSlice) Less(i int, j int) bool {
	return t[i] < t[j]
}

func (t tagSlice) Swap(i int, j int) {
	t[i], t[j] = t[j], t[i]
}

func menu(client *x.Nc, m *x.IMsg) {
	var str string
	str += fmt.Sprintf("Halo, %s 👋\nBot ini masih dalam tahap beta\n\n»» Host: Localhost\n»» Language: GoLang\n»» Library: waSocket\n\n", m.PushName)
	var tags map[string][]item
	for _, list := range x.GetList() {
		if tags == nil {
			tags = make(map[string][]item)
		}
		if _, ok := tags[list.Tags]; !ok {
			tags[list.Tags] = []item{}
		}
		tags[list.Tags] = append(tags[list.Tags], item{Name: list.Cmd, Prefix: list.Prefix})
	}

	var keys tagSlice
	for key := range tags {
		keys = append(keys, key)
	}

	sort.Sort(keys)

	for _, key := range keys {
		str += fmt.Sprintf("*%s MENU*\n", strings.ToUpper(key))
		for _, e := range tags[key] {
			var prefix string
			if e.Prefix {
				prefix = m.Prefix[:1]
			} else {
				prefix = ""
			}
			for _, nm := range e.Name {
				str += fmt.Sprintf("▹ %s%s\n", prefix, nm)
			}
		}
		str += "\n"
	}
	txt := str + "\n© Developed by Msaepul"
	var isMessageProtobuf = waProto.ContextInfo_ExternalAdReplyInfo_IMAGE
	client.SendText(m.From, strings.TrimSpace(txt), &waProto.ContextInfo{
		ExternalAdReply: &waProto.ContextInfo_ExternalAdReplyInfo{
			Title:                 proto.String("Impuls Bot || Beta Test"),
			Body:                  proto.String("Simple Whatsapp Bot"),
			MediaType:             &isMessageProtobuf,
			ThumbnailUrl:          proto.String("https://i.pinimg.com/1200x/05/85/a9/0585a9547d40b1908e30ed3d352d9116.jpg"),
			MediaUrl:              proto.String("https://wa.me/stickerpack/inc.dev"),
			SourceUrl:             proto.String("https://www.pixiv.net/en/tags/柊シノア"),
			ShowAdAttribution:     proto.Bool(false),
			RenderLargerThumbnail: proto.Bool(true),
		}})
	// m.Reply(txt)
}

func init() {
	x.NewCmd(&x.ICmd{
		Name:   "menu",
		Cmd:    []string{"menu"},
		Tags:   "main",
		Prefix: true,
		Exec:   menu,
	})
}
