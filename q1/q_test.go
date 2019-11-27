package main

import "testing"

var a *ChessBoard

func TestMain(m *testing.M) {
	a = new(ChessBoard)
	a.x = [limit_col][limit_row]point{}

}
func Test_checkLineWin(t *testing.T) {
	r := [7]uint8{1, 1, 0, 0, 1, 1, 1}
	t.Log(a.checkLineWin(r))
}

func Test_calcWin(t *testing.T) {

}
