/*
This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later
version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see
<https://www.gnu.org/licenses/>.

*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var EnvironmentFilePath *string
var Variable *string

var EnvironmentFile *os.File
var EnvironmentFileContent []string
var EnvironmentFileScanner *bufio.Scanner
var CurrentError error
var EnvironmentVariables map[string]string

func ExitApplication(code int) {
	if code < 0 {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	os.Exit(code)
}

func init() {
	EnvironmentFilePath = flag.String("env-file", "no-file", "This option is mandatory. Full path (including the file) in which environment variables in format <Key>=<Value> are stored.")
	Variable = flag.String("var", "no-variable", "This option is mandatory. Case sensitive key to search for in the environment file.")
	EnvironmentVariables = make(map[string]string)
}

func main() {

	flag.Parse()

	if *EnvironmentFilePath == "no-file" && *Variable == "no-variable" {
		ExitApplication(-1) // TODO: Specify and document specific exit code
	} else if *EnvironmentFilePath == "no-file" && *Variable != "no-variable" {
		ExitApplication(-1) // TODO: Specify and document specific exit code
	} else if *EnvironmentFilePath != "no-file" && *Variable == "no-variable" {
		ExitApplication(-1) // TODO: Specify and document specific exit code
	}

	EnvironmentFile, CurrentError = os.Open(*EnvironmentFilePath)

	if CurrentError != nil {
		// TODO: Write actual error to log / system wide log
		ExitApplication(-1) // TODO: Specify and document specific exit code
	}

	defer EnvironmentFile.Close()
	EnvironmentFileScanner = bufio.NewScanner(EnvironmentFile)
	for EnvironmentFileScanner.Scan() {
		EnvironmentFileContent = append(EnvironmentFileContent, EnvironmentFileScanner.Text())
	}

	for i := 0; i < len(EnvironmentFileContent); i++ {
		var keyValuePair []string = strings.SplitN(EnvironmentFileContent[i], "=", 2)
		if len(keyValuePair) != 2 {
			ExitApplication(-1) // TODO: Specify and document specific exit code
		}
		EnvironmentVariables[keyValuePair[0]] = keyValuePair[1]
	}

	fmt.Fprintf(os.Stdout, "%s", EnvironmentVariables[*Variable])
}
