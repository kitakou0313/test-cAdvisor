package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func writeToFile(data string) {
	file, err := os.OpenFile("/var/log/myapp/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ファイルを開く際にエラーが発生しました:", err)
		return
	}
	defer file.Close()

	// ファイルにデータを書き込む
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("ファイルに書き込む際にエラーが発生しました:", err)
		return
	}

	fmt.Println("ファイルにデータを追記しました")
}

func readMeminfo() string {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println("ファイルの読み込みエラー:", err)
		return ""
	}

	// 読み込んだデータを文字列に変換
	return string(data)
}

func readMemoryStat() string {
	data, err := os.ReadFile("/sys/fs/cgroup/memory/memory.stat")
	if err != nil {
		fmt.Println("ファイルの読み込みエラー:", err)
		return ""
	}

	// 読み込んだデータを文字列に変換
	return string(data)
}

func main() {
	mallocInterval, err := strconv.Atoi(os.Getenv("MALLOC_INTERVAL_MILLISECOND"))
	mallocSize, err := strconv.Atoi(os.Getenv("MALLOC_INTERVAL_MIBIBYTE"))

	if err != nil {
		panic(err)
	}

	timerForMalloc := time.NewTicker(time.Millisecond * time.Duration(mallocInterval))
	memory := make(map[int][]byte)

	i := 0

	for range timerForMalloc.C {
		memory[i] = make([]byte, mallocSize*1024)
		i += 1
		writeToFile(readMeminfo())
		writeToFile(readMemoryStat())
	}

}
