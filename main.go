package main

import (
	termbox "github.com/nsf/termbox-go"
	"go-tetris/block"
	"go-tetris/board"
	"go-tetris/game"
	"time"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.Flush()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	game.InitGame()

	var b *block.Block

	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			if b != nil {
				board.Erase(b)
				collid := game.CheckDownCollid(b)
				if !collid {
					b.MoveToDown()
				}
				board.PutOn(b)
				if collid {
					b = nil
					board.RemoveLines()
				}
			}
		}
	}()

loop:
	for {
		if game.StartGame && b == nil {
			b = block.NewBlock()
			board.PutOn(b)
		}

		select {
		case ev := <-eventQueue:
			if !game.KeyHandler(ev, b) {
				break loop
			}
		default:
			game.DrawBoard()
			termbox.Flush()

			time.Sleep(50 * time.Millisecond)

		}
	}
}
