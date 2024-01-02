package main

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/useWhisper/util"
	"io"
	"log/slog"
	"os"
	"os/exec"
)

func init() {
	setLog()
}
func main() {
	root := util.GetVal("Whisper", "srt")
	level := util.GetVal("Whisper", "level")
	location := util.GetVal("Whisper", "location")
	language := util.GetVal("Whisper", "language")
	files := GetFileInfo.GetAllFileInfo(root, "mp4")
	for _, file := range files {
		slog.Info("文件", slog.String("文件名", file.FullPath))
		//whisper true.mp4 --model base --language English --model_dir /Users/zen/Whisper --output_format srt
		cmd := exec.Command("whisper", file.FullPath, "--model", level, "--model_dir", location, "--output_format", "srt", "--language", language, "--output_dir", root, "--verbose", "True")
		err := util.ExecCommand(cmd)
		if err != nil {
			slog.Error("当前字幕生成错误", slog.String("命令原文", fmt.Sprint(cmd)), slog.String("错误原文", err.Error()))
		}
	}
}

func setLog() {
	opt := slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}
	file := "Process.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0770)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}
