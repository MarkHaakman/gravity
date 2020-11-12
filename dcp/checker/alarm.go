package checker

import "github.com/siddontang/go-log/log"

type AlarmManager interface {
	Alarm(r Result)
}

type ConsoleAlarmManager struct {
}

func (ConsoleAlarmManager) Alarm(r Result) {
	log.Infof("result: %+v", r)
}

type ChanAlarmManager struct {
	Output chan Result
}

func (c ChanAlarmManager) Alarm(r Result) {
	c.Output <- r
}
