---
title: "ギッハブ大作戦 - 故郷への長い道"
---
`jj git clone` で GitHub からクローンしてきたレポジトリですから、push は `jj git push` で出来そうな気がします。

```
C:> jj git push
No branches point to the specified revisions.
Nothing changed.
C:>
```

ダメでした。ログを見ると、`main` はちゃんとありますが、clone した場所がちょっと変です。clone した時点を指しているように見えます。

```
C:> jj log
@  xmlsppuq iyahaya@nifty.com 2024-01-31 12:34:51.000 +09:00 4ecfc545
│  (no description set)
◉  xmkxylpk iyahaya@nifty.com 2024-01-31 02:22:26.000 +09:00 f835aec9
│  jujutsu book: add 07_edit
    ：
◉  zlsztkok iyahaya@nifty.com 2024-01-29 15:11:45.000 +09:00 3f66e36f
│  Add new book for `jj-book`: books/c1e309aea68960/
◉  poyqqryp iyahaya@nifty.com 2024-01-29 15:08:37.000 +09:00 3f37139d
│  Makefile: add entry: new-book
◉  nzytmsts iyahaya@nifty.com 2023-12-16 14:39:31.000 +09:00 main 0f62c769
│  バッチファイル入門: 日付書式のdddd が日本語の曜日となる点も補足
~
```

jj における branch は分岐全体を指すものではなく、特定のコミットを指すポインターになります。これはコミットを積み重ねると勝手に移動するものではなく、次のようなコマンドで移動させる必要があります。

```
C:> jj branch set -r @- main
```

`-r @-` は移動先のリビジョンを「現在の作業コピーの親」とします。`@` が現在の作業コピーで、`-` がその親を意味します[^current-branch]

[^current-branch]: 現在の作業コピーは空だったり、仕掛り中だったりして、あまり push したくない場合が多いですよね。

なお、クローンで複製したものではない、`jj git init` で作成したレポジトリの場合は、初回だけ

+ リモートレポジトリとの関連付け:  
    `jj git remote add origin (URL)`
+ ブランチの設定:  
    `jj branch create -r @- main`

など実行が必要になります。
