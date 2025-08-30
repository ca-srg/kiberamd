# Kiberag Export Tool

Kibela GraphQL APIを使用してすべてのノートをmarkdownファイルとしてエクスポートするツールです。

## 必要な環境変数

- `KIBELA_TOKEN`: Kibela APIアクセストークン
- `KIBELA_TEAM`: 対象チーム名

## インストール

```bash
go build -o kiberag-export
```

## 使用方法

```bash
# 環境変数を設定
export KIBELA_TOKEN="your_token"
export KIBELA_TEAM="your_team"

# 実行
./kiberag-export
```

エクスポートされたファイルは `markdown/` ディレクトリに保存されます。

## ファイル名形式

エクスポートされるファイルは以下の形式で保存されます：
- `YYYY-MM-DD_タイトル.md`

## メタデータ

各markdownファイルには以下のメタデータが含まれます：
- title: ノートのタイトル
- id: ノートのID
- author: 作成者のアカウント名
- date: 公開日時
- category: カテゴリ（フォルダパスから抽出）
- reference: KibelaページへのURL