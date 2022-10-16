FROM golang:1.18-bullseye

WORKDIR /coreapp
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
# 在里面在build出一个二进制文件出来
RUN go build -o ./coreapp
EXPOSE 9000
CMD ["./coreapp"]