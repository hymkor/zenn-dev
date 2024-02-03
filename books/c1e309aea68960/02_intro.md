---
title: "バージョン管理システム Jujutsu とは"
---
Jujutsu とは、なにやらネタっぽい名前ですが、Google所属の開発者 martinvonz 氏によるちゃんとしたプロダクトです。

ただ、名前が災いしてか、ネット検索しても「呪術廻戦」ばかりヒットしてしてしまい、なかなか有用な情報が得られません *(「`martinvonz/jj`」が一番ヒットしやすいかなぁ)*。ですが、機械翻訳を頼れば、英語の公式情報でも十分です。

+ ホームページ：[Jujutsu docs](https://martinvonz.github.io/jj/latest)
+ レポジトリ：[github.com/martinvonz/jj](https://github.com/martinvonz/jj)
+ Jujutsu の紹介記事：[jj init — Sympolymathesy, by Chris Krycho](https://v5.chriskrycho.com/essays/jj-init/)

とはいえ、ある程度、噛みくだいた日本語の説明もほしいところです。そこで本書では、簡単に「Jujutsu を使って、どういうメリットがあるのか。どう使えばいいのか」を説明し、利用のとっかかりになるものを目指したいと思います。

まず、ざっくりした特徴は次のとおりです。

+ 前の章で説明したとおり、コミット漏れを防ぐシステムとなっている
    + 現在の作業コピーも最新コミットであるかのように扱われる
    + 別のブランチに移動した時に仕掛り中の変更が消えるということはない
    + ファイルを新規作成したり、変更すると、即反映される
    + 逆に含めないようにする時が、ひと手間必要
+ `git add` に相当する操作がないということはコミットを適切な範囲に編集することができないのではないか？  
    + No. `jj split` でコミットを適宜分割できる。しかも編集エディターが結構使いやすい！
+ 「ファイルの変更」の他、「ユーザの操作」も履歴管理される
    + ファイルの変更ログは `jj log`、 操作ログは `jj op log`
    + Git の場合は操作ごとにそれぞれの取り消しの方法があったが、Jujutsu は `jj undo` で統一されている
+ Git 連携
    + `jj git clone URL` で、既存の Git レポジトリもそのまま Jujutsu で使える。無論一方通行ではなく、GitHub への push も可能。
+ 安全な同時レプリケーション (Safe, concurrent replication)
    + ロックファイルを作らず、作業ログによるマージを行うことにより、rsync、NFS、Dropbox などによる同期へある程度の対応しているそうです。ただし、Gitバックエンドが完全にロックフリーになっているわけではないため、まだレポジトリ破損の可能性は残っています。詳しくは [本家ドキュメントの Concurrency](https://martinvonz.github.io/jj/v0.13.0/technical/concurrency/)を参照のこと
