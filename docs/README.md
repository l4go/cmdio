# golib/cmdio ライブラリ

コマンドの標準入出力から、[io.ReadWriteCloser](https://golang.org/pkg/io/#ReadWriteCloser)互換のI/Oを提供するライブラリです。
入出力をまとめて１つのI/Oとして管理できるようになります。

* [cmdio.CmdPipe](CmdPipe.md)
  * 外部コマンドを実行し、その標準入出力に対応するI/Oを提供します。
* [cmdio.StdPipe](StdPipe.md)
  * 実行したプロセスの標準入出力に対応するI/Oを提供します。
