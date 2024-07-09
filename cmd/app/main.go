package main

import "github.com/dnevsky/veryGoodProject/internal/app"

const configsDir = "configs"

// @title TestApp API
// @version 1.0
// @description API сервер для тестово задания

// @host localhost:8000
// @BasePath /api/v1
func main() {
	app.Run(configsDir)
}
