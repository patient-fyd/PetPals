// animation.go
//
// 提供动画帧加载与播放的功能，通过加载 assets 文件夹中的图片，
// 实现宠物的动画展示。

package pet

import (
	"fyne.io/fyne/v2"
	"os"
	"path/filepath"
	"sort"
)

// loadFrames 加载 assets 文件夹中的 PNG 文件作为动画帧
func loadFrames() ([]fyne.Resource, error) {
	var frames []fyne.Resource
	var paths []string

	// 遍历 assets 文件夹，收集 PNG 文件路径
	err := filepath.Walk("assets", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 遍历时出现错误
		}
		if !info.IsDir() && filepath.Ext(path) == ".png" {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// 对路径进行排序以确保帧顺序一致
	sort.Strings(paths)

	// 加载排序后的 PNG 资源
	for _, path := range paths {
		res, err := fyne.LoadResourceFromPath(path)
		if err != nil {
			return nil, err
		}
		frames = append(frames, res)
	}
	return frames, nil
}
