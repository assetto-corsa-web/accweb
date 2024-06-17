package instance

func (l *LiveState) GetCar(cId int) *CarState {
	if c, ok := l.Cars[cId]; ok {
		return c
	}

	return nil
}

func (l *LiveState) GetDriver(cId int) *DriverState {
	if c, ok := l.connections[cId]; ok {
		return c
	}

	return nil
}
