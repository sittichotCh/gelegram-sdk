package bot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func RegisterServer(port int) error {
	server := gin.Default()
	server.POST("", func(context *gin.Context) {
		body := WebHookRequest{}
		if err := context.BindJSON(&body); err != nil {
			log.Error().Err(err).Msg("cannot be bind request body")
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		go handler(&body)
		context.JSON(http.StatusOK, nil)
	})

	p := fmt.Sprintf(":%d", port)
	return server.Run(p)
}

func handler(request *WebHookRequest) {
	command := request.GetCommand()
	fn, found := handlerList[command]
	if !found {
		fmt.Printf("command %s is not found\n", command)
		return
	}
	fn(request)
}
