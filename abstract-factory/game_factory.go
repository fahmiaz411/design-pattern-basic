package abstractfactory

type GameFactory struct {
	Level
	Arena
	Health
}

func (g *GameFactory) createLevel() int {

	switch g.Level {
	case LevelEasy:
		return 10
	case LevelMedium:
		return 50
	case LevelHard:
		return 100
	}

	return 0
}

func (g *GameFactory) createArena() int {
	switch g.Arena {
	case ArenaEasy:
		return 5
	case ArenaMedium:
		return 10
	case ArenaHard:
		return 15
	}

	return 0
}

func (g *GameFactory) createHealth() int {
	switch g.Health {
	case HealthEasy:
		return 100
	case HealthMedium:
		return 80
	case HealthHard:
		return 60
	}

	return 0
}

func Easy() *GameFactory {
	return &GameFactory{
		Level:  LevelEasy,
		Arena:  ArenaEasy,
		Health: HealthEasy,
	}
}

func Medium() *GameFactory {
	return &GameFactory{
		Level:  LevelMedium,
		Arena:  ArenaMedium,
		Health: HealthMedium,
	}
}
func Hard() *GameFactory {
	return &GameFactory{
		Level:  LevelHard,
		Arena:  ArenaHard,
		Health: HealthHard,
	}
}