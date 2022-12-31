package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

type TelegramAPI struct {
	Url string
}

func NewTelegramAPI(url string) *TelegramAPI {
	return &TelegramAPI{
		Url: url,
	}
}

func (r *TelegramAPI) getSendMessageUrl() string {
	return fmt.Sprintf("%s/sendMessage", r.Url)
}

func (r *TelegramAPI) SendMessageContext(context context.Context, request *SendMessageRequest) (*SendMessageResponse, error) {
	client := resty.New()

	resp, err := client.R().SetBody(request).Post(r.getSendMessageUrl())
	if err != nil {
		return nil, err
	}

	logResponse(resp.Body())
	messageBody := SendMessageResponse{}
	if err := json.Unmarshal(resp.Body(), &messageBody); err != nil {
		return nil, err
	}
	if !messageBody.Ok {
		return nil, &ErrorMessage{
			Code:    messageBody.ErrorCode,
			Message: messageBody.Description,
		}
	}

	return &messageBody, nil
}

func (r *TelegramAPI) SendMessage(request *SendMessageRequest) (*SendMessageResponse, error) {
	return r.SendMessageContext(context.Background(), request)
}

func logResponse(bytes []byte) {
	log.Info().Str("body", string(bytes)).Msg("response body")
}
