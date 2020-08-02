// acmetag is a tool to programmatically interact with acme(1) tag
// bar.  It provides two flag:
//
// 	* -g prints the tag content (mnemonic: get)
// 	* -c clears the tag content (mnemonic: clear)
//
// Any other argument (if passed) will be appended to the tag bar.
//
// Of course, you can combine the flags:
//
// 	acmetag -g -c fmt
//
// BUG(op) it cannot change the text before the | character.
// AFAIK that's not possible
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"9fans.net/go/acme"
)

var (
	cl = flag.Bool("c", false, `Clear the tag`)
	gt = flag.Bool("g", false, `Get the content of the tag`)
)

func open() (*acme.Win, error) {
	winid := os.Getenv("winid")
	id, err := strconv.Atoi(winid)
	if err != nil {
		return nil, err
	}
	win, err := acme.Open(id, nil)
	return win, err
}

func usage() {
	me := os.Args[0]
	fmt.Println(me, "- manage acme(1) tag")
	fmt.Println("Usage:", me, " [-cg] [entries...]")
	fmt.Println(" where entries are words to be added to the acme tag bar")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	win, err := open()
	if err != nil {
		os.Exit(1)
	}
	defer win.CloseFiles()

	if *gt {
		tag, err := win.ReadAll("tag")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(tag))
	}

	if *cl {
		win.Ctl("cleartag")
	}

	sep := ""
	for _, arg := range flag.Args() {
		_, err = win.Write("tag", []byte(sep+arg))
		sep = " "
	}
}
