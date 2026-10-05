package code

import "code/internal"

func GenDiff(file1, file2, format string) (string, error) {
	return internal.GenDiff(file1, file2, format)
}
