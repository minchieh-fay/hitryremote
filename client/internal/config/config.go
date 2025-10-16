package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config 应用配置
type Config struct {
	ServerAddr string `json:"server_addr"` // 服务器地址
	ClientID   string `json:"client_id"`   // 客户端ID
	Version    string `json:"version"`     // 版本号
	AutoStart  bool   `json:"auto_start"`  // 自动启动
	LogLevel   string `json:"log_level"`   // 日志级别
}

// New 创建新配置
func New() *Config {
	return &Config{
		ServerAddr: "127.0.0.1:10001",
		ClientID:   "client-" + generateID(),
		Version:    "1.0.0",
		AutoStart:  false,
		LogLevel:   "info",
	}
}

// Load 加载配置
func (c *Config) Load() error {
	configPath := c.getConfigPath()
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return c.Save() // 如果文件不存在，创建默认配置
		}
		return err
	}

	return json.Unmarshal(data, c)
}

// Save 保存配置
func (c *Config) Save() error {
	configPath := c.getConfigPath()

	// 确保目录存在
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// getConfigPath 获取配置文件路径
func (c *Config) getConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".hitryremote", "client.json")
}

// generateID 生成客户端ID
func generateID() string {
	// 简单的ID生成，实际项目中可以使用更复杂的算法
	return "12345"
}
