package helpers

import (
	"fmt"
	"github.com/wplib/deploywp/only"
	"strings"
)

// Usage: {{ if ExecParseOutput $output "uid=%s" "mick" ... }}YES{{ end }}
func Contains(s interface{}, substr interface{}) bool {
	var ret bool

	for range only.Once {
		sp := ReflectString(s)
		if sp == nil {
			break
		}

		ssp := ReflectString(substr)
		if ssp == nil {
			break
		}

		ret = strings.Contains(*sp, *ssp)
	}

	return ret
}


// Usage: {{ Sprintf "uid=%s" "mick" ... }}
func Sprintf(format interface{}, a ...interface{}) string {
	var ret string

	for range only.Once {
		p := ReflectString(format)
		if p == nil {
			break
		}
		ret = fmt.Sprintf(*p, a...)
	}

	return ret
}
