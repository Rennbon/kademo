package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type point byte

const (
	limit_row, limit_col       = 7, 6
	blank                point = 0x00
	white                point = 0x01
	black                point = 0x02

	icon_blank = "⬜️"
	icon_white = "⚪️"
	icon_black = "⚫️"

	str_white = "白棋"
	str_black = "黑棋"
)

var (
	Err_not_exist = errors.New("Coordinates do not exist.")

	win_line = []byte{0x0f, 0x1e, 0x3c, 0x78}
)

type Chesser interface {
	//打印
	PrintM()
	//落子
	SetOneChess(col, row uint8, chess point)
	//随机数
	randPoint() (col, row uint8)
	//剩余可落子数
	remain() int
	//计算胜利
	calcWin(p point) bool
	//清空棋子
	clear()
}
type ChessBoard struct {
	x      [limit_col][limit_row]point
	blanks []uint8
}

func newBoard() Chesser {
	cb := &ChessBoard{
		blanks: make([]uint8, limit_row*limit_col),
	}
	for i := uint8(0); i < limit_row*limit_col; i++ {
		cb.blanks[i] = i
	}
	return cb
}

func (o *ChessBoard) PrintM() {
	for _, col := range o.x {
		for _, row := range col {
			o.print(row)
		}
		fmt.Println()
	}
}
func (o *ChessBoard) print(p point) {
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
func (o *ChessBoard) clear() {
	o.x = [limit_col][limit_row]point{}
	o.blanks = make([]uint8, limit_row*limit_col)
	for i := uint8(0); i < limit_row*limit_col; i++ {
		o.blanks[i] = i
	}
}
func (o *ChessBoard) checkCoordinates(col, row uint8) error {
	if col > limit_col || row > limit_row {
		return Err_not_exist
	}
	return nil
}
func (o *ChessBoard) checkLineWin(bits [7]uint8) bool {
	sum := uint8(0)
	for i := 0; i < 7; i++ {
		if bits[i] == 1 {
			sum += 1 << uint(i)
		}
	}
	for _, v := range win_line {
		if v&sum == v {
			return true
		}
	}
	return false
}
func (o *ChessBoard) calcWin(p point) bool {
	// 0 1 1 1 1 1 1 1 -> byte
	//check line
	bits := [7]uint8{}
	// 6横6组数据 7有效长度
	//fmt.Println("---------横--------")
	for _, row := range o.x {
		for k, pi := range row {
			if pi == p {
				bits[k] = 1
			} else {
				bits[k] = 0
			}
		}
		if o.checkLineWin(bits) {
			fmt.Println("横线方向")
			return true
		}
	}
	bits = [7]uint8{}
	//7竖7组 6有效长度
	//fmt.Println("---------竖--------")
	for j := 0; j < limit_row; j++ {
		for i := 0; i < limit_col; i++ {
			if o.x[i][j] == p {
				bits[i] = 1
			} else {
				bits[i] = 0
			}
		}
		if o.checkLineWin(bits) {
			fmt.Println("竖线方向")
			return true
		}
	}
	//单子左下右下 "^"二叉树样串串，第一排穿光，然后两侧各串头尾"v"
	//fmt.Println("---------X--------")
	for turn := 0; turn < limit_col/2+1; turn++ {
		for i := 0; i < limit_row; i++ {
			bits1 := [7]uint8{}
			bits2 := [7]uint8{}
			m, n := 0, 0
			for j := turn; j < limit_col; j++ {
				if turn > 0 && i > 0 && i < limit_row-1 {
					continue
				}
				l := i + m
				r := i + n
				if l >= 0 && l < limit_row {
					if o.x[j][l] == p {
						bits1[j] = 1
					} else {
						bits1[j] = 0
					}
				}
				if r >= 0 && r < limit_row {
					if o.x[j][r] == p {
						bits2[j] = 1
					} else {
						bits2[j] = 0
					}
				}
				m += 1
				n -= 1
			}
			//fmt.Println("loc1:", bits1)
			//fmt.Println("loc2:", bits2)
			if o.checkLineWin(bits1) {
				fmt.Println("↘️")
				return true
			}
			if o.checkLineWin(bits2) {
				fmt.Println("↗️")
				return true
			}
		}
	}
	return false
}

func (o *ChessBoard) randPoint() (col, row uint8) {
	rand.Seed(time.Now().Unix())
	randIdx := uint8(rand.Intn(len(o.blanks)))

	v := o.blanks[randIdx]
	row = v / limit_col
	col = v % limit_col
	return
}
func (o *ChessBoard) remain() int {
	//fmt.Println(len(o.blanks))
	return len(o.blanks)
}
func (o *ChessBoard) removeByVal(val uint8) {
	curIdx := len(o.blanks)
	rmIdx := 0
	for k, v := range o.blanks {
		if v == val {
			rmIdx = k
			break
		}
	}
	if rmIdx == 0 {
		o.blanks = o.blanks[1:]
	} else if rmIdx == curIdx {
		o.blanks = o.blanks[:rmIdx-1]
	} else {
		o.blanks = append(o.blanks[:rmIdx], o.blanks[rmIdx+1:]...)
	}
}

func (o *ChessBoard) SetOneChess(col, row uint8, chess point) {
	if err := o.checkCoordinates(row, col); err != nil {
		fmt.Println(err)
		return
	}
	o.x[col][row] = chess
	k := row*limit_col + col

	o.removeByVal(k)
}
func main() {
	cb := newBoard()
	fmt.Println("-------1.1-------")
	cb.PrintM()
	fmt.Println("-------1.2-------")
	cb.SetOneChess(0, 3, black)
	cb.SetOneChess(3, 2, black)
	cb.SetOneChess(2, 3, white)
	cb.SetOneChess(3, 5, white)
	cb.PrintM()
	fmt.Println("-------2&3-------")
	cb.clear()
	first := true
	competitor := white
	compStr := str_white
	fg := true
	for i := 1; cb.remain() > 0; i++ {
		fmt.Printf("-----next:%d------\r\n", i)
		if first {
			competitor = white
		} else {
			competitor = black
		}
		col, row := cb.randPoint()
		//fmt.Println(col, row)
		cb.SetOneChess(col, row, competitor)
		first = first == false
		time.Sleep(time.Millisecond * 10)
		cb.PrintM()
		if cb.calcWin(competitor) {
			if competitor == black {
				compStr = str_black
			}
			fmt.Println(compStr, "胜利")
			fg = false
			break
		}
	}
	if fg {
		fmt.Println("平局")
	}
	cb.PrintM()
}
