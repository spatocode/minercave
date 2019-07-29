package net

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spatocode/minercave/utils"
)

const (
	hashSize = 32
)

// StratumRequest message from stratum server
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
	ID                string
	Notification      string
	ExtraNonce1       string
	ExtraNonce2Length float64
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

type AuthorizeResponse struct {
	ID     uint
	Result bool
	Error  StratumError
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
	PrevHash    string
	CurrentHash string
	Ntime       string
	Nbits       string
	Nonce       uint
}

// Job contains information about the job generated
type Job struct {
	ID           string
	Valid        bool
	ExtraNonce1  string
	ExtraNonce2  uint
	Height       int
	Target       *big.Int
	Coinbase1    string
	Coinbase2    string
	MerkleBranch []string
	Block        Block
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
	log.Printf("Using pool %v", stratum.url)
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

		response, err := stratum.handleRawMessage([]byte(rawMsg))
		if err != nil {
			utils.LOG_ERR("Error handling raw message: %s\n", err)
			continue
		}

		stratum.dispatchHandler(response)
	}
}

func (stratum *Stratum) handleRawMessage(message []byte) (interface{}, error) {
	stratum.mutex.Lock()
	defer stratum.mutex.Unlock()

	var (
		method string
		id     uint
		obj    map[string]json.RawMessage
	)

	err := json.Unmarshal(message, &obj)
	if err != nil {
		return nil, err
	}

	value, ok := obj["method"]
	if ok == true {
		err = json.Unmarshal(value, &method)
		if err != nil {
			return nil, err
		}
	}

	err = json.Unmarshal(obj["id"], &id)
	if err != nil {
		return nil, err
	}

	if id == stratum.authID {
		var (
			result bool
			errObj []interface{}
		)

		err = json.Unmarshal(obj["result"], &result)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(obj["error"], &errObj)
		if err != nil {
			return nil, err
		}

		authResp := &AuthorizeResponse{ID: id, Result: result}
		if errObj != nil {
			errNum, _ := errObj[0].(float64)
			errStr, _ := errObj[1].(string)
			authResp.Error.Num = uint(errNum)
			authResp.Error.Str = errStr
		}

		return authResp, nil
	}

	if id == stratum.subID {
		var (
			result      []json.RawMessage
			subDetail   [][]string
			extranonce1 string
			extranonce2 float64
		)

		err = json.Unmarshal(obj["result"], &result)
		if err != nil {
			return nil, err
		}

		if len(result) == 0 {
			return nil, err
		}

		err = json.Unmarshal(result[0], &subDetail)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(result[1], &extranonce1)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(result[2], &extranonce2)
		if err != nil {
			return nil, err
		}

		subResp := &SubscribeResponse{
			ExtraNonce1:       extranonce1,
			ExtraNonce2Length: extranonce2,
		}

		for i := 0; i < len(subDetail); i++ {
			if subDetail[i][0] == "mining.notify" {
				subResp.Notification = subDetail[i][0]
				subResp.ID = subDetail[i][1]
			}
		}

		return subResp, nil
	}

	switch method {
	case "mining.notify":
		var params []interface{}

		err = json.Unmarshal(obj["params"], &params)
		if err != nil {
			return nil, err
		}

		notifResp := &NotificationResponse{
			JobID:     params[0].(string),
			PrevHash:  params[1].(string),
			Coinbase1: params[2].(string),
			Coinbase2: params[3].(string),
			Version:   params[5].(string),
			Nbits:     params[6].(string),
			Ntime:     params[7].(string),
			CleanJobs: params[8].(bool),
		}
		return notifResp, nil

	case "mining.set_difficulty":
		var params []interface{}

		err = json.Unmarshal(obj["params"], &params)
		if err != nil {
			return nil, err
		}

		difficulty, _ := params[0].(int)
		stratum.diff = float64(difficulty)

		stratReq := &StratumRequest{}
		stratReq.Method = method

		var stratParam []string
		stratParam = append(stratParam, strconv.FormatFloat(stratum.diff, 'E', -1, 32))
		stratReq.Params = stratParam
		return stratReq, nil

	default:
		stratResp := &StratumResponse{}
		err = json.Unmarshal(message, &stratResp)
		if err != nil {
			return nil, err
		}

		return stratResp, nil
	}
}

func (stratum *Stratum) dispatchHandler(response interface{}) {
	switch response.(type) {
	case *SubscribeResponse:
		stratum.subscribeHandler(response)
	case *AuthorizeResponse:
		stratum.authorizeHandler(response)
	case *NotificationResponse:
		stratum.notificationHandler(response)
	case *StratumResponse:
		stratum.stratumResponseHandler(response)
	case *StratumRequest:
		stratum.stratumRequestHandler(response)
	default:
		stratum.basicHandler(response)
	}
}

func (stratum *Stratum) stratumResponseHandler(response interface{}) {
	log.Printf("stratumResponseHandler")
}

func (stratum *Stratum) stratumRequestHandler(response interface{}) {
	log.Printf("stratumRequestHandler")
}

func (stratum *Stratum) notificationHandler(response interface{}) {
	stratum.mutex.Lock()
	defer stratum.mutex.Unlock()

	resp := response.(*NotificationResponse)
	stratum.currentJob.ID = resp.JobID
	stratum.currentJob.Valid = resp.CleanJobs
	//stratum.currentJob.Height
	stratum.currentJob.Coinbase1 = resp.Coinbase1
	stratum.currentJob.Coinbase2 = resp.Coinbase2
	stratum.currentJob.Block.Version = resp.Version
	stratum.currentJob.Block.Nbits = resp.Nbits
	stratum.currentJob.Block.Ntime = resp.Ntime
	stratum.currentJob.Block.PrevHash = resp.PrevHash
	log.Printf("New notification received")
}

func (stratum *Stratum) subscribeHandler(response interface{}) {
	resp := response.(*SubscribeResponse)
	stratum.currentJob.ExtraNonce1 = resp.ExtraNonce1
	stratum.currentJob.ExtraNonce2 = uint(resp.ExtraNonce2Length)
	log.Printf("Subscription successful")
}

func (stratum *Stratum) authorizeHandler(response interface{}) {
	stratum.mutex.Lock()
	defer stratum.mutex.Unlock()

	resp := response.(*AuthorizeResponse)
	if resp.ID == stratum.authID {
		if resp.Result {
			log.Printf("Logged in as %s", stratum.user)
		} else {
			log.Printf("Authorization failure")
		}
	}

	log.Printf("Authorization successful")
}

func (stratum *Stratum) basicHandler(response interface{}) {
	fmt.Println("basicHandler")
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
