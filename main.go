package main

import (
	"github.com/gdamore/tcell"
	"log"
	"os"
	"time"
)

func main() {

	blackStyle := tcell.StyleDefault.Background(tcell.ColorBlack)
	aliveColor := blackStyle.Foreground(tcell.ColorWhite)
	deadColor := blackStyle.Foreground(tcell.ColorBlack)

	s := initialize()
	s.SetStyle(aliveColor)

	const (
		l = false
		O = true
	)

	g := grid{ // Raucci's p38
		{l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, O, O, l, l, l, l, l, l, l, l, l, l, l, l, O, O, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, l, O, l, l, l, l, l, l, l, l, l, l, l, O, O, O, l, l, l, l, O, O, O, l, l, l, l, l, l, l, l, l, l, l, O, l, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, O, O, O, l, l, l, l, l, l, l, l, l, l, O, l, l, O, l, l, O, l, l, O, l, l, l, l, l, l, l, l, l, l, O, O, O, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, O, O, O, l, l, l, l, l, l, l, l, l, O, l, O, l, O, l, l, O, l, O, l, O, l, l, l, l, l, l, l, l, l, O, O, O, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, l, O, O, O, O, l, O, O, l, l, O, O, l, O, O, O, O, l, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, O, O, O, l, l, l, l, l, l, l, l, l, l, O, O, O, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, l, l, l, l, l, l, l, l, l, l, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, O, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, O, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, l, l, l, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, l, l, l, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, O, O, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, O, O, l, l, l, l, l, l},
		{l, l, l, l, l, l, l, O, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, O, l, l, l, l, l, l},
	}

	g.print(s, aliveColor, deadColor)
	s.Show()
	go checkExit(s)

	for {
		err := g.update()
		time.Sleep(150 * time.Millisecond)
		if err != nil {
			log.Fatal(err)
		}
		g.print(s, aliveColor, deadColor)
		s.Show()
	}
}

func checkExit(s tcell.Screen) {
	defer func() {
		s.Fini()
		os.Exit(1)
	}()
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyRune && ev.Rune() == 'q' || ev.Rune() == 'Q' {
				return
			}
		}
	}
}
