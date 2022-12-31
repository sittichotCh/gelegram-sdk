package bot

type HandlerFn func(*WebHookRequest)

var handlerList = map[string]HandlerFn{}

func AddHandler(command string, fn HandlerFn) {
	handlerList[command] = fn
}
