// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// The Rootly API host. Defaults to https://api.rootly.com
func GetApiHost(ctx *pulumi.Context) string {
	return config.Get(ctx, "rootly:apiHost")
}

// The Rootly API Token. Generate it from your account at https://rootly.com/account
func GetApiToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "rootly:apiToken")
}