//go:build darwin

//
// 提供 MacOS 平台的透明窗口设置。
// 使用 Objective-C 和 Cocoa 实现 MacOS 透明效果。

package ui

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#include <Cocoa/Cocoa.h>

// 设置当前应用程序的第一个窗口为透明
void MakeWindowTransparent() {
    NSApplication *app = [NSApplication sharedApplication];
    NSArray *windows = [app windows];
    if ([windows count] > 0) {
        NSWindow *window = [windows objectAtIndex:0];
        [window setOpaque:NO];
        [window setBackgroundColor:[NSColor clearColor]];
    }
}
*/
import "C"

// makeWindowTransparent 通过 Cocoa API 设置窗口透明
func makeWindowTransparent() {
	C.MakeWindowTransparent()
}
