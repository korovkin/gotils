package gotils

import (
	"bytes"
	"encoding/json"
	"log"
	"runtime"
	"strings"
)

func DebugToJSONStringNoIndentEscaped(o interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")
	err := encoder.Encode(o)
	if err == nil {
		return buffer.String()
	}
	return "{ToJSONString:ERROR}"
}

func DebugRuntimeCallerFuncion(targetName string) string {
	pc := make([]uintptr, 20)
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])
	isReturnNext := false
	for {
		frame, more := frames.Next()
		funcShortName := "failed_to_parse"
		comps := strings.Split(frame.Function, ".")
		if len(comps) > 0 {
			funcShortName = comps[len(comps)-1]
		}

		if isReturnNext {
			return funcShortName
		}

		// return the next one on top of this one:
		if targetName == funcShortName {
			isReturnNext = true
		}

		if !more {
			break
		}
	}

	log.Println("ERROR: DebugCallerFuncion:", targetName)
	return StringNone
}
