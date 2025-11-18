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

// func readStat(statfile string, ptype string) (string, error) {
// 	bytes, err := os.ReadFile(statfile)
// 	stat := strings.ReplaceAll(string(bytes), "\n", "")
// 	if err != nil {
// 		return "", fmt.Errorf("failed reading statfile %s", statfile)
// 	}

// 	return fmt.Sprintf("{\"type\":\"%s\",\"stat\":\"%s\"}", ptype, stat), nil
// }

func main() {
	start := time.Now()
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

		stat, err := procfs.ParseStat(filepath.Join("/proc", pid, "stat"), "process")
		if err != nil {
			log.Printf("Failed to read stat: %s", err)
			continue
		}
		builder.WriteString(stat + "\n")

		tids, _ := os.ReadDir(filepath.Join("/proc", pid, "task"))
		for _, tid := range tids {
			stat, err := readStat(filepath.Join("/proc", pid, "task", tid.Name(), "stat"), "thread")
			if err != nil {
				log.Printf("Failed to read stat: %s", err)
				continue
			}
			builder.WriteString(stat + "\n")
		}
	}

	fmt.Println(builder.String())

	elapsed := time.Since(start)
	fmt.Printf("total time: %s\n", elapsed)
}
