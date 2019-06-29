package app

import (
	"context"
	"log"

	pb "chatclient/chatpb"

	"google.golang.org/grpc"
)

// App type
type App struct {
	Client pb.ChatServerClient
}

// NewApp constructs a new server app
func NewApp(host string) *App {
	app := &App{}
	app.initChatClient(host)
	return app
}

func (app *App) initChatClient(host string) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	app.Client = pb.NewChatServerClient(conn)
}

// SendMessage sends message
func (app *App) SendMessage(sender, text string) {
	resp, err := app.Client.SendMessage(
		context.Background(),
		&pb.SendMessageRequest{Message: &pb.Message{Sender: sender, Text: text}})
	if err != nil {
		log.Fatalf("failed to get response: %v", err)
	}
	for i, message := range resp.GetMessages() {
		log.Printf("%v: %v", i, message)
	}
}
