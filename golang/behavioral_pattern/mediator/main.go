package main

func main() {
	stationManager := newStationManager()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}

	freighTrain := &FreighTrain{
		mediator: stationManager,
	}

	//it can arrive
	passengerTrain.arrive()

	//it cannot arrive, set train as queue
	freighTrain.arrive()

	// st train departure, now platform is free, permit train queue
	passengerTrain.depart()

}
