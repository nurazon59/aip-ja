package main

import (
	"strings"
	"testing"
)

func TestSplitFrontmatter(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		wantFM  string
		wantBod string
	}{
		{
			name:    "with frontmatter",
			src:     "---\nid: 1\n---\nbody\n",
			wantFM:  "id: 1\n",
			wantBod: "body\n",
		},
		{
			name:    "no frontmatter",
			src:     "# title\nbody\n",
			wantFM:  "",
			wantBod: "# title\nbody\n",
		},
		{
			name:    "frontmatter only no body",
			src:     "---\nid: 1\n---\n",
			wantFM:  "id: 1\n",
			wantBod: "",
		},
		{
			name:    "unterminated frontmatter",
			src:     "---\nid: 1\nbody\n",
			wantFM:  "",
			wantBod: "---\nid: 1\nbody\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm, body := splitFrontmatter([]byte(tt.src))
			if string(fm) != tt.wantFM {
				t.Errorf("frontmatter: got %q want %q", fm, tt.wantFM)
			}
			if string(body) != tt.wantBod {
				t.Errorf("body: got %q want %q", body, tt.wantBod)
			}
		})
	}
}

func TestCompareStructures(t *testing.T) {
	tests := []struct {
		name    string
		en      string
		ja      string
		wantOK  bool
		wantSub string
	}{
		{
			name:   "identical structure",
			en:     "---\nid: 1\n---\n# Title\n\nText [a](https://a.example).\n\n```go\nfmt.Println(1)\n```\n",
			ja:     "---\nid: 1\n---\n# タイトル\n\n本文 [a](https://a.example).\n\n```go\nfmt.Println(1)\n```\n",
			wantOK: true,
		},
		{
			name:    "frontmatter differs",
			en:      "---\nid: 1\n---\n# Title\n",
			ja:      "---\nid: 2\n---\n# タイトル\n",
			wantSub: "frontmatter mismatch",
		},
		{
			name:    "code block body differs",
			en:      "# T\n\n```go\nfmt.Println(1)\n```\n",
			ja:      "# T\n\n```go\nfmt.Println(2)\n```\n",
			wantSub: "code block #1 body differs",
		},
		{
			name:    "code block lang differs",
			en:      "# T\n\n```go\nx\n```\n",
			ja:      "# T\n\n```python\nx\n```\n",
			wantSub: "code block #1 lang differs",
		},
		{
			name:    "code block count differs",
			en:      "# T\n\n```go\na\n```\n\n```go\nb\n```\n",
			ja:      "# T\n\n```go\na\n```\n",
			wantSub: "code block count differs",
		},
		{
			name:    "heading structure differs",
			en:      "# A\n## B\n### C\n",
			ja:      "# A\n### B\n## C\n",
			wantSub: "heading structure differs",
		},
		{
			name:    "missing ref def",
			en:      "[link][x]\n\n[x]: https://x.example\n",
			ja:      "[link][x]\n",
			wantSub: "missing in ja",
		},
		{
			name:    "ref def destination differs",
			en:      "[link][x]\n\n[x]: https://en.example\n",
			ja:      "[link][x]\n\n[x]: https://ja.example\n",
			wantSub: "destination differs",
		},
		{
			name:    "link count differs",
			en:      "See [a](https://a.example) and [b](https://b.example).\n",
			ja:      "See [a](https://a.example).\n",
			wantSub: "link count differs",
		},
		{
			name:   "anchor-only links ignored",
			en:     "See [a](#foo).\n",
			ja:     "See [a](#bar).\n",
			wantOK: true,
		},
		{
			name:   "autolink matches",
			en:     "Visit <https://example.com>.\n",
			ja:     "訪問 <https://example.com>.\n",
			wantOK: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			en := extractStructure([]byte(tt.en))
			ja := extractStructure([]byte(tt.ja))
			diffs := compareStructures(en, ja)
			if tt.wantOK {
				if len(diffs) != 0 {
					t.Errorf("expected no diffs, got: %v", diffs)
				}
				return
			}
			if len(diffs) == 0 {
				t.Errorf("expected diff containing %q, got none", tt.wantSub)
				return
			}
			found := false
			for _, d := range diffs {
				if strings.Contains(d, tt.wantSub) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected diff containing %q, got: %v", tt.wantSub, diffs)
			}
		})
	}
}
