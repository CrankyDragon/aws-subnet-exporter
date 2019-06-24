package cmd

import (
	"fmt"
	"strings"
)

// Output ..
type Output struct {
	data map[string][]string
}

// NewOutput ..
func NewOutput() *Output {
	return &Output{
		data: make(map[string][]string),
	}
}

// Add ...
func (o *Output) Add(key, val string) {
	o.data[key] = append(o.data[key], val)
}

// Render ...
func (o *Output) Render() {
	for key, val := range o.data {
		fmt.Printf("export VPC_%s_SUBNETS='%s'\n", strings.ToUpper(key), strings.Join(val, ","))
	}
}
