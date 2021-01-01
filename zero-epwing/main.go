package main

import (
	"log"
	"os"

	zig "github.com/FooSoft/zero-epwing-go"
)

func main() {
	book, err := zig.Load(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fp, err := os.Create(os.Args[2])
	defer fp.Close()

	for _, subbook := range book.Subbooks {
		for _, entry := range subbook.Entries {
			fp.WriteString(entry.Heading)
			fp.WriteString("\n")
		}
	}
}
