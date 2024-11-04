//go:build linux

//
// 提供 Linux 平台的透明窗口设置。
// 使用 wmctrl 工具设置窗口置顶或透明。

package ui

import (
	"fyne.io/fyne/v2"
	"os/exec"
)

func makeWindowTransparent(win fyne.Window) {
	cmd := exec.Command("wmctrl", "-r", ":ACTIVE:", "-b", "add,above")
	cmd.Run()
}
