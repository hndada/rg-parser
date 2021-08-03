# rg-parser
 Rhythm game file parsers

## Usage
    go get github.com/hndada/rg-parser

## To do
* Add available extensions
    - [ ] BMS files (.bms, ...)
    - [ ] o2jam files (.ojn, .ojm)
    - [ ] stepmania files

* Move nofloat, delimiter tag from comment to `struct tag` 

### Protocol
1. Write `Format` struct and its field structs, tags except `json` manually
2. go run gencode.go, then paste to `format.go`
3. go run genjsontag.go, then paste to `Format` struct
 
## Reference
[.osu (file format)](https://osu.ppy.sh/help/wiki/osu!_File_Formats/Osu_(file_format))

[.osr (file format)](https://osu.ppy.sh/help/wiki/osu!_File_Formats/Osr_(file_format))

[The OJN Documentation](https://open2jam.wordpress.com/the-ojn-documentation/)
