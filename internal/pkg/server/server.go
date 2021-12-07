package server

import "os/exec"

type Server struct {
	ID     int
	Path   string
	Cfg    AccWebConfigJson
	AccCfg accConfigFiles

	cmd *exec.Cmd
}

func (s *Server) Start() {}

func (s *Server) Stop() {}

func (s *Server) GetProcessID() int {
	if s.isRunning() {
		return s.cmd.Process.Pid
	}

	return 0
}

func (s *Server) Save() {}

func (s *Server) isRunning() bool {
	return s.cmd != nil && s.cmd.Process != nil && s.cmd.Process.Pid != 0
}
