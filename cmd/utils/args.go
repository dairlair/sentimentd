package utils

import (
	"fmt"
	"strconv"
)

// Allows to run args processing when args is references for some entities
// Example:
// ```shell script
// sentimentd brain rm 1 2 3 // This command will remove brains with ID: 1, 2, 3
// ```
//
func IterateArgs(args []string, f func (int64)) {
	for _, arg := range args {
		id, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			fmt.Printf("Error: %s is invalid reference", arg)
			continue
		}
		f (id)
	}
}
