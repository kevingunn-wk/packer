// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package lxd

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
	OutputImage         *string           `mapstructure:"output_image" required:"false" cty:"output_image"`
	ContainerName       *string           `mapstructure:"container_name" cty:"container_name"`
	CommandWrapper      *string           `mapstructure:"command_wrapper" required:"false" cty:"command_wrapper"`
	Image               *string           `mapstructure:"image" required:"true" cty:"image"`
	Profile             *string           `mapstructure:"profile" cty:"profile"`
	InitSleep           *string           `mapstructure:"init_sleep" required:"false" cty:"init_sleep"`
	PublishProperties   map[string]string `mapstructure:"publish_properties" required:"false" cty:"publish_properties"`
	LaunchConfig        map[string]string `mapstructure:"launch_config" required:"false" cty:"launch_config"`
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
		"output_image":               &hcldec.AttrSpec{Name: "output_image", Type: cty.String, Required: false},
		"container_name":             &hcldec.AttrSpec{Name: "container_name", Type: cty.String, Required: false},
		"command_wrapper":            &hcldec.AttrSpec{Name: "command_wrapper", Type: cty.String, Required: false},
		"image":                      &hcldec.AttrSpec{Name: "image", Type: cty.String, Required: false},
		"profile":                    &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"init_sleep":                 &hcldec.AttrSpec{Name: "init_sleep", Type: cty.String, Required: false},
		"publish_properties":         &hcldec.BlockAttrsSpec{TypeName: "publish_properties", ElementType: cty.String, Required: false},
		"launch_config":              &hcldec.BlockAttrsSpec{TypeName: "launch_config", ElementType: cty.String, Required: false},
	}
	return s
}
