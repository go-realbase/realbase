package server

import (
	"net/http"

	"github.com/neutrinoapp/neutrino/src/common/client"
	"github.com/neutrinoapp/neutrino/src/common/config"
)

func init() {

}

func Initialize() (*http.Server, error) {
	redisClient := client.GetNewRedisClient()
	clientMessageProcessor := NewClientMessageProcessor()
	natsClient := client.NewNatsClient(config.Get(config.KEY_QUEUE_ADDR))

	_, server, wsClient, interceptor, err := NewWebSocketServer()
	if err != nil {
		return nil, err
	}

	natsProcessor := NatsMessageProcessor{natsClient, wsClient}
	natsProcessor.Process()

	rpcProcessor := RpcMessageProcessor{wsClient}
	rpcProcessor.Process()

	wsProcessor := WsMessageProcessor{interceptor, redisClient, clientMessageProcessor, wsClient}
	wsProcessor.Process()

	return server, nil
}
