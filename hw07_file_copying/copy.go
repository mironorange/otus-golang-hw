package main

import (
	"errors"
	"io"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Если не удалось открыть файл, то и скопировать его нельзя
	dst, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	dstStats, _ := dst.Stat()

	// Offset больше, чем размер файла - невалидная ситуация
	if offset > dstStats.Size() {
		return ErrOffsetExceedsFileSize
	}

	// Limit равный нулю, говорит об отсутствии ограничений
	if limit == 0 {
		limit = dstStats.Size()
	}

	bar := pb.Full.Start64(dstStats.Size())

	// Если записать файл нельзя, то и операция не выполнится
	src, err := os.Create(toPath)
	srcWriter := bar.NewProxyWriter(src)
	if err != nil {
		return ErrUnsupportedFile
	}

	dst.Seek(offset, io.SeekStart)
	totalBytesCopied := 0

	buf := make([]byte, 1*1024)
	for {
		readSize, err := dst.Read(buf)
		if totalBytesCopied+readSize > int(limit) {
			readSize = int(limit) - totalBytesCopied
		}
		writeSize, _ := srcWriter.Write(buf[0:readSize])
		totalBytesCopied += writeSize
		if totalBytesCopied == int(limit) {
			break
		}
		if err == io.EOF {
			break
		}
	}

	bar.Finish()
	defer dst.Close()
	defer src.Close()
	return nil
}

func iOpenAndReadAllFile(dstFilename string) ([]byte, error) {
	file, err := os.Open(dstFilename)
	if err != nil {
		return nil, err
	}
	buf, err := io.ReadAll(file)
	file.Close()
	return buf, err
}
