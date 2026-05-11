# 自社でのAIP採用

**注:** これをより良くするためのツールを開発中である。進捗については[このGitHub issue][]を参照されたい。

AIPはGoogleで生まれ、GoogleのAPIを書くGooglerを対象としていたが、文書化されたガイダンスの多くはGoogleの外部でも有用である。このドキュメントでは、Googleで働いていなくても、自社のAPIを文書化する方法としてAIPを採用する方法を説明する。

## 問題

民間組織が、APIに関するガイダンスを世界と共有したくない、または共有する必要がない場合がある。例えば、あなたの会社（Acme, Inc.）では、リソースを `string acme_id` という特別なフィールドで識別しているとする。

このルールはAIPとして書くことはできるが、全員と共有する理由はない。あなたと開発チームだけのためのものだからである。しかし、これを将来のAIPと競合させたくない（例えば、このAIPをAIP-1234にすると、将来そのAIPが作成されて競合する可能性がある）。ではどうすればいいのか？

## 9000ブロック

Unicode仕様と同様に、AIP番号の9000ブロック（すなわちAIP-9000からAIP-9999）を「内部使用」用に予約している。これは、これらのAIPが公開ドキュメント化されることはなく、民間企業が自社のAPIガイダンスのために自由に使用できることを意味する。その結果、通常通りAIPを作成し、9000ブロックの番号を割り当てるだけでよい。

## カスタムAIPドメイン

コピーして9000ブロックのAIPリポジトリとして使用できる、フォーク可能な別個のリポジトリを準備中である。GitHub Pagesでレンダリングされるページは、そのブロックのAIPのみを提供し、その他のページはすべて[aip.dev][]にリダイレクトされる。

つまり、独自のAIPドメイン（例：`aip.example.com`）を作成し、それをフォークしたGitHubリポジトリに向けることができる。これにより、既知のAIPはすべてリダイレクトされ、内部のAIPは自身のリポジトリから提供される。これを設定すれば、すべてのAIPをそのドメイン名（例：`aip.example.com/1234`）で引用でき、常に正しい場所に誘導される。

## AIPのフォーク

必要に応じてAIPプロジェクトをフォークし、自身のドメインで実行できる。これにより、組織のニーズに合わせてAIPをカスタマイズできる。または、インフラストラクチャを使用して独自のAIPセットを作成することもできる。

### URLの更新

AIPインフラストラクチャを別のリポジトリでGitHub Pageとして実行するには、新しいリポジトリで正しく動作するように `_config.yaml` ファイルを更新する必要がある。

カスタムドメインとCNAMEがAIP用に作成されている場合、`url` プロパティのみを新しいドメインに更新する必要がある。

```
url: https://aip.dev
```

新しいドメインを作成しない場合は、`baseurl` プロパティを `_config.yaml` に追加する必要がある。このプロパティには、`url` のドメインに追加される可能性のある追加のパス情報を含める。

例えば、GitHubユーザー jdoe123 がAIPプロジェクトを my-aips というリポジトリにフォークしたとする。このユーザーがmasterブランチからコンテンツを提供する場合、GitHub pagesへのURLは `https://jdoe123.github.io/my-aips/` となる。これに伴う `_config.yaml` は以下のように構成される：

```
url: https://jdoe123.github.io
baseurl: /my-aips
```

これらの値がGitHub Pagesでどのように使用されるかの詳細については、Jekyllでこれらの構成を標準化した[リリースノート][]を参照されたい。

### ナビゲーションの構成

ナビゲーションヘッダーとバーは、それぞれ `_data/header.yaml` と `_data/nav.yaml` ファイルに基づいて動的に生成される。

ナビゲーションバーのスキーマは `assets/schemas/nav-schema.yaml` で確認できる。2種類のナビゲーションコンポーネントをサポートしている：`static_group` と `matter_group`。`static_group` メニューコンポーネントは、ページやリポジトリのコンテンツに関係なく、常に同じナビゲーション要素を表示する。`matter_group` コンポーネントは、ドメイン内のAIPまたはサイトの現在のページに基づいて動的に生成される。`matter_group` の構成は `assets/schemas/nav-components.yaml#defintions/matter_group` スキーマで確認できる。

ヘッダーは特別にレンダリングされた `static_group` コンポーネントである。スキーマは `assets/schemas/static_group.yaml` で確認できる。

### テスト構成

`tests` フォルダには、データファイルを検証するnpmテストモジュールが含まれている。これらのテストを実行するにはnpmとmochaが必要である。これらがインストールされると、`npm test` コマンドでテストを実行できる。

<!-- prettier-ignore-start -->
[this github issue]: https://github.com/googleapis/aip/issues/98
[npm]: https://www.npmjs.com/get-npm
[mocha]: https://www.npmjs.com/package/mocha
[release notes]: https://jekyllrb.com/news/2016/10/06/jekyll-3-3-is-here/#2-relative_url-and-absolute_url-filters
<!-- prettier-ignore-end -->
