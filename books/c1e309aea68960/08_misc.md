---
title: "その他"
---
あとは、簡単に自分がやったことがあるものだけを軽くピックアップします。

### 過去コミットの「ログだけ」直す

```
jj desc (変更ID)
```

変更ID: `zlyuxkws` の `Add 05_work` を `jujutsu book: add 05_work` に直したい場合、`jj desc zly` で OK です。 テキストエディターが起動して、ログの編集ができます。ID は全部つづる必要はなく、紫に着色している数文字だけを省略系として利用できます。

### 過去のコミットの編集

```
jj edit (変更ID)
```

作業コピーを ID で指定したコミットへ切り替えます。

+ ここで変更すると、指定したコミットにただちに反映されます
+ rebase 的なものも自動的に行われるようです。
    + 変更の衝突が起きた場合は、コミットに `conflict` のマークがつきます  
      (`git log` で確認できます)
    + `jj edit (衝突した変更のID)` で衝突したコミットに切り替え、git と同じ要領で衝突箇所をテキストエディターで修正します
+ 修正が終わったら、再度 `jj edit (最新ID)` (および、必要に応じて `jj new`) で、最新コミットに戻ります[^better]

[^better]: もっとよいやり方がありそうな気もするんですが

### だいたい git と同じもの

|  git                    | jj           |
|-------------------------|--------------|
| `git restore`           | `jj restore` |
| `git diff`              | `jj diff`    |
| `git commit --amend -a` | `jj squash`  |

くわしくは [Git comparison - Jujutsu docs &gt; Command equivalence table](https://martinvonz.github.io/jj/latest/git-comparison/#command-equivalence-table)


### リビジョン指定(revset)

`jj log`、`jj show` などで ID 指定に使える式です。


| 指定          | 意味                       |
|---------------|----------------------------|
| `@`           | 現在の作業コピー           |
| `foo-`        | foo の親                   |
| `foo+`        | foo の子                   |
| `foo::`       | foo の子孫                 |
| `::foo`       | foo の先祖                 |
| `tags()`      | 全てのタグ                 |
| `latest(foo)` | foo のうちの最も新しいもの |
| `root()`      | ルートコミット             |

`latest(tags())::` という表現もできます。くわしくは [Revset language - Jujutsu docs](https://martinvonz.github.io/jj/latest/revsets/)

### ファイルを管理対象外にする

1. `.gitignore` にそのファイルを含むパターンを加える  
   ※ `.jjignore` というものは v0.14 時点ではまだない[^ignored_files]
2. `jj untrack (ファイル名)`

[^ignored_files]: https://martinvonz.github.io/jj/v0.14.0/working-copy/#ignored-files
