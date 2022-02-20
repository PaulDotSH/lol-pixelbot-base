package main

//hExternal bot

import (
	"badPixelbotTry/scripts"
	"badPixelbotTry/settings"
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
	"os"
	"time"
)

//make a funtion that gets the distance, so you only cast spells on enemies in distance

func Abs(x int) int {
	if x < 0 { //Do some bitwise stuff maybe
		return x * -1
	} else {
		return x
	}
}

func init() {
	//Get enemy champ pos every x ms
	go func() {
		ticker := time.NewTicker(time.Millisecond * settings.GetChampionDelay)
		defer ticker.Stop()
		for {
			EnemyChampPos[0], EnemyChampPos[1] = GetEnemyChampion()
			<-ticker.C
		}
	}()

	//GetMousePos
	go func() {
		ticker := time.NewTicker(time.Millisecond * settings.GetMouseDelay)
		defer ticker.Stop()
		for {
			for Attacking {
				time.Sleep(15 * time.Millisecond)
			}
			LastMousePos[0], LastMousePos[1] = robotgo.GetMousePos()
			<-ticker.C
		}
	}()

	//Get AS every x ms
	//go func() {
	//	ticker := time.NewTicker(time.Millisecond*settings.GetASDelay)
	//	defer ticker.Stop()
	//	for {
	//		AttackSpeed,err:=GetAttackSpeed()
	//		if err == nil {
	//			fmt.Println("ATTACK SPEED:",AttackSpeed)
	//		} else {
	//			fmt.Println("Err as:",err)
	//		}
	//		<-ticker.C
	//	}
	//}()

	//RunScript
	go func() {
		ticker := time.NewTicker(time.Millisecond * settings.RunScriptDelay)
		defer ticker.Stop()
		for {
			//HERE
			//RunIfUnlocked(scripts.Orb,EnemyChampPos,&run,LastMousePos,&Attacking)
			RunIfUnlocked(scripts.Zilean, EnemyChampPos, &run, LastMousePos, &Attacking)
		}
	}()

	//Check Run
	//go func() {
	//	ticker := time.NewTicker(time.Second)
	//	defer ticker.Stop()
	//	for {
	//		if run==true {
	//			fmt.Println("RUN IS TRUE")
	//		} else {
	//			fmt.Println("RUN IS FALSE")
	//		}
	//		<-ticker.C
	//	}
	//}()
}

func main() {
	add()
}

var locked = false
var run = false

func add() {
	fmt.Println("hook add...")
	s := hook.Start()
	defer hook.End()

	for {
		i := <-s
		if i.Kind == hook.KeyHold {
			switch i.Rawcode {
			case settings.ComboKey:
				run = true

			}
		}

		if i.Kind == hook.KeyUp {
			switch i.Rawcode {
			case settings.ComboKey:
				run = false
			}
		}

		if i.Kind == hook.KeyDown {
			switch i.Rawcode {
			//TODO: make it change "scripts" called somehow
			case settings.ComboKey:
				robotgo.KeyDown("c")
			case settings.ExitKey:
				fmt.Println("Exiting")
				os.Exit(0)
			//TODO: Check somehow if it already ran
			case settings.RangeKey:
				go robotgo.KeyDown("c")
			case settings.OpenMenuKey:
				fmt.Println("Should open menu")
				return
			case settings.DebugKey:
				go func() {
					if EnemyChampPos[0] != -1 {
						robotgo.Move(EnemyChampPos[0], EnemyChampPos[1])
						fmt.Println(EnemyChampPos[0], EnemyChampPos[1])
					}
				}()
				fmt.Println("Enemy champ pos:", EnemyChampPos)
			}
		}

		if i.Kind != hook.KeyDown {
			continue
		}
		fmt.Println(i.Kind, i.When, i.Keycode, i.Keychar)

	}
}

func RunIfUnlocked(f func(*[2]int, *bool, *[2]int, *bool), enemyChampPos *[2]int, run *bool, lastMousePos *[2]int, attacking *bool) {
	if locked == true {
		return
	}
	locked = true
	f(enemyChampPos, run, lastMousePos, attacking)
	locked = false
}
