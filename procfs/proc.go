package procfs

import (
	"errors"
	"fmt"
)

type Proc struct {
	ptype string
	stat  ProcStat
}

func ParseProc(procfs_path string, pid string, ptype string) (Proc, error) {

	stat, errStat := ParseStat(fmt.Sprintf("%s/%s/stat", procfs_path, pid))

	return Proc{
		ptype: ptype,
		stat:  stat,
	}, errors.Join(errStat)
}
