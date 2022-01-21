package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/l4go/cmdio"
	"github.com/l4go/lineio"
	"github.com/l4go/task"
	"github.com/l4go/timer"
)

func main() {
	log.Println("START")
	defer log.Println("END")

	m := task.NewMission()

	signal_ch := make(chan os.Signal, 1)
	signal.Notify(signal_ch, syscall.SIGINT, syscall.SIGTERM)

	cat_io, err := cmdio.Exec(m, "/bin/cat", "-n")
	if err != nil {
		defer log.Println("Error:", err)
		return
	}
	go read_worker(m.New(), cat_io)
	go func(cm *task.Mission) {
		defer cat_io.Close()
		write_worker(cm, cat_io)
	}(m.New())

	select {
	case <-m.Recv():
	case <-signal_ch:
		m.Cancel()
	}
}

func read_worker(m *task.Mission, r io.Reader) {
	defer m.Done()
	log.Println("start: reader worker")
	defer log.Println("end: reader worker")

	line_r := lineio.NewReader(r)
	for {
		var ln []byte
		var ok bool
		select {
		case ln, ok = <-line_r.Recv():
		case <-m.RecvCancel():
			return
		}
		if !ok {
			break
		}

		fmt.Println(">", string(ln))
	}
	if err := line_r.Err(); err != nil {
		log.Println("Error:", err)
		return
	}
}

const MAX_MSG = 5
const MSG_TICK = 500 * time.Millisecond

func write_worker(m *task.Mission, w io.Writer) {
	defer m.Done()
	log.Println("start: write worker")
	defer log.Println("end: write worker")
	wt := timer.NewTimer()

	defer wt.Stop()
	for i := 0; i < MAX_MSG; i++ {
		wt.Start(MSG_TICK)
		select {
		case <-wt.Recv():
		case <-m.RecvCancel():
			return
		}

		if _, err := io.WriteString(w, "test message\n"); err != nil {
			log.Println("Error:", err)
			return
		}
	}
}
