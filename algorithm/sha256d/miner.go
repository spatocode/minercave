package sha256d

import (
	"time"

	"github.com/spatocode/minercave/net"
)

type HashRate struct {
	Rate    float64
	MinerID int
}

type miningJob struct {
	Header []byte
	Offset int
}

type Miner struct {
	validShares   uint
	staleShares   uint
	invalidShares uint
	devices       int
	startTime     uint
	HashRate      chan *HashRate
	miningJob     chan *miningJob
	pool          *net.Stratum
}

func NewMiner(cfg *net.Config) (miner *Miner) {
	stratum := net.StratumClient(cfg)
	miner = &Miner{
		devices:   cfg.Threads,
		pool:      stratum,
		startTime: uint(time.Now().Unix()),
	}
	miner.miningJob = make(chan *miningJob, miner.devices)

	return
}

func (miner *Miner) createJob() {
	miner.pool.Connect()
}

func (miner *Miner) Mine() {
	go miner.createJob()
	for i := 0; i < miner.devices; i++ {

	}
}
