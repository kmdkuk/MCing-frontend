# ビルドステージ
FROM golang:1.24 as builder

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールと依存関係をコピー
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションコードをコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mcing-backend .


# 実行ステージ
FROM gcr.io/distroless/base-debian11

# 作業ディレクトリを設定
WORKDIR /app

# ビルドしたバイナリをコピー
COPY --from=builder /app/mcing-backend .

# ポートを公開
EXPOSE 80

# アプリケーションを実行
CMD ["./mcing-backend"]
