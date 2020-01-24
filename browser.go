package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type BrowserItem struct {
	Path     string
	Items    []BrowserItem
	Mimetype string
}

type Browser struct {
	view  *tview.Flex
	root  BrowserItem
	items map[string]BrowserItem
}

func (b *Browser) Setup() {
	b.view = tview.NewFlex().
		SetDirection(tview.FlexRow)
	location := tview.NewTextView().
		SetText("location/a/place/you/know")

	b.view.AddItem(location, 0, 2, false)

	rootDir := "."
	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	boot := tview.NewTreeNode("..").SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	root.AddChild(boot)

	b.view.AddItem(tree, 0, 16, true)
}

func (b *Browser) Open(loc string) error {
	var browserItems []BrowserItem

	err := requestPath(loc, &browserItems)
	if err != nil {
		logger.Log("requestPath: %+v", err)
		return err
	}
	for i, v := range browserItems {
		logger.Log("%d:%s:%s", i, v.Path, v.Mimetype)
	}

	return nil
}

func (b *Browser) GetLoc(loc string) BrowserItem {
	parts := strings.Split(loc, "/")
	for i, p := range parts {
	}
}
