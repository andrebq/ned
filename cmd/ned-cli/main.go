package main

import (
	"fmt"
	"log"

	"github.com/andrebq/ned/api"
	"github.com/gdamore/tcell"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

type (
	point struct {
		line int
		col  int
	}

	textView struct {
		screen    tcell.Screen
		firstCell point
		firstLine int
		lastLine  int
	}
)

func newTextView(screen tcell.Screen, topLeft point) *textView {
	_, h := screen.Size()
	return &textView{
		screen:    screen,
		firstCell: topLeft,
		firstLine: 0,
		lastLine:  h - topLeft.line,
	}
}

func (t *textView) writeLine(idx int, line string) {
	if !t.lineVisible(idx) {
		return
	}
	point := t.firstCell
	point.line += idx - t.firstLine
	for _, r := range line {
		t.screen.SetContent(point.col, point.line, r, nil, tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack))
		point.col++
	}
}

func (t *textView) flush() {
	t.screen.Show()
}

func (t *textView) lineVisible(l int) bool {
	return l-t.firstLine <= t.lastLine
}

func (t *textView) resize() {
	_, h := t.screen.Size()
	t.lastLine = t.firstLine + h
	t.flush()
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		log.Print("Unable to initialize screen", err)
		screen.Fini()
		return
	}
	defer screen.Fini()

	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack))
	screen.Clear()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tv := newTextView(screen, point{0, 0})

	grpcCli, err := connectToBackend()
	if err != nil {
		log.Print("unable to connect to backend", err)
		return
	}

	buffers := api.NewBuffersClient(grpcCli)

	go func(screen tcell.Screen) {
		for ev := screen.PollEvent(); ev != nil; ev = screen.PollEvent() {
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyCtrlQ:
					cancel()
				}
			case *tcell.EventResize:
				tv.resize()
			}
			tv.writeLine(2, fmt.Sprintf("Event: %#v", ev))
			tv.writeLine(3, fmt.Sprintf("first/last lines %v,%v", tv.firstLine, tv.lastLine))
			tv.flush()
		}
	}(screen)

	go func() {
		entries, err := buffers.WatchLines(context.Background(), &api.BufferIdentity{})
		if err != nil {
			cancel()
			return
		}
		for {
			line, err := entries.Recv()
			if err != nil {
				cancel()
				return
			}
			tv.writeLine(1, line.Contents)
			tv.flush()
		}
	}()

	tv.writeLine(0, "hello world")
	tv.flush()

	<-ctx.Done()
}

func connectToBackend() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:18080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
