# 基础镜像
# docker pull golang:1.21.5-bookworm
# docker run -idt --name=test golang:1.22.0-bookworm bash
FROM golang:1.22.0-bookworm
# go env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 备份原始安装源ls

RUN cp /etc/apt/sources.list.d/debian.sources /etc/apt/sources.list.d/debian.sources.bak
# 修改为国内源
RUN sed -i 's/deb.debian.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources


# 重试
RUN echo 'APT::Acquire::Retries "10";' > /etc/apt/apt.conf.d/80-retries
RUN apt update
RUN apt full-upgrade -y

# 安装基础软件
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o translate main.go
RUN apt install -y sqlite3 git build-essential python3 python3-pip vim nano less locales ca-certificates
# 本地化
RUN rm /etc/localtime
RUN ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN locale-gen zh_CN.utf8
RUN echo "export LC_ALL=zh_CN.UTF-8">> /etc/profile
# 设置python
RUN pip3 config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
RUN pip3 install -U openai-whisper --break-system-packages
RUN git clone https://github.com/zhangyiming748/useWhisper.git /root/useWhisper
WORKDIR /root/useWhisper
git checkout v0.0.2
# git checkout -b branch_name tag_name
# 准备软件
CMD ["bash"]
#docker build -t trans:v1 .
#docker run -itd --name=trans1 -v /d/srt:/srt -e APPID=20230419001647901 -e KEY=rsNTVeBCtQ1sF7RSmFVq trans:v1 ash
#docker run -idt --name=trans -v /d/srt:/srt -e APPID={your baidu appid} -e KEY={your baidu key} trans:v1 ash
