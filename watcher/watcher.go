package watcher

import (
	"os/exec"

	"github.com/acround/cronCommand/config"
	"github.com/robfig/cron"
)

type Watcher struct {
	cron *cron.Cron
}

func (w Watcher) Run() {
	w.cron.Run()
}
func New(cfg *config.Config) (watcher *Watcher, err error) {
	watcher = new(Watcher)
	watcher.cron = cron.New()
	err = watcher.cron.AddFunc(cfg.Common.Schedule, func() {
		cmd := exec.Command(cfg.Common.Command, cfg.Common.Args...)
		_, err := cmd.Output()
		if err != nil {
			watcher = nil
		}
	})
	if err != nil {
		watcher = nil
	}
	return watcher, err
}
