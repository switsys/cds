package cdsclient

import (
	"context"
	"log"
	"time"

	"github.com/ovh/cds/sdk"
)

func (c *client) EventsListen(ctx context.Context, chanSSEvt chan<- SSEvent) {
	for ctx.Err() == nil {
		if err := c.RequestSSEGet(ctx, "/events", chanSSEvt); err != nil {
			log.Println("EventsListen", err)
		}
		time.Sleep(1 * time.Second)
	}
}
func (c *client) WebsocketEventsListen(ctx context.Context, chanMsgToSend <-chan sdk.WebsocketFilter, chanMsgReceived chan<- sdk.WebsocketEvent) {
	for ctx.Err() == nil {
		if err := c.RequestWebsocket(ctx, "/ws", chanMsgToSend, chanMsgReceived); err != nil {
			log.Printf("websocket error: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}
}
