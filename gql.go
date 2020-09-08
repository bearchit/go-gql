package gql

import (
	"bytes"
	"fmt"
	"text/template"
)

type variables map[string]interface{}

func Must(op string, args ...interface{}) string {
	r, err := New(op, args...)
	if err != nil {
		panic(err)
	}
	return r
}

func New(op string, args ...interface{}) (string, error) {
	vars, err := parseArgs(args...)
	if err != nil {
		return "", err
	}

	return query(op, vars)
}

func query(op string, vars variables) (string, error) {
	tpl, err := template.New("query").Parse(op)
	if err != nil {
		return "", err
	}

	var r bytes.Buffer
	if err := tpl.Execute(&r, vars); err != nil {
		return "", err
	}

	return r.String(), nil
}

func parseArgs(args ...interface{}) (variables, error) {
	if len(args)%2 != 0 {
		return nil, fmt.Errorf("args was not consisted of key, value pairs")
	}

	r := make(variables)
	for i := 0; i < len(args); i += 2 {
		k, ok := args[i].(string)
		if !ok {
			return nil, fmt.Errorf("key type must be a string: %v", i)
		}

		r[k] = args[i+1]
	}
	return r, nil
}
