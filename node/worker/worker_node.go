package worker

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"yu/config"
	. "yu/node"
	"yu/storage/kv"
	"yu/tripod"
)

type Worker struct {
	Name           string
	Port           string
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
		Port:           ":" + cfg.ServesPort,
		NodeKeeperAddr: nkAddr,
		land:           land,
		metadb:         metadb,
	}, nil

}

func (w *Worker) HandleHttp() {
	r := gin.Default()

	r.GET(HeartbeatPath, func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
		logrus.Debugf("accept heartbeat from %s", c.ClientIP())
	})

	r.Run(w.Port)
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
	tripodsInfo := make(map[string][]string)
	for triName, tri := range w.land.Tripods {
		tripodsInfo[triName] = tri.TripodMeta().AllExeNames()
	}
	return &WorkerInfo{
		Name:           w.Name,
		Port:           w.Port,
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
