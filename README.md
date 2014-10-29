# goyish-voicetext

[VoiceText Web API](https://cloud.voicetext.jp/webapi) Golang wrapper library

## Installation
```
go get github.com/zaneli/goyish-voicetext/voicetext
```

## Usage
```
import (
	"github.com/zaneli/goyish-voicetext/voicetext"
	"log"
	"os"
)

client := voicetext.NewClient(<YOUR_API_KEY>)
content, err := client.Tts("Hello world", voicetext.Show, nil)

if err != nil {
	log.Fatal(err)
}

file, err := os.Create("hello.wav")
if err != nil {
	log.Fatal(err)
}
file.Write(content)
```

```
client := voicetext.NewClient(<YOUR_API_KEY>)
content, err := client.Tts(
	"Hello world",
	voicetext.Haruka,
	voicetext.TtsOptions().Emotion(voicetext.Happiness).EmotionLevel(2).Pitch(50).Speed(150).Volume(120))

if err != nil {
	log.Fatal(err)
}

file, err := os.Create("hello.wav")
if err != nil {
	log.Fatal(err)
}
file.Write(content)
```
