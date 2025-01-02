---
title: "インストールと初期設定"
---
Linux, macOS, Windows についてはコンパイル済みのバイナリが GitHub の Releases にて配布されています。

- [Releases · jj-vcs/jj](https://github.com/jj-vcs/jj/releases)

同梱されている jj , jj.exe を環境変数PATHで指定されたディレクトリにコピーすれば Ok です。また、サポートしているパッケージマネージャもあります。

+ [Installation and Setup - Jujutsu docs](https://jj-vcs.github.io/jj/latest/install-and-setup/)

jujutsu は Rust製プロダクトですので、

```
$ cargo binstall --strategies crate-meta-data jj-cli
```

でインストールできるとあります。その他、apt-get, nix, Homebrew, MacPorts などでもサポートされているようです。また、Windows では WinGet と scoop installer が対応しています。

- `winget install jj-vcs.jj`
- `scoop install jj`

インストールができたら、初期設定だけやっておきましょう。Git みたいにユーザ名とメールアドレスを登録しておきます。

```
$ jj config set --user user.name "(ユーザ名)"
$ jj config set --user user.email "(Eメールアドレス)"
```
