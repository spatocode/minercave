package net

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/spatocode/minercave/utils"
)

const (
	hashSize = 32
)

// StratumRequest message to stratum server
type StratumRequest struct {
	ID     uint     `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

// StratumResponse message from statum server
type StratumResponse struct {
	Method string           `json:"method"`
	ID     uint             `json:"id"`
	Params []string         `json:"params"`
	Result *json.RawMessage `json:"result"`
	Error  StratumError     `json:"error"`
}

// StratumError struct for response error
type StratumError struct {
	Num    uint
	Str    string
	Result *json.RawMessage
}

type SubscribeResponse struct {
	ID                 uint
	SubscriptionDetail []string
	ExtraNonce1        string
	ExtraNonce2Length  float64
}

type NotificationResponse struct {
	JobID        string
	PrevHash     string
	Coinbase1    string
	Coinbase2    string
	MerkleBranch []string
	Version      string
	Nbits        string
	Ntime        string
	CleanJobs    bool
}

// Pool struct for stratum pool settings
type Pool struct {
	URL       string `json:"url"`
	User      string `json:"user"`
	Pass      string `json:"password"`
	KeepAlive bool   `json:"keepalive"`
	RigID     string `json:"rig-id"`
}

// Config struct for configuration settings
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
	extranonce1   []byte
	currentJob    Job
	startTime     uint
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

	stratum.startTime = uint(time.Now().Unix())

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

		response, err = stratum.handleRawMessage([]byte(rawMsg))
		if err != nil {
			utils.LOG_ERR("Error handling raw message: %s\n", err)
			continue
		}

	}
}

func (stratum *Stratum) handleRawMessage(message []byte) (interface{}, error) {

}

func (stratum *Stratum) dispatchHandler(response StratumResponse) {
	switch response.ID {
	case 1:
		stratum.subscribeHandler(response)
	case 2:
		stratum.authorizeHandler(response)
	default:
		stratum.basicHandler(response)
	}
}

func (stratum *Stratum) subscribeHandler(response StratumResponse) {
	result, err := response.Result.MarshalJSON()
	if err != nil {
		return
	}

	fmt.Println(result)

	/*subscribeResp := SubscribeResponse {
		ID: response.ID,
		ExtraNonce1: response.Result
	}
	msg, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Subscription response: %s .......\n", msg)
	stratum.currentJob = response.Result*/
}

func (stratum *Stratum) authorizeHandler(response StratumResponse) {
	msg, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Authorize response: %s ........\n", msg)
}

func (stratum *Stratum) basicHandler(response StratumResponse) {

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

	stratum.startTime = uint(time.Now().Unix())

	return nil
}

// Subscribe subscribes client to stratum server for receiving mining jobs
func (stratum *Stratum) Subscribe() error {
	message := StratumRequest{
		ID:     stratum.id,
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

// Authorize authorizes miner for mining
func (stratum *Stratum) Authorize() error {
	message := StratumRequest{
		ID:     stratum.id,
		Method: "mining.authorize",
		Params: []string{stratum.user, stratum.pass},
	}

	stratum.authID = message.ID
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
		return nil
	}

	return nil
}
