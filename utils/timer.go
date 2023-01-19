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
		pomidorTimeSeconds: 24,
		secondTimer:        59,
		countMinute:        0,
		mode:               0,
		run:                false,
	}
}

func (t *Timer) Clean() {
	t.pomidorTimeSeconds = 24
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
				t.pomidorTimeSeconds = 24
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
	switch t.mode {
	case 0:
		t.pomidorTimeSeconds = 24
		break
	case 1:
		t.pomidorTimeSeconds = 4
		break
	case 2:
		t.pomidorTimeSeconds = 14
		break
	}
	t.secondTimer = 60
	t.countMinute = 0
}

func (t *Timer) startTimer() {
	time.Sleep(time.Second)
	switch t.mode {
	case POMIDOR:
		t.nextTick()
		return
	case SHORT_BREAK:
		t.nextTick()
		return
	case LONG_BREAK:
		t.nextTick()
		return
	}
}

func (t *Timer) nextTick() {
	if t.run {
		if t.pomidorTimeSeconds <= 0 && t.countMinute >= 59 {
			t.nextTimer()
			t.run = false
			return
		}
		t.countMinute += 1
		t.secondTimer -= 1
		flag := ""
		flagSecond := ""
		if t.pomidorTimeSeconds < 10 {
			flag = "0"
		}
		if t.secondTimer < 10 {
			flagSecond = "0"
		}
		if t.countMinute >= 60 {
			t.secondTimer = 59
			t.countMinute = 0
			flagSecond = ""
			t.pomidorTimeSeconds -= 1
		}
		t.TimerLable.SetText(flag + strconv.Itoa(t.pomidorTimeSeconds) + ":" + flagSecond + strconv.Itoa(t.secondTimer))
		t.TimerLable.Show()
	}
}
