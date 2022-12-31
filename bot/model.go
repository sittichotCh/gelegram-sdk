package bot

import "fmt"

func (r *WebHookRequest) GetCommand() string {
	if r.Message != nil {
		return r.Message.Text
	}
	if r.CallbackQuery != nil {
		return r.CallbackQuery.Data
	}

	return ""
}

type WebHookRequest struct {
	UpdateId      int             `json:"update_id"`
	Message       *MessageRequest `json:"message"`
	CallbackQuery *CallbackQuery  `json:"callback_query"`
}

type MessageRequest struct {
	MessageId int              `json:"message_id"`
	From      FromRequest      `json:"from"`
	Chat      ChatRequest      `json:"chat"`
	Date      int              `json:"date"`
	Text      string           `json:"text"`
	Entities  []*EntityRequest `json:"entities"`
}

type CallbackQuery struct {
	Id           string         `json:"id"`
	From         FromRequest    `json:"from"`
	Message      MessageRequest `json:"message"`
	ChatInstance string         `json:"chat_instance"`
	Data         string         `json:"data"`
}

type CallbackQueryFrom struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}

type FromRequest struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}

type ChatRequest struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
}

type EntityRequest struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

func NewSendMessageRequest(chatId int64, text string) *SendMessageRequest {
	return &SendMessageRequest{
		ChatId:      chatId,
		Text:        text,
		ReplyMarkup: &ReplyMarkup{},
	}
}

func (r *SendMessageRequest) SetInlineKeyboard(inlineKeyboard [][]*InlineKeyboard) {
	r.ReplyMarkup.InlineKeyboard = inlineKeyboard
}

func (r *SendMessageRequest) SetReplyKeyboardMarkup(keyboardMarkup [][]*KeyboardButton) {
	r.ReplyMarkup.ReplyKeyboardMarkup = keyboardMarkup
}

type SendMessageRequest struct {
	ChatId      int64        `json:"chat_id"`
	Text        string       `json:"text"`
	ReplyMarkup *ReplyMarkup `json:"reply_markup,omitempty"`
}

type ReplyMarkup struct {
	InlineKeyboard      [][]*InlineKeyboard `json:"inline_keyboard,omitempty"`
	ReplyKeyboardMarkup [][]*KeyboardButton `json:"keyboard,omitempty"`
}

type InlineKeyboard struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type SendMessageResponse struct {
	Ok          bool           `json:"ok"`
	Result      ResultResponse `json:"result"`
	ErrorCode   int            `json:"error_code"`
	Description string         `json:"description"`
}

type ResultResponse struct {
	MessageId int         `json:"message_id"`
	From      FromRequest `json:"from"`
	Chat      ChatRequest `json:"chat"`
	Date      int         `json:"date"`
	Text      string      `json:"text"`
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *ErrorMessage) Error() string {
	return fmt.Sprintf("%d %s", r.Code, r.Message)
}
