package abstractfactory

import "fmt"

func init() {
	gameEasy := NewGame(Easy())
	fmt.Println(gameEasy)
}