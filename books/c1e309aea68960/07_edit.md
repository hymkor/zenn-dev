---
title: "過去のコミットの編集"
---
雑にやったので、コミットが乱れてしまいました。ちょっと直しましょう。

### コミットログだけを直す

まず、コミットログが汚い。これを直すのはすごく簡単です。

変更ID: `zlyuxkws` の `Add 05_work` を `jujutsu book: add 05_work` に直したい場合、`jj desc zly` で OK です。 テキストエディターが起動して、ログの編集ができます。ID は全部つづる必要はなく、紫に着色している数文字だけを省略系として利用できます。楽でいいですね。

### コミットの統合

また、すでにコミットしたコミットに、今の修正を統合することも簡単です。`jj squash` で現在の作業コピーを、親のコミットに組み込めます。

```
$ jj st
Working copy changes:
M 07_edit.md
Working copy : kqtxqzoy 0375f7cc (no description set)
Parent commit: xmkxylpk 941ca129 jujutsu book: add 07_edit

$ jj log
@  kqtxqzoy iyahaya@nifty.com 2024-01-30 22:24:07.000 +09:00 0375f7cc
│  (no description set)
◉  xmkxylpk iyahaya@nifty.com 2024-01-30 22:20:01.000 +09:00 941ca129
│  jujutsu book: add 07_edit
◉  prswsnss iyahaya@nifty.com 2024-01-30 21:29:09.000 +09:00 fcbc52ef
│  jujutsu book: add 06_split

$ jj squash
Working copy now at: ztpmmkmq 0207d3bc (empty) (no description set)
Parent commit      : xmkxylpk 2d2ab331 jujutsu book: add 07_edit

$ jj st
The working copy is clean
Working copy : ztpmmkmq 0207d3bc (empty) (no description set)
Parent commit: xmkxylpk 2d2ab331 jujutsu book: add 07_edit
```
