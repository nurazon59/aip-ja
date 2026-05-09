---
name: translate
description: Translate Google AIP Markdown from English to Japanese while preserving technical standards wording, Markdown structure, AIP links, RFC 2119 requirement keywords, and API terminology. Use for AIP-ja translation work and review of English-to-Japanese technical documentation.
---

# AIP Translation

Use this skill when translating Google AIP Markdown into Japanese.

## Goal

Produce publication-quality Japanese Markdown that is faithful to the original
AIP, readable for Japanese API designers, and safe to publish after review.

AIP translation is not pure localization. Preserve specification precision,
recognizable API terminology, and Markdown/link integrity over fully naturalized
Japanese.

## Workflow

1. Read the source Markdown and any nearby existing Japanese translations for
   local style.
2. Translate prose and visible heading/list text.
3. Preserve frontmatter, Markdown structure, reference definitions, anchors,
   file paths, URLs, dates, code identifiers, enum values, method names, field
   names, and examples.
4. Review the output specifically for requirement keywords, links, code blocks,
   and terminology consistency.

## Hard Requirements

- Output only translated Markdown when asked for a translation.
- Do not add code fences around the whole document.
- Preserve YAML frontmatter keys and scalar values exactly.
- Preserve Markdown structure: headings, blank lines, lists, tables, code blocks,
  emphasis, links, anchors, and reference definitions.
- Do not translate reference definition labels or targets.
- For reference-style links in prose, translate the visible text but keep the
  original label:
  - `[Declarative clients][]` -> `[宣言的クライアント][Declarative clients]`
  - `[AIP-131][]` stays `[AIP-131][]`
- Keep AIP IDs in prose as `AIP-123` without zero-padding unless the source
  intentionally uses another form in code or paths.
- Do not invent explanatory notes or add missing context.

## Requirement Keywords

AIPs define RFC 2119 keywords in AIP-8. In source AIPs these keywords are
lower-case and bold. Preserve the English keyword and add the Japanese meaning
outside the bold span:

| Source           | Translation                          |
| ---------------- | ------------------------------------ |
| `**must**`       | `**must**（しなければならない）`     |
| `**must not**`   | `**must not**（してはならない）`     |
| `**should**`     | `**should**（するべきである）`       |
| `**should not**` | `**should not**（するべきではない）` |
| `**may**`        | `**may**（してもよい）`              |

Do not translate normative `may` as `かもしれない`.

When requirement keywords appear inside code examples, comments, or quoted text,
preserve the local context. Do not mechanically rewrite code identifiers or enum
values.

## Terminology

Use AIP-9 as the source of truth for common AIP terminology. Keep technical terms
recognizable. When in doubt, preserve the English term rather than inventing a
new Japanese translation.

| English                | Japanese                  |
| ---------------------- | ------------------------- |
| API                    | API                       |
| AIP                    | AIP                       |
| API backend            | APIバックエンド           |
| API consumer           | APIコンシューマー         |
| API definition         | API定義                   |
| API frontend           | APIフロントエンド         |
| API interface          | APIインターフェース       |
| API method             | APIメソッド               |
| API producer           | APIプロデューサー         |
| API product            | APIプロダクト             |
| API request            | APIリクエスト             |
| API service            | APIサービス               |
| API service definition | APIサービス定義           |
| API service endpoint   | APIサービスエンドポイント |
| API service name       | APIサービス名             |
| API title              | APIタイトル               |
| API version            | APIバージョン             |
| Network API            | ネットワークAPI           |
| client                 | クライアント              |
| declarative client     | 宣言的クライアント        |
| user                   | ユーザー                  |
| resource               | リソース                  |
| method                 | メソッド                  |
| field                  | フィールド                |
| message                | メッセージ                |
| enum                   | enum                      |
| annotation             | アノテーション            |
| Protocol Buffers       | Protocol Buffers          |
| proto / `.proto`       | proto / `.proto`          |
| resource-oriented      | リソース指向              |
| API surface            | APIサーフェス             |
| standard method        | 標準メソッド              |
| custom method          | カスタム メソッド         |
| long-running operation | long-running operation    |
| management plane       | 管理プレーン              |
| data plane             | データプレーン            |
| control plane          | コントロールプレーン      |
| critical path          | クリティカルパス          |
| provision              | プロビジョニングする      |
| configure              | 構成する                  |
| retrieve / retrieval   | 取得する / 取得           |
| interface with         | インターフェースする      |
| heterogeneous          | 異種混在                  |
| availability           | 可用性                    |
| throughput             | スループット              |
| latency                | レイテンシ                |
| blob store             | blob store                |
| push / pull            | push / pull               |
| output only            | Output only               |
| input only             | Input only                |
| immutable              | immutable                 |
| required               | required                  |
| optional               | optional                  |

## Style

- Use concise formal Japanese, mainly `である` style.
- Keep sentence meaning precise even if the Japanese becomes slightly technical.
- Avoid awkward mixed-language fragments such as `ユーザーFacing`; use
  `ユーザー向け`.
- Do not over-translate identifiers, API concepts, or standards terminology.
- Prefer `APIサーフェス` over `API表面`.
- Prefer `blob store` over forced translations such as `ブロブストア` unless
  local context already standardizes a Japanese form.
- Preserve product names, library names, domains, and protocol names.
- Code blocks are normally left unchanged except for prose comments when the
  surrounding document clearly translates comments.

## Review Checklist

- YAML frontmatter unchanged.
- Requirement keywords use the approved bilingual form.
- Reference definitions are byte-for-byte structurally intact.
- Reference-style links in prose still resolve to the original labels.
- AIP links and paths remain relative and valid.
- Code identifiers, enum values, field names, method names, and annotations are
  not translated.
- Terminology is consistent with the table above.
