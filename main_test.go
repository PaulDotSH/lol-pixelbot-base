package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"testing"
)

//0.02-0.06 depends if pixels are found (bcs of adding them to the slice)
func BenchmarkGetAllPointsByColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		robotgo.FindEveryColor(robotgo.UintToHex(16777215),robotgo.CaptureScreen())
	}
	fmt.Println()
}

//0.22 seconds, oof
func BenchmarkGetAS(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		GetAttackSpeed()
	}
}

func BenchmarkGetEnemyChamps(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		GetEnemyChampions()
	}
}

