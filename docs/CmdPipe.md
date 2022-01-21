# type CmdPipe
外部コマンドを実行し、その標準入出力に対応するio.ReadWriteCloser互換のI/Oを提供します。

## import
```go
import "github.com/l4go/cmdio"
```
vendoringして使うことを推奨します。

## 利用サンプル

[example](../examples/ex_cmdio/ex_cmdio.go)

## メソッド概略

### func Exec(cc task.Canceller, cmd string, arg ...string) (\*CmdPipe, error)
指定した外部コマンドを実行し、成功時には*CmdPipeを生成します。
失敗時には、nil以外のerrorの値を返します。

### func (self \*CmdPipe) Read(p []byte) (int, error)
io.Readerと同等の機能を提供します。

### func (self \*CmdPipe) Write(p []byte) (int, error)
io.Writerと同等の機能を提供します。

### func (self \*CmdPipe) Close() error
io.Closerと同等の機能を提供します。

### func (self *CmdPipe) RecvWait() <-chan struct{}
実行したプロセスが終了したときにcloseするchanを返します。
select文で非同期にプロセスの終了処理を行うために使います。

### func (self *CmdPipe) Wait() error
実行したプロセスの終了を待ちます。

### func (self \*CmdPipe) Process() *os.Process
実行したプロセスの\*os.Processを返します。

### func (self \*CmdPipe) Signal(sig os.Signal) error
実行したプロセスにos.Signalを送ります。

### func (self \*CmdPipe) ReaderClose() error
標準入力側のみCloseします。

### func (self \*CmdPipe) WriterClose() error
標準出力側のみCloseします。
