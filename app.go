package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// func (a *App) Greet(name string) {
// 	// runtime.RequestNotificationAuthorization(a.ctx)

// 	runtime.SendNotification(a.ctx, runtime.NotificationOptions{
// 		ID:    "notif-1",
// 		Title: "Hello",
// 		Body:  "Greetings human!",
// 	})
// }
