package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"go-tetris/block"
	"go-tetris/board"
	"go-tetris/game"
	"go-tetris/utils"
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
			if game.IsEnd() {
				game.StartGame = false
				game.KeyLock = true
			}
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
			if !game.StartGame {
				utils.PrintString(16, 15, "Start :  Ctrl + s")
				utils.PrintString(16, 17, "End : Ctrl + x")
			}

			if game.StartGame {
				buf := fmt.Sprintf("Level : %v", game.Level)
				utils.PrintString(60, 15, buf)
				buf = fmt.Sprintf("Score : %v", game.Score)
				utils.PrintString(60, 17, buf)
			}
			termbox.Flush()
			if game.StartGame && game.KeyLock {
				game.KeyLock = false
			}

			time.Sleep(utils.FlushInternal)

		}
	}
}
