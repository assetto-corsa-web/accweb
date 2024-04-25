package events

import (
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager/events/callback"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager/events/logparser"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager/events/logwriter"
	"github.com/assetto-corsa-web/accweb/internal/pkg/server_manager/events/windowsadv"
)

func InitializeAll(sm *server_manager.Service) {

	logwriter.Register(sm)
	logparser.Register(sm)
	windowsadv.Register(sm)
	callback.Register(sm)

}
