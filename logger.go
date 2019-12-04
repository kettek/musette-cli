package main

import (
	"fmt"

	"github.com/rivo/tview"
)

type Logger struct {
	messages []string
	view     *tview.Table
}

func (l *Logger) Setup() {
	l.view = tview.NewTable().SetSelectable(true, false)
}

func (l *Logger) Log(s string, p ...interface{}) {
	l.messages = append(l.messages, fmt.Sprintf(s, p...))

	cell := tview.NewTableCell(l.messages[len(l.messages)-1]).
		SetSelectable(false)

	l.view.SetCell(len(l.messages)-1, 0, cell)
	l.view.ScrollToEnd()
}
