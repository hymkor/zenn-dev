---
title: "いまのところ出来ないこと"
---
v0.14.0 時点で出来なくて困ったことをあげます。そのうち、きっとなんとかしてもらえるでしょう(他力本願)

### 新規タグの作成

[Git compatibility - Jujutsu docs](https://martinvonz.github.io/jj/v0.14.0/git-compatibility/#supported-features)
> * **Tags: Partial.** You can check out tagged commits by name (pointed to be either annotated or lightweight tags), but you cannot create new tags.

タグは部分的サポートで、GitHub より読み込むことはできますが、新規作成はできないようです。今のところ、ブラウザでGitHub側でタグを打って、それを `jj git fetch` で取り込むしか無いようです。

### git describe

プログラムにバージョン文字列(`(最後に打たれたタグ)-(その後積まれたリビジョンの数)-(最新コミットのハッシュ)`)を組み込むために、Makefile に

```
VERSION:=$(shell git describe --tags >nul || echo v0.0.0)
```

などとしたりしますが、現状、対応するものはなさそうです。ただ、

```
C:> jj log -r "latest(tags())::"
```

で、最後に打たれたタグ〜現在のコミットまでのログを出すことができるので、それをスクリプトで解析して等価なことは可能です[^jjtagdesc]

[^jjtagdesc]: [hymkor/jjtagdesc: An emurator of \`git describe --tags\` on the DVCS Jujutsu](https://github.com/hymkor/jjtagdesc)

### 改行コード変換 (Windows)

Git for Windows のような改行コード変換はありません。[^crlf]

+ `.gitattributes` のようなテキストファイルとバイナリファイルを区別するような動作はない
+ `core.autocrlf` のような設定もない (常に変換しない)

Linux で LF で登録された改行は Windows では LF でチェックアウトされますし、Windows で登録した CRLF は Linux でも CRLF でチェックアウトされます。

Windows で Linux 向けソースのメンテをしてきたチームで運用する場合は注意が必要です。[^crlf2]

[^crlf]: [Working on Windows &gt; Line endings are not converted](https://martinvonz.github.io/jj/v0.14.0/windows/#line-endings-are-not-converted)
[^crlf2]: 私見ですが、テキストファイルでどういう改行コードを使うかはテキストエディターの設定で調整すべきものだと考えます

### git push -f

GitHub に push してしまうと、リモートだけでなく、ローカルのコミットにも immutable なフラグがつき、もう編集不能になります。

immutable なフラグを外せればよいのですが、外し方が今のところ不明ですし、かつ外せても `git push -f` 的なことはできないように思われます。

こうなってしまうと、もう、一旦他のディレクトリで `git clone` して、そこで過去コミットを編集してから `git push -f` をするしかないのかもしれません。そうして GitHub 側の内容をうまく直せたら、jj のローカルレポジトリで `jj git fetch` すれば直った内容を取り込むことができます。
