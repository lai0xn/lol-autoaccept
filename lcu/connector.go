package lcu

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/gorilla/websocket"
)

type Connector struct {
	Port       string
	Password   string
	Authtoken  string
	HttpClient http.Client
	WsClient   websocket.Dialer
}

var ACCEPT bool

func encodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte("riot:" + password))
}

func RiotConnector() (*Connector, error) {
	port, password, err := open_lockfile()
	if err != nil {
		return nil, err
	}
	return &Connector{
		Port:     port,
		Password: password,
		HttpClient: http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		WsClient: websocket.Dialer{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Authtoken: encodePassword(password),
	}, nil
}

func (c *Connector) GetSummonerInfo() SunnmonerInfo {
	req, err := http.NewRequest("GET", BASE_URL+":"+c.Port+SUMMONER_INFO_ENDPOINT, nil)
	req.Header.Add("Authorization", "Basic "+c.Authtoken)
	if err != nil {
		return SunnmonerInfo{}
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return SunnmonerInfo{}
	}
	data, err := io.ReadAll(res.Body)
	var Summoner SunnmonerInfo
	json.Unmarshal(data, &Summoner)
	return Summoner
}

func (c *Connector) WsConnect() error {
	header := http.Header{}
	header.Add("Authorization", "Basic "+c.Authtoken)

	conn, _, err := c.WsClient.Dial(LCU_WS_URI+":"+c.Port, header)
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.WriteMessage(websocket.TextMessage, []byte(`[5, "OnJsonApiEvent"]`))

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				os.Exit(0)
			}
			if ACCEPT && strings.Contains(string(message), "Found") {
				c.HttpClient.Post(BASE_URL+":"+c.Port+ACCEPT_QUEUE_URI, "application/json", nil)
			}

		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	log.Println("Received interrupt signal, closing WebSocket connection...")
	conn.Close()
	return nil
}
