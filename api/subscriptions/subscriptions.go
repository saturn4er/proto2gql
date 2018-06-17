package subscriptions

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type SubscriptionsManager struct {
	upgrader websocket.Upgrader
}

func (s *SubscriptionsManager) Run() error {

}

func (s *SubscriptionsManager) Close() error {

}

func (s *SubscriptionsManager) Handler(w http.ResponseWriter, r *http.Request) {
	// Establish a WebSocket connection
	var ws, err = s.upgrader.Upgrade(w, r, nil)

	// Bail out if the WebSocket connection could not be established
	if err != nil {
		return
	}

	// Close the connection early if it doesn't implement the graphql-ws protocol
	if ws.Subprotocol() != "graphql-ws" {
		ws.Close()
		return
	}

	// Establish a GraphQL WebSocket connection
	conn := NewConnection(ws, ConnectionConfig{
		EventHandlers: ConnectionEventHandlers{
			Close: func(conn Connection) {
				subscriptionManager.RemoveSubscriptions(conn)

				delete(connections, conn)
			},
			StartOperation: func(
				conn Connection,
				opID string,
				data *StartMessagePayload,
			) []error {
				return subscriptionManager.AddSubscription(conn, &Subscription{
					ID:            opID,
					Query:         data.Query,
					Variables:     data.Variables,
					OperationName: data.OperationName,
					Connection:    conn,
					SendData: func(data *DataMessagePayload) {
						conn.SendData(opID, data)
					},
				})
			},
			StopOperation: func(conn Connection, opID string) {
				subscriptionManager.RemoveSubscription(conn, &Subscription{
					ID: opID,
				})
			},
		},
	})
	connections[conn] = true
}

func New() *SubscriptionsManager {
	var upgrader = websocket.Upgrader{
		CheckOrigin:  func(r *http.Request) bool { return true },
		Subprotocols: []string{"graphql-ws"},
	}

	return &SubscriptionsManager{
		upgrader: upgrader,
	}
}
