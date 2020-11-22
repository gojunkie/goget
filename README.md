# Requirements
- Go >= [1.0](https://golang.org/dl/)

# Installation
```SH
$ go get github.com/gojunkie/goget
```

# Quick Start
```Go
package main

import (
  "log"
  "os"

  "github.com/gojunkie/goget"
)

func main() {
  var (
		maxConn  = 8
		filePath = "sample.mp4"
		fileURL  = "https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4"
	)

	svc := goget.New(maxConn)
	ch, err := svc.Download(fileURL, filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for info := range ch {
		// on download progress code
	}
}
```

Downloaded file will be save in specified directory.
This package also comes with CLI that you can use.

```SH
$ go install github.com/gojunkie/goget/cmd/goget
$ goget https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4 -n 2 -o $(pwd)/tmp/sample.mp4
```
