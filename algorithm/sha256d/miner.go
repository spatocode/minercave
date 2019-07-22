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
	Devices    int
	HashRate   chan *HashRate
	MiningWork chan *miningWork
	Client     *net.Stratum
}

func (miner *Miner) Mine() {
	miner.MiningWork = make(chan *miningWork, miner.Devices)
}
