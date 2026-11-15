package internal

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func zipFiles(zipName string, paths ...string) error {
	// #nosec 304
	zipOutput, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer func() { _ = zipOutput.Close() }()

	archive := zip.NewWriter(zipOutput)
	defer func() { _ = archive.Close() }()

	for _, targetPath := range paths {
		info, err := os.Stat(targetPath)
		if err != nil {
			return fmt.Errorf("failed to stat %s: %w", targetPath, err)
		}

		if info.IsDir() {
			baseDir := filepath.Dir(targetPath)
			err = filepath.WalkDir(targetPath, func(filePath string, d os.DirEntry, walkErr error) error {
				if walkErr != nil {
					return walkErr
				}
				if d.IsDir() {
					return nil
				}

				fileInfo, err := d.Info()
				if err != nil {
					return err
				}

				relPath, err := filepath.Rel(baseDir, filePath)
				if err != nil {
					return err
				}

				return addFileToZip(archive, filePath, relPath, fileInfo)
			})
			if err != nil {
				return fmt.Errorf("failed to walk directory %s: %w", targetPath, err)
			}
		} else {
			relName := filepath.Base(targetPath)
			if err := addFileToZip(archive, targetPath, relName, info); err != nil {
				return err
			}
		}
	}

	return nil
}

func addFileToZip(archive *zip.Writer, srcPath, internalPath string, info os.FileInfo) error {
	// #nosec 304
	file, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	// Create a zip header from the file info, which preserves permissions and mod time
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// FileInfoHeader only uses the base name by default; override Name with the relative path
	header.Name = filepath.ToSlash(internalPath)

	// Use Deflate method so files are compressed (FileInfoHeader defaults to Store/uncompressed)
	header.Method = zip.Deflate

	writer, err := archive.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
