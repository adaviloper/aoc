/**
 * Package print
 */
package printUtils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
    "github.com/adaviloper/aoc/config"
)

func Danger(msg string) {
	r, g, b, _ := HexToRGB(config.Cfg.Theme.Danger)
    color.RGB(r, g, b).Println(msg)
}

func Info(msg string) {
	r, g, b, _ := HexToRGB(config.Cfg.Theme.Info)
    color.RGB(r, g, b).Println(msg)
}

func Success(msg string) {
	r, g, b, _ := HexToRGB(config.Cfg.Theme.Success)
    color.RGB(r, g, b).Println(msg)
}

func Warn(msg string) {
	r, g, b, _ := HexToRGB(config.Cfg.Theme.Warn)
    color.RGB(r, g, b).Println(msg)
}

func HexToRGB(s string) (int, int, int, bool) {                            
    s = strings.TrimSpace(s)
    s = strings.TrimPrefix(s, "#")
    if len(s) == 3 {
        s = fmt.Sprintf("%c%c%c%c%c%c", s[0], s[0], s[1], s[1], s[2], s[2])
    }
    if len(s) != 6 {
    	return HexToRGB("#cdd6f4")
    }
    r, err1 := strconv.ParseUint(s[0:2], 16, 8)
    g, err2 := strconv.ParseUint(s[2:4], 16, 8)
    b, err3 := strconv.ParseUint(s[4:6], 16, 8)
    if err1 != nil || err2 != nil || err3 != nil {
    	return HexToRGB("#cdd6f4")
    }

    return int(r), int(g), int(b), true
}                                                                          
