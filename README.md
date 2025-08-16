# C0fee API

コーヒー豆管理 API - Go 製 RESTful API サーバー

## 技術スタック

- **言語**: Go 1.22.2
- **フレームワーク**: Echo v4.12.0
- **データベース**: PostgreSQL
- **ORM**: GORM
- **ストレージ**: MinIO/S3 互換
- **アーキテクチャ**: Clean Architecture

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
