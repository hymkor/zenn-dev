---
title: "それでも僕たちは git add -p -e したい"
---
本bookのテキストを編集していたら

```
$ jj st
Working copy changes:
M books\c1e309aea68960\04_init.md
A books\c1e309aea68960\05_work.md
A books\c1e309aea68960\06_split.md
M books\c1e309aea68960\config.yaml
Working copy : zlukvlvx 99d5dabf (no description set)
Parent commit: uuzumwyw d8e5a922 jujutsu book: add 04_init
```

ちょっと、あちらこちらへ手を加えすぎてしまいました。 git なら個別に `git add` → `git commit` を繰り返して 3 つのコミットに分けることができます。

jj の場合は、現在の作業コピーが 1 コミットになっているので、これを3つに分ける形になります。どうすればいいかというと、`jj split` というコマンドを実行するだけです。

![jj split の実行イメージ](/images/jjsplit.png)

テキストエディターが起動するかと思いきや、独自の範囲の選択画面になりました。

画面では `[ ]` の後にファイル名がならんでいます。`[ ]` が `( )` になっている箇所が現在のカーソル位置のようです。`l` キーを押下すると、ファイル内の変更範囲などが展開されます。どうやら

+ `q` … コミット範囲の選択を破棄して終了
+ `c` … コミット範囲の選択を採用して終了
+ `SPACE` … カーソル位置の選択をコミット範囲に入れる or 外す
+ `↓`, `j` … カーソルを下へ移動
+ `↑`, `k` … カーソルを上へ移動
+ `→`, `l` … おりたたんでいる箇所を開く
+ `←`, `h` … おりたたんでいる箇所を閉じる

という操作で、分離する前半のコミット範囲の選択を行うことができるようです。`git add -p -e` とかだと「`-` の記号を空白にする」「`+` で始まる行を削除する」ということをテキストエディターで行わないといけなかったので、ミスも発生していたわけですが、`jj` のエディターなら安心です。

とりあえず、`04_init.md` だけ選択状態 `[x]` にして、`c` で終了します。
すると、テキストファイルが起動しますので、選択範囲についてのコミットログを書きます。

```
JJ: Enter commit description for the first part (parent).

JJ: This commit contains the following changes:
JJ:     M books\c1e309aea68960\04_init.md

JJ: Lines starting with "JJ: " (like this one) will be removed.
```

```
$ jj split
Using default editor ':builtin'; you can change this by setting ui.diff-editor
First part: zlukvlvx f164507f Fix: 04_init.md
Second part: zlyuxkws ac9c80d5 (no description set)
Working copy now at: zlyuxkws ac9c80d5 (no description set)
Parent commit      : zlukvlvx f164507f Fix: 04_init.md

$ jj st
Working copy changes:
A books\c1e309aea68960\05_work.md
A books\c1e309aea68960\06_split.md
M books\c1e309aea68960\config.yaml
A images\jjsplit.png
Working copy : zlyuxkws 6162265a (no description set)
Parent commit: zlukvlvx f164507f Fix: 04_init.md

$ jj log
@  zlyuxkws iyahaya@example.com 2024-01-30 20:21:30.000 +09:00 371e3fdb
│  (no description set)
◉  zlukvlvx iyahaya@example.com 2024-01-30 20:20:47.000 +09:00 f164507f
│  Fix: 04_init.md
◉  uuzumwyw iyahaya@example.com 2024-01-30 03:29:17.000 +09:00 d8e5a922
│  jujutsu book: add 04_init
：
```

予定どおり `04_init.md` への変更だけ、独立したコミットになりました。`05_work.md` に対しても同じことを繰り返せば、3つのコミットに分割が完了です。

