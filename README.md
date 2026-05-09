# AIP JA

This repository hosts an unofficial Japanese translation site for
google.aip.dev.

This project is not affiliated with, sponsored by, or endorsed by Google or the
original AIP project maintainers.

## License

This repository uses multiple licenses:

- Site implementation code, local scripts, and original repository
  configuration are licensed under the MIT License. See [LICENSE](./LICENSE).
- Original AIP content from google.aip.dev is licensed under the Creative
  Commons Attribution 4.0 International License.
- Japanese translations and other adaptations of AIP content are licensed under
  the Creative Commons Attribution 4.0 International License.
- Code samples derived from google.aip.dev are licensed under the Apache
  License 2.0. See [LICENSE-Apache-2.0.txt](./LICENSE-Apache-2.0.txt).
- The Hugo Book theme is included as a submodule and remains under its own MIT
  license.

For content licensing details and attribution, see
[LICENSE-content.md](./LICENSE-content.md).

## Syncing English Content

`content/en` is regenerated from the `upstream/google.aip.dev` submodule by
`scripts/sync-en.go`.

Do not make local-only edits directly under `content/en`. Put files that should
override upstream content under `content-overrides/en` using the same relative
path. The sync script copies upstream content first, then applies those
overrides.
