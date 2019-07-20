package app

import (
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
)

const (
	APP_ID = "minercave"
	APP_NAME = "MinerCave"
	APP_DESC = "MinerCave CPU miner"
	APP_VERSION = "1.0.0"
	APP_COPYRIGHT = "Copyright (C) 2019 Ekene Izukanne"
	APP_VER_MAJOR = 1
	APP_VER_MINOR = 0
	APP_VER_PATCH = 0
)


func init() {
	c := color.New(color.FgCyan, color.Bold)
	c.Println(` 
	 _      _   _   _   _   _ _ _ _   _ _ _ _   _ _ _ _       _   _        _  _ _ _ _
	| \    / | | | | \ | | |  _ _ _| | _ _ _ | |  _ _ _|     / \  \ \    / / |  _ _ _|
	|  \  /  | | | |  \| | | |_ _ _  | |_ _ _/ | |          /_ _\  \ \  / /  | |_ _ _
	| |\  /| | | | | |\  | |  _ _ _| |  _ _ \  | |         / _ _ \  \ \/ /   |  _ _ _|
	| | \/ | | | | | | \ | | |_ _ _  | |   | | | |_ _ _   / /   \ \  \  /    | |_ _ _
	|_|    |_| |_| |_|  \| |_ _ _ _| |_|   |_| |_ _ _ _| /_/     \_\  \/     |_ _ _ _|
	`)
}


func Exec() {
	printStartup()
}


func printStartup() {
	printVersionInfo()
	printMemoryInfo()
	printCPUInfo()
}


func printVersionInfo() {
	c := color.New(color.FgWhite, color.Bold)
	c.Printf("	SOFTWARE		")

	e := color.New(color.FgMagenta, color.Bold)
	e.Printf("%s v%s\n", APP_NAME, APP_VERSION)
}

func printMemoryInfo() {
	memory, _ := mem.VirtualMemory()
	c := color.New(color.FgWhite, color.Bold)
	c.Printf("	MEMORY			")

	e := color.New(color.FgMagenta, color.Bold)
	e.Printf("Total: %vGB, Free: %v, UsedPercent: %v%%\n", memory.Total/1000000000, memory.Free, uint64(memory.UsedPercent))
}


func printCPUInfo(){
	cpu, _ := cpu.Info()
	c := color.New(color.FgWhite, color.Bold)
	c.Printf("	CPU			")

	e := color.New(color.FgMagenta, color.Bold)
	e.Printf("%s (%v)\n", cpu[0].ModelName, cpu[0].Cores)
}

func printMinerInfo() {
	
}