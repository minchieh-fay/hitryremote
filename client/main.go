package main

import (
	"embed"
	"hitryremote-client/internal/config"
	"hitryremote-client/internal/logger"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 创建应用实例
	app := NewApp()

	// 初始化配置
	cfg := config.New()
	if err := cfg.Load(); err != nil {
		log.Printf("加载配置失败: %v", err)
	}

	// 初始化日志
	logger := logger.New()
	logger.Info("HiTryRemote 客户端启动")

	// 创建应用
	err := wails.Run(&options.App{
		Title:  "HiTryRemote 客户端",
		Width:  1200,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour:   &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:          app.startup,
		OnDomReady:         app.domReady,
		OnBeforeClose:      app.beforeClose,
		OnShutdown:         app.shutdown,
		Logger:             nil,
		LogLevel:           wails.INFO,
		LogLevelProduction: wails.ERROR,
		StartHidden:        false,
		HideWindowOnClose:  false,
		RGBA:               &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Windows: &options.Windows{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			WebviewUserDataPath:  "",
			ZoomFactor:           1.0,
		},
		Mac: &options.Mac{
			TitleBar: &options.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           options.DefaultMacAppearance,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &options.AboutInfo{
				Title:   "HiTryRemote 客户端",
				Message: "基于 QUIC 协议的高性能代理客户端",
				Icon:    nil,
			},
		},
		Linux: &options.Linux{
			Icon:                nil,
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    options.WebviewGpuPolicyAlways,
			ProgramName:         "hitryremote-client",
		},
		Debug: options.OpenInspectorOnStartup,
	})

	if err != nil {
		logger.Error("应用启动失败: " + err.Error())
		os.Exit(1)
	}
}
