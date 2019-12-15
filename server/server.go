package server

import (
	"net/http"
	"strconv"

	"github.com/donbattery/snake-hub-backend/logger"
	"github.com/spf13/viper"
)

var (
	log = logger.New("HTTP Server")
)

// Start the server
func Start() error {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.WithField("Request from", req.RemoteAddr).Info("200")
		res.WriteHeader(200)
		res.Write([]byte("Hasta la vista Kigyo\n"))
	})
	return http.ListenAndServe(":"+strconv.Itoa(viper.GetInt("port")), nil)
}
