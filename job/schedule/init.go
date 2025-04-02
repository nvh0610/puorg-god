package schedule

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

var cronTabb = cron.New()

// "0 */5 * * * *"
func RegisterScheduler(exp string, cmd func()) (cron.EntryID, error) {
	return cronTabb.AddFunc(exp, cmd)
}

func NewScheduleTimeout(d time.Duration, cmd func()) *time.Timer {
	return time.AfterFunc(d, cmd)
}

func NewScheduleInterval(d time.Duration, cmd func()) (cron.EntryID, error) {
	return cronTabb.AddFunc(fmt.Sprintf("@every %s", d.String()), cmd)
}

func UnRegisterScheduler(entryId int) {
	cronTabb.Remove(cron.EntryID(entryId))
}

func Start() {
	cronTabb.Start()
}

func Stop() {
	cronTabb.Stop()
}
