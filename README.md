# Who Was That - Server

Go + Echo + AWS Lambda を使用したバックエンド API サーバー

## 🚀 技術スタック

### コア技術

- **言語**: Go 1.24.0
- **Web フレームワーク**: Echo v4.13.4
- **デプロイメント**: AWS Lambda (+ ローカル開発)
- **データベース**: PostgreSQL 15
- **キャッシュ**: Redis 7
- **コンテナ化**: Docker & Docker Compose

### 開発ツール

- **ホットリロード**: Air (cosmtrek/air)
- **データベース管理**: Adminer
- **環境管理**: Docker Compose
- **依存関係管理**: Go Modules

## 🔧 開発環境セットアップ

### 前提条件

- Go 1.24.0 以上
- Docker & Docker Compose
- Air (ホットリロード用)

### 初回セットアップ

1. **リポジトリのクローン**

   ```bash
   git clone git@github.com:hiroki1238/who-was-that-server.git
   cd who-was-that-server
   ```

2. **Air のインストール（ホットリロード用）**

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

3. **環境変数の設定**

   ```bash
   cp .env.local.example .env.local
   # 必要に応じて環境変数を編集
   ```

4. **依存関係のインストール**
   ```bash
   go mod tidy
   ```

### 🚀 開発サーバーの起動

#### 方法 1: Docker Compose を使用（推奨）

```bash
# データベース、Redis、Adminer を起動
docker-compose up -d db redis adminer

# Goアプリケーションをホットリロードで起動
air
```

#### 方法 2: 完全に Docker Compose で起動

```bash
# すべてのサービスを起動
docker-compose up
```

#### 方法 3: ローカル直接起動

```bash
# データベースとRedisのみDocker で起動
docker-compose up -d db redis

# Goアプリケーションをローカル実行
go run main.go
```

### 🌐 アクセス先

| サービス     | URL                          | 説明             |
| ------------ | ---------------------------- | ---------------- |
| API Server   | http://localhost:8080        | メイン API       |
| Health Check | http://localhost:8080/health | ヘルスチェック   |
| Adminer      | http://localhost:8081        | データベース管理 |
| PostgreSQL   | localhost:5432               | データベース     |
| Redis        | localhost:6379               | キャッシュ       |

### 🔍 データベース接続情報

```
Host: localhost
Port: 5432
Database: who_was_that_db
Username: postgres
Password: password
```

## 📁 プロジェクト構成

```
.
├── .air.toml              # Air設定ファイル
├── .cursorrules           # Cursor エディタ設定
├── .env.local             # 環境変数（開発用）
├── docker-compose.yml     # Docker Compose設定
├── Dockerfile.dev         # 開発用Dockerファイル
├── go.mod                 # Go依存関係
├── go.sum                 # Go依存関係のチェックサム
├── main.go                # メインアプリケーション
├── CLAUDE.md              # Claude AI開発ガイド
├── README.md              # このファイル
├── internal/              # 内部パッケージ
│   └── config/           # 設定関連
└── scripts/              # スクリプト
    └── init.sql          # DB初期化スクリプト
```

## 🔗 API エンドポイント

### ヘルスチェック

```bash
GET /health
```

### 認証

```bash
POST /auth/login      # ログイン（実装予定）
POST /auth/register   # 新規登録（実装予定）
```

### API

```bash
GET /api/hello       # サンプルエンドポイント
```

## 🛠️ 開発コマンド

### 基本コマンド

```bash
# 開発サーバー起動（ホットリロード）
air

# 通常のGo実行
go run main.go

# ビルド
go build -o bin/main .

# テスト実行
go test ./...

# 依存関係の更新
go mod tidy
```

### Docker コマンド

```bash
# 全サービス起動
docker-compose up

# データベースのみ起動
docker-compose up -d db

# ログ確認
docker-compose logs -f

# サービス停止
docker-compose down

# データベースリセット
docker-compose down -v
```

## 🔐 環境変数

### 必要な環境変数

```env
# データベース設定
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=who_was_that_db

# Redis設定
REDIS_HOST=localhost
REDIS_PORT=6379

# AWS設定（本番環境）
AWS_REGION=ap-northeast-1
AWS_LAMBDA_FUNCTION_NAME=who-was-that-api
```

## 🚀 デプロイメント

### AWS Lambda デプロイ

1. **ビルド**

   ```bash
   GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
   ```

2. **Lambda 関数の作成**
   ```bash
   # terraform または AWS CLI を使用
   # 詳細は who-was-that-infra リポジトリを参照
   ```

### ローカル Lambda テスト

```bash
# Lambda環境変数を設定
export AWS_LAMBDA_FUNCTION_NAME=test

# アプリケーション起動
go run main.go
```

## 📊 監視・デバッグ

### データベース管理

- Adminer: http://localhost:8081
- 接続情報: postgres/password@localhost:5432/who_was_that_db

### ログ

```bash
# アプリケーションログ
docker-compose logs -f app

# データベースログ
docker-compose logs -f db

# すべてのログ
docker-compose logs -f
```

### パフォーマンス監視

```bash
# メモリ使用量
docker stats

# CPU使用量
top -p $(pgrep main)
```

## 🧪 テスト

### 単体テスト

```bash
go test ./...
```

### API テスト

```bash
# ヘルスチェック
curl http://localhost:8080/health

# サンプルAPI
curl http://localhost:8080/api/hello
```

## 🤝 開発ガイドライン

- `.cursorrules` ファイルに従ってコーディングを行う
- `CLAUDE.md` を参照して Claude AI との協働を行う
- Go 言語のベストプラクティスに従う
- 適切なエラーハンドリングを実装する
- 単体テストを書く

## 📚 関連リポジトリ

- **フロントエンド**: [who-was-that-client](https://github.com/hiroki1238/who-was-that-client)
- **インフラ**: [who-was-that-infra](https://github.com/hiroki1238/who-was-that-infra)

## 🐛 トラブルシューティング

### よくある問題

1. **データベースに接続できない**

   ```bash
   # Dockerコンテナが起動しているか確認
   docker-compose ps
   ```

2. **ホットリロードが動作しない**

   ```bash
   # Air を再インストール
   go install github.com/cosmtrek/air@latest
   ```

3. **ポートが使用中**
   ```bash
   # 使用中のプロセスを確認
   lsof -i :8080
   ```

## 📝 ライセンス

MIT License

## 👥 貢献

プルリクエストやイシューの作成を歓迎します。
