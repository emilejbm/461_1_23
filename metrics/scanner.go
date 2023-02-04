package metrics

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func CallPython(url string) float64 {
	app := "./venv/bin/python3"
	args := []string{"metrics/scanner.py", url}
	exec_output := exec.Command(app, args...)
	fmt.Println("scanner: starting process - ", exec_output)
	stdout, err := exec_output.CombinedOutput()
	if err != nil {
		fmt.Println("scanner: process error - ", fmt.Sprint(err)+": "+string(stdout))
		return 0
	}

	value, err := strconv.ParseFloat(strings.TrimSpace(string(stdout)), 64)
	if err != nil {
		fmt.Println("scanner: conversion error - ", fmt.Sprint(err))
		return 0
	}

	fmt.Println("scanner: retrieved value - ", value)
	return value
}
