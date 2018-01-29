package main

import "fmt"
import "akhenakh/statgo"

type metrics struct{
	Cpu *statgo.CPUStats
	Memory *statgo.MemStats
	Net []*statgo.NetIOStats
	Processes *statgo.ProcessStats
}

func NewMetrics (s *statgo.Stat) *metrics{
	m := new(metrics);
	m.Cpu = s.CPUStats()
	m.Memory = s.MemStats()
	m.Net = s.NetIOStats()
	m.Processes = s.ProcessStats()

	return m
}

func main() {
    m := NewMetrics(statgo.NewStat())
    fmt.Println(m)
}

