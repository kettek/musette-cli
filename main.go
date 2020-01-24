package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	app                  *tview.Application
	httpClient           *http.Client
	loginView            tview.Primitive
	playerView           tview.Primitive
	playerControllerView tview.Primitive
	playerList           PlayerList
	browser              Browser
	connect              Connect
	logger               Logger
	config               Config
	pages                *tview.Pages
)

func main() {
	app = tview.NewApplication()

	// Set up http client
	cookieJar, _ := cookiejar.New(nil)
	httpClient = &http.Client{
		Jar: cookieJar,
	}

	// Set up UI
	logger.Setup()
	logger.Log("Musette CLI v0 started")

	browser.Setup()

	connect.Setup()

	config.LoadFromFile("musette.yaml")

	playerList.Setup()

	createLogin()
	createPlayer()

	fullView := tview.NewGrid().
		SetRows(0, 3).
		SetColumns(20, 0).
		SetBorders(true)

	fullView.AddItem(browser.view, 0, 0, 1, 1, 0, 0, true)
	fullView.AddItem(playerView, 0, 1, 1, 1, 0, 0, true)
	fullView.AddItem(logger.view, 1, 0, 1, 2, 0, 0, true)

	pages = tview.NewPages().
		AddPage("playerView", playerView, true, true).
		AddPage("browserView", browser.view, true, true).
		AddPage("fullView", fullView, true, true).
		AddPage("loginView", loginView, true, true).
		AddPage("loggerView", logger.view, true, true).
		AddPage("connectView", connect.view, true, true).
		SwitchToPage("fullView")

	if config.Server != "" {
		connectTo(config.Server)
	} else {
		pages.SwitchToPage("connectView")
		app.SetFocus(connect.view)
	}

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

func createLogin() {
	var container *tview.Grid
	var userField, passwordField *tview.InputField
	var loginButton *tview.Button

	container = tview.NewGrid().
		SetColumns(0).
		SetRows(0, 0, 0)

	container.
		SetBorder(true).
		SetTitle("Login Required")

	userField = tview.NewInputField().
		SetLabel("Username: ").
		SetFieldWidth(16).
		SetDoneFunc(func(key tcell.Key) {
			app.SetFocus(passwordField)
		})
	passwordField = tview.NewInputField().
		SetLabel("Password: ").
		SetFieldWidth(16).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyUp {
				app.SetFocus(userField)
			} else {
				app.SetFocus(loginButton)
			}
		})

	loginButton = tview.NewButton("Login").
		SetSelectedFunc(func() {
			resp := struct {
				Status  int16
				Message string
			}{}

			err := postAPI(fmt.Sprintf("%s/api/auth/login", config.Server), url.Values{
				"username": {userField.GetText()},
				"password": {passwordField.GetText()},
			}, &resp)

			if err != nil {
				logger.Log("argh: %v+", err)
				//pages.SwitchToPage("connectView")
				pages.SwitchToPage("loggerView")
				return
			}
			// TODO: Check if 401 && INVALID
			if resp.Status == 401 {
				logger.Log("still invalid")
			} else if resp.Status == 200 {
				logger.Log("Successfully connected.")
				browser.Open("/")
				pages.SwitchToPage("fullView")
			} else {
				logger.Log("what: %+v", resp)
				pages.SwitchToPage("loggerView")
			}

		})
	loginButton.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		if e.Key() == tcell.KeyUp {
			app.SetFocus(passwordField)
			return nil
		}
		return e
	})

	container.AddItem(userField, 0, 0, 1, 1, 0, 0, true)
	container.AddItem(passwordField, 1, 0, 1, 1, 0, 0, true)
	container.AddItem(loginButton, 2, 0, 1, 1, 0, 0, true)

	loginView = container
}

func createPlayer() {
	grid := tview.NewGrid().
		SetRows(3, 0).
		SetColumns(0)

	createPlayerController()

	grid.AddItem(playerControllerView, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(playerList.table, 1, 0, 1, 1, 0, 0, true)

	playerView = grid
}

func createPlayerController() {
	text := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("controller")

	playerControllerView = text
}

/*func createPlayerList() {
	table := tview.NewTable().
		SetSelectable(true, false)

	for r := 0; r < 99; r++ {
		for c := 0; c < 5; c++ {
			color := tcell.ColorWhite

			cell := tview.NewTableCell("").
				SetTextColor(color).
				SetAlign(tview.AlignLeft)
			switch c {
			case 0:
				cell.SetText(fmt.Sprintf("%d", r))
			case 1:
				cell.SetText("track name")
				cell.SetExpansion(100)
			case 2:
				cell.SetText("artist")
			case 3:
				cell.SetText("album")
			case 4:
				cell.SetText("+")
			}

			table.SetCell(r, c, cell)
		}
	}

	playerListView = table
}*/
