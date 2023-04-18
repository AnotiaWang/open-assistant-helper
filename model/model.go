package model

var ConfigKey = "config"
var ChatGPTClientKey = "chatgpt"

type Config struct {
	ApiKey   string `json:"apiKey"`
	OaCookie string `json:"oaCookie"`
	Language string `json:"language"`
}

type OALabel string

const (
	Spam           OALabel = "spam"
	LangMismatch   OALabel = "lang_mismatch"
	FailsTask      OALabel = "fails_task"
	Quality        OALabel = "quality"
	Helpfulness    OALabel = "helpfulness"
	Creativity     OALabel = "creativity"
	Humor          OALabel = "humor"
	Toxicity       OALabel = "toxicity"
	Violence       OALabel = "violence"
	NotAppropriate OALabel = "not_appropriate"
	PII            OALabel = "pii"
	HateSpeech     OALabel = "hate_speech"
	SexualContent  OALabel = "sexual_content"
)

type OARandomTaskLabel struct {
	Name        OALabel `json:"name"`
	Widget      string  `json:"widget"`
	HelpText    *string `json:"help_text,omitempty"`
	DisplayText string  `json:"display_text"`
}

type OAConversationMessage struct {
	Id                string                 `json:"id"`
	Lang              string                 `json:"lang"`
	Text              string                 `json:"text"`
	Emojis            map[string]interface{} `json:"emojis"`
	UserId            string                 `json:"user_id"`
	Synthetic         bool                   `json:"synthetic"`
	UserEmojis        interface{}            `json:"user_emojis"` // Unknown type in TypeScript
	IsAssistant       bool                   `json:"is_assistant"`
	UserIsAuthor      *bool                  `json:"user_is_author,omitempty"`
	FrontendMessageId string                 `json:"frontend_message_id"`
}

type OAConversation struct {
	Messages []OAConversationMessage `json:"messages"`
}

type OALabelAssistantReplyTask struct {
	Id              string              `json:"id"`
	Mode            string              `json:"mode"`
	Type            string              `json:"type"`
	Reply           string              `json:"reply"`
	Labels          []OARandomTaskLabel `json:"labels"`
	MessageId       string              `json:"message_id"`
	Disposition     string              `json:"disposition"`
	Conversation    OAConversation      `json:"conversation"`
	ValidLabels     []OALabel           `json:"valid_labels"`
	MandatoryLabels []OALabel           `json:"mandatory_labels"`
}

type OALabelPrompterReplyTask struct {
	Id              string              `json:"id"`
	Mode            string              `json:"mode"`
	Type            string              `json:"type"`
	Reply           string              `json:"reply"`
	Labels          []OARandomTaskLabel `json:"labels"`
	MessageId       string              `json:"message_id"`
	Disposition     string              `json:"disposition"`
	Conversation    OAConversation      `json:"conversation"`
	ValidLabels     []OALabel           `json:"valid_labels"`
	MandatoryLabels []OALabel           `json:"mandatory_labels"`
	UserId          string              `json:"userId"`
}

type OALabelInitialPromptTask struct {
	Id              string              `json:"id"`
	Mode            string              `json:"mode"`
	Type            string              `json:"type"`
	Labels          []OARandomTaskLabel `json:"labels"`
	Prompt          string              `json:"prompt"`
	MessageId       string              `json:"message_id"`
	Disposition     string              `json:"disposition"`
	Conversation    OAConversation      `json:"conversation"`
	ValidLabels     []OALabel           `json:"valid_labels"`
	MandatoryLabels []OALabel           `json:"mandatory_labels"`
	UserId          string              `json:"userId"`
}

type OAAssistantReplyTask struct {
	Id           string         `json:"id"`
	Type         string         `json:"type"`
	Conversation OAConversation `json:"conversation"`
	UserId       string         `json:"userId"`
}

type OAPrompterReplyTask struct {
	Id           string         `json:"id"`
	Hint         *string        `json:"hint,omitempty"`
	Type         string         `json:"type"`
	Conversation OAConversation `json:"conversation"`
	UserId       string         `json:"userId"`
}

type OARankAssistantRepliesTask struct {
	Id              string                  `json:"id"`
	Type            string                  `json:"type"`
	Replies         []string                `json:"replies"`
	Conversation    OAConversation          `json:"conversation"`
	ReplyMessages   []OAConversationMessage `json:"reply_messages"`
	MessageTreeId   string                  `json:"message_tree_id"`
	RevealSynthetic bool                    `json:"reveal_synthetic"`
	RankingParentId string                  `json:"ranking_parent_id"`
}

type OARandomTaskResponse struct {
	Id     string      `json:"id"`
	Task   interface{} `json:"task"` // Unknown type in TypeScript
	UserId string      `json:"userId"`
}

type PostBodyLabelAssistantContent struct {
	Labels    map[OALabel]float32 `json:"labels"`
	MessageId string              `json:"message_id"`
	Text      string              `json:"text"`
}
type PostBodyLabelAssistant struct {
	Id         string                        `json:"id"`
	Lang       string                        `json:"lang"`
	UpdateType string                        `json:"update_type"`
	Content    PostBodyLabelAssistantContent `json:"content"`
}
type PostBodyLabelPrompterContent struct {
	MessageId string              `json:"message_id"`
	Text      string              `json:"text"`
	Labels    map[OALabel]float32 `json:"labels"`
}
type PostBodyLabelPrompter struct {
	Id         string                       `json:"id"`
	Lang       string                       `json:"lang"`
	UpdateType string                       `json:"update_type"`
	Content    PostBodyLabelPrompterContent `json:"content"`
}

type PostBodyAssistantContent struct {
	Text string `json:"text"`
}
type PostBodyAssistant struct {
	Id         string                   `json:"id"`
	Lang       string                   `json:"lang"`
	UpdateType string                   `json:"update_type"`
	Content    PostBodyAssistantContent `json:"content"`
}
type PostBodyPrompterContent struct {
	Text string `json:"text"`
}
type PostBodyPrompter struct {
	Id         string                  `json:"id"`
	Lang       string                  `json:"lang"`
	UpdateType string                  `json:"update_type"`
	Content    PostBodyPrompterContent `json:"content"`
}

type PostBodyRankAssistantContent struct {
	NotRankable bool  `json:"not_rankable"`
	Ranking     []int `json:"ranking"`
}
type PostBodyRankAssistant struct {
	Id         string                       `json:"id"`
	Lang       string                       `json:"lang"`
	UpdateType string                       `json:"update_type"`
	Content    PostBodyRankAssistantContent `json:"content"`
}

type PostBodyLabelInitialPromptContent struct {
	Labels    map[OALabel]float32 `json:"labels"`
	MessageId string              `json:"message_id"`
	Text      string              `json:"text"`
}
type PostBodyLabelInitialPrompt struct {
	Id         string                            `json:"id"`
	Lang       string                            `json:"lang"`
	UpdateType string                            `json:"update_type"`
	Content    PostBodyLabelInitialPromptContent `json:"content"`
}
