package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type syncTask struct {
	upstreamDir string
	outputDir   string
}

func (s *syncTask) run() error {
	files, err := s.collectMarkdownFiles()
	if err != nil {
		return err
	}

	if err := s.prepareOutputDir(); err != nil {
		return err
	}

	for _, rel := range files {
		src := filepath.Join(s.upstreamDir, rel)
		dst := filepath.Join(s.outputDir, rel)
		if err := s.copyFile(src, dst); err != nil {
			return err
		}
	}

	fmt.Printf("synced %d files from %s to %s\n", len(files), s.upstreamDir, s.outputDir)
	return nil
}

func (s *syncTask) prepareOutputDir() error {
	if err := os.RemoveAll(s.outputDir); err != nil {
		return err
	}
	return os.MkdirAll(s.outputDir, 0o755)
}

func (s *syncTask) collectMarkdownFiles() ([]string, error) {
	var files []string

	err := filepath.WalkDir(s.upstreamDir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(s.upstreamDir, path)
		if err != nil {
			return err
		}

		if !isContentMarkdown(rel) {
			return nil
		}

		files = append(files, filepath.ToSlash(rel))
		return nil
	})

	return files, err
}

func isContentMarkdown(path string) bool {
	return strings.HasSuffix(path, ".md") &&
		(strings.HasPrefix(path, "aip/") || strings.HasPrefix(path, "pages/"))
}

func (s *syncTask) copyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}

	if _, err := io.Copy(out, in); err != nil {
		_ = out.Close()
		return err
	}

	return out.Close()
}

func main() {
	task := &syncTask{
		upstreamDir: "upstream/google.aip.dev",
		outputDir:   "content/en",
	}

	if err := task.run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
