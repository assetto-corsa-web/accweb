package server

func CopyServerSettings(id int) error {
	server := GetServerById(id)

	if server == nil {
		return ServerNotFound
	}

	server.Id = 0
	server.Settings.ServerName += " (copy)"

	if err := SaveServerSettings(server); err != nil {
		return err
	}

	return nil
}
