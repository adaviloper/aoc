/**
 * Package print
 */
package printUtils

import "fmt"

const (
    Reset  = "\x1b[0m"
    Red    = "\x1b[31m"
    Green  = "\x1b[32m"
    Yellow = "\x1b[33m"
    Orange = "\x1b[38;5;208m"
    Bold   = "\x1b[1m"
)

func Danger(msg string) {
    fmt.Println(Bold + Red + msg + Reset)
}

func Info(msg string) {
    fmt.Println(Bold + Yellow + msg + Reset)
}

func Success(msg string) {
    fmt.Println(Bold + Green + msg + Reset)
}

func Warn(msg string) {
    fmt.Println(Bold + Orange + msg + Reset)
}

