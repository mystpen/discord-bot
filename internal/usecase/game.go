package usecase

import (
	"math/rand"
)

type GameUsecase struct {
}

// New GameUsecase
func NewGameUsecase() *GameUsecase {
	return &GameUsecase{}
}

// game implementation
func (g *GameUsecase) RollDice() string {
	res := rand.Intn(6) + 1
	diceSides := map[int]string{
		1: `    	    _________ 
		   /        /|
		  /________/ |
		  |        | |
		  |   o    | |
		  |        | /
		  |________|/ `,

		2: `    	    _________ 
		   /        /|
		  /________/ |
		  | o      | |
		  |        | |
		  |     o  | /
		  |________|/ `,

		3: `    	    _________ 
		   /        /|
		  /________/ |
		  | o      | |
		  |   o    | |
		  |     o  | /
		  |________|/ `,

		4: `    	    _________ 
		   /        /|
		  /________/ |
		  | o   o  | |
		  |        | |
		  | o   o  | /
		  |________|/ `,

		5: `    	    _________ 
		   /        /|
		  /________/ |
		  | o   o  | |
		  |   o    | |
		  | o   o  | /
		  |________|/ `,

		6: `    	    _________ 
		   /        /|
		  /________/ |
		  | o o o  | |
		  |        | |
		  | o o o  | /
		  |________|/ `,
	}
	return "```" + diceSides[res] + "```"
}
