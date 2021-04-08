package main

import (
	service "moneysaverapi/main/service"
)

func main() {
	service.StartAPIServer(service.ApiService)
}
