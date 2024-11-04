// pet.go
//
// 定义 Pet 结构体和宠物的初始化方法。包含宠物的动画、交互、
// 提醒和语音功能的初始化逻辑。

package pet

import (
	"fyne.io/fyne/v2"
	"time"
)

type Pet struct {
	img              *DraggableImage
	frames           []fyne.Resource
	position         fyne.Position
	animationRun     bool // 控制动画运行的变量
	reminderTicker   *time.Ticker
	reminderInterval time.Duration
	// TODO: 添加更多属性，例如状态、计时器等
}

func NewPet(win fyne.Window) (*Pet, error) {
	// 加载动画帧
	frames, err := loadFrames()
	if err != nil || len(frames) == 0 { // 检查 frames 是否非空
		return nil, err
	}

	// 创建可拖拽的图片对象
	img := NewDraggableImage(frames[0])

	pet := &Pet{
		img:          img,
		frames:       frames,
		animationRun: true,
	}

	// 启动动画
	pet.startAnimation()

	// 启动交互功能
	pet.setupInteraction(win)

	// TODO: 启动提醒功能
	// TODO: 启动语音支持

	return pet, nil
}

func (p *Pet) CanvasObject() fyne.CanvasObject {
	return p.img
}

func (p *Pet) startAnimation() {
	go func() {
		for p.animationRun { // 检查动画是否应该继续运行
			for _, res := range p.frames {
				if !p.animationRun {
					break
				}
				p.img.Resource = res
				p.img.Refresh()
				time.Sleep(100 * time.Millisecond) // 控制动画速度
			}
		}
	}()
}

// 停止动画的函数，可以在应用关闭或切换时调用
func (p *Pet) stopAnimation() {
	p.animationRun = false
}
