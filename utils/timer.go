package utils

import (
	"strconv"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type Mode int

const (
	POMIDOR     Mode = iota
	SHORT_BREAK Mode = iota
	LONG_BREAK  Mode = iota
)

type Timer struct {
	pomidorTimeSeconds int
	secondTimer        int
	countMinute        int
	mode               Mode
	TimerLable         *gtk.Label
	run                bool
}

func NewTimer(timerLable *gtk.Label) *Timer {
	return &Timer{
		TimerLable:         timerLable,
		pomidorTimeSeconds: 25 * 60 * 60,
		secondTimer:        59,
		countMinute:        0,
		mode:               0,
		run:                false,
	}
}

func (t *Timer) Clean() {
	t.pomidorTimeSeconds = 25 * 60 * 60
	t.mode = 0
	t.secondTimer = 59
	t.countMinute = 0
	t.TimerLable.Show()
}

func (t *Timer) Start(ch chan int) {
	for {
		select {
		case x := <-ch:
			if x == 0 {
				t.run = true
			} else if x == 1 {
				t.run = false
			} else if x == 2 {
				t.run = false
				t.pomidorTimeSeconds = 25 * 60 * 60
				t.mode = 0
				t.secondTimer = 59
				t.countMinute = 0
				t.TimerLable.Show()
			}
		default:
			t.startTimer()
		}
	}
}

func (t *Timer) nextTimer() {
	if t.mode+1 > 2 {
		t.mode = 0
	} else {
		t.mode += 1
	}
	t.secondTimer = 60
	t.countMinute = 0
}

func (t *Timer) startTimer() {
	switch t.mode {
	case POMIDOR:
		t.pomidorTimeSeconds = 25 * 60 * 60
		if t.run {
			if t.pomidorTimeSeconds == 0 {
				t.nextTimer()
			}

			time.Sleep(time.Second)
			t.pomidorTimeSeconds -= 1
			t.countMinute += 1
			t.secondTimer -= 1
			flag := ""
			flagSecond := ""
			minutes := int(t.pomidorTimeSeconds / 60 / 60)
			if minutes < 10 {
				flag = "0"
			}
			if t.secondTimer < 10 {
				flagSecond = "0"
			}
			if t.countMinute == 60 {
				t.secondTimer = 59
				flagSecond = ""
			}
			t.TimerLable.SetText(flag + strconv.Itoa(minutes) + ":" + flagSecond + strconv.Itoa(t.secondTimer))
			t.TimerLable.Show()
		}
		return
	case SHORT_BREAK:
		t.pomidorTimeSeconds = 5 * 60 * 60
		if t.run {
			if t.pomidorTimeSeconds == 0 {
				t.nextTimer()
			}

			time.Sleep(time.Second)
			t.pomidorTimeSeconds -= 1
			t.countMinute += 1
			t.secondTimer -= 1
			flag := ""
			flagSecond := ""
			minutes := int(t.pomidorTimeSeconds / 60 / 60)
			if minutes < 10 {
				flag = "0"
			}
			if t.secondTimer < 10 {
				flagSecond = "0"
			}
			if t.countMinute == 60 {
				t.secondTimer = 59
				flagSecond = ""
			}
			t.TimerLable.SetText(flag + strconv.Itoa(minutes) + ":" + flagSecond + strconv.Itoa(t.secondTimer))
			t.TimerLable.Show()
		}
		return
	case LONG_BREAK:
		t.pomidorTimeSeconds = 15 * 60 * 60
		if t.run {
			if t.pomidorTimeSeconds == 0 {
				t.nextTimer()
			}

			time.Sleep(time.Second)
			t.pomidorTimeSeconds -= 1
			t.countMinute += 1
			t.secondTimer -= 1
			flag := ""
			flagSecond := ""
			minutes := int(t.pomidorTimeSeconds / 60 / 60)
			if minutes < 10 {
				flag = "0"
			}
			if t.secondTimer < 10 {
				flagSecond = "0"
			}
			if t.countMinute == 60 {
				t.secondTimer = 59
				flagSecond = ""
			}
			t.TimerLable.SetText(flag + strconv.Itoa(minutes) + ":" + flagSecond + strconv.Itoa(t.secondTimer))
			t.TimerLable.Show()
		}
		return
	}
}
