package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type TelegramAPI struct {
	config config
}

type config struct {
	APIUrl     string `envconfig:"API_URL"`
	WebhookUrl string `envconfig:"WEBHOOK_URL"`
}

func NewTelegramAPI() *TelegramAPI {
	config := config{}
	if err := envconfig.Process("TELEGRAM", config); err != nil {
		panic(err)
	}
	return &TelegramAPI{
		config: config,
	}
}

func (r *TelegramAPI) SetWebhookUrl() error {
	if r.config.WebhookUrl == "" {
		return nil
	}

	url := fmt.Sprintf("%s/setWebhook", r.SetWebhookUrl)
	resp, err := resty.New().R().SetQueryParams(map[string]string{
		"url": r.config.WebhookUrl,
	}).Get(url)
	if err != nil {
		return err
	}

	log.Info().Interface("response", string(resp.Body())).Msg("set webhook url success")
	return nil
}

func (r *TelegramAPI) getSendMessageUrl() string {
	return fmt.Sprintf("%s/sendMessage", r.config.APIUrl)
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
	return r.SendMessageContext(context.TODO(), request)
}

func logResponse(bytes []byte) {
	log.Info().Str("body", string(bytes)).Msg("response body")
}
