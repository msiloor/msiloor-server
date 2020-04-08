package main

import (
	SpringBoot "github.com/go-spring/go-spring/spring-boot"
	_ "github.com/go-spring/go-spring/starter-gin"
	_ "github.com/go-spring/go-spring/starter-web"
)

func main() {
	SpringBoot.RunApplication("config/")
}
