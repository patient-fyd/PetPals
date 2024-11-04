//go:build windows

//
// 提供 Windows 平台的透明窗口设置。
// 使用 Windows API 设置窗口透明效果。

package ui

import (
	"fyne.io/fyne/v2"
	"syscall"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	getConsoleWindow     = user32.NewProc("GetConsoleWindow")
	setWindowLongPtr     = user32.NewProc("SetWindowLongPtrW")
	setLayeredWindowAttr = user32.NewProc("SetLayeredWindowAttributes")
)

const (
	GWL_EXSTYLE   = -20
	WS_EX_LAYERED = 0x80000
	LWA_ALPHA     = 0x02
)

func makeWindowTransparent(win fyne.Window) {
	hwnd, _, _ := getConsoleWindow.Call()
	setWindowLongPtr.Call(hwnd, GWL_EXSTYLE, uintptr(WS_EX_LAYERED))
	setLayeredWindowAttr.Call(hwnd, 0, 200, LWA_ALPHA)
}
