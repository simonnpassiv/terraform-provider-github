package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/integrations/terraform-provider-githubalt/v6/github"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: github.Provider})
}
