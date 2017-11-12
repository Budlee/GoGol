package game_of_life

import "testing"

func TestNew(t *testing.T) {
	var gameBoard [][]int = [][]int{{1, 1, 0, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	var gol (GameOfLifeBoard) = New(&gameBoard)
	if gol == nil {
		t.Error("Return from New should not be nil")
	}
}

func TestWrapAroundEvolve(t *testing.T) {
	var gameBoard [][]int = [][]int{{1, 1, 0, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	var desiredEvolvedBoard [][]int = [][]int{{1, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}}
	var gol (GameOfLifeBoard) = New(&gameBoard)
	compareGOLBoards(t, &desiredEvolvedBoard, gol.Evolve())
}

func TestGliderEvlove(t *testing.T) {
	var gameBoard [][]int = [][]int{{0, 1, 0, 0, 0}, {0, 0, 1, 1, 0}, {0, 1, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	var desiredEvolvedBoard [][]int = [][]int{{0, 0, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 1, 1, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}}
	var gol (GameOfLifeBoard) = New(&gameBoard)
	gol.Evolve()
	compareGOLBoards(t, &desiredEvolvedBoard, gol.Evolve())
}

func compareGOLBoards(t *testing.T, desiredGOLBoard *[][]int, actualGOLBoard *[][]int) {
	var dBoard [][]int = *desiredGOLBoard
	var aBoard [][]int = *actualGOLBoard
	for x := 0; x <= len(dBoard)-1; x++ {
		for y := 0; y <= len(dBoard)-1; y++ {
			if dBoard[x][y] != aBoard[x][y] {
				t.Errorf("Evolved gameboard at x->%d and y->%d does not match desiredEvolvedBoard", x, y)
				return
			}
		}
	}
}
