# type StdPipe
実行しているプロセスの標準入出力に対応するio.ReadWriteCloser互換のI/Oを提供します。

## import
```go
import "github.com/l4go/cmdio"
```
vendoringして使うことを推奨します。

## 利用サンプル

[example](../examples/ex_cmdio2/ex_cmdio2.go)

## メソッド概略

### func StdDup() (\*StdPipe, error)
実行したプロセスの標準入出力を`dup()`(システムコール)を実行して、成功時には*CmdPipeを返します。
失敗時には、nil以外のerrorの値を返します。

### func (self \*StdPipe) Read(p []byte) (int, error)
io.Readerと同等の機能を提供します。

### func (self \*StdPipe) Write(p []byte) (int, error)
io.Writerと同等の機能を提供します。

### func (self \*StdPipe) Close() error
io.Closerと同等の機能を提供します。

### func (self \*StdPipe) ReaderClose() error
標準入力側のみCloseします。

### func (self \*StdPipe) WriterClose() error
標準出力側のみCloseします。
