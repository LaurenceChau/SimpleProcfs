package procfs

import (
	"bytes"
	"os"
)

type ProcStatCode int

const (
	PID ProcStatCode = iota
	Comm
	State
	PPID
	PGRP
	Session
	TTY
	TPGID
	Flags
	MinFlt
	CMinFlt
	MajFlt
	UTime
	STime
	CUTime
	CSTime
	Priority
	Nice
	NumThreads
	Itrealvalue
	Starttime
	VSize
	RSS
	RSSLimit
	StartCode
	EndCode
	StartStack
	Kstkesp
	Kstkeip
	Signal
	Blocked
	SigIgnore
	SigCatch
	WChan
	NSwap
	CNSwap
	ExitSignal
	Processor
	RTPriority
	Policy
	DelayacctBlkioTicks
	GuestTime
	CGuestTime
	StartData
	EndData
	StartBrk
	ArgStart
	ArgEnd
	EnvStart
	EnvEnd
	ExitCode
)

const userHZ = 100

type ProcStat struct {
	PID                 int
	PType               string
	Comm                string
	State               string
	MinFlt              uint
	CMinFlt             uint
	MajFlt              uint
	CMajFlt             uint
	UTime               uint
	STime               uint
	CUTime              int
	CSTime              int
	NumThreads          int
	Starttime           uint64
	VSize               uint
	Processor           uint
	DelayAcctBlkIOTicks uint64
	RSS                 int
}

func ParseStat(statfile string, ptype string) (ProcStat, error) {
	data, err := ReadFileNoStat(statfile)
	if err != nil {
		return ProcStat{}, err
	}

	stats := bytes.Fields(data)

	return ProcStat{
		PID:   toInt[int](stats[PID]),
		PType: ptype,
		Comm:  string(stats[Comm]),
		State: string(stats[State]),
	}, nil
}

func (s ProcStat) VirtualMemory() uint {
	return s.VSize
}

func (s ProcStat) ResidentMemory() int {
	return s.RSS * os.Getpagesize()
}

func (s ProcStat) StartTime() (float64, error) {
	return (float64(s.Starttime) / userHZ), nil
}

func (s ProcStat) CPUUTime() float64 {
	return float64(s.UTime) / userHZ
}

func (s ProcStat) CPUSTime() float64 {
	return float64(s.STime) / userHZ
}
