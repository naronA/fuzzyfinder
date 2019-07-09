package gui

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func Gui() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	// 終了コマンド
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	// メインループのエラー監視
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	// score.NeedlemanWunsch("ATACABACCCC", "ACBCCCC", true)
	// score.NeedlemanWunsch("AC", "ATAC", true)
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	// 表示される文字数 / 2
	if v, err := g.SetView("hello1", maxX/2-8, maxY*2/3, maxX/2+8, maxY*2/3+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!!!!!")
	}

	if v, err := g.SetView("hello2", maxX/2-8, maxY/3, maxX/2+8, maxY/3+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!!!!!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
