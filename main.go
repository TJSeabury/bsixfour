package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	inPathPtr := flag.String("in", "", "The input path.")
	outPathPtr := flag.String("out", "", "The output path.")
	htmlPtr := flag.Bool("html", true, "Output as html or txt file.")
	rawPtr := flag.Bool("raw", false, "Enable direct output.")

	flag.Parse()

	// Ensure inPath ok
	if *inPathPtr == "" {
		panic(errors.New("Must specify a read file path!"))
	}
	openPath, err := filepath.Abs(*inPathPtr)
	check(err)
	file, err := os.Open(openPath)
	check(err)

	// ok, lets extract all the useful parts from inPath
	// see: https://regex101.com/r/qZouIJ/1
	reg := regexp.MustCompile(`(?m:(.*?)\\(\w+)\.(\w+)$)`)
	match := reg.FindAllStringSubmatch(file.Name(), -1)
	if match == nil {
		panic(errors.New("No file extension found!"))
	}
	path := match[0][1] // first match, first group, the directory path
	name := match[0][2] // first match, second group, the file name
	ext := match[0][3]  // first match, third group, the extension

	// Ensure outPath ok or create from inPath
	var savePath string
	if *outPathPtr == "" {
		savePath = path + string(os.PathSeparator) + name
	} else {
		savePath = *outPathPtr
	}

	imgBytes, err := ioutil.ReadAll(file)
	check(err)

	var b64Data []byte
	var htmlString string
	b64String, err := imgToBase64(imgBytes, ext)
	if *rawPtr == false {
		if *htmlPtr == true {
			htmlString = "<img src=\"data:image/png;base64," + b64String + "\" alt=\"\">"
			b64Data = []byte(htmlString)
			savePath += ".html"
		} else {
			b64Data = []byte(b64String)
			savePath += ".txt"
		}
		err = os.WriteFile(savePath, b64Data, 0644)
		check(err)
	} else {
		fmt.Println(b64String)
	}
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
	return base64.StdEncoding.EncodeToString(b), err
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
