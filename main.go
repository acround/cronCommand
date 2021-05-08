package main

import (
	"flag"
	"fmt"
	"github.com/acround/cronCommand/config"
	"github.com/robfig/cron"
	"log"
	"os/exec"
)

func main() {
	configPath := flag.String("c", "config.toml", "Path to `toml` configuration file")
	c, err := config.GetConfigFromFile(*configPath)
	if err != nil {
		log.Fatalf(WrapLog("couldn't load config: %s"), err)
		return
	}
	r := cron.New()
	if c.Common.WithSeconds {
		r = cron.New(cron.WithSeconds())
	}
	_, err = r.AddFunc(c.Common.Schedule, func() {
		cmd := exec.Command(c.Common.Command, c.Common.Arg1, c.Common.Arg2, c.Common.Arg3)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
	})
	if err != nil {
		log.Fatalf(WrapLog("couldn't add the func: %s"), err)
		return
	}
	r.Run()
}
func WrapLog(msg string) string {
	return fmt.Sprintf("[APP] %s", msg)
}
