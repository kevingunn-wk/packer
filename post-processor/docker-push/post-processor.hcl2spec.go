// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package dockerpush

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Login               *bool             `cty:"login"`
	LoginUsername       *string           `mapstructure:"login_username" cty:"login_username"`
	LoginPassword       *string           `mapstructure:"login_password" cty:"login_password"`
	LoginServer         *string           `mapstructure:"login_server" cty:"login_server"`
	EcrLogin            *bool             `mapstructure:"ecr_login" cty:"ecr_login"`
	AccessKey           *string           `mapstructure:"aws_access_key" required:"false" cty:"aws_access_key"`
	SecretKey           *string           `mapstructure:"aws_secret_key" required:"false" cty:"aws_secret_key"`
	Token               *string           `mapstructure:"aws_token" required:"false" cty:"aws_token"`
	Profile             *string           `mapstructure:"aws_profile" required:"false" cty:"aws_profile"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"login":                      &hcldec.AttrSpec{Name: "login", Type: cty.Bool, Required: false},
		"login_username":             &hcldec.AttrSpec{Name: "login_username", Type: cty.String, Required: false},
		"login_password":             &hcldec.AttrSpec{Name: "login_password", Type: cty.String, Required: false},
		"login_server":               &hcldec.AttrSpec{Name: "login_server", Type: cty.String, Required: false},
		"ecr_login":                  &hcldec.AttrSpec{Name: "ecr_login", Type: cty.Bool, Required: false},
		"aws_access_key":             &hcldec.AttrSpec{Name: "aws_access_key", Type: cty.String, Required: false},
		"aws_secret_key":             &hcldec.AttrSpec{Name: "aws_secret_key", Type: cty.String, Required: false},
		"aws_token":                  &hcldec.AttrSpec{Name: "aws_token", Type: cty.String, Required: false},
		"aws_profile":                &hcldec.AttrSpec{Name: "aws_profile", Type: cty.String, Required: false},
	}
	return s
}
