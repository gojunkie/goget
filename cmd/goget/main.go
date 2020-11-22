package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	pb "github.com/cheggaaa/pb/v3"
	"github.com/gojunkie/goget"
)

var (
	fs       = flag.NewFlagSet("goget", flag.ExitOnError)
	maxConn  = fs.Int("n", runtime.NumCPU(), "Max connections")
	fileName = fs.String("o", "", "Output file path")
)

const usage = `Usage: goget [url] [options]

Options:
  -n Max connections
  -o Output file path`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}
	err := fs.Parse(os.Args[2:])
	handleError(err)

	fileURL := os.Args[1]
	if *fileName == "" {
		dir, err := os.Getwd()
		handleError(err)
		*fileName = filepath.Join(dir, filepath.Base(fileURL))
	}

	svc := goget.New(*maxConn)
	ch, err := svc.Download(fileURL, *fileName)
	handleError(err)

	var bar *pb.ProgressBar
	for p := range ch {
		if bar == nil {
			bar = pb.Start64(p.TotalSize)
			bar.Set(pb.Bytes, true)
		}
		bar.Add64(p.ChunkSize)
	}

	bar.Finish()
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
