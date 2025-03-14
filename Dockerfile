# ベースイメージとしてgolang:1.24-bullseyeを使用
FROM golang:1.24-bullseye

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY main.go .
COPY static/ ./static/
COPY templates/ ./templates/

# 必要なディレクトリ構造を作成
RUN mkdir -p /app/pictures

# Go アプリケーションをビルド
RUN go build -o photo-viewer main.go

# エクスポートするポートを指定
EXPOSE 8080

# アプリケーションを実行
CMD ["./photo-viewer"]
