---
title: "不具合・トラブルシューティング"
---
### Windows で指定のテキストエディターが起動しない

テキストエディターの起動パスは、[Editor - Settings - Jujutsu docs](https://martinvonz.github.io/jj/v0.15.1/config/#editor) に記載のとおり

1. 環境変数 $JJ\_EDITOR
2. 設定ファイル(TOML)の ui.editor
3. 環境変数 $VISUAL
4. 環境変数 $EDITOR

の順で探します。

起動パス中の空白はパラメータ区切りとして認識されるため、`C:/Program Files/Notepad++/notepad++.exe` などは `C:/Program` というエディター名と誤認識されます。その場合は、一旦メモ帳などをエディターに指定してから、`jj config edit --user` で設定ファイルを開いて次のように指定してください。

```
[ui]
editor = ["C:/Program Files/Notepad++/notepad++.exe",
    "-multiInst", "-notabbar", "-nosession", "-noPlugin"]
```

### Windows用の vim でコミットログの編集ができない

v0.15 前後で、コミットログを書くためのファイル名が `\\?\C:\...` 形式に正規化されるようになったようです[^rust-canonical] 。このパスに `?` が含まれるため、vim.exe はワイルドカードと誤認識し、ファイル名展開に失敗してしまうのが原因と考えられます。

[^rust-canonical]: Rust 標準のパス正規化関数 fs::canonicalize の仕様が原因のようです。

回避策としてワイルドカード展開を抑制するオプション --literal を与えれば Ok です。

```
jj config set --user ui.editor "%USERPROFILE%\scoop\apps\vim\current\vim.exe --literal"
```

### v0.18.0 にて、jj split の画面が乱れる

`jj split` の編集処理を担う外部パッケージ [scm-record] が v0.3.0 へ更新された影響で、差分テキストのCR や TAB をうまく表示できなくなったようです。

+ [Screen corruption in `--interactive` with tabs · Issue #3944 · martinvonz/jj](https://github.com/martinvonz/jj/issues/3944) 
+ [bug: tab characters not rendered correctly · Issue #2 · arxanas/scm-record](https://github.com/arxanas/scm-record/issues/2)
+ [Built-in diff editor doesn't redraw screen correctly in the presence of tabs · Issue #4001 · martinvonz/jj](https://github.com/martinvonz/jj/issues/4001)

[scm-record] 側ではこの不具合を修正するパッチが既にマージされていますが、2025年9月5日現在まだ [scm-record] や jj の最新リリースに含まれていません

+ [fix: force tabs to fixed size by firestack · Pull Request #37 · arxanas/scm-record](https://github.com/arxanas/scm-record/pull/37)

2024年9月5日現在の対策としては二種類考えられます。

#### (1) バージョン v0.17.1 を使い続ける

一番、手早い方法です。 scoop-installer でバージョンを戻す場合は `scoop reset jj@0.17.1` を実行すれば Ok です。ただし、最新機能が使えないため、アーリーアダプタ気分が台無しです。

#### (2) scm-record 部分を別途ビルドする

Rust の開発環境がある場合は [scm-record] を実行ファイル scm-diff-editor を自分でビルドして利用するという方法があります。

以下、Windows の場合の手順です

```
C:> git clone https://github.com/arxanas/scm-record.git
C:> cd scm-record\scm-diff-editor
C:> cargo build --release
C:> copy ..\target\scm-diff-editor.exe (%PATH%の通ったディレクトリ)
```

`jj config edit --user` で scm-diff-editor を使う設定に追加する

```
[ui]
diff-editor = "scm-diff-editor"
merge-editor = "scm-diff-editor"

[merge-tools.scm-diff-editor]
program = "scm-diff-editor.exe"
edit-args = ["--dir-diff", "$left", "$right"]
```

### jj split で、削除されたファイルをコミットに含めることができない

削除されたファイルにチェックを入れても、チェックしていなかったように無視されてしまいます。
issue もあがっていました。

+ [The builtin diff editor for `jj split` mishandles truncating a file to 0 bytes · Issue #3526 · martinvonz/jj](https://github.com/martinvonz/jj/issues/3526)
+ [file deletion ignored by `jj split` · Issue #3702 · martinvonz/jj](https://github.com/martinvonz/jj/issues/3702)

回避策としては、`jj commit (削除したファイル名)` で別々に分割コミットして、あとから `jj squash` で一つのコミットにマージするのが一番早そうです。

### Linux の実行ファイル名が自動でコミットに含まれがちで困る

実行ファイル名は「拡張子のない、英小文字だけのファイル名になりがち」と想定して、次のようなエントリをグローバルの .gitignore に追加するとよいかもしれません

```
[a-z]
[a-z][a-z]
[a-z][a-z][a-z]
[a-z][a-z][a-z][a-z]
[a-z][a-z][a-z][a-z][a-z]
[a-z][a-z][a-z][a-z][a-z][a-z]
[a-z][a-z][a-z][a-z][a-z][a-z][a-z]
[a-z][a-z][a-z][a-z][a-z][a-z][a-z][a-z]
![a-z]/
![a-z][a-z]/
![a-z][a-z][a-z]/
![a-z][a-z][a-z][a-z]/
![a-z][a-z][a-z][a-z][a-z]/
![a-z][a-z][a-z][a-z][a-z][a-z]/
![a-z][a-z][a-z][a-z][a-z][a-z][a-z]/
![a-z][a-z][a-z][a-z][a-z][a-z][a-z][a-z]/
```

[scm-record]: https://github.com/arxanas/scm-record
