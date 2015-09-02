textbelt.com API in GO  [![GoDoc](https://godoc.org/github.com/dietsche/textbelt?status.png)](https://godoc.org/github.com/dietsche/textbelt)
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
        yourPhone := "123-456-7890"
        if err := texter.Text(yourPhone, "txt msg from go!"); err != nil {
            painc(err)
        }
    }
