package master

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/host"
	peerstore "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	maddr "github.com/multiformats/go-multiaddr"
	"net/http"
	"strconv"
	"sync"
	"time"
	"yu/config"
	. "yu/node"
	"yu/storage/kv"
)

var MasterInfoKey = []byte("master-info")

type Master struct {
	sync.RWMutex
	info    *MasterInfo
	p2pHost host.Host
	metadb  kv.KV
	port    string
	ctx     context.Context

	// key: NodeKeeper's IP
	tickers map[string]*time.Ticker
	timeout time.Duration
}

func NewMaster(cfg *config.MasterConf) (*Master, error) {
	metadb, err := kv.NewKV(&cfg.DB)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	p2pHost, err := makeP2pHost(ctx, cfg)
	if err != nil {
		return nil, err
	}
	data, err := metadb.Get(MasterInfoKey)
	if err != nil {
		return nil, err
	}

	var info *MasterInfo
	if data == nil {
		info = &MasterInfo{
			NodeKeepersInfo: make(map[string]NodeKeeperInfo),
		}
		infoByt, err := info.EncodeMasterInfo()
		if err != nil {
			return nil, err
		}
		err = metadb.Set(MasterInfoKey, infoByt)
		if err != nil {
			return nil, err
		}
	} else {
		info, err = DecodeMasterInfo(data)
		if err != nil {
			return nil, err
		}
	}

	info.P2pID = p2pHost.ID().String()

	timeout := time.Duration(cfg.Timeout) * time.Second

	return &Master{
		info:    info,
		p2pHost: p2pHost,
		metadb:  metadb,
		tickers: make(map[string]*time.Ticker),
		timeout: timeout,
		ctx:     ctx,
		port:    ":" + cfg.ServesPort,
	}, nil
}

func (m *Master) ConnectP2PNetwork(cfg *config.MasterConf) error {
	m.p2pHost.SetStreamHandler(protocol.ID(cfg.ProtocolID), m.handleStream)

	for _, addrStr := range cfg.ConnectAddrs {
		addr, err := maddr.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		peer, err := peerstore.AddrInfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		err = m.p2pHost.Connect(m.ctx, *peer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Master) HandleHttp() {
	r := gin.Default()

	r.POST(WatchNodeKeepersPath, func(c *gin.Context) {
		m.watchNodeKeepers(c)
	})

	r.Run(m.port)
}

// check the health of NodeKeepers
func (m *Master) CheckHealth() {
	for {
		for nk, ticker := range m.tickers {
			<-ticker.C
			m.Lock()
			delete(m.info.NodeKeepersInfo, nk)
			delete(m.tickers, nk)
			m.Unlock()
		}
	}
}

func (m *Master) watchNodeKeepers(c *gin.Context) {
	var nkInfo NodeKeeperInfo
	err := c.ShouldBindJSON(&nkInfo)
	if err != nil {
		c.String(
			http.StatusBadRequest,
			fmt.Sprintf("NodeKeeperInfo decode failed: %s", err.Error()),
		)
		return
	}
	nkIP := c.ClientIP()
	m.RLock()
	oldNkInfo, ok := m.info.NodeKeepersInfo[nkIP]
	m.RUnlock()
	if !ok || !nkInfo.Equals(oldNkInfo) {
		m.Lock()
		m.info.NodeKeepersInfo[nkIP] = nkInfo
		workerId := len(m.info.NodeKeepersInfo)
		m.tickers[nkIP].Reset(m.timeout)
		m.Unlock()
		c.String(http.StatusOK, strconv.Itoa(workerId))
	} else {
		c.String(http.StatusOK, "")
	}
}
