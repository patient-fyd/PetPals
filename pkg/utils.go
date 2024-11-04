// utils.go
//
// 提供通用的对话框显示封装函数，包括信息对话框、
// 错误对话框和确认对话框等功能。

package pkg

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// ShowMessage 显示信息对话框
func ShowMessage(win fyne.Window, title, message string) {
	dialog.ShowInformation(title, message, win)
}

// ShowError 显示错误信息对话框
func ShowError(win fyne.Window, title, message string) {
	dialog.ShowError(fmt.Errorf(message), win)
}

// ShowConfirmation 显示确认对话框
func ShowConfirmation(win fyne.Window, title, message string, onConfirm func(bool)) {
	dialog.ShowConfirm(title, message, onConfirm, win)
}
