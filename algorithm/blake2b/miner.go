package blake2b

import (
	"github.com/spatocode/minercave/net"
)

type HashRate struct {
	Rate float64
	MinerID  int
}

type MiningWork struct {
	Header []byte
	Offset int
}

type Miner struct {
	HashRate   chan *HashRate
	MiningWork chan *MiningWork
	Client            *net.Stratum
}


func (miner *Miner) Mine() {
	
}