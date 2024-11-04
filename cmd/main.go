// main.go
//
// PetPals 的主入口文件，负责初始化应用程序、窗口和宠物组件。
//
// 包含应用启动和窗口设置的核心逻辑。

package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/patient-fyd/PetPals/internal/pet"
	"github.com/patient-fyd/PetPals/internal/ui"
	"log"
)

func main() {
	myApp := app.NewWithID("com.example.desktoppet")
	myWindow := myApp.NewWindow("桌面宠物")

	// 初始化宠物
	desktopPet, err := pet.NewPet(myWindow)
	if err != nil {
		log.Fatalf("Failed to initialize pet: %v", err)
	}

	// 确保在程序退出时停止提醒
	defer desktopPet.StopReminders()

	// 设置窗口属性
	ui.SetupWindow(myWindow, desktopPet.CanvasObject())

	// 显示窗口并运行应用
	myWindow.ShowAndRun()

}
