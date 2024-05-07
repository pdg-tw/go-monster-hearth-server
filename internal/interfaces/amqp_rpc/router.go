package amqprpc

import (
	"sync"

	"github.com/pdg-tw/go-monster-hearth-server/internal/application"
	"github.com/pdg-tw/go-monster-hearth-server/pkg/rabbitmq/rmq_rpc/server"
)

var hdlOnce sync.Once
var amqpRpcRouter map[string]server.CallHandler

// NewRouter -.
func NewRouter(t *application.TranslationUseCase) map[string]server.CallHandler {

	hdlOnce.Do(func() {
		amqpRpcRouter = make(map[string]server.CallHandler)
		{
			newTranslationRoutes(amqpRpcRouter, t)
		}
	})

	return amqpRpcRouter
}
