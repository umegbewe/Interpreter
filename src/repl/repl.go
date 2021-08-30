package repl

import (
	"bufio"
	"fmt"
	"io"
	"lexer"
	"token"
)

const PROMPT = ">> "
