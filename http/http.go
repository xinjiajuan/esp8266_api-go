package http

import (
	"DHT22-temperature_databases-go/config"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type JsonStruct struct {
	devid int
	hum   float64
	tem   float64
}

type HandlerServer struct {
	ServerDatabases config.DatabasesConfig
}

func MakeHttpServer(config config.Yaml) (*http.Server, config.TlsConfig) {
	serverInfo := HandlerServer{}
	serverInfo.ServerDatabases = config.Databases
	webserver := http.Server{
		Addr:    ":" + strconv.Itoa(config.Server.Listen),
		Handler: serverInfo,
	}
	tls := config.Server.Tls
	return &webserver, tls
}
func RunHttpAPIServer(webserver *http.Server, tlsConfig config.TlsConfig) {
	if tlsConfig.Enable {
		go webserver.ListenAndServeTLS(tlsConfig.CertFile, tlsConfig.KeyFile)
	} else {
		go webserver.ListenAndServe()
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	//程序堵塞
	case <-sigs: //检测Ctrl+c退出程序命令
		fmt.Println("instance gracefully...")
		webserver.Shutdown(context.Background()) //平滑关闭Http Server线程
		fmt.Println("exited safely!")
	}
}
func (webserver HandlerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/dht22" {
		//println(r.Header.Get("Content-Type"))
		if r.Header.Get("Content-Type") == "application/json" {
			r.Body = http.MaxBytesReader(w, r.Body, 1048576)
			var req JsonStruct
			er := json.NewDecoder(r.Body).Decode(&req)
			if er != nil {
				panic(er.Error())
			}
			println(req.devid)
			println(req.hum)
			println(req.tem)
		} else {
			fmt.Fprint(w, "url error")
		}

	}
}
