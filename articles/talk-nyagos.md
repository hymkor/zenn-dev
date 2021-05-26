---
title: "できた！うかつにもしゃべってしまう nyagos が"
emoji: "📘"
type: "tech"
topics: ["go", "shell", "nyagos", "lua"]
published: true
---

- [Windows 10は好きな文章を合成音声で簡単に喋らせることができる - ASCII.jp](https://ascii.jp/elem/000/004/055/4055975/amp/)

サンプルは PowerShell ですが…

```ps1
$x=New-Object -ComObject SAPI.SpVoice
$x.Speak("早く人間に成りたい")
```

あ、COM のインスタンス作ってるじゃないですか。なら、nyagos の内蔵 Lua でもいけるんちゃうか？

```lua
-- talk.lua
nyagos.alias.talk = function(s)
    local sapi = nyagos.create_object("SAPI.SpVoice")
    for i=1,#s do
        sapi:Speak(s[i])
    end
    sapi:_release()
end
```

これを `talk.lua` という名前でセーブしてっと

```
$ lua_f talk.lua
$ talk 早く人間に成りたい
```

でけた！（音声を聞かせられないのが残念です。かといって、音声ファイルを UP するのもなぁ）

よし、これを応用すれば、プロンプトが表示されるタイミングで「Ready」とか言わせることで「蒼き流星SPTレイズナー」ごっこが出来るはず…  
( あ、なんか、タイムリーなことに YouTube のサンライズチャンネルで [第一話](https://www.youtube.com/watch?v=hmVRQ_sOCQY) が公開されていましたよ！)

```lua
-- vmax.lua
share.sapi = nyagos.create_object("SAPI.SpVoice")
nyagos.prompt = function(this)
    share.sapi:Speak("Ready")
    return default_prompt(this)
end
```

どや！

```
$ lua_f vmax.lua
```

びみょう…発音・音程はともかく…「Ready」と言った後、カーソルが点滅するまで１秒ほど妙な間があって、普段使いはとてもできない。

だが、天は見捨てなかった！

> なおデフォルトでは、同期呼び出しなので、発声が終わるまでSpeakメソッドは戻ってこない。これだとスクリプトの実行に差し障るというのであれば、非同期呼び出しとして発声が終わる前に制御を戻すこともできる。それには、後ろにもう1つ引数を追加して1（SVSFlagsAsync）をつける

おぉ

```lua
share.sapi = nyagos.create_object("SAPI.SpVoice")
nyagos.prompt = function(this)
    share.sapi:Speak("Ready",1)
    return default_prompt(this)
end
```

待ち時間がなくなった！やったね！
（音程はともかく、なんかいけそう）

まぁ、こういうお遊びはともかくとして、長時間処理を裏でやっている時に終わったことを知らせるアラームとかに使うのにいいかもしれませんね！