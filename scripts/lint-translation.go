package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type codeBlock struct {
	lang string
	body string
}

type structure struct {
	frontmatter []byte
	codeBlocks  []codeBlock
	headings    []int
	links       []string
	refDefs     map[string]string
}

var mdParser = goldmark.DefaultParser()

func splitFrontmatter(src []byte) (fm []byte, body []byte) {
	first := bytes.IndexByte(src, '\n')
	if first < 0 {
		return nil, src
	}
	if !bytes.Equal(bytes.TrimRight(src[:first], "\r"), []byte("---")) {
		return nil, src
	}
	rest := src[first+1:]
	idx := 0
	for idx <= len(rest) {
		nl := bytes.IndexByte(rest[idx:], '\n')
		var line []byte
		var nextIdx int
		if nl < 0 {
			line = rest[idx:]
			nextIdx = len(rest)
		} else {
			line = rest[idx : idx+nl]
			nextIdx = idx + nl + 1
		}
		if bytes.Equal(bytes.TrimRight(line, "\r"), []byte("---")) {
			return rest[:idx], rest[nextIdx:]
		}
		if nl < 0 {
			break
		}
		idx = nextIdx
	}
	return nil, src
}

func extractStructure(src []byte) structure {
	fm, body := splitFrontmatter(src)
	s := structure{frontmatter: fm, refDefs: map[string]string{}}

	reader := text.NewReader(body)
	ctx := parser.NewContext()
	doc := mdParser.Parse(reader, parser.WithContext(ctx))

	_ = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch v := n.(type) {
		case *ast.FencedCodeBlock:
			lang := string(v.Language(body))
			var buf bytes.Buffer
			lines := v.Lines()
			for i := 0; i < lines.Len(); i++ {
				seg := lines.At(i)
				buf.Write(seg.Value(body))
			}
			s.codeBlocks = append(s.codeBlocks, codeBlock{lang: lang, body: buf.String()})
		case *ast.Heading:
			s.headings = append(s.headings, v.Level)
		case *ast.Link:
			dest := string(v.Destination)
			if !strings.HasPrefix(dest, "#") {
				s.links = append(s.links, dest)
			}
		case *ast.AutoLink:
			s.links = append(s.links, string(v.URL(body)))
		}
		return ast.WalkContinue, nil
	})

	for _, ref := range ctx.References() {
		s.refDefs[string(ref.Label())] = string(ref.Destination())
	}
	return s
}

func compareStructures(en, ja structure) []string {
	var diffs []string

	if !bytes.Equal(en.frontmatter, ja.frontmatter) {
		diffs = append(diffs, "frontmatter mismatch")
	}

	if len(en.codeBlocks) != len(ja.codeBlocks) {
		diffs = append(diffs, fmt.Sprintf("code block count differs: en=%d ja=%d", len(en.codeBlocks), len(ja.codeBlocks)))
	} else {
		for i := range en.codeBlocks {
			if en.codeBlocks[i].lang != ja.codeBlocks[i].lang {
				diffs = append(diffs, fmt.Sprintf("code block #%d lang differs: en=%q ja=%q", i+1, en.codeBlocks[i].lang, ja.codeBlocks[i].lang))
			}
			if en.codeBlocks[i].body != ja.codeBlocks[i].body {
				diffs = append(diffs, fmt.Sprintf("code block #%d body differs", i+1))
			}
		}
	}

	if !sameIntSlice(en.headings, ja.headings) {
		diffs = append(diffs, fmt.Sprintf("heading structure differs: en=%v ja=%v", en.headings, ja.headings))
	}

	var labels []string
	seen := map[string]bool{}
	for l := range en.refDefs {
		if !seen[l] {
			labels = append(labels, l)
			seen[l] = true
		}
	}
	for l := range ja.refDefs {
		if !seen[l] {
			labels = append(labels, l)
			seen[l] = true
		}
	}
	sort.Strings(labels)
	for _, label := range labels {
		enDest, enOk := en.refDefs[label]
		jaDest, jaOk := ja.refDefs[label]
		switch {
		case enOk && !jaOk:
			diffs = append(diffs, fmt.Sprintf("link ref [%s] missing in ja", label))
		case !enOk && jaOk:
			diffs = append(diffs, fmt.Sprintf("link ref [%s] unexpected in ja", label))
		case enDest != jaDest:
			diffs = append(diffs, fmt.Sprintf("link ref [%s] destination differs: en=%q ja=%q", label, enDest, jaDest))
		}
	}

	if d := compareMultiset("link", en.links, ja.links); d != "" {
		diffs = append(diffs, d)
	}

	return diffs
}

func sameIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compareMultiset(name string, a, b []string) string {
	if len(a) != len(b) {
		return fmt.Sprintf("%s count differs: en=%d ja=%d", name, len(a), len(b))
	}
	counts := map[string]int{}
	for _, x := range a {
		counts[x]++
	}
	for _, x := range b {
		counts[x]--
	}
	var diff []string
	for k, v := range counts {
		if v != 0 {
			diff = append(diff, fmt.Sprintf("%s=%+d", k, v))
		}
	}
	if len(diff) == 0 {
		return ""
	}
	sort.Strings(diff)
	return fmt.Sprintf("%s set differs: %s", name, strings.Join(diff, ", "))
}

func collectPairs(enRoot, jaRoot string) ([][2]string, error) {
	var pairs [][2]string
	err := filepath.WalkDir(enRoot, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(enRoot, path)
		if err != nil {
			return err
		}
		slashed := filepath.ToSlash(rel)
		if !strings.HasSuffix(slashed, ".md") {
			return nil
		}
		if !strings.HasPrefix(slashed, "aip/") {
			return nil
		}
		jaPath := filepath.Join(jaRoot, rel)
		if _, err := os.Stat(jaPath); err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				return nil
			}
			return err
		}
		pairs = append(pairs, [2]string{path, jaPath})
		return nil
	})
	return pairs, err
}

func lintPair(enPath, jaPath string) ([]string, error) {
	enSrc, err := os.ReadFile(enPath)
	if err != nil {
		return nil, err
	}
	jaSrc, err := os.ReadFile(jaPath)
	if err != nil {
		return nil, err
	}
	return compareStructures(extractStructure(enSrc), extractStructure(jaSrc)), nil
}

func findProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := cwd
	for {
		if _, err := os.Stat(filepath.Join(dir, "hugo.toml")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("project root (containing hugo.toml) not found from %s", cwd)
		}
		dir = parent
	}
}

func main() {
	root, err := findProjectRoot()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	pairs, err := collectPairs(filepath.Join(root, "content/en"), filepath.Join(root, "content/ja"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	failed := 0
	for _, p := range pairs {
		diffs, err := lintPair(p[0], p[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error linting %s: %v\n", p[1], err)
			os.Exit(2)
		}
		if len(diffs) > 0 {
			fmt.Printf("FAIL %s\n", p[1])
			for _, d := range diffs {
				fmt.Printf("  - %s\n", d)
			}
			failed++
		}
	}

	fmt.Printf("lint-translation: checked %d pairs, %d failed\n", len(pairs), failed)
	if failed > 0 {
		os.Exit(1)
	}
}
