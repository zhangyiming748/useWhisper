# 基础镜像
# docker pull golang:1.21.5-bookworm
# docker run -idt --name=test golang:1.22.0-bookworm bash
FROM zhangyiming748/use-whisper:v0.0.3
VOLUME /srt
WORKDIR /root/useWhisper
ENTRYPOINT ["go", "run", "main.go"]
#docker build -t trans:v1 .
#docker run -itd --name=trans1 -v /d/srt:/srt zhangyiming748/use-whisper:v0.0.3 bash
#docker run -idt --name=trans -v /d/srt:/srt -e APPID={your baidu appid} -e KEY={your baidu key} trans:v1 ash
# docker run -itd --name=trans1 -v /d/srt:/srt use-whisper:v0.0.3