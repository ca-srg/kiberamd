# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

Kibela GraphQL APIを使用してすべてのノートをmarkdownファイルとしてエクスポートするCLIツールです。

## 開発環境

- Go 1.22.5
- 依存パッケージ: godotenv, machinebox/graphql, cobra
- **vendor は使用禁止**

## コマンド

### インストール
```bash
go install github.com/ca-srg/kiberamd@latest
```

### ローカル開発でのビルド
```bash
go build -o kiberamd
```

### 実行
```bash
# 環境変数を設定
export KIBELA_TOKEN="your_token"
export KIBELA_TEAM="your_team"

# または .env ファイルを作成
echo 'KIBELA_TOKEN=your_token' > .env
echo 'KIBELA_TEAM=your_team' >> .env

# 実行（インストール済みの場合）
kiberamd

# または、ローカルビルドの場合
./kiberamd
```

### ライブラリ調査
ライブラリの使用方法を調べる際は必ず `go doc` コマンドを使用：
```bash
go doc github.com/machinebox/graphql
```

## アーキテクチャ

### パッケージ構成

- **main.go**: CLIエントリーポイント（cobra使用）
  - 環境変数の読み込み、クライアント初期化、エクスポート実行を行う

- **internal/config**: 設定管理
  - 環境変数（KIBELA_TOKEN, KIBELA_TEAM）から設定を読み込む
  - `.env` ファイルをサポート（godotenv使用）

- **internal/kibela**: Kibela API クライアント
  - GraphQL APIを使用してノートを取得
  - ページネーション実装（100件ずつ取得）
  - 主要な型: `Client`, `Note`, `Author`, `Folder`, `PageInfo`

- **internal/export**: エクスポート処理
  - ノートをmarkdownファイルに変換
  - ファイル名生成（日付プレフィックス + サニタイズされたタイトル）
  - カテゴリ抽出ロジック（フォルダパスから第3階層を抽出）

### 主要な処理フロー

1. `config.Load()` で環境変数を読み込み
2. `kibela.NewClient()` でAPIクライアントを初期化
3. `exporter.ExportAllNotes()` でエクスポート実行
   - `client.ProcessNotesInBatches()` でバッチ単位（100件ずつ）でノートを取得・処理（メモリ効率のためストリーミング処理）
   - 各ノートに対して `saveNoteAsMarkdown()` を実行
     - `generateFilename()` でファイル名を生成
     - `convertToMarkdown()` でmarkdownに変換
     - `os.WriteFile()` でファイルに書き込み

### カテゴリ抽出ロジックの特徴

- フォルダパスの第3階層（インデックス2）をカテゴリとして使用
- 第1階層が優先カテゴリ（個人メモ、日報、手順など）の場合はそれを使用
- 特定フォルダの専用マッピング（例: 施策仕様書 → 仕様書）
- カテゴリが決定できない場合は `CategoryNotFoundError` を返し、スキップして処理を継続

## コーディング規約

### エラーハンドリング

以下の形式で記述すること：
```go
err := hoge()
if err != nil {
    return fmt.Errorf("failed to hoge: %w", err)
}
```

**NG例:**
```go
if err := hoge(); err != nil {
    return err
}
```

### ファイル名サニタイズ

- macOS、Linux、Windows で問題となる文字を除去
- UTF-8文字境界を考慮した安全な長さ制限
- `truncateUTF8()` 関数を使用して文字が壊れないように切り詰め

## 注意事項

- module名は `github.com/ca-srg/kiberamd` （READMEの表記とは異なる）
- 出力先ディレクトリは `markdown/`
- GraphQLクエリは100件ずつページネーションで取得
- カテゴリが決定できないノートはスキップされる（エラーではなく警告）