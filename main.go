package main

import (
	"fmt"
	"net/http"
	"simple-webserver/api"
)

func main() {
	// Устанавливаем порт HTTP Web Server, на котором он будет работать
	port := ":8080"
	// Логируем его запуск
	fmt.Printf("HTTP Web Server successfully started & running on port >> %s \n", port)
	// Вызываем создание сервера
	srv := api.NewServer()
	// Сервим наш сервак на порту
	http.ListenAndServe(port, srv)
}
