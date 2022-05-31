package abstractfactory

type Game struct {
	Stage    int
	Point    int
	Health   int
	MaxLevel int
	MaxArena int
}

func (g *Game) Next() {

}

type GameFactoryIn interface {
	createLevel() int
	createArena() int
	createHealth() int
}

func NewGame(factory GameFactoryIn) *Game {
	return &Game{
		Stage:    1,
		Point:    0,
		Health:   factory.createHealth(),
		MaxLevel: factory.createLevel(),
		MaxArena: factory.createArena(),
	}
}