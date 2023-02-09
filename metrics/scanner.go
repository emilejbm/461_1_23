package metrics

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func ScanRepo(url string) (float64, error) {
	app := "./venv/bin/python3"
	args := []string{"metrics/scanner.py", url}
	exec_output := exec.Command(app, args...)
	fmt.Println("scanner: starting process - ", exec_output)
	stdout, err := exec_output.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("scanner: process error - %s: %s", err.Error(), string(stdout))
	}

	value, err := strconv.ParseFloat(strings.TrimSpace(string(stdout)), 64)
	if err != nil {
		return 0, fmt.Errorf("scanner: conversion error - %s", err.Error())
	}

	fmt.Println("scanner: retrieved value - ", value)
	return value, nil
}
