# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Running the Application

- Start server: `go run cmd/server/main.go` (runs on port 8080)
- Start services: `docker compose up` (PostgreSQL on 5432, MinIO on 9000/9001)

### Database Operations

- Run migrations: `go run cmd/db/migrate/main.go up`
- Check migration status: `go run cmd/db/migrate/main.go status`
- Create new migration: `goose create -dir infrastructure/db/migrations {{migration_name}} go`
- Seed database: `go run cmd/db/seed/main.go` (interactive - choose table or "all")

### Build and Dependencies

- Install dependencies: `go mod tidy`
- Build: `go build cmd/server/main.go`

## Current Architecture (Clean Architecture)

### レイヤー構成と責務

```
Controller Layer (DTO)
    ↓
UseCase Layer (Entity)
    ↓
Repository Layer (Model)
    ↓
Database
```

### Domain Layer

#### Domain Entities (`/domain/entity/`)

- **目的**: ビジネスドメインの核となるオブジェクト
- **特徴**:
  - ビジネスルールとロジックを内包
  - データベース技術に依存しない
  - Clean Architecture の中心部分
- **使用箇所**: UseCase 層で主に使用
- **例**: `entity.Bean`, `entity.User`, `entity.BeanRating`

#### Domain Repository Interfaces (`/domain/repository/`)

- **目的**: データアクセスの抽象化
- **特徴**:
  - Repository 層の実装に依存しない
  - Domain Entity を扱う
  - Dependency Inversion の実現
- **例**: `IBeanRepository`, `IUserRepository`

#### Controller Layer

- **扱うデータ**: DTO (Data Transfer Object)
- **責務**: HTTP リクエスト/レスポンスの処理、バリデーション、ルーティング
- **場所**: `/controller`
- **重要**: model パッケージを参照しない

#### UseCase Layer

- **扱うデータ**: Entity (Domain Entity)
- **責務**: ビジネスロジック、ドメインルールの実装
- **場所**: `/usecase`
- **重要**: model パッケージを直接参照しない
- **依存**: Domain Repository Interface に依存（実装には依存しない）

#### Repository Layer

- **扱うデータ**: Model (Database Model)
- **責務**: データベースへのアクセス、永続化
- **場所**: `/repository`
- **実装**: Domain Repository Interface を実装
- **変換**: Entity ↔ Model の変換を担当

### Domain-Driven Design (DDD) 要素

#### Entities vs Models

**Domain Entity (`/domain/entity/`)**:

- ビジネスドメインの概念を表現
- アプリケーションのコアロジック
- データベース技術に非依存
- UseCase 層で使用

**Database Model (`/model/`)**:

- データベーステーブルの構造を表現
- GORM タグによる OR Mapping
- データベース技術に依存
- Repository 層で使用

#### Repository Pattern

```go
// Domain Repository Interface
type IBeanRepository interface {
    GetById(bean *entity.Bean, id uint) error
    Create(bean *entity.Bean) error
    Update(bean *entity.Bean) error
}

// Implementation
type beanRepository struct {
    db *gorm.DB
}

func (br *beanRepository) GetById(domainBean *entity.Bean, id uint) error {
    var modelBean model.Bean
    // DB操作
    err := br.db.First(&modelBean, id).Error

    // Model → Entity 変換
    *domainBean = *entity_model.ModelBeanToEntity(&modelBean)
    return err
}
```

### データフロー例

#### Create Bean API

```
1. Controller: HTTP Request → DTO
2. Controller: DTO → Entity (dto_entity converter)
3. UseCase: Entity でビジネスロジック実行
4. Repository: Entity → Model (entity_model converter)
5. Repository: Model でDB操作
6. Repository: Model → Entity (entity_model converter)
7. UseCase: Entity → DTO (dto_entity converter)
8. Controller: DTO → HTTP Response
```

### コンバーター構成

```
/common/converter/
├── entity_model/        # Repository層で使用 (Entity ↔ Model)
│   ├── bean_entity_converter.go       # EntityBeanToModel, ModelBeanToEntity
│   ├── user_entity_converter.go       # EntityUserToModel, ModelUserToEntity
│   ├── bean_rating_entity_converter.go
│   └── ...
├── dto_entity/          # UseCase層で使用 (DTO ↔ Entity)
│   ├── bean_dto_converter.go          # DtoBeanToEntity, EntityBeanToDto
│   ├── user_dto_converter.go          # DtoUserToEntity, EntityUserToDto
│   └── ...
└── model_dto/           # レガシーサポート（直接変換）
    └── bean_converter.go              # ConvertBeanToOutput（非推奨）
```

#### コンバーター命名規則

**Entity ↔ Model (`entity_model/`)**:

- `EntityToModel`: Entity → Model
- `ModelToEntity`: Model → Entity

**DTO ↔ Entity (`dto_entity/`)**:

- `DtoToEntity`: DTO → Entity
- `EntityToDto`: Entity → DTO
- `EntityToBeanSummary`: Entity → 要約 DTO

## Key Architecture Principles

1. **Clean Architecture + DDD**:

   - Domain 層がアプリケーションの中心
   - 外部技術（DB、Framework）に依存しない
   - 依存関係が内側（Domain）に向かう

2. **Layer Separation**:

   - Controller = DTO only
   - UseCase = Entity only (no model package imports)
   - Repository = Model + Entity (conversion)

3. **Data Flow**:

   - HTTP JSON ↔ DTO ↔ Entity ↔ Model ↔ Database

4. **Dependency Inversion**:

   - UseCase は Repository Interface に依存
   - Repository Implementation は UseCase に依存しない

5. **Conversion**:
   - 各層間で専用コンバーター関数を使用
   - 循環参照を避ける

## Domain Entities vs Models

### Entity/Model 対応関係

| Domain Entity       | Database Model     | 説明           |
| ------------------- | ------------------ | -------------- |
| `entity.Bean`       | `model.Bean`       | コーヒー豆情報 |
| `entity.User`       | `model.User`       | ユーザー情報   |
| `entity.BeanRating` | `model.BeanRating` | 豆の評価       |
| `entity.Country`    | `model.Country`    | 生産国         |
| `entity.Roaster`    | `model.Roaster`    | 焙煎業者       |

### Required Associations (always exist):

- User, Roaster, Country, RoastLevel

### Optional Associations (may not exist):

- Area, Farm, Farmer, ProcessMethod, Varieties

### ID Types:

- **User**: UUID
- **Other entities**: uint

### Data Types:

- **Pointer types**: Area, Farm, Farmer, ProcessMethod (optional)
- **Value types**: User, Roaster, Country, RoastLevel (required)

## Technical Stack

- **Language**: Go 1.22.2
- **Framework**: Echo v4.12.0
- **Database**: PostgreSQL with GORM ORM
- **Storage**: MinIO (S3-compatible) for file storage
- **Authentication**: JWT-based with custom middleware
- **Validation**: go-playground/validator

## Important Implementation Notes

1. **Domain Layer Isolation**:

   - Domain entities は外部技術に依存しない
   - `entity` package は `model` package を import しない
   - `usecase` package は `model` package を直接 import しない

2. **Repository Pattern Implementation**:

   - UseCase は Repository Interface のみに依存
   - Repository Implementation で Entity ↔ Model 変換
   - Database 操作は Repository 内部に隠蔽

3. **Converters**:

   - Handle ID and association together (not separately)
   - 必須フィールドは常に設定（条件付き設定を避ける）
   - Optional fields: Use pointer types for optional associations

4. **Data Consistency**:

   - Entity 作成時は必須フィールドを全て設定
   - Model への変換時は ID の整合性を保つ
   - 外部キー制約違反を防ぐ

5. **Error Handling**: Each layer handles and converts errors appropriately

6. **Performance**: Use GORM Preload to avoid N+1 problems

## Refactoring History

### Domain Layer 導入 (2025/08/30)

**目的**: Repository 層以外では model package を使わず entity を使用

**主な変更**:

- Domain Entity と Repository Interface の追加
- Converter 構造の整理 (`entity_model/`, `dto_entity/`, `model_dto/`)
- UseCase 層の Entity 化
- Repository 層での Entity ↔ Model 変換実装

**アーキテクチャ向上**:

- Clean Architecture + DDD パターンの実装
- 依存関係の逆転（Dependency Inversion）
- ビジネスロジックとデータベース技術の分離
