package scripts

import (
	"context"
	"github.com/go-vgo/robotgo"
	"time"
)

func Zilean(enemyChampPos *[2]int, run *bool, lastMousePos *[2]int, attacking *bool) {
	if *run == false {
		return
	}
	if enemyChampPos[0] == -1 {
		*attacking = true
		robotgo.Click("right")
		*attacking = false
		time.Sleep(200 * time.Millisecond)
		return
	}
	*attacking = true
	robotgo.Move(enemyChampPos[0], enemyChampPos[1])
	robotgo.KeyTap("q")
	robotgo.KeyTap("q")
	time.Sleep(290 * time.Millisecond)
	robotgo.KeyTap("w")
	time.Sleep(80 * time.Millisecond)
	robotgo.KeyTap("q")
	robotgo.Move(lastMousePos[0], lastMousePos[1])
	*attacking = false

	time.Sleep(1 * time.Second)
}

func Rengar() {
	robotgo.KeyTap("q")
	time.Sleep(10 * time.Millisecond)
	robotgo.KeyTap("e")
	time.Sleep(10 * time.Millisecond)
	robotgo.KeyTap("w")
	time.Sleep(350 * time.Millisecond)
	robotgo.KeyTap("q")
}

func Xerath(enemyChampPos *[2]int, run *bool, lastMousePos *[2]int, attacking *bool) {
	if *run == false {
		return
	}
	if enemyChampPos[0] == -1 {
		*attacking = true
		robotgo.Click("right")
		*attacking = false
		time.Sleep(200 * time.Millisecond)
		return
	}
	*attacking = true
	robotgo.Move(enemyChampPos[0], enemyChampPos[1])
	robotgo.KeyUp("q")
	time.Sleep(50 * time.Millisecond)
	robotgo.KeyTap("w")
	time.Sleep(50 * time.Millisecond)
	robotgo.KeyTap("e")
	time.Sleep(16 * time.Millisecond)
	robotgo.Move(lastMousePos[0], lastMousePos[1])
	*attacking = false
	time.Sleep(1 * time.Second)
}

func Spin4Fun(enemyChampPos *[2]int, run *bool, lastMousePos *[2]int, attacking *bool) {
	if *run == false {
		return
	}
	robotgo.Click("right")
	time.Sleep(80 * time.Millisecond)
}

//Do AS calculations

func Orb(enemyChampPos *[2]int, run *bool, lastMousePos *[2]int, attacking *bool) {
	if *run == false {
		return
	}
	if enemyChampPos[0] == -1 {
		*attacking = true
		robotgo.Click("right")
		*attacking = false
		time.Sleep(200 * time.Millisecond)
		return
	}
	*attacking = true
	robotgo.Move(enemyChampPos[0], enemyChampPos[1])
	robotgo.Click("right")
	time.Sleep(270 * time.Millisecond) //Time it takes get mouse back (cancel auto)

	*attacking = false
	robotgo.Move(lastMousePos[0], lastMousePos[1])

	robotgo.Click("right")
	time.Sleep(270 * time.Millisecond) //move

	return
}

//Vayne
//0.65 -
//2.5 270,270

//ToFIX
func RunFuncTimeout(f func(str string), duration time.Duration, str string) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			f(str)
		}
	}

}
