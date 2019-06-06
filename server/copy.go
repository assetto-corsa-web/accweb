package server

func CopyServerSettings(id int) error {
	server := GetServerById(id, true)

	if server == nil {
		return ServerNotFound
	}

	server.Id = 0
	server.PID = 0
	server.Cmd = nil
	server.Settings.ServerName += " (copy)"

	if err := SaveServerSettings(server); err != nil {
		return err
	}

	return nil
}
