package app

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/spatocode/minercave/base"
	"github.com/spatocode/minercave/net"
	"github.com/spatocode/minercave/utils"
)

const (
	APP_ID        = "minercave"
	APP_NAME      = "MinerCave"
	APP_DESC      = "MinerCave CPU miner"
	APP_VERSION   = "1.0.0"
	APP_COPYRIGHT = "Copyright (C) 2019 Ekene Izukanne"
	APP_VER_MAJOR = 1
	APP_VER_MINOR = 0
	APP_VER_PATCH = 0
)

var config net.Config

func init() {
	title := color.New(color.FgCyan, color.Bold)
	title.Println(` 
	 _      _   _   _   _   _ _ _ _   _ _ _ _   _ _ _ _       _   _        _  _ _ _ _
	| \    / | | | | \ | | |  _ _ _| | _ _ _ | |  _ _ _|     / \  \ \    / / |  _ _ _|
	|  \  /  | | | |  \| | | |_ _ _  | |_ _ _/ | |          /_ _\  \ \  / /  | |_ _ _
	| |\  /| | | | | |\  | |  _ _ _| |  _ _ \  | |         / _ _ \  \ \/ /   |  _ _ _|
	| | \/ | | | | | | \ | | |_ _ _  | |   | | | |_ _ _   / /   \ \  \  /    | |_ _ _
	|_|    |_| |_| |_|  \| |_ _ _ _| |_|   |_| |_ _ _ _| /_/     \_\  \/     |_ _ _ _|
	`)
}

func Exec(config *net.Config) {
	printVersionInfo()
	printMemoryInfo()
	printCPUInfo()
	printMinerInfo(config)

	base.Connect(config)
}

func Configure(config *net.Config) {
	configfile := "config.json"
	configfile, _ = filepath.Abs(configfile)

	file, err := os.Open(configfile)
	if err != nil {
		utils.LOG_ERR("File error: %s\n", err.Error())
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&config); err != nil {
		utils.LOG_ERR("Configuration error: %s\n", err.Error())
	}
}

func printVersionInfo() {
	title := color.New(color.FgWhite, color.Bold)
	title.Printf("	SOFTWARE		")

	value := color.New(color.FgMagenta, color.Bold)
	value.Printf("%s v%s\n", APP_NAME, APP_VERSION)
}

func printMemoryInfo() {
	memory, _ := mem.VirtualMemory()
	title := color.New(color.FgWhite, color.Bold)
	title.Printf("	MEMORY			")

	value := color.New(color.FgMagenta, color.Bold)
	value.Printf("Total: %vMB, Free: %v, UsedPercent: %v%%\n", memory.Total/1000000, memory.Free, uint64(memory.UsedPercent))
}

func printCPUInfo() {
	cpu, _ := cpu.Info()
	title := color.New(color.FgWhite, color.Bold)
	title.Printf("	CPU			")

	value := color.New(color.FgMagenta, color.Bold)
	value.Printf("%s (%v)\n", cpu[0].ModelName, cpu[0].Cores)
}

func printMinerInfo(config *net.Config) {
	color.New(color.FgWhite, color.Bold).Printf("	CRYPTOCURRENCY		")
	color.New(color.FgMagenta, color.Bold).Printf("%s\n", config.Currency)

	color.New(color.FgWhite, color.Bold).Printf("	THREADS			")
	color.New(color.FgMagenta, color.Bold).Printf("%v\n", config.Threads)

	color.New(color.FgWhite, color.Bold).Printf("	POOL 			")
	color.New(color.FgMagenta, color.Bold).Printf("%s\n", config.Pool.Url)
}
