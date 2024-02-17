package cmd1

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

func GetDescription() string {
	return "Scan files located on the local file-system with Xray."
}

func GetArguments() []components.Argument {
	return []components.Argument{{Name: "source pattern", Description: `Specifies the local file system path of the files to be scanned.
		You can specify multiple files by using wildcards, Ant pattern or a regular expression.
		If you have specified that you are using regular expressions, then the first one used in the argument must be enclosed in parenthesis.`}}
}
