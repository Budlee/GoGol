package game_of_life

import "sync"

type GameOfLifeBoard interface {
	Evolve() *[][]int
}

type golBoard struct {
	board     *[][]int
	drawBoard *[][]int
}

func zeroOutBoard(board *[][]int) {
	for y := range *board {
		for x := range (*board)[y] {
			(*board)[y][x] = 0
		}
	}
}

func mapCoordinates(y int, x int, golBoard *[][]int) (dy int, dx int) {
	dy = y
	dx = x
	if y < 0 {
		dy = len(*golBoard) - 1
	} else if y > len(*golBoard)-1 {
		dy = 0
	}
	if x < 0 {
		dx = len((*golBoard)[dy]) - 1
	} else if x > len((*golBoard)[dy])-1 {
		dx = 0
	}
	return
}

func countNeighbours(y int, x int, golBoard *[][]int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			sy, sx := mapCoordinates(y+i, x+j, golBoard)
			count += (*golBoard)[sy][sx]
			if count == 4 {
				return count
			}
		}
	}
	return count
}

func cNeighbours(gb *[][]int, xyIn <-chan struct {
	y int
	x int
}, xyOut chan<- struct {
	y int
	x int
	c int
}) {
	for xy := range xyIn {
		count := countNeighbours(xy.y, xy.x, gb)
		xyOut <- struct {
			y int
			x int
			c int
		}{y: xy.y, x: xy.x, c: count}
	}
}

func aRules(countChanIn <-chan struct {
	y int
	x int
	c int
}, gb *golBoard, wg *sync.WaitGroup) {
	for xyc := range countChanIn {
		applyRulesToGolBoard(xyc.y, xyc.x, xyc.c, (*(*gb).board)[xyc.y][xyc.x] == 1, (*gb).drawBoard)
		wg.Done()
	}
}

func applyRulesToGolBoard(y int, x int, count int, live bool, golBoard *[][]int) {
	if live {
		if count < 2 {
			(*golBoard)[y][x] = 0
		} else if count > 3 {
			(*golBoard)[y][x] = 0
		} else {
			(*golBoard)[y][x] = 1
		}
	} else {
		if count == 3 {
			(*golBoard)[y][x] = 1
		}
	}
}

func (g *golBoard) Evolve() *[][]int {
	var b [][]int = *g.board

	xyInChan := make(chan struct {
		y int
		x int
	})
	countChan := make(chan struct {
		y int
		x int
		c int
	})
	var wg sync.WaitGroup
	go cNeighbours(g.board, xyInChan, countChan)
	go aRules(countChan, g, &wg)
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[y]); x++ {
			wg.Add(1)
			xyInChan <- struct {
				y int
				x int
			}{y: y, x: x}
		}
	}
	wg.Wait()
	close(xyInChan)
	close(countChan)
	tmpGolRef := g.board
	g.board = g.drawBoard
	g.drawBoard = tmpGolRef
	zeroOutBoard(g.drawBoard)
	return g.board
}

func New(inBoard *[][]int) GameOfLifeBoard {
	return &golBoard{board: inBoard, drawBoard: copyIncomingBoard(inBoard)}
}

func copyIncomingBoard(inBoard *[][]int) *[][]int {
	var b [][]int = *inBoard
	drawOnBoard := make([][]int, len(b))
	for i := range drawOnBoard {
		drawOnBoard[i] = make([]int, len(b[i]))
	}
	return &drawOnBoard
}
