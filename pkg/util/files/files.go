package files

import (
	"io"
	"os"
)

// FileExists 判断所给路径文件/文件夹是否存在
func FileExists(firePath string) bool {
	_, err := os.Stat(firePath) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(firePath string) bool {
	s, err := os.Stat(firePath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(firePath string) bool {
	return !IsDir(firePath)
}

// WriteFileByCoverContent 覆盖内容写入文件
func WriteFileByCoverContent(firePath string, content string) (err error) {
	f, err := os.OpenFile(firePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if nil != err {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)
	_, err = io.WriteString(f, content)
	if err != nil {
		return err
	}
	return nil
}

// WriteFileByAddContent 追加写入文件
func WriteFileByAddContent(firePath, content string) (err error) {
	// 打开文件并设置文件指针到文件末尾
	file, err := os.OpenFile(firePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// 写入内容到文件
	if _, err := io.WriteString(file, content); err != nil {
		return err
	}

	// 关闭文件
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}
