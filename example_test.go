package game_of_life

import "fmt"

func displayGOL(pGOLBoard *[][]int) {
	var golBoard [][]int = *pGOLBoard
	for y := 0; y < len(golBoard); y++ {
		for x := 0; x < len((golBoard[y])); x++ {
			fmt.Printf("%d", golBoard[y][x])
		}
		fmt.Println()
	}

}

func ExampleEvolveAndDisplay() {
	var initialGameBoard [][]int = [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	var gol (GameOfLifeBoard) = New(&initialGameBoard)
	var golBoard = *gol.Evolve()
	for y := 0; y < len(golBoard); y++ {
		for x := 0; x < len((golBoard[y])); x++ {
			fmt.Printf("%d", golBoard[y][x])
		}
		fmt.Println()
	}

	//Output
	//00000
	//00100
	//00100
	//00100
	//00000
}
