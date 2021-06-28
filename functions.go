package expr

import "fmt"

type Function func(args ...interface{}) (interface{}, error)

var functions map[string]Function = map[string]Function{
	"sum": func(args ...interface{}) (interface{}, error) {
		var ret float64
		for _, arg := range args {
			argInstance, ok := arg.(float64)
			if !ok {
				return nil, fmt.Errorf("param is not a number")
			}

			ret += argInstance
		}

		return ret, nil
	},
}
