---
title: "最初のチェンジ！"
---
### ファイルの追加

jujutsu では `git add` のような操作は不要で、作業ディレクトリにファイルをコピーしてくるなり、既存のファイルを変更するだけで、管理対象として認識されます。

```
$ jj status
The working copy is clean
Working copy : tusovlyu 6ee1456d (empty) (no description set)
Parent commit: zzzzzzzz 00000000 (empty) (no description set)

$ copy %USERPROFILE%\.nyagos .
C:\Users\hymkor\.nyagos -> .nyagos

$ jj status
Working copy changes:
A .nyagos
Working copy : tusovlyu a5d5ecab (no description set)
Parent commit: zzzzzzzz 00000000 (empty) (no description set)

$
```

### コミットの登録

`git commit` には「コミットログを書く」という操作と「コミットを登録する」という二つの機能がありましたが、jj では両者は分離されています。

`jj describe` もしくは `jj desc` は、作業コピーのログを書くコマンドです。実行するとエディター[^editor]が起動して、作業コピーのログを更新できます。これは何回でも出来ますし、実行したからといって、それで現在の変更を確定させるわけではありません。

```
$ vim .nyagos

$ jj desc
Working copy now at: tusovlyu f6479066 Add .nyagos
Parent commit      : zzzzzzzz 00000000 (empty) (no description set)

$ jj log
@  tusovlyu iyahaya@nifty.com 2024-01-30 19:08:48.000 +09:00 f6479066
│  Add .nyagos
◉  zzzzzzzz root() 00000000

$
```

そして `jj new` で作業コピーをコミットとして確定します。作業コピーのログがそのままコミットログとなりますので、特にエディターも起動せず、すぐ終了します。そして、作業コピーは無変更状態というかたちになります。

```

$ jj new
Working copy now at: xyyypnuy cf9ee955 (empty) (no description set)
Parent commit      : tusovlyu f6479066 Add .nyagos

$ jj status
The working copy is clean
Working copy : xyyypnuy cf9ee955 (empty) (no description set)
Parent commit: tusovlyu f6479066 Add .nyagos

$ jj log
@  xyyypnuy iyahaya@nifty.com 2024-01-30 19:10:31.000 +09:00 cf9ee955
│  (empty) (no description set)
◉  tusovlyu iyahaya@nifty.com 2024-01-30 19:08:48.000 +09:00 f6479066
│  Add .nyagos
◉  zzzzzzzz root() 00000000

$
```

`jj describe` と `jj new` に機能が分かれているのにはメリットがあります。git などの場合、一つのコミットの内容が大きくなると、機能を実装するタイミングとコミットするタイミングが離れてしまうことがあります。そうすると修正内容の詳細を忘れてしまうことがあります。

jj のように `jj describe` が分離されていると、「修正は続くが、とりあえず今まで出来た分の情報を忘れないうちに書いておく」ということができます。

なお、`jj describe` と `jj new` とセットで行う `jj commit` というコマンドもあります。

[^editor]: git と同じで、環境変数EDITORで指定されたエディターが呼び出されます。
