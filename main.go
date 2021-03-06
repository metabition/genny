package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(`genny has moved to https://github.com/cheekybits/genny

Run this:

  cd .. ; rm -rf $GOPATH/src/github.com/metabition/genny ; rm $GOPATH/bin/genny ; go get github.com/cheekybits/genny ; cd $GOPATH/src/github.com/cheekybits/genny ; pwd

Or do this:

  1. Delete $GOPATH/src/github.com/metabition/genny
  2. Delete $GOPATH/bin/genny
  3. go get github.com/cheekybits/genny

Sorry, we won't move it again :)
    `)
	os.Exit(1)
}
