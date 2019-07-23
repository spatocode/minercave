package base

import (
	"github.com/spatocode/minercave/algorithm/sha256d"
	"github.com/spatocode/minercave/net"
)

var numOfDevice int

func Connect(cfg *net.Config) {
	if cfg.Threads > 1 {
		numOfDevice = cfg.Threads
	} else {
		numOfDevice = 1
	}

	hashrateChan := make(chan *sha256d.HashRate, numOfDevice)
	stratum := net.StratumClient(cfg)
	stratum.Connect()
	for i := 0; i < numOfDevice; i++ {
		miner := &sha256d.Miner{
			Devices:  numOfDevice,
			HashRate: hashrateChan,
			Pool:     stratum,
		}
		miner.Mine()
	}

	/*hashrateReport := make([]float64, numOfDevice)
	for {
		for i := 0; i < numOfDevice; i++ {
			report := <-hashrateChan
			hashrateReport[report.MinerID] = report.Rate
		}

		var totalhashrate float64
		for minerID, hashrate := range hashrateReport {
			utils.LOG_INFO("hashrate %d: %.1fH/s\n", minerID, hashrate)
			totalhashrate += hashrate
		}
		utils.LOG_INFO("total hashrate: %.1f MH/s\n", totalhashrate)
	}*/
}
