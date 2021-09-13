package main

import (
	"math"
)

type Robot struct {
	room [][]int
	positionX int
	positionY int
	direction int
}

func (r Robot) move() bool {
	switch r.direction {
	case 0:
		if len(r.room[r.positionX]) >= r.positionY + 1 && r.room[r.positionX][r.positionY + 1] != 0 {
			r.positionY++
			return true
		} else {
			return false
		}
	case 90:
		if r.positionX != 0 && r.room[r.positionX - 1][r.positionY] != 0 {
			r.positionX--
			return true
		} else {
			return false
		}
	case 270:
		if len(r.room) >= r.positionX + 1 && r.room[r.positionX + 1][r.positionY] != 0 {
			r.positionX++
			return true
		} else {
			return false
		}
	case 180:
		if r.positionY != 0 && r.room[r.positionX][r.positionY - 1] != 0 {
			r.positionY--
			return true
		} else {
			return false
		}
	default:
		panic(nil)
	}
}

func (r Robot) turnLeft() {
	r.direction = int(math.Min(0, float64(r.direction - 90)))
}

func (r Robot) turnRight() {
	r.direction = int((r.direction + 90) % 360)
}

func (r Robot) clean() {

}

//room = [
//  [1,1,1,1,1,0,1,1],
//  [1,1,1,1,1,0,1,1],
//  [1,0,1,1,1,1,1,1],
//  [0,0,0,1,0,0,0,0],
//  [1,1,1,1,1,1,1,1]
//],
// row = 1, column = 3

