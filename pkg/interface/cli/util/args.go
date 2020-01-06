package util

// Allows to run args processing when args is references for some entities
// Example:
// ```shell script
// sentimentd brain rm 1 2 3 // This command will remove brains with ID: 1, 2, 3
// ```
//
func IterateArgs(args []string, f func (string)) {
	for _, arg := range args {
		f (arg)
	}
}
