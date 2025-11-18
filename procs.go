package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"procfs/procfs"
	"strconv"
	"strings"
	"time"
)

func getMetrics() string {
	procs, err := os.ReadDir("/proc")
	if err != nil {
		log.Fatalf("Failed to read path: %v", err)
	}
	var builder strings.Builder

	for _, proc := range procs {
		pid := proc.Name()
		if _, err := strconv.Atoi(pid); err != nil {
			continue
		}

		proc, err := procfs.ParseProc("/proc", pid, "process")
		if err != nil {
			log.Printf("Failed to read stat: %s", err)
			continue
		}

		tids, _ := os.ReadDir(fmt.Sprintf("/proc/%s/task"))
		for _, tid := range tids {
			stat, err := procfs.ParseStat(filepath.Join("/proc", pid, "task", tid.Name(), "stat"), "thread")
			if err != nil {
				log.Printf("Failed to read stat: %s", err)
				continue
			}
			builder.WriteString(stat + "\n")
		}
	}

	return builder.String()
}

func main() {
	start := time.Now()

	procInfo := getMetrics()
	fmt.Println(procInfo)

	elapsed := time.Since(start)
	fmt.Printf("total time: %s\n", elapsed)
}
