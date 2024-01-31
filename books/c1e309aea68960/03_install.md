---
title: "インストールと初期設定"
---
Linux, macOS, Windows についてはコンパイル済みのバイナリが GitHub の Releases にて配布されています。

- [Releases · martinvonz/jj](https://github.com/martinvonz/jj/releases)

同梱されている jj , jj.exe を環境変数PATHで指定されたディレクトリにコピーすれば Ok です。また、サポートしているパッケージマネージャもあります。

+ [Installation and Setup - Jujutsu docs](https://martinvonz.github.io/jj/v0.13.0/install-and-setup/)

jujutsu は Rust製プロダクトですので、

```
$ cargo binstall --strategies crate-meta-data jj-cli
```

でインストールできるとあります。その他、apt-get, nix, Homebrew, MacPorts などでもサポートされているようです。また、Windows では scoop の main bucket にも入っていたので

```
C:> scoop install jj
```

で導入可能でした。chocolatey は、2024-01-29時点ではまだのようです。

インストールができたら、初期設定だけやっておきましょう。Git みたいにユーザ名とメールアドレスを登録しておきます。

```
C:> jj config set --user user.name "(ユーザ名)"
C:> jj config set --user user.email "(Eメールアドレス)"
```