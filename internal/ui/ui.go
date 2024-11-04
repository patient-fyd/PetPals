// ui.go
//
// 定义 UI 窗口设置和通用对话框显示功能。
// 包含窗口内容设置和关于对话框显示的逻辑。

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// SetupWindow 设置窗口内容和透明背景
func SetupWindow(win fyne.Window, content fyne.CanvasObject) {
	// 设置窗口内容
	win.SetContent(content)

	// 调用平台特定的透明设置函数
	makeWindowTransparent()
}

// showAboutDialog 显示关于对话框
func showAboutDialog(win fyne.Window) {
	dialog.ShowInformation("关于", "桌面宠物应用\n版本：1.0.0\n作者：fyd", win)
}
