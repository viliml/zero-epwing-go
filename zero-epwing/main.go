package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	zig "github.com/viliml/zero-epwing-go"
)

type Entry struct {
	Heading string `json:"heading"`
	Text    string `json:"text"`
}

type Subbook struct {
	Title     string  `json:"title"`
	Copyright string  `json:"copyrignt"`
	Entries   []Entry `json:"entries"`
}

type Book struct {
	DiscCode string    `json:"discCode"`
	CharCode string    `json:"charCode"`
	Subbooks []Subbook `json:"subbooks"`
}

func outputBook(bookSrc *zig.Book, path string, pretty bool) error {
	bookDst := Book{DiscCode: bookSrc.DiscCode, CharCode: bookSrc.CharCode}
	for _, subbookSrc := range bookSrc.Subbooks {
		subbookDst := Subbook{Title: subbookSrc.Title, Copyright: subbookSrc.Copyright}
		for _, entrySrc := range subbookSrc.Entries {
			entryDst := Entry{entrySrc.Heading, entrySrc.Text}
			subbookDst.Entries = append(subbookDst.Entries, entryDst)
		}

		bookDst.Subbooks = append(bookDst.Subbooks, subbookDst)
	}

	var (
		data []byte
		err  error
	)

	if pretty {
		data, err = json.MarshalIndent(bookDst, "", "\t")
	} else {
		data, err = json.Marshal(bookDst)
	}

	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0644)
}

func main() {
	var (
		entriesPath   = flag.String("entries-path", "", "output path for dictionary entries")
		entriesPretty = flag.Bool("entries-pretty", false, "pretty-print dictionary entries")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: zero-epwing [options] path\n")
		fmt.Fprintf(os.Stderr, "Parameters:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(2)
	}

	book, err := zig.Load(args[0])
	if err != nil {
		log.Fatal(err)
	}

	if err := outputBook(book, *entriesPath, *entriesPretty); err != nil {
		log.Fatal(err)
	}
}
