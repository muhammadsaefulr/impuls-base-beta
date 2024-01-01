package config

import "sync"

var (
	mu    sync.Mutex
	Name  = "waSocket Bot"
	Login = "code"
	Bot   = "6287815556349"
	Owner = []string{"6288219406742","6287815556349"}
	Self  = false
	LolSite = "https://api.lolhuman.xyz/"
	LolKey = "5f38494f3555283d0446abdf"
)

func SetName(newName string) {
	mu.Lock()
	defer mu.Unlock()
	Name = newName
}

func SetSelf(new bool) {
	mu.Lock()
	defer mu.Unlock()
	Self = new
}
