# ビルド用ステージ
FROM golang:latest AS builder

# 作業ディレクトリを指定
WORKDIR /app

# ソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# 実行用ステージ
FROM alpine:latest

# ポートを公開
EXPOSE 8080

# ビルド用ステージからバイナリをコピー
COPY --from=builder /app/myapp .

# コンテナ内で実行するコマンドを指定
CMD ["./myapp"]
