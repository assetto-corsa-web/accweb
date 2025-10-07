package logwriter

import (
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/assetto-corsa-web/accweb/internal/pkg/event"
	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
	"github.com/assetto-corsa-web/accweb/internal/pkg/instance"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/sirupsen/logrus"
)

const (
	logDir        = "logs"
	logTimeFormat = "20060102_150405"
	logExt        = ".log"
)

var sM *server_manager.Service
var withTs bool
var lm sync.Map

func Register(sm *server_manager.Service) {
	sM = sm
	withTs = sm.Config().Log.WithTimestamp

	event.Register(handleEvent)
}

func handleEvent(data event.Eventer) {
	switch ev := data.(type) {
	case event.EventInstanceBeforeStart:
		if v, ok := lm.Load(ev.InstanceId); ok {
			v.(*os.File).Close()
		}

		i, err := sM.GetServerByID(ev.InstanceId)
		if err != nil {
			logrus.Error("instance not found")
		}

		f, err := createLogFile(i)
		if err != nil {
			logrus.Error("failed to create instance log")
		}
		lm.Store(ev.InstanceId, f)

	case event.EventInstanceBeforeStop:
		v, ok := lm.Load(ev.InstanceId)
		if !ok {
			return
		}
		v.(*os.File).Close()
		lm.Delete(ev.InstanceId)

	case event.EventInstanceOutput:
		f, ok := lm.Load(ev.InstanceId)
		if !ok {
			return
		}

		data := string(ev.Output) + "\n"

		if withTs {
			data = ev.Timestamp.Format(time.RFC3339Nano) + ": " + data
		}

		f.(*os.File).Write([]byte(data))

		logrus.
			WithFields(logrus.Fields{
				"instanceId": ev.InstanceId,
				"log":        string(ev.Output),
			}).
			Debug("instance log")
	}
}

func createLogFile(s *instance.Instance) (*os.File, error) {
	if err := helper.CreateIfNotExists(path.Join(s.Path, logDir), 0755); err != nil {
		return nil, err
	}

	filename := fmt.Sprintf("logs_%s_%s%s", time.Now().Format(logTimeFormat), s.GetID(), logExt)

	return os.Create(path.Join(s.Path, logDir, filename))
}
