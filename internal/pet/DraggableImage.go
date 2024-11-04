// DraggableImage.go
//
// DraggableImage 提供了支持拖动、点击、悬停等事件的图片组件，
// 用于在应用中实现可交互的宠物形象。

package pet

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

// DraggableImage 是一个可以被拖动并支持鼠标事件的图片类型
type DraggableImage struct {
	*canvas.Image
	dragging          bool
	offset            fyne.Position
	onTapped          func() // 左键点击回调函数
	onTappedSecondary func() // 右键点击（双击）回调函数
	onMouseIn         func() // 鼠标进入回调函数
	onMouseOut        func() // 鼠标离开回调函数
}

// NewDraggableImage 创建一个新的 DraggableImage 对象
func NewDraggableImage(res fyne.Resource) *DraggableImage {
	img := &DraggableImage{
		Image:             canvas.NewImageFromResource(res),
		onTapped:          nil,
		onTappedSecondary: nil,
		onMouseIn:         nil,
		onMouseOut:        nil,
	}
	img.FillMode = canvas.ImageFillOriginal
	return img
}

// 设置点击事件的回调函数
func (img *DraggableImage) SetOnTapped(callback func()) {
	img.onTapped = callback
}

// 设置右键点击（双击）事件的回调函数
func (img *DraggableImage) SetOnTappedSecondary(callback func()) {
	img.onTappedSecondary = callback
}

// 设置鼠标进入事件的回调函数
func (img *DraggableImage) SetOnMouseIn(callback func()) {
	img.onMouseIn = callback
}

// 设置鼠标离开事件的回调函数
func (img *DraggableImage) SetOnMouseOut(callback func()) {
	img.onMouseOut = callback
}

// Dragged 实现拖拽事件
func (img *DraggableImage) Dragged(event *fyne.DragEvent) {
	if img.dragging { // 确保拖动只在 dragging 状态时生效
		img.Move(event.Position.Subtract(img.offset))
	}
}

// DragEnd 停止拖动
func (img *DraggableImage) DragEnd() {
	img.dragging = false
}

// MouseDown 实现 Mouseable 接口，处理鼠标按下事件
func (img *DraggableImage) MouseDown(event *desktop.MouseEvent) {
	img.offset = event.Position
	img.dragging = true
	if event.Button == desktop.MouseButtonPrimary && img.onTapped != nil {
		img.onTapped() // 触发左键点击事件
	} else if event.Button == desktop.MouseButtonSecondary && img.onTappedSecondary != nil {
		img.onTappedSecondary() // 触发右键点击（双击）事件
	}
}

// MouseUp 停止拖动
func (img *DraggableImage) MouseUp(event *desktop.MouseEvent) {
	img.dragging = false
}

// MouseIn 实现 Hoverable 接口，处理鼠标进入事件
func (img *DraggableImage) MouseIn() {
	if img.onMouseIn != nil {
		img.onMouseIn() // 触发鼠标进入事件
	}
}

// MouseOut 实现 Hoverable 接口，处理鼠标离开事件
func (img *DraggableImage) MouseOut() {
	if img.onMouseOut != nil {
		img.onMouseOut() // 触发鼠标离开事件
	}
}
