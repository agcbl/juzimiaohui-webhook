package model

const (
	Unknown = 0
	Attachment = 1
	Audio = 2
	Contact = 3
	ChatHistory = 4
	Emoticon = 5
	Image = 6
	Text = 7
	Location = 8
	MiniProgram = 9
	Money = 10
	Recalled = 11
	Url = 12
	Video = 13
)

type MessagePayload struct {
	Text string `json:"text,omitempty"`
	VoiceUrl string `json:"voiceUrl,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
	VideoUrl string `json:"videoUrl,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`
	Url string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
	FileUrl string `json:"fileUrl,omitempty"`
	Content string `json:"content,omitempty"`
}

type WechatMessage struct {
	MessageId string `json:"messageId"`
	ChatId string `json:"chatId"`
	RoomTopic string `json:"roomTopic,omitempty"`
	RoomId string `json:"roomId,omitempty"`
	ContactName string `json:"contactName"`
	ContactId string `json:"contactId"`
	Payload MessagePayload `json:"payload"`
	Type int `json:"type"`
	Timestamp int `json:"timestamp"`
}

func (p *WechatMessage) GetContent() string {
	switch p.Type {
	case Text:
		return p.Payload.Text
	case Image:
		return p.Payload.ImageUrl
	case Url:
		return p.Payload.Url
	case Video:
		return p.Payload.VideoUrl
	case Audio:
		return p.Payload.VoiceUrl
	default:
		return ""
	}
}

type WechatMessageData struct {
	Data WechatMessage `json:"data"`
}
