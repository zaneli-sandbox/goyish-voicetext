package voicetext

import "strconv"

type ttsOptions struct {
	emotion      *string
	emotionLevel *int
	pitch        *int
	speed        *int
	volume       *int
}

func TtsOptions() *ttsOptions {
	return &ttsOptions{}
}

func (o *ttsOptions) Emotion(emotion string) *ttsOptions {
	o.emotion = &emotion
	return o
}

func (o *ttsOptions) EmotionLevel(emotionLevel int) *ttsOptions {
	o.emotionLevel = &emotionLevel
	return o
}

func (o *ttsOptions) Pitch(pitch int) *ttsOptions {
	o.pitch = &pitch
	return o
}

func (o *ttsOptions) Speed(speed int) *ttsOptions {
	o.speed = &speed
	return o
}

func (o *ttsOptions) Volume(volume int) *ttsOptions {
	o.volume = &volume
	return o
}

func (o *ttsOptions) addOption(params map[string]string) {
	if o.emotion != nil {
		params["emotion"] = *o.emotion
	}
	if o.emotionLevel != nil {
		params["emotion_level"] = strconv.Itoa(*o.emotionLevel)
	}
	if o.pitch != nil {
		params["pitch"] = strconv.Itoa(*o.pitch)
	}
	if o.speed != nil {
		params["speed"] = strconv.Itoa(*o.speed)
	}
	if o.volume != nil {
		params["volume"] = strconv.Itoa(*o.volume)
	}
}
