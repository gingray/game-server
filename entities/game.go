package entities

type Game struct {

}

func NewGame() *Game {
	return &Game{}
}

func (self *Game) Fetch (evt string) {

}

func (self *Game) Broadcast() string {
	return ""
}