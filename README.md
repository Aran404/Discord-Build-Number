# Discord-Build-Number
Elegant way of getting the discord build number.

```go
package main

import (
	build "Build/Core/Build"
	"fmt"
)

func main() {
	builder, err := build.New("")

	if err != nil {
		panic(err)
	}

	buildNum, err := builder.GetBuildNumber()

	if err != nil {
		panic(err)
	}

	fmt.Println(buildNum)
}
```
```
239004
```
