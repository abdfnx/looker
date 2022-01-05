package looker

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func chkStat(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}

	if d.IsDir() {
		return os.ErrPermission
	}

	return nil
}

func hasExt(file string) bool {
	i := strings.LastIndex(file, ".")
	if i < 0 {
		return false
	}

	return strings.LastIndexAny(file, `:\/`) < i
}

func findExecutable(file string, exts []string) (string, error) {
	if len(exts) == 0 {
		return file, chkStat(file)
	}

	if hasExt(file) {
		if chkStat(file) == nil {
			return file, nil
		}
	}

	for _, e := range exts {
		if f := file + e; chkStat(f) == nil {
			return f, nil
		}
	}

	return "", os.ErrNotExist
}

func LookPath(file string) (string, error) {
	var exts []string

	x := os.Getenv(`PATHEXT`)

	if x != "" {
		for _, e := range strings.Split(strings.ToLower(x), `;`) {
			if e == "" {
				continue
			}

			if e[0] != '.' {
				e = "." + e
			}

			exts = append(exts, e)
		}
	} else {
		exts = []string{".com", ".exe", ".bat", ".cmd"}
	}

	if strings.ContainsAny(file, `:\/`) {
		if f, err := findExecutable(file, exts); err == nil {
			return f, nil
		} else {
			return "", &exec.Error{file, err}
		}
	}

	path := os.Getenv("path")

	for _, dir := range filepath.SplitList(path) {
		if f, err := findExecutable(filepath.Join(dir, file), exts); err == nil {
			return f, nil
		}
	}

	return "", &exec.Error{file, exec.ErrNotFound}
}
