textbelt.com API for GO  [![GoDoc](https://godoc.org/github.com/dietsche/textbelt?status.png)](https://godoc.org/github.com/dietsche/textbelt) [![Build Status](https://travis-ci.org/dietsche/textbelt.svg)](https://travis-ci.org/dietsche/textbelt)
==================

This library sends text messages using http://textbelt.com

## Installing

### go get
    $ go get github.com/dietsche/textbelt

### Example Code
    import (
        "github.com/dietsche/textbelt"
    )
    
    func main() {
        texter := textbelt.New()
        phoneToText := "123-456-7890"
        if err := texter.Text(phoneToText, "txt msg from go!"); err != nil {
            painc(err)
        }
    }
