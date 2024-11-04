// reminders.go
//
// 定义定时提醒功能，使用 Ticker 定时触发提醒事件。
// 提供 startReminders 和 stopReminders 控制提醒的启动与停止。

package pet

import (
	"time"
)

// startReminders 启动定时提醒
func (p *Pet) startReminders() {
	if p.reminderInterval == 0 {
		p.reminderInterval = 1 * time.Hour // 默认提醒间隔为1小时
	}

	p.reminderTicker = time.NewTicker(p.reminderInterval)
	go func() {
		for {
			select {
			case <-p.reminderTicker.C:
				p.showReminder()
			}
		}
	}()
}

// stopReminders 停止定时提醒
func (p *Pet) StopReminders() {
	if p.reminderTicker != nil {
		p.reminderTicker.Stop()
	}
}

// showReminder 显示提醒信息
func (p *Pet) showReminder() {
	// TODO: 实现提醒显示逻辑，例如弹出气泡或改变宠物表情
}
