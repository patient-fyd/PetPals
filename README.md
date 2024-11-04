# Desktop Pet - PatPals

**桌面宠物应用程序**  
使用Go语言和Fyne库开发的桌面宠物应用程序。这个项目创建了一个可爱的桌面小助手，它可以与你互动，添加乐趣并为你提供各种提醒服务。

> 本人毕设计划用fyne写，这是个练手项目，然后目前还没空（准备考研，先创建项目，有空就写写）

## 项目概述

**PatPals** 是一个桌面宠物应用，旨在提供简单的互动功能和丰富的视觉效果。宠物可以在桌面上随意拖动，支持动画、透明背景和置顶显示，未来还可以添加更多个性化互动功能。



## 功能特点

- **动画支持**：通过加载多个帧图像来实现宠物的动画效果。
- **拖拽功能**：可以使用鼠标随意拖动宠物。
- **透明背景**：支持透明背景（在不同系统上可能效果有所差异）。
- **窗口置顶**：始终保持在其他窗口之上，方便你随时与宠物互动。
- **右键菜单**：提供快捷菜单以便快速退出或扩展其他功能。



## 环境依赖

- **Go 1.18或更高版本**
- **Fyne库**：用于UI和窗口管理


## 使用说明

### 启动宠物

运行`main.go`文件后，桌面宠物窗口将会显示在屏幕上，默认无边框、透明背景并始终保持在最前。

### 动画效果

桌面宠物会自动播放动画，通过不断切换图片帧实现动态效果。默认帧图像存储在`assets/`目录下。

### 拖拽功能

在宠物的窗口上按住鼠标左键即可拖动宠物到任意位置。

### 右键菜单

右键点击宠物可以打开快捷菜单，目前提供退出选项，未来可扩展更多功能。



## 项目结构

```plaintext
PatPals/
├── assets/               # 动画帧图像目录
│   ├── frame1.png
│   ├── frame2.png
│   └── frame3.png
├── cmd/
│   └── main.go                    # 项目入口，启动应用并初始化窗口和宠物
├── internal/
│   ├── pet/
│   │   ├── DraggableImage.go      # 支持拖动和鼠标事件的图片组件
│   │   ├── animation.go           # 宠物动画帧的加载与播放
│   │   ├── interaction.go         # 宠物交互逻辑（点击、悬停等）
│   │   ├── pet.go                 # Pet 结构体定义与初始化
│   │   ├── reminders.go           # 定时提醒功能
│   │   └── speech.go              # 宠物的语音提示功能
│   └── ui/
│       ├── ui.go                  # 窗口设置和通用 UI 组件（关于对话框）
│       ├── window_darwin.go       # MacOS 平台特定的透明窗口设置
│       ├── window_windows.go      # Windows 平台特定的透明窗口设置
│       └── window_linux.go        # Linux 平台特定的透明窗口设置
└── pkg/
│   └── utils.go                   # 常用的对话框显示封装函数
├── go.mod                # Go模块配置文件
└── main.go               # 主程序文件
```


## 项目架构图

```mermaid
graph TD

%% Main entry point
A[main.go:main] -->|Initialize App & Window| B[NewPet]
A -->|Setup Window| C[SetupWindow]
A -->|Defer Cleanup| J[StopReminders]

%% Pet initialization and functionality
B -->|Load Frames| D[loadFrames]
B -->|Create Draggable Image| E[NewDraggableImage]
B -->|Start Animation| F[startAnimation]
B -->|Setup Interaction| G[setupInteraction]
B -->|Start Reminders| H[startReminders]
B -->|Enable Speech| I[speak]

%% Draggable Image functionality
E -->|Setup Mouse Events| G1[SetOnTapped]
E --> G2[SetOnTappedSecondary]
E --> G3[SetOnMouseIn]
E --> G4[SetOnMouseOut]

%% Animation Loop
F -->|Update Image Resource| E
F -->|Loop Through Frames| F

%% Reminder System
H -->|Trigger Reminder on Interval| K[showReminder]

%% UI Setup and Dialog
C -->|Make Window Transparent| L[makeWindowTransparent]
C -->|About Dialog| M[showAboutDialog]

%% Utility Functions
M -->|Show Information Dialog| N[ShowMessage]

%% Styling nodes
style A fill:#f9f,stroke:#333,stroke-width:4px
style B fill:#ccf,stroke:#333,stroke-width:2px
style C fill:#ccf,stroke:#333,stroke-width:2px
style E fill:#cfc,stroke:#333,stroke-width:2px
style F fill:#cfc,stroke:#333,stroke-width:2px
style H fill:#cfc,stroke:#333,stroke-width:2px
style I fill:#cfc,stroke:#333,stroke-width:2px
style K fill:#fcc,stroke:#333,stroke-width:2px
style M fill:#ccf,stroke:#333,stroke-width:2px
```

## 未来计划

- **更多互动功能**：添加宠物的点击反馈、表情变化等。
- **定时提醒**：提供休息提醒、喝水提醒等健康小助手功能。
- **语音支持**：通过语音反馈或对话功能，增强互动体验。
- **个性化皮肤**：支持不同宠物皮肤，提供多样化选择。


## 贡献

欢迎为本项目贡献代码！如果你有任何改进建议或发现问题，请在GitHub提交`issue`或`pull request`。



## 许可证

本项目采用 [MIT License](LICENSE) 许可证。


## 致谢

感谢 [Fyne](https://fyne.io) 提供的强大UI库，使桌面宠物应用的开发变得更加简单高效。



希望你能喜欢这个桌面宠物项目！ 😊 如果有任何问题，请随时联系我！
