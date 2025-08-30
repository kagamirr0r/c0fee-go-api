# C0fee API

コーヒー豆管理 API - Go 製 RESTful API サーバー

## アーキテクチャ

### レイヤー構成

```
Controller Layer (DTO)
    ↓
UseCase Layer (Entity)
    ↓
Repository Layer (Model)
    ↓
Database
```

### 各レイヤーの責務

#### Controller Layer

- **扱うデータ**: DTO (Data Transfer Object)
- **責務**: HTTP リクエスト/レスポンスの処理、バリデーション、ルーティング
- **場所**: `/controller`
- **変換**: DTO ↔ Entity (UseCase 層とのやり取り)

#### UseCase Layer

- **扱うデータ**: Entity (Domain Entity)
- **責務**: ビジネスロジック、ドメインルールの実装
- **場所**: `/usecase`
- **変換**: Entity ↔ Model (Repository 層とのやり取り)
- **特徴**: model パッケージを直接参照しない

#### Domain Layer

- **扱うデータ**: Entity (Domain Entity) & Repository Interface
- **責務**: ドメインモデルの定義、ビジネスルールの表現、インターフェースの定義
- **場所**: `/domain`
- **構成**:
  - `/domain/entity/` - ドメインエンティティ（ビジネスオブジェクトの定義）
  - `/domain/repository/` - リポジトリインターフェース（データアクセス抽象化）
- **特徴**:
  - 他のレイヤーに依存しない（依存関係の逆転）
  - ビジネスロジックの中核を担う
  - インフラストラクチャ層の実装詳細から独立

#### Repository Layer

- **扱うデータ**: Model (Database Model)
- **責務**: データベースへのアクセス、永続化
- **場所**: `/repository`
- **変換**: Model ↔ Entity (UseCase 層とのやり取り)

### データフロー

```
HTTP Request (JSON)
    ↓
Controller: JSON → DTO
    ↓
UseCase: DTO → Entity → ビジネスロジック → Entity
    ↓
Repository: Entity → Model → Database操作 → Model → Entity
    ↓
UseCase: Entity → DTO
    ↓
Controller: DTO → JSON
    ↓
HTTP Response (JSON)
```

### コンバーター構成

```
/common/converter/
├── entity_model/        # Repository層で使用 (Entity ↔ Model)
│   ├── bean_converter.go
│   ├── user_converter.go
│   └── ...
├── dto_entity/          # UseCase層で使用 (DTO ↔ Entity)
│   ├── bean_converter.go
│   ├── user_converter.go
│   └── ...
└── model_dto/           # 将来の拡張用（現在は未使用）
```

## 技術スタック

- **言語**: Go 1.22.2
- **フレームワーク**: Echo v4.12.0
- **データベース**: PostgreSQL
- **ORM**: GORM
- **ストレージ**: MinIO/S3 互換
- **アーキテクチャ**: Clean Architecture

## プロジェクト構成

```
c0fee-api/
├── controller/          # Controller Layer (DTO)
├── usecase/            # UseCase Layer (Entity)
├── repository/         # Repository Layer (Model)
├── domain/             # Domain Layer (Core Business Logic)
│   ├── entity/         # Domain Entities
│   │   ├── bean.go     # Bean エンティティ
│   │   ├── user.go     # User エンティティ
│   │   ├── country.go  # Country エンティティ
│   │   └── ...         # その他のエンティティ
│   └── repository/     # Repository Interfaces
│       ├── bean_repository.go     # Bean 操作インターフェース
│       ├── user_repository.go     # User 操作インターフェース
│       └── ...                    # その他のリポジトリインターフェース
├── model/              # Database Models (GORM)
├── dto/                # Data Transfer Objects
├── common/
│   └── converter/      # データ変換ロジック
├── infrastructure/     # 外部サービス (S3, etc.)
└── cmd/
    ├── server/         # サーバー起動
    └── db/            # DB操作コマンド
```

### Domain Layer 詳細

#### Domain Entities (`/domain/entity/`)

- **目的**: ビジネスロジックの中核となるオブジェクトを定義
- **特徴**:
  - データベース実装に依存しない純粋なビジネスオブジェクト
  - ビジネスルールや制約を表現
  - 他のレイヤーからの依存を受けない
- **例**: Bean, User, Country, Roaster など

#### Repository Interfaces (`/domain/repository/`)

- **目的**: データアクセスの抽象化
- **特徴**:
  - UseCase 層がデータアクセス方法を知る必要がない
  - 依存関係の逆転により、ドメイン層が外部実装に依存しない
  - テスタビリティの向上
- **実装**: `/repository/` で具体的な実装を提供

## 開発ガイドライン

### 1. レイヤー間のルール

- Controller は**DTO のみ**扱う
- UseCase は**Entity のみ**扱う（model パッケージを参照しない）
- Repository は**Model のみ**扱う
- Domain は**純粋なビジネスロジック**のみ（外部実装に依存しない）

### 2. Domain Layer のルール

- **Entity**:
  - データベース実装に依存しない
  - ビジネスルールを表現する
  - 他のレイヤーのパッケージを import しない
- **Repository Interface**:
  - データアクセスの抽象化を提供
  - 具体的な実装詳細を含まない
  - Entity のみを扱う（Model は使わない）

### 3. 変換処理

- 各レイヤー間の変換は専用の converter 関数を使用
- 循環参照を避けるため、必要最小限の関連データのみ変換

### 3. エラーハンドリング

- 各レイヤーで適切なエラーハンドリングを実装
- Repository 層のエラーは UseCase 層で適切に変換

## API エンドポイント

### Bean 関連

- `GET /beans/:id` - Bean 詳細取得
- `POST /beans` - Bean 作成
- `PUT /beans/:id` - Bean 更新
- `GET /users/:id/beans` - ユーザーの Bean 一覧取得

### User 関連

- `GET /users/:id` - ユーザー詳細取得
- `POST /users` - ユーザー作成

## セットアップ

```bash
# 依存関係のインストール
go mod download

# データベースマイグレーション
go run cmd/db/migrate/main.go

# サーバー起動
go run cmd/server/main.go
```

## 環境変数

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=c0fee
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
S3_BUCKET_NAME=your-bucket-name
```

## 重要な設計決定

### Domain Driven Design の原則

**依存関係の逆転**:

- Domain 層は他のどの層にも依存しない
- UseCase 層が Domain 層の interface に依存
- Repository 層が Domain 層の interface を実装

**ドメインエンティティの独立性**:

- Entity は純粋なビジネスオブジェクト
- データベース、フレームワーク、外部ライブラリに依存しない
- ビジネスルールとデータの整合性を保証

**Repository パターン**:

- データアクセスの抽象化
- Domain 層で interface を定義、Repository 層で実装
- テストでのモック化が容易

### 必須関連エンティティ vs オプショナル

**必須関連エンティティ**（常に存在する）:

- User
- Roaster
- Country
- RoastLevel

**オプショナル関連エンティティ**（存在しない場合がある）:

- Area
- Farm
- Farmer
- ProcessMethod
- Varieties

### ID 型の使い分け

- **User**: UUID
- **その他のエンティティ**: uint

### データ型の使い分け

- **ポインタ型**: Area, Farm, Farmer, ProcessMethod（オプショナル）
- **値型**: User, Roaster, Country, RoastLevel（必須）

## DTO 命名規則

このプロジェクトでは、一貫性と可読性を保つために以下の命名規則を採用しています：

| 種類        | 命名規則   | 用途                                | 例                                  |
| ----------- | ---------- | ----------------------------------- | ----------------------------------- |
| **Input**   | `~Input`   | リクエスト JSON の 1 次情報         | `BeanInput`, `BeanRatingInput`      |
| **Ref**     | `~Ref`     | リクエスト JSON の 2 次情報（参照） | `CountryRef`, `RoasterRef`, `IdRef` |
| **Output**  | `~Output`  | レスポンス JSON の構造体            | `BeanOutput`, `BeanRatingOutput`    |
| **Summary** | `~Summary` | レスポンス JSON の構成要素          | `IdNameSummary`, `BeanSummary`      |

### 命名規則の詳細

#### Request 側（Input）

- **`~Input`**: API リクエストで受け取る主要なデータ構造
- **`~Ref`**: リクエスト内で他のエンティティを参照するための ID 構造体

#### Response 側（Output）

- **`~Output`**: API レスポンスとして返す完全なデータ構造
- **`~Summary`**: レスポンス内で使用される部分的なデータ構造（例：関連エンティティの要約情報）

### 使用例

```go
// Input（リクエスト）
type BeanInput struct {
    Name          *string          `json:"name"`
    Country       CountryRef       `json:"country"`       // 参照データ
    Roaster       RoasterRef       `json:"roaster"`       // 参照データ
    BeanRating    *BeanRatingInput `json:"bean_rating"`   // 1次情報
}

// Output（レスポンス）
type BeanOutput struct {
    ID            uint            `json:"id"`
    Name          *string         `json:"name"`
    User          IdNameSummary   `json:"user"`          // 構成要素
    Roaster       IdNameSummary   `json:"roaster"`       // 構成要素
    BeanRatings   []BeanRatingOutput `json:"bean_ratings"` // 完全なデータ
}
```

## データベース操作

### マイグレーション

```bash
# マイグレーション状況確認
goose postgres "user=c0fee-user password=c0fee-pass dbname=c0fee host=localhost port=5432 sslmode=disable" status

# 新しいマイグレーション作成
goose create -dir infrastructure/db/migrations {{migration_name}} go
```

## プロジェクト構造

```
c0fee-api/
├── cmd/                    # エントリーポイント
├── common/                 # 共通機能
├── controller/             # コントローラー層
├── dto/                    # データ転送オブジェクト
├── infrastructure/         # インフラストラクチャ層
├── model/                  # ドメインモデル
├── repository/             # リポジトリ層
├── router/                 # ルーティング
└── usecase/               # ユースケース層
```
