//动态规划4子棋
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	blank = 0
	white = 1
	black = 2

	icon_blank = "⬜️"
	icon_white = "⚪️"
	icon_black = "⚫️"

	win       = 4
	str_white = "白棋"
	str_black = "黑棋"
	board_row = 6
	board_col = 7
)

var (
	Err_not_exist = errors.New("Coordinates do not exist.")

	win_line = []byte{0x0f, 0x1e, 0x3c, 0x78}
)

type ChessBoard struct {
	Board  [][]int
	Blanks []int
	result [][]int
}

func main() {
	cb := initBoard(board_row, board_col)
	t1 := time.Now()

	player := 1
	isWin := false

	for i := 1; cb.next(); i++ {
		fmt.Printf("-----next:%d------\r\n", i)
		c, r := cb.randPoint()
		cb.SetOneChess(r, c, player)
		isWin = cb.IsWin(player)
		if isWin {

			break
		}
		player = 3 - player
		//cb.PrintM()

	}
	cb.PrintM()
	if isWin {

		p := str_white
		if player == black {
			p = str_black
		}
		fmt.Println(p, "赢了")
	} else {
		fmt.Println("平局")
	}
	fmt.Println(time.Now().Sub(t1).Nanoseconds())
}

func initBoard(row, col int) *ChessBoard {
	cb := &ChessBoard{
		Blanks: make([]int, board_row*board_col),
		Board:  make([][]int, row),
	}
	for i := 0; i < row; i++ {
		cb.Board[i] = make([]int, col)
	}
	for i := 0; i < board_row*board_col; i++ {
		cb.Blanks[i] = i
	}
	return cb
}

//dp方法
func (o *ChessBoard) IsWin(chess int) bool {

	//先算横向
	for i := 0; i < board_row; i++ {
		f_row := make([]int, board_col)
		for j := 0; j < board_col; j++ {
			if o.Board[i][j] == chess {
				if j == 0 {
					f_row[j] = 1
				} else {
					f_row[j] = f_row[j-1] + 1
					if f_row[j] == win {
						return true
					}
				}
			}
		}
	}
	//然后纵向
	for i := 0; i < board_col; i++ {
		f_col := make([]int, board_row)
		for j := 0; j < board_row; j++ {
			if o.Board[j][i] == chess {
				if j == 0 {
					f_col[j] = 1
				} else {
					f_col[j] = f_col[j-1] + 1
					if f_col[j] == win {
						return true
					}
				}
			}
		}
	}
	//斜线1
	f1 := make([][]int, board_row)
	for i := 0; i < board_row; i++ {
		f1[i] = make([]int, board_col)
		for j := 0; j < board_col; j++ {
			if o.Board[i][j] == chess {
				if i == 0 || j == 0 {
					f1[i][j] = 1
				} else {
					f1[i][j] = f1[i-1][j-1] + 1
					if f1[i][j] == win {
						return true
					}
				}
			}
		}
	}
	//斜线2
	f2 := make([][]int, board_row)
	for i := 0; i < board_row; i++ {
		f2[i] = make([]int, board_col)
		for j := 0; j < board_col; j++ {
			if o.Board[i][j] == chess {
				if i == 0 || j == board_col-1 {
					f2[i][j] = 1
				} else {
					f2[i][j] = f2[i-1][j+1] + 1
					if f2[i][j] == win {
						return true
					}
				}
			}
		}
	}
	return false
}
func (o *ChessBoard) SetOneChess(row, col int, chess int) {
	o.Board[row][col] = chess
}

func (o *ChessBoard) randPoint() (row, col int) {
	rand.Seed(time.Now().Unix())
	randIdx := rand.Intn(len(o.Blanks))

	v := o.Blanks[randIdx]
	col = v / board_col
	row = v % board_col
	if randIdx == 0 {
		o.Blanks = o.Blanks[1:]
	} else if randIdx == len(o.Blanks) {
		o.Blanks = o.Blanks[:randIdx-1]
	} else {
		o.Blanks = append(o.Blanks[:randIdx], o.Blanks[randIdx+1:]...)
	}
	return
}
func (o *ChessBoard) next() bool {
	return len(o.Blanks) > 0
}
func (o *ChessBoard) print(p int) {
	switch p {
	case blank:
		fmt.Print(icon_blank)
	case white:
		fmt.Print(icon_white)
	case black:
		fmt.Print(icon_black)
	default:
		return
	}
}
func (o *ChessBoard) PrintM() {
	for _, col := range o.Board {
		for _, row := range col {
			o.print(row)
		}
		fmt.Println()
	}
}
