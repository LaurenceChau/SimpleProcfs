package procfs

import (
	"errors"
)

type Proc struct {
	Ptype string
	Stat  ProcStat
}

func ParseProc(procfs_path string, pid string, ptype string) (Proc, error) {
	stat, errStat := ParseStat(procfs_path + "/" + pid + "/stat")

	return Proc{
		Ptype: ptype,
		Stat:  stat,
	}, errors.Join(errStat)
}
