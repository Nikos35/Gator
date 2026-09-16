package main

import (
	"fmt"

	"github.com/Nikos35/Gator/internal/config"
)

func main() {
	cfg, _ := config.Read()

	cfg.SetUser("nko")
	cfg, _ = config.Read()

	fmt.Println(cfg.DataBaseURL)
	fmt.Println(cfg.CurrentUserName)
}
