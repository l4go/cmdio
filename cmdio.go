package cmdio

import (
	"io"
	"os"
	"os/exec"
	"sync/atomic"
	"syscall"

	"github.com/l4go/task"
)

type CmdPipe struct {
	cmd     *exec.Cmd
	pr      io.ReadCloser
	pw      io.WriteCloser
	cc      task.Canceller
	do_wait uint32
	err     error
}

func (self *CmdPipe) Read(p []byte) (int, error) {
	return self.pr.Read(p)
}

func (self *CmdPipe) Write(p []byte) (int, error) {
	return self.pw.Write(p)
}

func Exec(cc task.Canceller, cmd string, arg ...string) (*CmdPipe, error) {
	self := &CmdPipe{
		cmd: exec.CommandContext(cc.AsContext(), cmd, arg...),
		cc:  task.NewCancel(),
	}
	atomic.StoreUint32(&self.do_wait, 0)

	var err error
	var pr io.ReadCloser
	var pw io.WriteCloser
	pr, err = self.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	pw, err = self.cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	self.cmd.Stderr = os.Stderr
	self.pr = pr
	self.pw = pw

	err = self.cmd.Start()
	if err != nil {
		return nil, err
	}

	return self, nil
}

func (self *CmdPipe) RecvWait() <-chan struct{} {
	return self.cc.RecvCancel()
}

func (self *CmdPipe) Wait() error {
	<-self.cc.RecvCancel()
	return self.err
}

func (self *CmdPipe) Process() *os.Process {
	return self.cmd.Process
}

func (self *CmdPipe) Signal(sig os.Signal) error {
	return self.cmd.Process.Signal(sig)
}

func (self *CmdPipe) ReaderClose() error {
	err := self.pr.Close()
	once_wait := atomic.SwapUint32(&self.do_wait, 1)
	if once_wait == 0 {
		go func() {
			self.err = self.cmd.Wait()
			self.cc.Cancel()
		}()
	}
	return err
}

func (self *CmdPipe) WriterClose() error {
	return self.pw.Close()
}

func (self *CmdPipe) Close() error {
	err := self.WriterClose()
	self.ReaderClose()
	self.cmd.Process.Kill()

	return err
}

type StdPipe struct {
	pr io.ReadCloser
	pw io.WriteCloser
}

func StdDup() (*StdPipe, error) {
	in, err := syscall.Dup(syscall.Stdin)
	if err != nil {
		return nil, err
	}
	out, err := syscall.Dup(syscall.Stdout)
	if err != nil {
		return nil, err
	}

	return &StdPipe{
		pr: os.NewFile(uintptr(in), "dup stdin"),
		pw: os.NewFile(uintptr(out), "dup stdout"),
	}, nil
}

func (self *StdPipe) Read(p []byte) (int, error) {
	return self.pr.Read(p)
}

func (self *StdPipe) Write(p []byte) (int, error) {
	return self.pw.Write(p)
}

func (self *StdPipe) Close() error {
	self.pw.Close()
	self.pr.Close()
	return nil
}

func (self *StdPipe) ReaderClose() error {
	return self.pr.Close()
}

func (self *StdPipe) WriterClose() error {
	return self.pw.Close()
}
