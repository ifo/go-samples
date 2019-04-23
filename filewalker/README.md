## filewalker

A very simple file walker.
It will count up the number of files of the type `.go` in the current directory and all sub directories by default.

### Running

You can run the file walker with `go run main.go`.
If you would like to pass arguments, `go run main.go -filetype ".php"` will show all php files in the current directory and all sub directories.

Including `-v` will give you more information.
Beware, it is very verbose.
