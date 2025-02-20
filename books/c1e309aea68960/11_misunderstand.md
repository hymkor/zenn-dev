---
title: "jj についての誤解"
---
### (誤解) jj には GitHub に相当するサービスがない

jj は Git 互換のプロトコルが使えますので、今使っている GitHub に普通に push できます。GitHub 側から見ると、git で作成されたコミットか jj で作成されたコミットが区別できないため、GitHub Action も普通に動作します。

### (誤解) jj から push すると、`git push -f` 相当になり危険

ひとりで使っている分には、一度、push されたコミットは immutable になって、そもそも変更できません。`git push -f` 相当のことを行うには、まず immutable を解除するところから始めなくてはいけません。

自分以外の誰かの手で GitHub 側のリモートブランチが進められてしまった時に `jj git push` を行うと次のようにエラーになります。別に知らぬ間に `git push -f` 相当のことになってしまうわけではありません。

```
$ jj git push
Changes to push to origin:
  Move forward bookmark master from 2160bd032eb3 to 4fbdabcab90e
Error: Refusing to push a bookmark that unexpectedly moved on the remote. Affected refs: refs/heads/master
Hint: Try fetching from the remote, then make the bookmark point to where you want it to be, and push again.
```

この競合状態を解決するには `jj git fetch` でリモートの変更を取り込んでから `jj bookmark` でブランチの先頭を再設定しなければいけません。

この時に敢えて上書きするような形に再設定し `jj git push` をしてしまうと、確かに実質 `git push -f` と同じになります。そこは確かに注意は必要です。
