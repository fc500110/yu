package worker

import (
	"bytes"
	"net/http"
	"yu/config"
	. "yu/node"
	"yu/storage/kv"
	"yu/tripod"
)

type Worker struct {
	Name           string
	HttpPort       string
	WsPort         string
	NodeKeeperAddr string
	land           *tripod.Land
	metadb         kv.KV
}

func NewWorker(cfg *config.WorkerConf, land *tripod.Land) (*Worker, error) {
	metadb, err := kv.NewKV(&cfg.DB)
	if err != nil {
		return nil, err
	}
	nkAddr := "localhost:" + cfg.NodeKeeperPort
	return &Worker{
		Name:           cfg.Name,
		HttpPort:       ":" + cfg.HttpPort,
		WsPort:         ":" + cfg.WsPort,
		NodeKeeperAddr: nkAddr,
		land:           land,
		metadb:         metadb,
	}, nil

}

// Register into NodeKeeper
func (w *Worker) RegisterInNk() error {
	infoByt, err := w.Info().EncodeMasterInfo()
	if err != nil {
		return err
	}
	_, err = w.postToNk(RegisterWorkersPath, infoByt)
	return err
}

func (w *Worker) Info() *WorkerInfo {
	tripodsInfo := make(map[string]TripodInfo)
	for triName, tri := range w.land.Tripods {
		execNames := tri.TripodMeta().AllExecNames()
		queryNames := tri.TripodMeta().AllQueryNames()
		tripodsInfo[triName] = TripodInfo{
			ExecNames:  execNames,
			QueryNames: queryNames,
		}
	}
	return &WorkerInfo{
		Name:           w.Name,
		ServesPort:     w.ServesPort,
		NodeKeeperAddr: w.NodeKeeperAddr,
		TripodsInfo:    tripodsInfo,
		Online:         true,
	}
}

func (w *Worker) postToNk(path string, body []byte) (*http.Response, error) {
	url := w.NodeKeeperAddr + path
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	cli := &http.Client{}
	return cli.Do(req)
}