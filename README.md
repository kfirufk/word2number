# fork

this a fork of https://github.com/donna-legal/word2number

# word2number

This is a library to convert words (three hundred-thousand) to numbers (300,000) in Go

# what changed

1. I added support of getting an array of numbers, so if you have `take one glass and pour one and a half cups of sugar in it`
it will provide an array of []float64{1,1.5} instead of summing it to two.
2. added support to get to replace number in word form to the actual number in a string and return it, so `take one glass and pour one and a half cups of sugar in it`  would be converted to `take 1 glass and pour 1.5 cups of sugar in it`
3. added support for `and a half` and `1/2` syntax (currently only in english, still didn't translate)

## Installation

Go get the package:

`go get https://github.com/kfirufk/word2number`

and add it as an import:

```golang
import "github.com/kfirufk/word2number"
```

## Usage

```golang
package main

import (
    "fmt"
    "github.com/kfirufk/word2number"
)

func main() {
    converter, err := word2number.NewConverter("en")
    if err != nil {
        panic(err)
    }
    var f float64
	var s string
	var fa []float64
    f = converter.Words2Number("two-thousand seventy-five")
    fmt.Println(f) // should return 2075
    f = converter.Words2Number("one-million two hundred thousand")
    fmt.Println(f) // should return 1200000
	s = converter.ReplaceNumbersInWordForm("take one glass and pour one and a half cups of sugar in it")
	// should return "take 1 glass and pour 1.5 cups of sugar in it"
	fa = converter.Words2NumberArray("take one glass and pour one and a half cups of sugar in it")
	// should return []float64{1,1.5}
	f = converter.Words2Number("take one glass and pour one and a half cups of sugar in it")
	// this should return 2.5 since it sums 1 and 1.5

}
```

## Now and the future

Look in the test cases what works and what doesn't.
Most things to the left of the decimal point should work.

Needs improvement:

* Decimal numbers. The simpler cases work just fine (eg. _one point three hundredths_ = 1.03), but there are quite a few failing test cases
* Things like _One point two billion_ doesn't check out at the moment either.

Extensions:

* More languages. It would be somewhat easy to extend with other languages that are constructed the same way, like Swedish or Spanish.
* Misspellings. One could extend the resources with common misspellings.
* Concatenations. Some people might write seventyfive instead of seventy-five. This is definitely necessary for Spanish and Swedish.

## Contribute

If you use this program and find short-comings. Please start a fork or add issues. We appreciate it deeply.
