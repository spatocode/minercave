package base

import (
	"log"
	"runtime"

	"github.com/spatocode/minercave/algorithm/sha256d"
	"github.com/spatocode/minercave/net"
)

func Connect(cfg *net.Config) {
	if cfg.Threads > 1 {
		runtime.GOMAXPROCS(cfg.Threads)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	hashrateChan := make(chan *sha256d.HashRate, cfg.Threads)

	miner := sha256d.NewMiner(cfg)
	miner.HashRate = hashrateChan
	miner.Mine()

	hashrateReport := make([]float64, cfg.Threads)
	for {
		for i := 0; i < cfg.Threads; i++ {
			report := <-hashrateChan
			hashrateReport[report.MinerID] = report.Rate
		}

		for minerID, hashrate := range hashrateReport {
			log.Printf("hashrate %d: %.1fH/s\n", minerID, hashrate)
		}
	}
}
