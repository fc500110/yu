package master

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	peerstore "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	maddr "github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	. "yu/blockchain"
	. "yu/common"
	"yu/config"
	. "yu/node"
	. "yu/txn"
)

func makeP2pHost(ctx context.Context, cfg *config.MasterConf) (host.Host, error) {
	r, err := loadNodeKeyReader(cfg)
	if err != nil {
		return nil, err
	}
	priv, _, err := crypto.GenerateKeyPairWithReader(cfg.NodeKeyType, cfg.NodeKeyBits, r)
	if err != nil {
		return nil, err
	}
	return libp2p.New(
		ctx,
		libp2p.Identity(priv),
		libp2p.ListenAddrStrings(cfg.P2pListenAddrs...),
	)
}

func loadNodeKeyReader(cfg *config.MasterConf) (io.Reader, error) {
	if cfg.NodeKey != "" {
		return bytes.NewBufferString(cfg.NodeKey), nil
	}
	if cfg.NodeKeyFile != "" {
		return os.Open(cfg.NodeKeyFile)
	}
	return rand.Reader, nil
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

func (m *Master) handleStream(s network.Stream) {
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	go m.readFromNetwork(rw)
	go m.writeToNetwork(rw)
}

// Read the data of blockchain from P2P network.
func (m *Master) readFromNetwork(rw *bufio.ReadWriter) {
	for {
		byt, err := rw.ReadBytes('\n')
		if err != nil {
			logrus.Errorf("Read data from P2P-network error: %s", err.Error())
			continue
		}
		tbody, err := DecodeTb(byt)
		if err != nil {
			logrus.Errorf("get transfer-body error: %s", err.Error())
			continue
		}
		err = m.handleTransferBody(tbody)
		if err != nil {
			logrus.Errorf("handle transfer-body error：%s", err.Error())
		}
	}
}

// Write and broadcast the data to P2P network.
func (m *Master) writeToNetwork(rw *bufio.ReadWriter) {
	for {
		select {
		case blocksBody := <-m.blocksBcChan:
			byt, err := blocksBody.Encode()
			if err != nil {
				logrus.Errorf("encode block-body error: %s", err.Error())
				continue
			}
			_, err = rw.Write(byt)
			if err != nil {
				logrus.Errorf("write block-body to P2P network error: %s", err.Error())
				continue
			}
			rw.Flush()
		case txnsBody := <-m.txnsBcChan:
			byt, err := txnsBody.Encode()
			if err != nil {
				logrus.Errorf("encode txns-body error: %s", err.Error())
				continue
			}
			_, err = rw.Write(byt)
			if err != nil {
				logrus.Errorf("write txns-body error: %s", err.Error())
				continue
			}
			rw.Flush()
		}
	}
}

func (m *Master) handleTransferBody(tbody *TransferBody) error {
	if m.RunMode == MasterWorker {
		// todo: forwards to worker
		return nil
	}
	switch tbody.Type {
	case BlockTransfer:
		var block IBlock
		err := tbody.DecodeBody(block)
		if err != nil {
			return err
		}
		m.blocksFromNetChan <- block
	case TxnsTransfer:
		var txns SignedTxns
		err := tbody.DecodeBody(&txns)
		if err != nil {
			return err
		}
		for _, txn := range txns {
			err = m.txPool.Insert(txn)
			if err != nil {
				return err
			}
		}
	}
	return nil
}