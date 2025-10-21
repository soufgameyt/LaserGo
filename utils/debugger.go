package utils

import (
	"encoding/json"
	"fmt"
)

type debugger struct{}

var colors = map[string]string{ // yes \x1b moment
	"RESET": "\x1b[0m",
	"INFO":  "\x1b[96;1m",
	"WARN":  "\x1b[93;1m",
	"ERROR": "\x1b[91;1m",
	"DEBUG": "\x1b[95;1m",
}

var DebuggerInst = &debugger{}

func (d *debugger) log(level string, msgs ...interface{}) {
	parts := make([]string, len(msgs))
	for i, m := range msgs {
		switch v := m.(type) {
		case string:
			parts[i] = v
		default:
			b, _ := json.Marshal(v)
			parts[i] = string(b)
		}
	}
	fmt.Println(fmt.Sprintf("%s[%s] %s%s", colors[level], level, join(parts, " "), colors["RESET"]))
}

func join(arr []string, sep string) string {
	res := ""
	for i, s := range arr {
		res += s
		if i < len(arr)-1 {
			res += sep
		}
	}
	return res
}

// as i cant use my own conventions, i'll just shitcode everything
func (d *debugger) Info(msgs ...interface{})  { d.log("INFO", msgs...) }
func (d *debugger) Warn(msgs ...interface{})  { d.log("WARN", msgs...) }
func (d *debugger) Error(msgs ...interface{}) { d.log("ERROR", msgs...) }
func (d *debugger) Debug(msgs ...interface{}) { d.log("DEBUG", msgs...) }
