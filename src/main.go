package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func printMeminfo() {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println("ファイルの読み込みエラー:", err)
		return
	}

	// 読み込んだデータを文字列に変換
	meminfoStr := string(data)

	// 改行で分割して各行を表示
	lines := strings.Split(meminfoStr, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}

func printMemoryStat() {
	data, err := os.ReadFile("/sys/fs/cgroup/memory/memory.stat")
	if err != nil {
		fmt.Println("ファイルの読み込みエラー:", err)
		return
	}

	// 読み込んだデータを文字列に変換
	meminfoStr := string(data)

	// 改行で分割して各行を表示
	lines := strings.Split(meminfoStr, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	mallocInterval, err := strconv.Atoi(os.Getenv("MALLOC_INTERVAL_MILLISECOND"))
	if err != nil {
		panic(err)
	}

	timerForMalloc := time.NewTicker(time.Millisecond * time.Duration(mallocInterval))
	memory := make(map[int][]byte)

	i := 0

	for range timerForMalloc.C {
		memory[i] = make([]byte, 1024*1024)
		i += 1
		printMeminfo()
		printMemoryStat()
	}

}
