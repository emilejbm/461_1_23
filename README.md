# 461_1_23
Team Members:
Mimi Chon, Anna Shen, Emile Baez, Ben Schwartz

# Building program
In command line, first run 
`go build -o "run"`
in order to create executable with 'run' as name.

# CLI
<`build`, `install`, `test`, `'URL_FILE'`> commands are recognized where 'URL_FILE' must be an absolute path to a file in the system.

# Architecture
![Architecture](resources/arch.jpg)   
This is the current architecture of our program. Each block represents a collection of functions towards a single functionality. Each color represents a package in Go.

# Commands
`go run . install` to install Go and Python dependencies  

To build and run the executable  
`go run . build`  
`./metric_cli URL_FILE`

Or just run directly  
`go run . URL_FILE`

# Used Stackoverflow articles
Module : Link  
cli/scanner : https://stackoverflow.com/questions/18159704/how-to-debug-exit-status-1-error-when-running-exec-command-in-golang  
worker : https://stackoverflow.com/questions/25306073/always-have-x-number-of-goroutines-running-at-any-time  
worker : https://stackoverflow.com/questions/55203251/limiting-number-of-go-routines-running  
scanner : https://stackoverflow.com/questions/2466735/how-to-sparsely-checkout-only-one-single-file-from-a-git-repository