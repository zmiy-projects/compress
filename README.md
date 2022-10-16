# compress


# install

```bash
go get -u github.com/ZmiyProjects/compress
```

# examples

```go
package main

import (
	"fmt"
	"github.com/ZmiyProjects/compress"
)

func main() {
	result := compress.Compress("aaaabbbcccccaaaa")
	fmt.Println(result) // a8b3c5
}
```