package board

import (
	"fmt"
	"math/rand"
	"time"
)

const maxSideways = 50

// Board is a 1-D board designed to hold queens. The index is the column
// and the value is the row of the queen
type Board struct {
	Queens []int
}

func RunHillClimbing(size int, allowSideways bool) {
	newBoard := CreateRandomBoard(size)
	newBoard.Print(false)
	// newBoard.Print(true)
	previousAttacking := -1
	moves := 0
	totalSideways := 0
	for true {
		lowestAttacking, _, _ := newBoard.FindLowestMove(true, false)
		if lowestAttacking < 0 {
			continueSearch := false
			if allowSideways {
				sidewaysCount := 0
				for sidewaysCount < maxSideways {
					lowestAttacking, _, _ = newBoard.FindLowestMove(true, true)
					if lowestAttacking < 0 {
						break
					}
					if lowestAttacking < previousAttacking {
						continueSearch = true
						break
					}
					sidewaysCount++
				}
				totalSideways += sidewaysCount
			}
			if !continueSearch {
				break
			}
			fmt.Printf("Went Sideways at %d and continuing at %d\n", previousAttacking, lowestAttacking)
		}
		moves++
		previousAttacking = lowestAttacking
	}
	newBoard.Print(false)
	if allowSideways {
		fmt.Printf("Went sideways %d times\n", totalSideways)
	}
	fmt.Printf("Number of moves: %d\n", moves)
}

func (b *Board) FindLowestMove(updateBoard, sideways bool) (int, int, int) {
	length := len(b.Queens)
	lowestAttacking := b.GetTotalAttackingQueens()
	lowestRow := -1
	lowestCol := -1
	var compare func(int, int) bool
	if sideways {
		compare = func(x, y int) bool { return x <= y }
	} else {
		compare = func(x, y int) bool { return x < y }
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if b.Queens[j] != i {
				original := b.Queens[j]
				b.Queens[j] = i
				attacking := b.GetTotalAttackingQueens()
				if compare(attacking, lowestAttacking) {
					lowestCol = j
					lowestRow = i
					lowestAttacking = attacking
				}
				b.Queens[j] = original
			}
		}
	}
	if lowestCol == -1 {
		return -1, -1, -1
	}
	if updateBoard {
		b.Queens[lowestCol] = lowestRow
	}
	return lowestAttacking, lowestRow, lowestCol
}

func (b *Board) GetTotalAttackingQueens() int {
	length := len(b.Queens)
	count := 0
	for col, row := range b.Queens {
		for i := col + 1; i < length; i++ {
			curRow := b.Queens[i]
			if row == curRow || abs(col-i) == abs(row-curRow) {
				count++
			}
		}
	}
	return count
}

func (b *Board) Print(oneMoveQueen bool) {
	length := len(b.Queens)
	if length > 50 {
		fmt.Println("Board is wayyy too big to print")
		fmt.Printf("Total Number of Attacking Queens: %d\n", b.GetTotalAttackingQueens())
		return
	}
	fmt.Print("|")
	for i := 0; i < length; i++ {
		fmt.Print(" - ")
	}
	fmt.Println("|")
	for i := 0; i < length; i++ {
		fmt.Print("|")
		for j := 0; j < length; j++ {
			if b.Queens[j] == i {
				if oneMoveQueen {
					fmt.Print(" QQ ")
				} else {
					fmt.Print(" Q ")
				}
			} else {
				if oneMoveQueen {
					original := b.Queens[j]
					b.Queens[j] = i
					attacking := b.GetTotalAttackingQueens()
					b.Queens[j] = original
					fmt.Printf(" %2d ", attacking)
				} else {
					fmt.Print(" x ")
				}
			}
		}
		fmt.Println("|")
	}
	fmt.Print("|")
	for i := 0; i < length; i++ {
		fmt.Print(" - ")
	}
	fmt.Println("|")
	fmt.Printf("Total Number of Attacking Queens: %d\n", b.GetTotalAttackingQueens())
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CreateRandomBoard(size int) *Board {
	b := Board{
		Queens: make([]int, 0, size),
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		b.Queens = append(b.Queens, rand.Intn(size))
	}
	return &b
}
