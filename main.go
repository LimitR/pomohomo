package main

import (
	"log"
	"os"
	"pomidoro/utils"
	"strings"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)
	b, err := gtk.BuilderNew()
	if err != nil {
		panic(err)
	}
	err = b.AddFromFile("ui.glade")
	if err != nil {
		panic(err)
	}
	obj, err := b.GetObject("main")
	if err != nil {
		panic(err)
	}
	win := obj.(*gtk.Window)

	obj, err = b.GetObject("tasks")
	if err != nil {
		panic(err)
	}
	tasks := obj.(*gtk.Box)

	obj, err = b.GetObject("inp")
	if err != nil {
		panic(err)
	}
	inp := obj.(*gtk.Entry)

	obj, err = b.GetObject("start")
	if err != nil {
		panic(err)
	}
	start := obj.(*gtk.Button)

	obj, err = b.GetObject("pause")
	if err != nil {
		panic(err)
	}
	pause := obj.(*gtk.Button)

	obj, err = b.GetObject("reset")
	if err != nil {
		panic(err)
	}
	reset := obj.(*gtk.Button)

	obj, err = b.GetObject("time")
	if err != nil {
		panic(err)
	}
	times := obj.(*gtk.Label)
	timer := utils.NewTimer(times)
	ch := make(chan int)
	go timer.Start(ch)
	start.Connect("button-press-event", func() {
		start.SetSensitive(false)
		pause.SetSensitive(true)
		reset.SetSensitive(false)
		go func() {
			ch <- 0
		}()
		win.ShowAll()
	})

	pause.Connect("button-press-event", func() {
		start.SetSensitive(true)
		pause.SetSensitive(false)
		reset.SetSensitive(true)
		time.Sleep(time.Millisecond)
		go func() {
			ch <- 1
		}()
		win.ShowAll()
	})

	reset.Connect("button-press-event", func() {
		start.SetSensitive(true)
		pause.SetSensitive(false)
		time.Sleep(time.Millisecond)
		go func() {
			ch <- 2
		}()
		win.ShowAll()
	})

	inp.Connect("activate", func() {
		s, _ := inp.GetText()
		flag := true
		inp.DeleteText(0, len(s))
		if len(strings.ReplaceAll(s, " ", "")) != 0 {
			l, _ := gtk.LabelNew(s)
			ch, _ := gtk.CheckButtonNew()
			ch.Connect("toggled", func() {
				if flag {
					t, _ := l.GetText()
					l.SetMarkup("<s>" + t + "</s>")
					win.ShowAll()
					flag = false
				} else {
					t, _ := l.GetText()
					l.SetText(t)
					win.ShowAll()
					flag = true
				}
			})

			box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
			l.SetMarginStart(10)
			box.Add(ch)
			box.Add(l)
			tasks.Add(box)
		}
		win.ShowAll()
	})

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.ShowAll()

	gtk.Main()
}

func soundInit() {
	f, err := os.Open("./audio/Wir_sind_Helden-Von_hier_an_blind.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
}
