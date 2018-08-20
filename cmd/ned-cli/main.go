package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rs/xid"

	"github.com/andrebq/ned/api"
	"github.com/gdamore/tcell"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

var (
	logFile = flag.String("log-out", "-", "File to write log output. stderr by default")
)

func newTextView(screen tcell.Screen, topLeft point) *textView {
	_, h := screen.Size()
	return &textView{
		screen:    screen,
		firstCell: topLeft,
		firstLine: 0,
		lastLine:  (h - 1) - topLeft.line,
	}
}

func (t *textView) writeLine(idx int, line string) {
	if !t.lineVisible(idx) {
		return
	}
	point := t.firstCell
	point.line += idx - t.firstLine
	t.writeAtPoint(point, line, tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack))
}

func (t *textView) writeStatus(text string) {
	point := t.statusBarFirstCell()
	t.writeAtPoint(point, text, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkBlue))
}

func (t *textView) writeAtPoint(point point, text string, style tcell.Style) point {
	log.WithField("point", point.String()).
		WithField("text", text).
		Info()
	for _, r := range text {
		t.screen.SetContent(point.col, point.line, r, nil, style)
		point.col++
	}
	return point
}

func (t *textView) flush() {
	t.screen.Show()
}

func (t *textView) lineVisible(l int) bool {
	return l-t.firstLine <= t.lastLine
}

func (t *textView) visibleTextLines() int {
	return (t.lastLine - t.firstLine)
}

func (t *textView) statusBarFirstCell() point {
	h := t.visibleTextLines()
	point := t.firstCell
	point.line += h
	return point
}

func (t *textView) resize() {
	_, h := t.screen.Size()
	t.lastLine = t.firstLine + (h - 1)
	t.flush()
}

func main() {
	flag.Parse()
	outputFile := openLogOutput()
	defer outputFile.Close()
	log.SetOutput(outputFile)

	log.SetFormatter(&log.JSONFormatter{})
	grpcCli, sessionClient := connectToBackend()
	_ = sessionClient // do something with the session object later

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	screen, textView := setupScreenAndTextView()
	defer screen.Fini()

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
				textView.resize()
			}
			textView.writeStatus(fmt.Sprintf("Event: %#v", ev))
			textView.flush()
		}
	}(screen)

	go func() {
		entries, err := buffers.WatchLines(context.Background(), &api.BufferIdentity{
			Path: "/buffers/main",
		})
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
			textView.writeLine(int(line.Number), line.Contents)
			textView.flush()
		}
	}()

	<-ctx.Done()
}

func openLogOutput() *os.File {
	if *logFile == "-" {
		return os.Stderr
	}
	file, err := os.OpenFile(*logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.WithError(err).WithField("log-out", *logFile).Fatal("Unable to open log file")
	}
	return file
}

func connectToBackend() (*grpc.ClientConn, api.SessionClient) {
	conn, err := grpc.Dial("localhost:19080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	sessionClient := api.NewSessionClient(conn)
	pong, err := sessionClient.Ping(context.Background(), &api.PingMessage{
		Nonce:    xid.New().String(),
		UnixNano: time.Now().UnixNano(),
	})

	if err != nil {
		log.Fatal(err)
	}

	if time.Duration(pong.PongUnixNano-pong.PingUnixNano) > time.Second {
		log.Fatal("more then one second to do a ping/pong")
	}
	return conn, sessionClient
}

func setupScreenAndTextView() (tcell.Screen, *textView) {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		screen.Fini()
		log.Fatal("Unable to initialize screen", err)
	}

	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack))
	screen.Clear()

	tv := newTextView(screen, point{0, 0})

	return screen, tv
}
