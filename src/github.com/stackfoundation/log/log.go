package log

import "fmt"

var debug bool

// JsonOutput Is JSON output enabled?
var JSONOutput bool

// Error Log error output
func Error(code string, message string, arguments ...interface{}) {
	if JSONOutput {
		fmt.Printf("{\"code\":\"%v\",\"message\":\"", code)
		fmt.Printf(message, arguments...)
		fmt.Printf("\"}")
		fmt.Println()
	} else {
		fmt.Printf(message, arguments...)
		fmt.Println()
	}
}

// Debug Log debug output
func Debug(code string, message string, arguments ...interface{}) {
	if debug || JSONOutput {
		if JSONOutput {
			fmt.Printf("{\"code\":\"%v\",\"message\":\"", code)
			fmt.Printf(message, arguments...)
			fmt.Printf("\"}")
			fmt.Println()
		} else {
			fmt.Printf(message, arguments...)
			fmt.Println()
		}
	}
}

// SetDebug Set debug mode on or off
func SetDebug(isDebug bool) {
	debug = isDebug
}
