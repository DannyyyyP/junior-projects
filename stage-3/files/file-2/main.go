package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
)

func main() {
	zipR, err := zip.OpenReader("task.zip")
	if err != nil {
		panic(err)
	}
	for _, file := range zipR.File {
		info := file.FileInfo()
		if info.IsDir() {
			continue
		}
		FindCell(file, info.Size())
	}
}

func FindCell(file *zip.File, size int64) {
	f, err := file.Open()
	if err != nil {
		panic(err)
	}
	buff := make([]byte, size)
	f.Read(buff)
	csvReader := csv.NewReader(bytes.NewReader(buff))
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	if len(records) > 1 {
		fmt.Println(records[4][2])
	}
}
