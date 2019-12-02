package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	browser := createBrowser()

	grid := tview.NewGrid().
		SetRows(0).
		SetColumns(10, 0).
		SetBorders(true)

	grid.AddItem(browser, 0, 0, 1, 1, 0, 0, true)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func createBrowser() tview.Primitive {
	rootDir := "."
	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	return tree
}
