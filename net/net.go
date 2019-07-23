package net

import (
	"encoding/json"
	"bufio"
	"io"
	"log"
	"math/big"
	"net"
	"sync"
	"strings"
	"time"
	"fmt"

	"github.com/spatocode/minercave/utils"
)

const (
	hashSize = 32
)

type StratumRequest struct {
	ID		uint		`json:"id"`
	Method	string 		`json:"method"`
	Params	[]string	`json:"params"`
}

type StratumResponse struct {
	Method	string				`json:"method"`
	ID		uint				`json:"id"`
	Result	*json.RawMessage	`json:"result"`
	Error	StratumError		`json:"error"`
}

type StratumError struct {
	Num		uint
	Str		string
	Result	*json.RawMessage
}

type Pool struct {
	URL       string `json:"url"`
	User      string `json:"user"`
	Pass      string `json:"password"`
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

// Block contains information about our blockchain header
type Block struct {
	Version     string
	PrevHash    []byte
	CurrentHash []byte
	Time        uint
	Bits        uint
	Nonce       uint
}

// Job contains information about the job generated
type Job struct {
	ID          string
	Valid       bool
	ExtraNonce1 uint
	ExtraNonce2 uint
	Height      int
	Target      *big.Int
	Block       Block
}

// Stratum contains information about a stratum pool connection
type Stratum struct {
	validShares   uint
	invalidShares uint
	latestJobTime uint
	socket        net.Conn
	reader        *bufio.Reader
	url           string
	user          string
	pass          string
	target        *big.Int
	extranonce1    []byte
	currentJob    Job
	startTime	  uint32
	mutex         sync.Mutex
	subID         uint
	submitID      []uint
	id            uint
	authID        uint
	diff          float64
}

// StratumClient registers a new stratum client
func StratumClient(cfg *Config) *Stratum {
	var host string
	if strings.HasPrefix(cfg.Pool.URL, "stratum+tcp://") {
		host = strings.TrimPrefix(cfg.Pool.URL, "stratum+tcp://")
	}
	stratum := &Stratum{url: host, user: cfg.Pool.User, pass: cfg.Pool.Pass}
	return stratum
}

// Connect simply connects to a stratum pool
func (stratum *Stratum) Connect() {
	log.Printf("Connecting to pool %v", stratum.url)
	conn, err := net.Dial("tcp", stratum.url)
	if err != nil {
		utils.LOG_ERR("DNS error: Bad network or host [%s]\n", stratum.url)
	}

	stratum.socket = conn
	stratum.id = 1
	stratum.authID = 2
	stratum.diff = 1
	stratum.reader = bufio.NewReader(stratum.socket)
	go stratum.Listen()

	err = stratum.Subscribe()
	if err != nil {
		utils.LOG_ERR("%s", err)
	}

	err = stratum.Authorize()
	if err != nil {
		utils.LOG_ERR("%s", err)
	}

	stratum.startTime = uint32(time.Now().Unix())

	time.Sleep(10000 * time.Minute)
}

// Listen always listens to incoming message from stratum server
func (stratum *Stratum) Listen() {
	for {
		rawMsg, err := stratum.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				utils.LOG_ERR("DNS error: Bad network or host [%s]\n", stratum.url)
				stratum.Reconnect()
			} else {
				utils.LOG_ERR("%s", err)
			}
			continue
		}

		response := StratumResponse{}
		err = json.Unmarshal([]byte(rawMsg), &response)
		if err != nil {
			utils.LOG_ERR("%s", err)
		}
		stratum.dispatch(response)
	}
}

func (stratum *Stratum) dispatch(response StratumResponse) {
	msg, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n",msg)
}

// Reconnect reconnects to the stratum server when lost
func (stratum *Stratum) Reconnect() error {
	conn, err := net.Dial("tcp", stratum.url)
	if err != nil {
		utils.LOG_ERR("DNS error: Bad network or host [%s]\n", stratum.url)
	}

	stratum.socket = conn
	stratum.reader = bufio.NewReader(stratum.socket)

	err = stratum.Subscribe()
	if err != nil {
		return err
	}

	return nil
}

// Subscribe subscribes client to stratum server for receiving mining jobs
func (stratum *Stratum) Subscribe() error {
	message := StratumRequest{
		ID: stratum.id,
		Method: "mining.subscribe",
		Params: []string{"minercave"},
	}

	stratum.subID = message.ID
	stratum.id++

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = stratum.socket.Write(jsonMsg)
	if err != nil {
		return err
	}

	_, err = stratum.socket.Write([]byte("\n"))
	if err != nil {
		return err
	}

	return nil
}

// Authorize authorizes workers for mining
func (stratum *Stratum) Authorize() error {
	
}
