// interaction.go
//
// 定义宠物的交互逻辑，包括左键、右键点击和鼠标悬停事件。
// 设置各种鼠标事件回调，丰富宠物的交互体验。

package pet

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// setupInteraction 为 Pet 组件设置交互事件
func (p *Pet) setupInteraction(win fyne.Window) {
	// 设置左键点击事件
	p.img.SetOnTapped(func() {
		// TODO: 实现宠物被点击时的反应，例如改变表情或显示信息
		println("宠物被左键点击")
	})

	// 设置右键点击事件
	p.img.SetOnTappedSecondary(func() {
		// 创建并显示右键菜单
		menu := fyne.NewMenu("",
			fyne.NewMenuItem("关于", func() {
				showAboutDialog(win)
			}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("退出", func() {
				win.Close()
			}),
		)

		// 显示右键菜单
		widget.ShowPopUpMenuAtPosition(menu, win.Canvas(), fyne.CurrentApp().Driver().AbsolutePositionForObject(p.img))
	})

	// 设置鼠标进入事件
	p.img.SetOnMouseIn(func() {
		println("鼠标进入宠物区域")
		// TODO: 实现鼠标悬停时的效果，例如高亮显示
	})

	// 设置鼠标离开事件
	p.img.SetOnMouseOut(func() {
		println("鼠标离开宠物区域")
		// TODO: 恢复鼠标离开后的效果
	})
}

// showAboutDialog 显示关于对话框
func showAboutDialog(win fyne.Window) {
	dialog.ShowInformation("关于", "桌面宠物应用\n版本：1.0.0\n作者：fyd", win)
}
