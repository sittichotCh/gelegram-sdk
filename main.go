package main

//const url = "https://api.telegram.org/bot5550135438:AAGkLM-WhXP14foQAQdZLFF9dBUD6NsSwRo"
//
//func main() {
//	bot.AddHandler("test", testHandler)
//
//	if err := bot.RegisterServer(8080); err != nil {
//		panic(err)
//	}
//}
//
//func testHandler(input *bot.WebHookRequest) {
//	api := bot.NewTelegramAPI(url)
//	request := bot.NewSendMessageRequest(input.Message.Chat.Id, "please, select ledger type")
//	resp, err := api.SendMessage(request)
//	if err != nil {
//		log.Error().Err(err).Msg("callback failed")
//		return
//	}
//
//	log.Info().Interface("response", resp).Msg("callback success")
//}
