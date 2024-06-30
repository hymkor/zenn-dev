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

`jj split` の編集処理を担う外部パッケージ [scm-record] が更新された影響で、差分テキストのタブを適切に取り扱えなくなったようです。

+ [Screen corruption in `--interactive` with tabs · Issue #3944 · martinvonz/jj](https://github.com/martinvonz/jj/issues/3944) 
+ [bug: tab characters not rendered correctly · Issue #2 · arxanas/scm-record](https://github.com/arxanas/scm-record/issues/2)
+ [Built-in diff editor doesn't redraw screen correctly in the presence of tabs · Issue #4001 · martinvonz/jj](https://github.com/martinvonz/jj/issues/4001)

[scm-record] 側では修正パッチがマージされましたが、2024年7月4日現在まだ修正版はリリースされていません([scm-record]の最新リリースの v0.3 には含まれていません)

+ [fix: force tabs to fixed size by firestack · Pull Request #37 · arxanas/scm-record](https://github.com/arxanas/scm-record/pull/37)

2024年6月29日現在の対策としては、v0.17.1 あたりのバージョンへ戻すのが早いようです[^scoop]

### jj split で、削除されたファイルをコミットに含めることができない

削除されたファイルにチェックを入れても、チェックしていなかったように無視されてしまいます。
issue もあがっていました。

+ [The builtin diff editor for `jj split` mishandles truncating a file to 0 bytes · Issue #3526 · martinvonz/jj](https://github.com/martinvonz/jj/issues/3526)
+ [file deletion ignored by `jj split` · Issue #3702 · martinvonz/jj](https://github.com/martinvonz/jj/issues/3702)

回避策としては、`jj commit (削除したファイル名)` で別々に分割コミットして、あとから `jj squash` で一つのコミットにマージするのが一番早そうです。


[scm-record]: https://github.com/arxanas/scm-record

[^scoop]: scoop-installer でバージョンを戻す場合は `scoop reset jj@0.17.1` を実行します。
