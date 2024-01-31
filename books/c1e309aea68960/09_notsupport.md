---
title: "いまのところ出来ないこと"
---
v0.13.0 時点で出来なくて困ったことをあげます。そのうち、きっとなんとかしてもらえるでしょう(他力本願)

### タグ関連

#### 新規タグの作成

[Git compatibility - Jujutsu docs](https://martinvonz.github.io/jj/v0.13.0/git-compatibility/#supported-features)
> * **Tags: Partial.** You can check out tagged commits by name (pointed to be either annotated or lightweight tags), but you cannot create new tags.

タグは部分的サポートで、GitHub より読み込むことはできますが、新規作成はできないようです。この結果、運用としてタグを使っているレポジトリでは移行が難しくなっています。

( 自分も zenn.dev のこのテキストのレポジトリはもう jj で行っているのですが、他の個人開発物のレポジトリでは、リリース作業で `git tag` の使用を前提しており、移行に踏み出せていない状況です )

#### git describe 的なもの

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
