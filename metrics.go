package main

import(
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

func main() {
    fmt.Println("hello")
    stat, _ := linuxproc.ReadStat("/proc/stat")
    fmt.Println(stat)
}