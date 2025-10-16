package main

import (
	"context"
	"hitryremote-client/internal/config"
	"hitryremote-client/internal/logger"
	"hitryremote-client/internal/quic"
)

// App struct
type App struct {
	ctx    context.Context
	quic   *quic.Client
	config *config.Config
	logger *logger.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		config: config.New(),
		logger: logger.New(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.logger.Info("应用启动")
}

// domReady is called after the frontend dom has been loaded
func (a *App) domReady(ctx context.Context) {
	a.logger.Info("前端 DOM 已加载")
}

// beforeClose is called when the app is about to quit,
// either by clicking the window's close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	a.logger.Info("应用即将关闭")
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	a.logger.Info("应用已关闭")
	if a.quic != nil {
		a.quic.Close()
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return "Hello " + name + ", It's show time!"
}

// GetConfig 获取配置
func (a *App) GetConfig() *config.Config {
	return a.config
}

// SetConfig 设置配置
func (a *App) SetConfig(cfg *config.Config) error {
	a.config = cfg
	return a.config.Save()
}

// ConnectToServer 连接到服务器
func (a *App) ConnectToServer(serverAddr string) error {
	if a.quic != nil {
		a.quic.Close()
	}

	a.quic = quic.NewClient(serverAddr, a.config)
	return a.quic.Connect()
}

// DisconnectFromServer 断开服务器连接
func (a *App) DisconnectFromServer() error {
	if a.quic != nil {
		return a.quic.Close()
	}
	return nil
}

// GetConnectionStatus 获取连接状态
func (a *App) GetConnectionStatus() string {
	if a.quic == nil {
		return "disconnected"
	}
	return a.quic.GetStatus()
}
