# compile

This is trying to be a compiler, maybe it will happen someday

tokenizer in `token.go`
ast definition in `ast.go`
parser in `parse.go`

`go generate` to rebuild the `*_string.go` files

For now, it just interprets arithmetic expressions, though the tokenizer
supports a little more than that.
