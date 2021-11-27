package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"testing"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (tfsdk.Provider, error){
	"scaffolding": func() (tfsdk.Provider, error) {
		return New(), nil
	},
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}
