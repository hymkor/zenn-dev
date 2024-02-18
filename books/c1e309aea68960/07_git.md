---
title: "ギッハブ大作戦 - 故郷への長い道"
---
### GitHub への push

`jj git clone` で GitHub からクローンしてきたレポジトリですから、push は `jj git push` で出来そうな気がします。

```
$ jj git push
No branches point to the specified revisions.
Nothing changed.

$
```

ダメでした。ログを見ると、`main` はちゃんとありますが、clone した場所がちょっと変です。clone した時点を指しているように見えます。

```
$ jj log
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
$ jj branch set -r @- main
```

`-r @-` は移動先のリビジョンを「現在の作業コピーの親」とします。`@` が現在の作業コピーで、`-` がその親を意味します[^current-branch]

※ Windows の PowerShell 上で行う場合、`@`マークは演算子と解釈されてしまうため、`"@-"` のように二重引用符で囲むか、`` `@-`` のように `` ` `` でエスケープする必要があります。

[^current-branch]: 現在の作業コピーは空だったり、仕掛り中だったりして、あまり push したくない場合が多いですよね。

### クローンしていないレポジトリの場合

なお、クローンで複製したものではない、`jj git init` で作成したレポジトリの場合は、初回だけ

+ リモートレポジトリとの関連付け:  
    `jj git remote add origin (URL)`
+ ブランチの設定:  
    `jj branch create -r @- main`

など実行が必要になります。

### 同一ワークディレクトリで jj/git を併用

既にある Git のワークディレクトリ上で

```
$ jj git init --git-repo=.
```

を実行すると、同じワークディレクトリで git と jj が併用できるようになります。[^wwg]

[^wwg]: [Working in a Git co-located repository](https://martinvonz.github.io/jj/v0.14.0/github/#working-in-a-git-co-located-repository)

この状態では「jj にはないが、git にあるコマンド」(`git tag`, `git describe --tag`)がそのまま使えますが、[git のカレントブランチがない状態][detached]になるため、gitの操作が一部制限されます。たとえば `git push` も次のようなエラーになります。

```
$ git push
fatal: You are not currently on a branch.
To push the history leading to the current (detached HEAD)
state now, use

    git push origin HEAD:<name-of-remote-branch>
```

これは一度 `jj branch track master@origin` を行えば、以後 `jj git push` で代替できるので問題ありません。

また、git側の状態は jj 側で認識されるため、一応、不整合な状態などは発生しないようです。`jj log` で確認すると、git の HEAD が `HEAD@git` とマークされており、一応ちゃんと認識されているようです。

[detached]: https://git-scm.com/docs/git-checkout#_detached_head

```
$ jj log
@  npzunqrv iyahaya@nifty.com 2024-02-16 23:41:53.000 +09:00 0a15bf3e
│  (empty) (no description set)
◉  vtopmqzn iyahaya@nifty.com 2024-02-16 23:41:53.000 +09:00 master HEAD@git 38c0cee6
│  update the manifest of the scoop-install for v1.1.3
~
```
