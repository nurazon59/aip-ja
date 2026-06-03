# AGENTS.md

このリポジトリ `aip-ja` は Google AIP (https://google.aip.dev) の日本語訳である。
本ファイルは OpenCode / Claude Code などのエージェントが本リポジトリで作業する際の
共通指示を提供する。**翻訳タスクではここに書かれたルールを厳格に守ること。**

## リポジトリ構成

```
.
├── content/
│   ├── en/aip/general/0000.md   原文 (upstream から sync-en で同期)
│   └── ja/aip/general/0000.md   日本語訳 (人手で翻訳・編集)
├── upstream/google.aip.dev/      原文の git submodule
├── scripts/                       Go 製の同期・lint スクリプト
└── .github/workflows/             CI
```

- `content/en/` は自動同期される。**手動で編集しない**。
- 翻訳作業は `content/ja/aip/<category>/NNNN.md` に対して行う。

## 翻訳タスクの絶対ルール

### 出力形式

- 翻訳依頼に対しては **翻訳後の Markdown 本文のみ** を出力する。説明文・前置き不要。
- ドキュメント全体を ```` ``` ```` で囲まない。
- YAML frontmatter のキー・スカラー値は **完全に保持**。翻訳しない。
- Markdown 構造 (見出し、空行、リスト、テーブル、コードブロック、強調、リンク、
  アンカー、reference definition) は **すべて保持**。

### リファレンス・リンク

- reference definition のラベルとターゲットは翻訳しない。
- 散文中の reference-style リンクは「表示テキストのみ訳し、ラベルは英語のまま保持」:
  - `[Declarative clients][]` → `[宣言的クライアント][Declarative clients]`
  - `[AIP-131][]` はそのまま `[AIP-131][]`
- AIP ID は散文では `AIP-123` 形式 (ゼロパディングなし)。コードやパスの中で
  原文が `aip-0123` 形式を使っている場合のみそれを維持する。
- 散文中で `[AIP-132][aip-132]` のような ref-style AIP リンクを使う場合、
  対応する `[aip-132]: ./0132.md` の reference definition がファイル内に
  存在することを必ず確認する。無ければ追加する。

### コードブロック

- フェンス付きコードブロック内のコメント (`//`, `#`, proto コメントなど) は **訳さない**。
  原文の言語のままにする。
- code identifier / enum 値 / フィールド名 / メソッド名 / アノテーションは **訳さない**。

### コールアウト

- `**Note:**` → `**注:**`
- `**Important:**` → `**重要:**`
- **半角コロン `:` を使う**。全角コロン `：` にしない。

### 推測の追加禁止

- 原文に無い説明や補足を勝手に足さない。
- 不明確な箇所も忠実に訳し、解釈で補わない。

## RFC 2119 要件キーワード

AIP-8 で定義される RFC 2119 キーワード。原文では小文字＋bold。
**日本語の bold + 英語原語の括弧併記** で訳す。

| 原文             | 訳                                       |
| ---------------- | ---------------------------------------- |
| `**must**`       | `**しなければならない**（must）`         |
| `**must not**`   | `**してはならない**（must not）`         |
| `**should**`     | `**すべきである**（should）`             |
| `**should not**` | `**すべきではない**（should not）`       |
| `**may**`        | `**してもよい**（may）`                  |

- 規範的な `may` を「かもしれない」と訳さない。
- 括弧は **全角丸括弧 `（）`** を使う。
- コード例・引用・コメント内のキーワードは文脈に応じて保持。機械的に置換しない。

### 動詞接続 (重要・頻出ミス)

動詞のあとに要件キーワードが続くときは、**動詞の活用形に融合させる**。
辞書形＋bold キーワードの直接連結は非文法的なので **絶対にやらない**。

| ❌ Bad                              | ✅ Good                       |
| ----------------------------------- | ----------------------------- |
| `定義する **しなければならない**`   | `定義**しなければならない**`  |
| `含める **しなければならない**`     | `含め**なければならない**`    |
| `異なる**してもよい**`              | `異なっ**てもよい**`          |
| `持つ**してはならない**`            | `持っ**てはならない**`        |
| `設定する**すべきである**`          | `設定**すべきである**`        |
| `作成して**してはならない**`        | `作成**してはならない**`      |
| `使用して**してはならない**`        | `使用**してはならない**`      |
| `配置され**てもよい**` (受身は OK) | `配置**してもよい**` (能動可) |

英語原語の括弧 `（must）` / `（may）` / `（should）` は bold 句の **末尾** に付く。
例: `**しなければならない**（must）`

#### 自己チェック (必須)

翻訳後、提出前に以下を **必ず** 確認:

1. `して**してはならない**` を grep。マッチしたら誤り。`**してはならない**` のみが正。
2. `する **しなければならない**` (動詞辞書形＋スペース) を grep。マッチしたら活用形に直す。
3. `（may）される` のような余分な助動詞が末尾に残っていないか確認。

## 用語表 (AIP-9 準拠)

迷ったら **英語のまま** 残す。新しい和訳を発明しない。

| 英語                     | 日本語                    |
| ------------------------ | ------------------------- |
| API                      | API                       |
| AIP                      | AIP                       |
| API backend              | APIバックエンド           |
| API consumer             | APIコンシューマー         |
| API definition           | API定義                   |
| API frontend             | APIフロントエンド         |
| API interface            | APIインターフェース       |
| API method               | APIメソッド               |
| API producer             | APIプロデューサー         |
| API product              | APIプロダクト             |
| API request              | APIリクエスト             |
| API service              | APIサービス               |
| API service definition   | APIサービス定義           |
| API service endpoint     | APIサービスエンドポイント |
| API service name         | APIサービス名             |
| API title                | APIタイトル               |
| API version              | APIバージョン             |
| Network API              | ネットワークAPI           |
| client                   | クライアント              |
| declarative client       | 宣言的クライアント        |
| user                     | ユーザー                  |
| customer                 | 顧客                      |
| resource                 | リソース                  |
| method                   | メソッド                  |
| field                    | フィールド                |
| message                  | メッセージ                |
| enum                     | enum                      |
| annotation               | アノテーション            |
| Protocol Buffers         | Protocol Buffers          |
| proto / `.proto`         | proto / `.proto`          |
| resource-oriented        | リソース指向              |
| API surface              | APIサーフェス             |
| standard method          | 標準メソッド              |
| custom method            | カスタム メソッド         |
| long-running operation   | long-running operation    |
| management plane         | 管理プレーン              |
| data plane               | データプレーン            |
| control plane            | コントロールプレーン      |
| critical path            | クリティカルパス          |
| provision                | プロビジョニングする      |
| configure                | 構成する                  |
| retrieve / retrieval     | 取得する / 取得           |
| interface with           | インターフェースする      |
| heterogeneous            | 異種混在                  |
| availability             | 可用性                    |
| throughput               | スループット              |
| latency                  | レイテンシ                |
| blob store               | blob store                |
| push / pull              | push / pull               |
| output only              | Output only               |
| input only               | Input only                |
| immutable                | immutable                 |
| required                 | required                  |
| optional                 | optional                  |

**`customer` と `user` の訳し分けに注意。** `customer` は「顧客」、`user` は「ユーザー」。
APIプロバイダから見たAPIプロダクトの利用者は通常 customer = 顧客。

## 文体

- **`である` 調**。常体・断定形。「です・ます」調にしない。
- 簡潔で技術文書らしい日本語。意味の正確さを優先し、流暢さのために情報を削らない。
- `ユーザーFacing` のような不自然な英日混在を避ける → `ユーザー向け`。
- 識別子・API 概念・標準用語は過度に和訳しない。
- 製品名・ライブラリ名・ドメイン名・プロトコル名は保持。

## ワークフロー

1. `content/en/aip/general/NNNN.md` (原文) を読む。
2. 同じディレクトリの近い番号の既存 `ja` 翻訳を 1〜2 ファイル覗いて文体・用語を合わせる。
3. 散文と表示テキストを訳す。frontmatter・構造・コード・リンクは保持。
4. **翻訳後、上記「動詞接続の自己チェック」を必ず実行**。
5. 可能なら `cd scripts && go run .` で lint-translation を実行し、構造ずれを潰す。
6. コミットメッセージは `AIP-NNN: <日本語タイトル>` または `fix(translation): ...` 形式。

## 触ってはいけないもの

- `content/en/` 以下のファイル (upstream から自動同期される)
- `upstream/` 配下の submodule
- `.github/workflows/sync-en.yml` などの CI 設定 (明示依頼なき限り)
- 既存翻訳ファイルの YAML frontmatter

## レビュー時の典型的な誤り

過去の修正で繰り返し見つかった問題:

- 動詞辞書形 ＋ bold キーワードの直接連結 (前述)
- `（may）される` のような末尾の余分な助動詞
- reference-style リンクのラベルを和訳してしまう
- frontmatter の値 (`category: protobuf` など) を訳してしまう
- コードブロック内のコメントを訳してしまう
- AIP ID をゼロパディング (`AIP-0215` と書く) してしまう
- 全角コロン `：` を `**注:**` 等で使ってしまう

## 言語

- ユーザーとの対話は **日本語**。
- コミットメッセージ・PR タイトルも **日本語**。
- コードコメントは英語 (既存スタイル踏襲)。
