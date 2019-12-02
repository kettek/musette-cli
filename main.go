package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	app                  *tview.Application
	loginView            tview.Primitive
	browserView          tview.Primitive
	playerView           tview.Primitive
	playerControllerView tview.Primitive
	playerListView       tview.Primitive
)

func main() {
	app = tview.NewApplication()

	createLogin()
	createBrowser()
	createPlayer()

	fullView := tview.NewGrid().
		SetRows(0).
		SetColumns(20, 0).
		SetBorders(true)

	fullView.AddItem(browserView, 0, 0, 1, 1, 0, 0, true)
	fullView.AddItem(playerView, 0, 1, 1, 1, 0, 0, true)

	pages := tview.NewPages().
		AddPage("playerView", playerView, true, true).
		AddPage("browserView", browserView, true, true).
		AddPage("fullView", fullView, true, true).
		AddPage("loginView", loginView, true, true).
		SwitchToPage("fullView")

	if err := app.SetRoot(pages, true).SetFocus(fullView).Run(); err != nil {
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
			app.Stop()
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

func createBrowser() {
	rootDir := "."
	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	boot := tview.NewTreeNode("..").SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	root.AddChild(boot)

	browserView = tree
}

func createPlayer() {
	grid := tview.NewGrid().
		SetRows(3, 0).
		SetColumns(0)

	createPlayerController()
	createPlayerList()

	grid.AddItem(playerControllerView, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(playerListView, 1, 0, 1, 1, 0, 0, true)

	playerView = grid
}

func createPlayerController() {
	text := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("controller")

	playerControllerView = text
}

func createPlayerList() {
	table := tview.NewTable()

	for r := 0; r < 18; r++ {
		for c := 0; c < 5; c++ {
			color := tcell.ColorWhite

			cell := tview.NewTableCell("").
				SetTextColor(color).
				SetAlign(tview.AlignLeft)
			switch c {
			case 0:
				cell.SetText("00")
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
}
