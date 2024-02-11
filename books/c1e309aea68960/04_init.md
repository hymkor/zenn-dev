---
title: "君のレポジトリを領域展開"
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

[Git compatibility - Jujutsu docs](https://martinvonz.github.io/jj/latest/git-compatibility/)によると、サポートされているのは以下だけとのことです。

+ ssh-agent
+ a password-less key ( only `~/.ssh/id_rsa`, `~/.ssh/id_ed25519` or `~/.ssh/id_ed25519_sk)`
+ a credential.helper

### 新規レポジトリの作成

GitHub にレポジトリがなく、新規にローカルで Jujutsu 管理を始める場合は、そのディレクトリに移動してから次のコマンドを発行します。[^init-git]

```
C:> jj git init
```

[^init-git]: jj v0.13.0 までは `jj init --git` というコマンドでしたが、v0.14.0 より Deprecated になり、`jj git init` を使うことになりました。

`git`サブコマンドではない `jj init` もありますが…

```
$ jj init
Error: The native backend is disallowed by default.
Hint: Did you mean to call `jj git init`?
Set `ui.allow-init-native` to allow initializing a repo with the native backend.
```

と怒られます。設定を変更すると、`git` なしで実行できるようですが、[そうした場合](https://martinvonz.github.io/jj/latest/git-comparison/#command-equivalence-table)、 jujutsu ネイティブのレポジトリが作成されます。今のところ遅い上に、clone がまだできないようです。

### 作業ディレクトリの状況の確認

`jj git init` の後、作業ディレクトリの状況は `jj status` もしくは省略形の `jj st` で確認できます。

![](/images/jj-init-and-st.png)

現在の作業コピーが `(empty)` となっているのはよいとして、その親も `(empty)` というのはちょっと変です。`jj log` で履歴を見ましょう。

![](/images/jj-init-and-log.png)

親ディレクトリは[ルートコミット](https://martinvonz.github.io/jj/latest/glossary/#root-commit)といって、全レポジトリのルートとなる仮想的なコミットです。ルートコミットはリビジョン指定[^r]の箇所に `root()` という関数を指定することで参照できます。

[^r]: jj のリビジョン指定は単なる名前指定ではなく、revset と呼ばれる式で範囲指定できるようになっています。詳しくは後述

ログの列は

1. 変更ID (色が変わっている部分が短縮名)
2. ユーザ名
3. 日時
4. (もしあれば)タグやブランチ名
5. コミットID (色が変わっている部分が短縮名)

の順にならんでいます。ユーザはコミットを指定するために、変更IDとコミットIDのID全部、短縮名のどちらを使ってもよいとされています。変更IDは一度決まれば変わりませんが、コミットIDはそのコミットを変更するたびに変わるものであるために変更IDを指定した方がよさそうです。
