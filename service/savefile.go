package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func Savefile(file *multipart.FileHeader, saveDir string) (string, error) {
	input, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer input.Close() // 确保文件句柄被关闭
	h := md5.New()
	io.Copy(h, input)
	re := h.Sum(nil)
	dst := saveDir + fmt.Sprintf("%x", re) + "__" + file.Filename
	output, err := os.Create(dst)
	if err != nil {
		print(err.Error())
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer output.Close()
	input.Seek(0, 0) //移动回流头部
	_, err = io.Copy(output, input)
	if err != nil {
		print(err.Error())
	}
	return dst, nil
}
