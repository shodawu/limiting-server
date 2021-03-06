package server

import (
	"fmt"
	"net/http"
	"time"
)

type ExeCmd interface {
	Exec()
}

type HitCount struct {
	ResetInterval time.Duration
	HitLimit      uint
	HitCurr       uint
	ResetCmd      ExeCmd
}

func (hc *HitCount) Exec() {
	hc.HitCurr = 0
	fmt.Println("reset request counts")
}

type Asking struct {
	Writer       http.ResponseWriter
	AskCmd       ExeCmd
	CurrentCount CurrentCount
	Resp         chan bool
}

func (a *Asking) Exec() {
	res := a.CurrentCount.Add()

	fmt.Fprint(a.Writer, res)
	a.Resp <- true
}

type CurrentCount interface {
	Add() string
}

type Limiter struct {
	CmdRev       chan ExeCmd
	HitCount     HitCount
	CurrentCount CurrentCount
}

func (l *Limiter) Add() string {
	if l.HitCount.HitCurr < l.HitCount.HitLimit {
		l.HitCount.HitCurr = l.HitCount.HitCurr + 1
		return fmt.Sprintf("%v", l.HitCount.HitCurr)
	}
	return "Error"
}

func (l *Limiter) HitHandler(w http.ResponseWriter, r *http.Request) {

	ask := &Asking{
		Writer:       w,
		Resp:         make(chan bool),
		CurrentCount: l,
	}

	l.CmdRev <- ask

	for {
		<-ask.Resp
		break
	}

}

func (l *Limiter) ResetPolling() {
	for {
		l.CmdRev <- &l.HitCount
		time.Sleep(time.Second * l.HitCount.ResetInterval)
	}
}

func (l *Limiter) CommandPolling() {
	for {
		cmd := <-l.CmdRev
		cmd.Exec()
	}
}
