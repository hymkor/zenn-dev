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
+ rebase 的なものも自動的に行われます。

過去のコミットを修正することにより後続のコミットで衝突が起きた場合、`jj log` でそのコミットに `conflict` のマークがつくことで確認できます。`jj edit (衝突した変更のID)` をした後、衝突箇所をテキストエディターで修正し、衝突を解消します。修正が終わったら、再度 `jj edit (最新ID)` もしくは `jj new --insert-after (最新ID)` などで、最新コミットに戻ります

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

### ファイルのサイズ制限を増やす

通常は 1MB 以上のサイズのファイルをワークディレクトリに置けないのですが

```
$ jj config set --repo snapshot.max-new-file-size 5MiB
```

などとすると、上限を5メガバイトに増やすことができます。 *ですが、これはやらない方がよいです* 

と言いますのは 1MB 以上に増やすと、誤ってバイナリファイルがコミットに含めてしまうリスクが増えるからです。素直に対象のファイルを .gitignore に加えるべきです。[^f]

[^f]: よくあるミスの連鎖は (1)サイズ増やす (2) makeで実行ファイルが作成されるが気付かない (3) そのまま `jj new`, `jj git push` (3) GitHub上に実行ファイルが乗ったのに気付く (4) ローカルのコミットから消して `git push -f` 相当のことをしたくなるが、push 済みのコミットは immutable なのでローカルからも消せない…という事態です。自分は仕方がないので、GitHub 側からコミットを削除して、それを `jj git fetch` で取り込んだ後、ローカルレポジトリの辻褄合わせをしました。

### Windows用の vim でコミットログの編集ができない

v0.15 前後で、コミットログを書くためのファイル名が `\\?\C:\...` 形式に正規化されるようになったようです[^rust-canonical] 。このパスに `?` が含まれるため、vim.exe はワイルドカードと誤認識し、ファイル名展開に失敗してしまうのが原因と考えられます。

[^rust-canonical]: Rust 標準のパス正規化関数 fs::canonicalize の仕様が原因のようです。

回避策としてワイルドカード展開を抑制するオプション --literal を与えれば Ok です。 `jj config edit --user` で

```
[ui]
editor = [ "C:/Users/hymkor/scoop/apps/vim/current/gvim.exe" , "--literal" ]
```

といった設定を追加しましょう。
