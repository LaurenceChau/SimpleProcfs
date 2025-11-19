package main

import (
	"fmt"
	"log"
	"os"
	"procs/procfs"
	"strings"
	"time"
)

func getMetrics(procfs_path string) string {
	procs, err := os.ReadDir("/proc")
	if err != nil {
		log.Fatalf("Failed to read path: %v", err)
	}
	var sb strings.Builder

	for _, p := range procs {
		if !procfs.IsDigitsOnly(p.Name()) {
			continue
		}

		proc, err := procfs.ParseProc(procfs_path, p.Name(), "process")
		if err != nil {
			log.Printf("Failed to read stat: %s", err)
			continue
		}
		fmt.Fprintf(&sb, "%s", proc.Ptype)

		task_path := procfs_path + "/" + p.Name() + "/task"
		tasks, _ := os.ReadDir(task_path)
		for _, t := range tasks {
			if !procfs.IsDigitsOnly(t.Name()) {
				continue
			}
			proc, err := procfs.ParseProc(task_path, t.Name(), "thread")
			if err != nil {
				log.Printf("Failed to read stat: %s", err)
				continue
			}
			fmt.Fprintf(&sb, "%s", proc.Ptype)
		}
	}

	return sb.String()
}

func main() {
	start := time.Now()

	procInfo := getMetrics("/proc")
	fmt.Println(procInfo)

	elapsed := time.Since(start)
	fmt.Printf("total time: %s\n", elapsed)
}
