package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(`genny has moved to https://github.com/cheekybits/genny

Run this:

  cd .. ; rm -r $GOPATH/src/github.com/metabition/genny ; rm $GOPATH/bin/genny ; go get https://github.com/cheekybits/genny
  
Or do this:

  1. Delete $GOPATH/src/github.com/metabition/genny
  2. Delete $GOPATH/bin/genny
  3. go get https://github.com/cheekybits/genny

Sorry, we won't move it again :)
    `)
	os.Exit(1)
}
