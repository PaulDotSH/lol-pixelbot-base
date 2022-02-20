package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"strconv"
	"strings"
)

var enemyChampBmp = robotgo.OpenBitmap("./data/enemy.png")
var EnemyChampPos  = &[2]int{-1,-1}
var AttackSpeed = -1.0
var LastMousePos = &[2]int{-1,-1}
var Attacking = false

func GetAttackSpeed() (float64, error) {
	//Settings for 50% hud scale at 1080p
	cap := robotgo.CaptureScreen(486,1026,36,20)
	robotgo.SaveBitmap(cap,"./data/as.bmp")
	text, err := robotgo.GetText("./data/as.bmp")
	if err != nil {
		fmt.Println(text,err)
		return -1,err
	}
	//fmt.Println("OCR GOT",text)
	builder := strings.Builder{}
	for i:=0; i<len(text); i++ {
		//if text[i] ascii code is >= 48 which is one and <=57 which is 9 or it is a dot
		asciiCode:=int(text[i])
		if (asciiCode>47 && asciiCode<58) || asciiCode==46 {
			builder.WriteByte(text[i])
		}
	}
	return strconv.ParseFloat(builder.String(),64)
}

func GetEveryPixelWithColor(color uint32) []robotgo.Point {
	return robotgo.FindEveryColor(robotgo.UintToHex(color),robotgo.CaptureScreen())
}

//This gets it from top to down, make another that finds every enemy champ and gets the closest one to mouse or to screen center or to ally champ
func GetEnemyChampions() []robotgo.Point {
	return robotgo.FindEveryBitmap(enemyChampBmp)
}

func GetEnemyChampion() (int,int) {
	x,y := robotgo.FindBitmap(enemyChampBmp)
	//Approximate pos
	if x!=-1 {
		return x+70,y+110
	}
	return -1,-1
}