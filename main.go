package main

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/useWhisper/sql"
	"github.com/zhangyiming748/useWhisper/util"
	"io"
	"log/slog"
	"os"
	"os/exec"
)

func init() {
	setLog()
	sql.SetEngine()
}
func main() {
	util.ExitAfterRun()
	root := util.GetVal("Whisper", "srt")
	level := util.GetVal("Whisper", "level")
	location := util.GetVal("Whisper", "location")
	language := util.GetVal("Whisper", "language")
	files := GetFileInfo.GetAllFileInfo(root, "mp4")
	for _, file := range files {
		slog.Info("文件", slog.String("文件名", file.FullPath))
		//whisper true.mp4 --model base --language English --model_dir /Users/zen/Whisper --output_format srt
		//cmd := exec.Command("whisper", file.FullPath, "--model", level, "--model_dir", location, "--language", language, "--output_dir", root, "--verbose", "True")
		cmd := exec.Command("whisper", file.FullPath, "--model", level, "--model_dir", location, "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", language, "--output_dir", root, "--verbose", "True")
		err := util.ExecCommand(cmd)
		if err != nil {
			slog.Error("当前字幕生成错误", slog.String("命令原文", fmt.Sprint(cmd)), slog.String("错误原文", err.Error()))
		}
	}
	in()
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

func in() {
	dir := util.GetVal("Whisper", "srt")
	files := GetFileInfo.GetAllFileInfo(dir, "srt")
	for _, file := range files {
		lines := util.ReadByLine(file.FullPath)
		for i := 0; i < len(lines); i += 4 {
			if i+3 > len(lines) {
				continue
			}
			w := new(sql.Whisper)
			w.Name = file.PurgeName
			w.No = lines[i]
			w.Duration = lines[i+1]
			w.Srt = lines[i+2]
			w.SetOne()
			slog.Info("入库前", slog.Any("结构体", w))
		}
	}
}
