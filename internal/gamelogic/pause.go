package gamelogic

import (
	"fmt"

	"github.com/bootdotdev/go-api-gate/courses/learn-pub-sub/exchanges-queues/internal/routing"
)

func (gs *GameState) HandlePause(ps routing.PlayingState) {
	defer fmt.Println("------------------------")
	fmt.Println()
	if ps.IsPaused {
		fmt.Println("==== Pause Detected ====")
		gs.pauseGame()
	} else {
		fmt.Println("==== Resume Detected ====")
		gs.resumeGame()
	}
}
