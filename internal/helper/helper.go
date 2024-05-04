package helper

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
	"strings"
	"time"

	mimeTypeLib "github.com/gabriel-vasile/mimetype"
)

func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveElement[T comparable](element T, sliceList []T) []T {
	for i, v := range sliceList {
		if v == element {
			return append(sliceList[:i], sliceList[i+1:]...)
		}
	}
	return sliceList
}

func FindDeletedElements[T comparable](old, new []T) []T {
	newMap := make(map[T]bool)
	for _, n := range new {
		newMap[n] = true
	}

	var deleted []T
	for _, o := range old {
		if !newMap[o] {
			deleted = append(deleted, o)
		}
	}

	return deleted
}

func ReadFile(file *multipart.FileHeader) (data []byte, fileNameS3 string, extension string, err error) {
	fileData, err := file.Open()
	if err != nil {
		return nil, "", "", err
	}
	defer fileData.Close()

	data, err = io.ReadAll(fileData)
	if err != nil {
		return nil, "", "", err
	}

	var fileRawName string
	now := time.Now().Unix()
	splitFileName := strings.Split(file.Filename, ".")
	if len(splitFileName) <= 1 {
		mimeType := mimeTypeLib.Detect(data)
		fileRawName = file.Filename
		extension = mimeType.Extension()
	} else {
		fileRawName = strings.Join(splitFileName[:len(splitFileName)-1], ".")
		extension = fmt.Sprintf(".%s", strings.Join(splitFileName[len(splitFileName)-1:], "."))
	}
	if fileRawName == "" {
		return nil, "", "", err
	}
	fileNameS3 = fmt.Sprintf("%s-%d%s", fileRawName, now, extension)
	return data, fileNameS3, extension[1:], nil
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsArray(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Array
}

func LowerCaseFieldName(field string) string {
	arrayString := strings.Split(field, ".")
	var newArrayString []string
	for i, v := range arrayString {
		if i == 0 {
			continue
		}
		bts := []byte(v)
		lc := bytes.ToLower([]byte{bts[0]})
		rest := bts[1:]
		newArrayString = append(newArrayString, string(bytes.Join([][]byte{lc, rest}, nil)))
	}
	return strings.Join(newArrayString, ".")
}
