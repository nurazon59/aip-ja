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
- Do not translate comments inside fenced code blocks. Leave code-block
  contents (including `//`, `#`, and proto comments) in the original
  language unless the entire surrounding section explicitly demonstrates
  translated code comments.
- Translate callout labels with a half-width colon to match the upstream
  form `**Note:**` / `**Important:**`:
  - `**Note:**` -> `**注:**`
  - `**Important:**` -> `**重要:**`
  - Do not use the full-width colon `：` for these labels.
- When using a reference-style AIP link in prose such as `[AIP-132][aip-132]`,
  ensure a matching `[aip-132]: ...` reference definition exists in the file.
  Carry over the original AIP's reference definitions when they are missing.

## Requirement Keywords

AIPs define RFC 2119 keywords in AIP-8. In source AIPs these keywords are
lower-case and bold. Translate the keyword into Japanese in bold and add the
English original in parentheses:

| Source           | Translation                          |
| ---------------- | ------------------------------------ |
| `**must**`       | `**しなければならない**（must）`     |
| `**must not**`   | `**してはならない**（must not）`     |
| `**should**`     | `**すべきである**（should）`       |
| `**should not**` | `**すべきではない**（should not）` |
| `**may**`        | `**してもよい**（may）`              |

Do not translate normative `may` as `かもしれない`.

When requirement keywords appear inside code examples, comments, or quoted text,
preserve the local context. Do not mechanically rewrite code identifiers or enum
values.

### Verb attachment

When a requirement keyword follows a verb in running prose, fuse it into the
verb's conjugated form. Do not concatenate the dictionary form of the verb
with the bolded keyword — the result is ungrammatical and obscures the
normative requirement.

| Bad                                    | Good                                |
| -------------------------------------- | ----------------------------------- |
| `定義する **しなければならない**`      | `定義**しなければならない**`        |
| `含める **しなければならない**`        | `含め**なければならない**`          |
| `異なる**してもよい**`                 | `異なっ**てもよい**`                |
| `持つ**してはならない**`               | `持っ**てはならない**`              |
| `設定する**すべきである**`             | `設定**すべきである**`              |

The bilingual `（must）` / `（may）` / `（should）` parenthesis stays at the
end of the resulting bold phrase, e.g. `**しなければならない**（must）`.

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
| customer               | 顧客                      |
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

## Review Checklist

- YAML frontmatter unchanged.
- Requirement keywords use the approved bilingual form.
- Reference definitions are byte-for-byte structurally intact.
- Reference-style links in prose still resolve to the original labels.
- AIP links and paths remain relative and valid.
- Code identifiers, enum values, field names, method names, and annotations are
  not translated.
- Terminology is consistent with the table above.
