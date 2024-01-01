# 基础镜像
# docker run -itd name=test golang:1.21.5-alpine3.18 ash
FROM golang:1.21.5-bookworm
# go env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 备份原始安装源
RUN cp /etc/apt/sources.list.d/debian.sources /etc/apt/sources.list.d/debian.sources.bak
# 修改为国内源
#RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources
RUN apk update
RUN apk upgrade
# 安装基础软件
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o translate main.go
RUN apt install -y sqlite3 bash git build-essential python3 python3-pip
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
RUN pip install -U openai-whisper --break-system-packages --root-user-action
# 准备软件
CMD ["bash"]