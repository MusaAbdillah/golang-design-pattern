package main

type Player struct {
	dress Dress
	playerType string
	lat int
	long int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: dressType,
		dress: dress,
	}
}

func (p *Player) newLocation(lat, long int){
	p.lat = lat
	p.long = long
} 
