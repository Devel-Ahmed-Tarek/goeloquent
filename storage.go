package goeloquent

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nfnt/resize"
)

// SaveFile يقوم بحفظ الملف من القارئ داخل المسار المحدد ويعيد المسار الكامل
func SaveFile(basePath, relativePath string, file io.Reader) (string, error) {
	fullPath := filepath.Join(basePath, relativePath)
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return "", err
	}
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}
	return fullPath, nil
}

// GetFile يقوم بقراءة محتوى الملف من المسار المحدد
func GetFile(basePath, relativePath string) ([]byte, error) {
	fullPath := filepath.Join(basePath, relativePath)
	return os.ReadFile(fullPath)
}

// SaveMediaFile تقوم بحفظ ملف وسائط (صورة/فيديو) وتوليد نسخة مصغرة إذا كان صورة
func SaveMediaFile(basePath, relativePath string, file io.Reader, filename string) (map[string]string, error) {
	result := make(map[string]string)
	ext := strings.ToLower(filepath.Ext(filename))
	mediaType := "others"
	if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
		mediaType = "images"
	} else if ext == ".mp4" || ext == ".avi" || ext == ".mov" {
		mediaType = "videos"
	}
	if relativePath == "" {
		relativePath = mediaType
	}
	// قراءة الملف بالكامل
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	data := buf.Bytes()
	origFileName := GenerateFileName(filename)
	origPath := filepath.Join(basePath, relativePath, origFileName)
	if err := os.MkdirAll(filepath.Dir(origPath), os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}
	err = os.WriteFile(origPath, data, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to save original file: %v", err)
	}
	result["original"] = origPath

	// إذا كان الملف صورة، نقوم بإنشاء نسخة مصغرة
	if mediaType == "images" {
		img, format, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			return nil, fmt.Errorf("failed to decode image: %v", err)
		}
		thumbImg := resize.Resize(200, 0, img, resize.Lanczos3)
		thumbFileName := strings.TrimSuffix(origFileName, ext) + "_thumb" + ext
		thumbPath := filepath.Join(basePath, relativePath, thumbFileName)
		thumbFile, err := os.Create(thumbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create thumbnail file: %v", err)
		}
		defer thumbFile.Close()
		if format == "png" {
			err = png.Encode(thumbFile, thumbImg)
		} else {
			err = jpeg.Encode(thumbFile, thumbImg, &jpeg.Options{Quality: 80})
		}
		if err != nil {
			return nil, fmt.Errorf("failed to encode thumbnail: %v", err)
		}
		result["thumbnail"] = thumbPath
	}
	return result, nil
}

// ListMediaFileVersions يبحث عن كل النسخ المتوفرة للملف
func ListMediaFileVersions(basePath, relativePath, fileName string) (map[string]string, error) {
	result := make(map[string]string)
	origPath := filepath.Join(basePath, relativePath, fileName)
	if _, err := os.Stat(origPath); err == nil {
		result["original"] = origPath
	}
	ext := filepath.Ext(fileName)
	base := fileName[:len(fileName)-len(ext)]
	thumbName := base + "_thumb" + ext
	thumbPath := filepath.Join(basePath, relativePath, thumbName)
	if _, err := os.Stat(thumbPath); err == nil {
		result["thumbnail"] = thumbPath
	}
	mediumName := base + "_medium" + ext
	mediumPath := filepath.Join(basePath, relativePath, mediumName)
	if _, err := os.Stat(mediumPath); err == nil {
		result["medium"] = mediumPath
	}
	return result, nil
}

// GetMediaFileVersion يقوم بقراءة محتوى النسخة المطلوبة من الملف
func GetMediaFileVersion(basePath, relativePath, fileName, versionKey string) ([]byte, error) {
	versions, err := ListMediaFileVersions(basePath, relativePath, fileName)
	if err != nil {
		return nil, err
	}
	path, ok := versions[strings.ToLower(versionKey)]
	if !ok {
		return nil, fmt.Errorf("version %s not found", versionKey)
	}
	return os.ReadFile(path)
}
