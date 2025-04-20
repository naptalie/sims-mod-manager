package fsutil

import (
	"io"
	"os"
	"path/filepath"
)

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Sync()
}

// CopyDir recursively copies a directory from src to dst
func CopyDir(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	
	if !info.IsDir() {
		return CopyFile(src, dst)
	}
	
	if err := os.MkdirAll(dst, info.Mode()); err != nil {
		return err
	}
	
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		
		if entry.IsDir() {
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	
	return nil
}

// EnsureDirExists creates a directory if it doesn't exist
func EnsureDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}