# FuzMatch, an approximate string matching in golang

### About

fuzmatch is a library inspired by the fuzzywuzzy python library. I've used the same function name so if you want to know more about how the function works, go see the [SeatGeek blog](http://chairnerd.seatgeek.com/fuzzywuzzy-fuzzy-string-matching-in-python/) post about fuzzywuzzy.


#### WARNING
The algorithm used in this library is the levenshtein distance but the fuzzywuzzy package use the python-Levenshtein who has a special function for the ratio where the replace operation costs 2 (in my algorithm it costs 1). So if you compare fuzzywuzzy and fuzmatch there mays be a few differences.


### Installation

    go get github.com/charlesvdv/fuzmatch

You need of course to set your GOPATH, otherwise you will have an error.

### Usage

A simple ratio.

    fuzmatch.Ratio("book", "back")
    "0.75"

A partial ratio.

    fuzmatch.PartialRatio("hello world!","hello")
    "1"

A token sort ratio.

    fuzmatch.TokenSortRatio("Rust vs Golang", "Golang vs Rust")
    "1"

A token set ratio.

    fuzmatch.TokenSetRatio("Rust from Mozilla vs Go from Google's employees", "Rust vs Go")
    "1"


If you want more informations on how the function works, go see the [SeatGeek blog](http://chairnerd.seatgeek.com/fuzzywuzzy-fuzzy-string-matching-in-python/).

### To-Do
- Benchmarks
- More unit tests

### Contribute

Pull requests, commits are welcome !
