package util

import "os"

func Files(paths []string) (r []string) {
	for _, path := range paths {
		if stat, err := os.Stat(path); err != nil || stat.IsDir() {
			continue
		}
		r = append(r, path)
	}
	return
}
