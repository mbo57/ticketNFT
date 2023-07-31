#go version
FROM golang:1.20-alpine3.17

#update & install (vimやbashを入れとくと便利）
RUN apk add --update && apk add git && apk add vim && apk add bash

# 作業ディレクトリ作成
RUN mkdir /go/src/app

# 作業ディレクトリ設定
WORKDIR /go/src/app

# ホットリロード用にairをインストール
RUN go install github.com/cosmtrek/air@latest

# airを実行
CMD ["air", "-c", ".air.toml"]
