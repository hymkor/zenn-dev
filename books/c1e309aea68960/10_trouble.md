---
title: "不具合・トラブルシューティング"
---
### Windows で指定のテキストエディターが起動しない

テキストエディターの起動パスは、[Editor - Settings - Jujutsu docs](https://jj-vcs.github.io/jj/v0.15.1/config/#editor) に記載のとおり

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

### ~~Windows用の vim でコミットログの編集ができない~~

→ [v0.25.0] にて解消されました (2025-01-02)  
  必要がない限り、`\\?\`形式のファイル名が使われなくなりました

[v0.25.0]: https://github.com/jj-vcs/jj/releases/tag/v0.25.0

### ~~v0.18 〜 v0.22 にて、jj split の画面が乱れる~~

→ [v0.23.0] にて解消されました (2024-11-08)  
  原因のパッケージ: scm-record が修正版 [v0.4.0][scm-record v0.4.0] に更新されました。

[scm-record v0.4.0]: https://github.com/arxanas/scm-record/releases/tag/v0.4.0
[v0.23.0]: https://github.com/jj-vcs/jj/releases/tag/v0.23.0

### jj split で、削除されたファイルをコミットに含めることができない

削除されたファイルにチェックを入れても、チェックしていなかったように無視されてしまいます。
issue もあがっていました。

+ [The builtin diff editor for `jj split` mishandles truncating a file to 0 bytes · Issue #3526 · jj-vcs/jj](https://github.com/jj-vcs/jj/issues/3526)
+ [file deletion ignored by `jj split` · Issue #3702 · jj-vcs/jj](https://github.com/jj-vcs/jj/issues/3702)

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
