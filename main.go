package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]

	openPath := args[0]
	savePath := args[1]

	file, err := os.Open(openPath)
	check(err)

	info, err := file.Stat()
	check(err)

	reg := regexp.MustCompile(`(?m:.*?\.(\w+)$)`)
	match := reg.FindAllStringSubmatch(info.Name(), -1)
	if match == nil {
		panic(errors.New("No file extension found!"))
	}
	ext := match[0][1] // first match, first group

	imgBytes, err := ioutil.ReadAll(file)
	check(err)

	b64String, err := imgToBase64(imgBytes, ext)

	b64Data := []byte(b64String)

	err = os.WriteFile(savePath, b64Data, 0644)
	check(err)

	fmt.Println("Success!")

}

func imgToBase64(b []byte, ext string) (string, error) {
	var err error = nil
	switch ext {
	case "png":
	case "jpeg":
	case "jpg":
		img, err := jpeg.Decode(bytes.NewReader(b))
		if err != nil {
			return "", err
		}
		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			return "", err
		}
		b = buf.Bytes()
	default:
		return "", errors.New("Invalid file type!")
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(b), err
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
