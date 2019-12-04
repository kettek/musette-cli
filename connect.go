package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Connect struct {
	view *tview.Grid
}

func connectTo(url string) {
	logger.Log("Connecting to %s", url)
	// Make an API request. If we receive a 401, show loginView, then:
	// request /api/auth/login, send POST with body {username:, password:}
	// on success, switch to fullView, otherwise back to login
	// if no 401, switch to fullView

	auth := struct {
		Status  int16
		Message string
	}{}

	err := requestAPI(fmt.Sprintf("%s/api/auth", url), &auth)
	if err != nil {
		logger.Log("%v+", err)
		//pages.SwitchToPage("connectView")
		pages.SwitchToPage("loggerView")
		return
	}

	if auth.Status == 401 {
		pages.SwitchToPage("loginView")
	} else {
		logger.Log("Unhandled %d status: %s", auth.Status, auth.Message)
		pages.SwitchToPage("loggerView")
		return
	}
}

func (c *Connect) Setup() {
	var serverField *tview.InputField
	var connectButton *tview.Button

	c.view = tview.NewGrid().
		SetColumns(0).
		SetRows(0, 0)

	c.view.
		SetBorder(true).
		SetTitle("Enter Server URL")

	serverField = tview.NewInputField().
		SetLabel("Server: ").
		SetFieldWidth(24).
		SetDoneFunc(func(key tcell.Key) {
			app.SetFocus(connectButton)
		})

	connectButton = tview.NewButton("Connect").
		SetSelectedFunc(func() {
			config.Server = serverField.GetText()
			connectTo(config.Server)
		})
	connectButton.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		if e.Key() == tcell.KeyUp {
			app.SetFocus(serverField)
			return nil
		}
		return e
	})

	c.view.AddItem(serverField, 0, 0, 1, 1, 0, 0, true)
	c.view.AddItem(connectButton, 1, 0, 1, 1, 0, 0, true)
}
