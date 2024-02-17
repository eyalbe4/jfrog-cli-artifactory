package main

import (
	"github.com/eyalbe4/jfrog-cli-artifactory/cli"
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
)

func main() {
	plugins.PluginMain(cli.GetJfrogCliArtifactoryApp())
}
