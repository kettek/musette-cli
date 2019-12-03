package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type PlayerList struct {
	table  *tview.Table
	tracks []Track
}

func (p *PlayerList) Setup() {
	p.table = tview.NewTable().SetSelectable(true, false)

	p.table.SetSelectedFunc(func(row, column int) {
		logger.Log("Change current song to %s(row %d)", p.tracks[row].trackName, row)
	})

	// temp
	for i := 0; i < 30; i++ {
		p.tracks = append(p.tracks, Track{
			trackNumber: int8(i),
			trackName:   "Name",
			trackArtist: "Artist",
			trackAlbum:  "Album",
		})
	}

	p.syncList()
}

func (p *PlayerList) syncList() {
	p.table.Clear()

	// Add tracks
	for i, t := range p.tracks {
		for c := 0; c < 5; c++ {
			color := tcell.ColorWhite
			cell := tview.NewTableCell("").
				SetTextColor(color).
				SetAlign(tview.AlignLeft)
			switch c {
			case 0:
				cell.SetText(fmt.Sprintf("%d", t.trackNumber))
			case 1:
				cell.SetText(t.trackName)
				cell.SetExpansion(100)
			case 2:
				cell.SetText(t.trackArtist)
			case 3:
				cell.SetText(t.trackAlbum)
			case 4:
				if t.selected {
					cell.SetText("x")
				} else {
					cell.SetText(" ")
				}
			}
			p.table.SetCell(i, c, cell)
		}
	}
}
