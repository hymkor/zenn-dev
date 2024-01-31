---
title: "その他"
---
あとは、簡単に自分がやったことがあるものだけを軽くピックアップします。

### 過去のコミットログだけを直す

```
jj desc (ID)
```

変更ID: `zlyuxkws` の `Add 05_work` を `jujutsu book: add 05_work` に直したい場合、`jj desc zly` で OK です。 テキストエディターが起動して、ログの編集ができます。ID は全部つづる必要はなく、紫に着色している数文字だけを省略系として利用できます。

### 過去のコミットの編集

```
jj edit (変更ID)
```

指定した ID の変更の編集モードになります。ここで変更すると、指定したコミットにただちに反映されます。rebase 的なものも自動的に行われるようです。もし、変更の衝突が起きた場合は、コミットに `conflict` のマークがつくので、`jj edit (衝突した変更のID)` で直せばよいようです。

### だいたい git と同じもの

|  git                    | jj           |
|-------------------------|--------------|
| `git restore`           | `jj restore` |
| `git diff`              | `jj diff`    |
| `git commit --amend -a` | `jj squash`  |

### リビジョン指定(revset)

| 指定          | 意味                       |
|---------------|----------------------------|
| `@`           | 現在の作業コピー           |
| `foo-`        | foo の親                   |
| `foo+`        | foo の子                   |
| `foo::`       | foo の子孫                 |
| `::foo`       | foo の先祖                 |
| `tags()`      | 全てのタグ                 |
| `latest(foo)` | foo のうちの最も新しいもの |

リビジョン指定は「式」のようなものなので、`latest(tags())::` というような表現もできる。

### ファイルを管理対象外にする

1. `.gitignore` にそのファイルを含むパターンを加える  
   ※ `.jjignore` というものは v0.13.0 時点ではまだない[^ignored_files]
2. `jj untrack (ファイル名)`

[^ignored_files]: https://martinvonz.github.io/jj/v0.13.0/working-copy/#ignored-files
