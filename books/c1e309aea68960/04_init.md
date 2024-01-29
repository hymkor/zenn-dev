---
title: "レポジトリを領域展開"
---
では、以降は実際に行う順番に Jujutsu の操作を説明します。

### GitHub の既存レポジトリのクローン

まずは、GitHub の既存レポジトリをローカルに落としてきましょう。

Git と互換系の処理は主に jj のサブコマンド git を使います。


```
C:> jj git clone https://github.com/hymkor/go-htnblog.git
Fetching into new repo in "\\?\C:\Users\hymkor\tmp\go-htnblog"
Working copy now at: xytwpzqm 610b0e40 (empty) (no description set)
Parent commit      : xqnlmwrl 7e60edf3 master | htnblog.exe: prevent from refering $EDITOR twice to edit draft
Added 25 files, modified 0 files, removed 0 files
C:>
```

ただ、残念ながら、git の認証方法によっては、失敗する場合があります。

```
C:> jj git clone git@github.com:hymkor/go-htnblog.git
Fetching into new repo in "\\?\C:\Users\hymkor\go-htnblog"
Error: Failed to authenticate SSH session: ; class=Ssh (23)
Hint: Jujutsu uses libssh2, which doesn't respect ~/.ssh/config. Does `ssh -F /dev/null` to the host work?
C:>
```

[Git compatibility - Jujutsu docs](https://martinvonz.github.io/jj/v0.13.0/git-compatibility/)によると、サポートされているのは以下だけとのことです。

+ ssh-agent
+ a password-less key ( only `~/.ssh/id_rsa`, `~/.ssh/id_ed25519` or `~/.ssh/id_ed25519_sk)`
+ a credential.helper

### 新規レポジトリの作成

GitHub にレポジトリがなく、新規にローカルで Jujutsu 管理を始める場合は、そのディレクトリに移動してから init サブコマンドを使います。

```
C:> jj init --git
```

今のところ、とりあえず `--git` オプションが必須のようです。なしで実行すると

```
C:> jj init
Error: The native backend is disallowed by default.
Hint: Did you mean to pass `--git`?
Set `ui.allow-init-native` to allow initializing a repo with the native backend.
C:>
```

と怒られます。これはおそらくですが、`--git` を最初につけておかないと、あとから GitHub へ push するのが困難になるからではないかと思われます。そのため、明示的に強制実行する形にしないと、git 連携なしには出来ないようにしているのでしょう 【要確認】

### 作業ディレクトリの状況の確認

`jj init` などの後、作業ディレクトリの状況は `jj status` もしくは省略形の `jj st` で確認できます。

```
C:> jj st
The working copy is clean
Working copy : tusovlyu 6ee1456d (empty) (no description set)
Parent commit: zzzzzzzz 00000000 (empty) (no description set)
```

現在の作業コピーが (empty) となっているのはよいとして、その親があって、それも (empty) というのはちょっとおかしく見えます。`jj log` で履歴を見ましょう。

```
C:> jj log
@  tusovlyu iyahaya@nifty.com 2024-01-30 18:36:07.000 +09:00 6ee1456d
│  (empty) (no description set)
◉  zzzzzzzz root() 00000000
C:>
```

親ディレクトリは `root()` という空コミットのようなもののようです。`git` では最初のコミット以前のゼロ状態を指定するのが何かと難しかったので、それに配慮した仕組みかと思われます。

ログの列は

1. 変更ID
2. ユーザ名
3. 日時
4. コミットID

の順にならんでいます。ユーザはコミットを指定するために、変更IDとコミットIDのどちらを使ってもよいとされています。変更IDは一度決まれば変わりませんが、コミットIDはそのコミットを変更するたびに変わるものであるために変更IDを指定した方がよさそうです。
