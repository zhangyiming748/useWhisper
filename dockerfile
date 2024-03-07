# 基础镜像
# docker pull golang:1.21.5-bookworm
# docker run -dit --name=test usewhisper:v0.0.1 -v /mnt/e/video/srt:/srt -e language=en bash
FROM usewhisper:v0.0.1
VOLUME /srt
WORKDIR /root/useWhisper
CMD ["srt"]
# docker build -t trans:v1 .
# docker run -itd --name=trans1 -v /d/srt:/srt zhangyiming748/use-whisper:v0.0.3 bash
# docker run -idt --name=trans -v /d/srt:/srt -e APPID={your baidu appid} -e KEY={your baidu key} trans:v1 ash
# docker run -itd --name=trans1 -v /d/srt:/srt use-whisper:v0.0.3