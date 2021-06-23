---
title: "Go言語によるWindows向けズボラ式文字コード判定法"
emoji: "ず"
type: "tech"
topics: ["go", "utf8", "windows"]
published: false
---

「[Windows と Unicode とボク](https://zenn.dev/zetamatta/books/b820d588f4856bcf836c)」でも記しました通り、Windows10 のメモ帳では UTF8 で保存した時、BOM コードを付けなくなりました。

結果として、UTF8 と ANSI を区別する確実な情報はなくなってしまったので、自前で判定しなくてはいけません。

本記事では、比較的、お手軽、かつ日本語環境に依存しない判別手順を提案したいと思います。

### 従来の方法の問題点

これまで nkf 等でよく行なわれていたのは、バイト列が ShiftJISの有効範囲を見たしているかどうかを判定するといったものでした。

が、この方法では ShiftJIS の有効範囲の判定が煩雑なうえ、コードが日本語環境ローカル依存になってしまいます。ダメです。

### 提案する方法

UTF8 の有効範囲で同じことをやればよいのです。有効範囲の判定は utf8.IsValid という標準関数があります。

だいたいのロジック

* テキストファイルを一行ずつ読みとる
* utf8.IsValid が true ならば、その行は UTF8 とする
* false の場合、その行を現在のコードページの文字列と過程して、  
    UTF16 への変換に成功したら非UTF8確定 (以降の行も全て非UTF8確定とする)
* 失敗したら判定失敗とする(バイナリかな？)

### サンプル

変換機能付き bufio.Scanner みたいなのを作ってみましょう。

判定方法は上記の通りですが、とりあえず、NonUTF8 to UTF8 の変換機能は別途必要です。

Windows には MultiByteToWideChar という API があって、これを用いれば、任意の文字コードと UTF16 との変換ができます。そして、Go の準標準ライブラリの golang.org/x/sys/windowsにそれを呼び出す定義があります。

これを使う手順としては

1. 1回目の呼び出しで、UTF16テキスト格納先として NULL を渡して、変換した結果を格納するのに必要なバッファのサイズを得る
2. 2回目の呼び出して、実際にUTF16へ変換する
3. UTF16 に変換したテキストを、UTF8 へ変換する

という段取りになります。これは一つの汎用関数として定義してしまいましょう。

```go
package mbcs

import (
    "bufio"
    "io"
    "unicode/utf8"

    "golang.org/x/sys/windows"
)

const _ACP = 0 // 現在のコードページを示す

func ansiToUtf8(mbcs []byte) (string, error) {
    // query buffer's size
    size, _ := windows.MultiByteToWideChar(
            _ACP, 0, &mbcs[0], int32(len(mbcs)), nil, 0)
    if size <= 0 {
            return "", windows.GetLastError()
    }

    // convert ansi to utf16
    utf16 := make([]uint16, size)
    rc, _ := windows.MultiByteToWideChar(
            _ACP, 0, &mbcs[0], int32(len(mbcs)), &utf16[0], size)
    if rc == 0 {
            return "", windows.GetLastError()
    }
    // convert utf16 to utf8
    return windows.UTF16ToString(utf16), nil
}
```

この関数は、現在のコードページの文字列を UTF8 へ変換するもので、既に UTF8 になっているものを通してしまうと、誤変換になるか、エラーになってしまいます。そのため、ちゃんと判定を行いましょう。

```go
type Filter struct {
    sc   *bufio.Scanner
    text string
    ansi bool
    err  error
}
```

これが文字コード判定機能付きの Scanner の型定義です。基本は bufio.Scanner をラップしますが、Scan() と Text() メソッドだけ、差し替えます。

```
func NewFilter(r io.Reader) *Filter {
    return &Filter{
        sc: bufio.NewScanner(r),
    }
}
```

コンストラクター。ゼロ値をうまいこと使って、サボります。

- メンバ Scanner : 内部で使う bufio.Scanner です
- メンバ text : 変換結果を格納する
- メンバ ansi : 非UTF8 であると確定していたら true。最初は不明なので false
- メンバ err : エラーーーーーーー

判定とか変換は全て Scan の中で行ってしまいます。

```go
func (f *Filter) Scan() bool {
    if !f.sc.Scan() {
        f.err = f.sc.Err()
        return false
    }
    line := f.sc.Bytes()
    if !f.ansi && utf8.Valid(line) {
        f.text = f.sc.Text()
    } else {
        f.text, f.err = ansiToUtf8(line)
        if f.err != nil {
            return false
        }
        f.ansi = true
    }
    return true
}
func (f *Filter) Text() string {
    return f.text
}

func (f *Filter) Err() error {
    return f.err
}
```

はたして、こんなんでちゃんと自動判別できるのでしょうか？サンプルコードで確認しましょう。

パッケージは暫定的に "mbcs" としてます。`go mod init mbcs` で初期化して、同じフォルダーに example.go として以下のコードを書きます。

そのままでは main と mbcs は共存できませんが、今回は main は go run 専用でよいので、一行目に `//+build ignore` を入れて go build の対象外にしましょう。

```go
//+build ignore

package main

import (
    "fmt"
    "os"

    "mbcs"
)

func main() {
    filter := mbcs.NewFilter(os.Stdin)
    for filter.Scan() {
        fmt.Println(filter.Text())
    }
    if err := filter.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
```

さて、いけるでしょうか？

テストその1:

```
$ nkf32 --guess sample1.txt
Shift_JIS (CRLF)
$ hexdump sample1.txt
53 68 69 66 74 4A 49 53 0D 0A 82 C5 8F 91 82 A2
82 BD 81 42 83 54 83 93 83 76 83 8B 83 65 83 4C
83 58 83 67 82 C5 82 B7 81 42 0D 0A 94 BB 92 E8
82 C5 82 AB 82 E9 82 A9 82 C8 0D 0A
$ go run example.go < sample1.txt
ShiftJIS
で書いた。サンプルテキストです。
判定できるかな
```

テストその2：

```
$ nkf32 --guess sample2.txt
UTF-8 (CRLF)
$ hexdump sample2.txt
55 54 46 38 0D 0A E3 81 A7 E6 9B B8 E3 81 84 E3
81 9F E3 80 82 E3 82 B5 E3 83 B3 E3 83 97 E3 83
AB E3 83 86 E3 82 AD E3 82 B9 E3 83 88 E3 81 A7
E3 81 99 E3 80 82 0D 0A E5 88 A4 E5 AE 9A E3 81
A7 E3 81 8D E3 82 8B E3 81 8B E3 81 AA 0D 0A
$ go run example.go < sample2.txt
UTF8
で書いた。サンプルテキストです。
判定できるかな
```

いけてるんちゃいます？？

なお、以上のコードは次のレポジトリより入手できます。

https://github.com/zetamatta/go-mbcs-to-utf8-sample
