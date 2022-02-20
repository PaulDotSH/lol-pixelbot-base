package ss

import (
	"github.com/kbinani/screenshot"
	"testing"
)

// 0.04s/op on my system
func BenchmarkScreenshot(b *testing.B) {
	db:=screenshot.GetDisplayBounds(0)
	for i := 0; i < b.N; i++ {
		ScreenshotRect(db)
	}
}