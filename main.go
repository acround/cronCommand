package main

import (
	"flag"
	"fmt"
	"github.com/acround/cronCommand/config"
	"github.com/acround/cronCommand/watcher"
	"log"
)

func main() {
	configPath := flag.String("c", "config.toml", "Path to `toml` configuration file")
	c, err := config.GetConfigFromFile(*configPath)
	if err != nil {
		log.Fatalf(WrapLog("couldn't load config: %s"), err)
		return
	}
	w, err := watcher.New(c)
	if err != nil {
		log.Fatalf(WrapLog("couldn't create cron: %s"), err)
		return
	}
	w.Run()
}
func WrapLog(msg string) string {
	return fmt.Sprintf("[APP] %s", msg)
}
