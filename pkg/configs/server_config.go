package configs

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/febzey/goapi/pkg/utils"
	"github.com/gorilla/mux"
)

func ServerConf(router *mux.Router) *http.Server {
	serverConnURL, _ := utils.ConnUrlBuilder("server")
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return &http.Server{
		Handler:     router,
		Addr:        serverConnURL,
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}

}
