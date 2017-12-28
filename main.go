package main

import (
	"log"

	git "gopkg.in/src-d/go-git.v4"
)

func main() {
	options := &git.CloneOptions{
		URL: "https://github.com/theclickpro/dep-vendor-test.git",
	}

	_, err := git.PlainClone("/tmp/dummy-path", false, options)
	if err != nil {
		log.Fatal(err)
	}
}
