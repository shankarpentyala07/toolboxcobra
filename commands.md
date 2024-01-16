export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

bin is the folder where packages are installed.

cobra-cli init toolbox

Add a command:
cobra-cli add [command name]
cobra-cli add ping
This creates a ping.go file in the directory where the cobra-cli add command is being run

Verify with go run main.go
Example:
 go run main.go 
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  toolbox [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ping        A brief description of your command

Flags:
  -h, --help     help for toolbox
  -t, --toggle   Help message for toggle

Use "toolbox [command] --help" for more information about a command.