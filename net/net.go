package net

import (
	"log"
	"sync"
)

const (
	HashSize = 32
)

type Pool struct {
	Url       string `json:"url"`
	User      string `json:"user"`
	Password  string `json:"password"`
	KeepAlive bool   `json:"keepalive"`
	RigID     string `json:"rig-id"`
}

type Config struct {
	Currency string `json:"cryptocurrency"`
	Threads  int    `json:"threads"`
	Log      bool   `json:"log"`
	Solo     bool   `json:"solo"`
	Address  string `json:"address"`
	Pool     Pool   `json:"pool"`
}

type Target [HashSize]byte

type Block struct {
	Version     string
	PrevHash    []byte
	CurrentHash []byte
	Time        uint
	Bits        uint
	Nonce       uint
}

type ExtraNonce struct {
	Size  uint
	Value uint
}

type Job struct {
	ID         string
	Valid      bool
	ExtraNonce *ExtraNonce
	Target     Target
	Block      Block
}

type Stratum struct {
	url        string
	user       string
	target     Target
	extranonce []byte
	currentJob Job
	mutex      sync.Mutex
}

func StratumClient(cfg *Config) *Stratum {
	stratum := &Stratum{url: cfg.Pool.Url, user: cfg.Pool.User}
	log.Printf("Miner created at %s", cfg.Pool.Url)
	return stratum
}
