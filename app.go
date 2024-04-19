package main

import (
	"context"
	"fmt"
	"ryzeaccept/lcu"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var AcceptOn bool

// App struct
type App struct {
	ctx       context.Context
	connector *lcu.Connector
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Setup() bool {
	conn, err := lcu.RiotConnector()
	if err != nil {
		return false
	}
	a.connector = conn

	go func() {
		a.connector.WsConnect()
	}()

	return true
}

func (a *App) ToggleAccept() {
	fmt.Println(lcu.ACCEPT)
	lcu.ACCEPT = !lcu.ACCEPT
}

func (a *App) Summoner() lcu.SunnmonerInfo {
	summonerData := a.connector.GetSummonerInfo()

	return summonerData
}

func (a *App) Reload() {
	runtime.WindowReload(a.ctx)
}

// Greet returns a greeting for the given name
