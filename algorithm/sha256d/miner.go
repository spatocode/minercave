package sha256d

import (
	"github.com/spatocode/minercave/net"
)

type HashRate struct {
	Rate    float64
	MinerID int
}

type miningWork struct {
	Header []byte
	Offset int
}

type Miner struct {
	validShares   uint
	staleShares   uint
	invalidShares uint
	Devices       int
	HashRate      chan *HashRate
	MiningWork    chan *miningWork
	Pool          *net.Stratum
}

func (miner *Miner) Mine() {
	miner.MiningWork = make(chan *miningWork, miner.Devices)
}
