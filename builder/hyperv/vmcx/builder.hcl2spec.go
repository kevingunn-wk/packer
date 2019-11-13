// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package vmcx

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName                *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType              *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug                    *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce                    *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError                  *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars                 map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars            []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	HTTPDir                        *string           `mapstructure:"http_directory" cty:"http_directory"`
	HTTPPortMin                    *int              `mapstructure:"http_port_min" cty:"http_port_min"`
	HTTPPortMax                    *int              `mapstructure:"http_port_max" cty:"http_port_max"`
	ISOChecksum                    *string           `mapstructure:"iso_checksum" required:"true" cty:"iso_checksum"`
	ISOChecksumURL                 *string           `mapstructure:"iso_checksum_url" cty:"iso_checksum_url"`
	ISOChecksumType                *string           `mapstructure:"iso_checksum_type" cty:"iso_checksum_type"`
	RawSingleISOUrl                *string           `mapstructure:"iso_url" required:"true" cty:"iso_url"`
	ISOUrls                        []string          `mapstructure:"iso_urls" cty:"iso_urls"`
	TargetPath                     *string           `mapstructure:"iso_target_path" cty:"iso_target_path"`
	TargetExtension                *string           `mapstructure:"iso_target_extension" cty:"iso_target_extension"`
	BootGroupInterval              *string           `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval"`
	BootWait                       *string           `mapstructure:"boot_wait" cty:"boot_wait"`
	BootCommand                    []string          `mapstructure:"boot_command" cty:"boot_command"`
	OutputDir                      *string           `mapstructure:"output_directory" required:"false" cty:"output_directory"`
	Type                           *string           `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect             *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                        *string           `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                        *int              `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername                    *string           `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword                    *string           `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName                 *string           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName        *string           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys         *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile              *string           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                         *bool             `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                     *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth                   *bool             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding      *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts           *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost                 *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort                 *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth            *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername             *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword             *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile       *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod          *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost                   *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort                   *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername               *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword               *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval           *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout            *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels               []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels                []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey                   []byte            `cty:"ssh_public_key"`
	SSHPrivateKey                  []byte            `cty:"ssh_private_key"`
	WinRMUser                      *string           `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword                  *string           `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                      *string           `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                      *int              `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout                   *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL                    *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure                  *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM                   *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	FloppyFiles                    []string          `mapstructure:"floppy_files" cty:"floppy_files"`
	FloppyDirectories              []string          `mapstructure:"floppy_dirs" cty:"floppy_dirs"`
	FloppyLabel                    *string           `mapstructure:"floppy_label" cty:"floppy_label"`
	DiskBlockSize                  *uint             `mapstructure:"disk_block_size" required:"false" cty:"disk_block_size"`
	RamSize                        *uint             `mapstructure:"memory" required:"false" cty:"memory"`
	SecondaryDvdImages             []string          `mapstructure:"secondary_iso_images" required:"false" cty:"secondary_iso_images"`
	AdditionalDiskSize             []uint            `mapstructure:"disk_additional_size" required:"false" cty:"disk_additional_size"`
	GuestAdditionsMode             *string           `mapstructure:"guest_additions_mode" required:"false" cty:"guest_additions_mode"`
	GuestAdditionsPath             *string           `mapstructure:"guest_additions_path" required:"false" cty:"guest_additions_path"`
	VMName                         *string           `mapstructure:"vm_name" required:"false" cty:"vm_name"`
	SwitchName                     *string           `mapstructure:"switch_name" required:"false" cty:"switch_name"`
	SwitchVlanId                   *string           `mapstructure:"switch_vlan_id" required:"false" cty:"switch_vlan_id"`
	MacAddress                     *string           `mapstructure:"mac_address" required:"false" cty:"mac_address"`
	VlanId                         *string           `mapstructure:"vlan_id" required:"false" cty:"vlan_id"`
	Cpu                            *uint             `mapstructure:"cpus" required:"false" cty:"cpus"`
	Generation                     *uint             `mapstructure:"generation" required:"false" cty:"generation"`
	EnableMacSpoofing              *bool             `mapstructure:"enable_mac_spoofing" required:"false" cty:"enable_mac_spoofing"`
	EnableDynamicMemory            *bool             `mapstructure:"enable_dynamic_memory" required:"false" cty:"enable_dynamic_memory"`
	EnableSecureBoot               *bool             `mapstructure:"enable_secure_boot" required:"false" cty:"enable_secure_boot"`
	SecureBootTemplate             *string           `mapstructure:"secure_boot_template" required:"false" cty:"secure_boot_template"`
	EnableVirtualizationExtensions *bool             `mapstructure:"enable_virtualization_extensions" required:"false" cty:"enable_virtualization_extensions"`
	TempPath                       *string           `mapstructure:"temp_path" required:"false" cty:"temp_path"`
	Version                        *string           `mapstructure:"configuration_version" required:"false" cty:"configuration_version"`
	KeepRegistered                 *bool             `mapstructure:"keep_registered" required:"false" cty:"keep_registered"`
	SkipCompaction                 *bool             `mapstructure:"skip_compaction" required:"false" cty:"skip_compaction"`
	SkipExport                     *bool             `mapstructure:"skip_export" required:"false" cty:"skip_export"`
	Headless                       *bool             `mapstructure:"headless" required:"false" cty:"headless"`
	ShutdownCommand                *string           `mapstructure:"shutdown_command" required:"false" cty:"shutdown_command"`
	ShutdownTimeout                *string           `mapstructure:"shutdown_timeout" required:"false" cty:"shutdown_timeout"`
	CloneFromVMCXPath              *string           `mapstructure:"clone_from_vmcx_path" cty:"clone_from_vmcx_path"`
	CloneFromVMName                *string           `mapstructure:"clone_from_vm_name" cty:"clone_from_vm_name"`
	CloneFromSnapshotName          *string           `mapstructure:"clone_from_snapshot_name" required:"false" cty:"clone_from_snapshot_name"`
	CloneAllSnapshots              *bool             `mapstructure:"clone_all_snapshots" required:"false" cty:"clone_all_snapshots"`
	DifferencingDisk               *bool             `mapstructure:"differencing_disk" required:"false" cty:"differencing_disk"`
	CompareCopy                    *bool             `mapstructure:"copy_in_compare" required:"false" cty:"copy_in_compare"`
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
		"packer_build_name":                &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":              &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                     &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                     &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":                  &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":            &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":       &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":                   &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_port_min":                    &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                    &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"iso_checksum":                     &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"iso_checksum_url":                 &hcldec.AttrSpec{Name: "iso_checksum_url", Type: cty.String, Required: false},
		"iso_checksum_type":                &hcldec.AttrSpec{Name: "iso_checksum_type", Type: cty.String, Required: false},
		"iso_url":                          &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"iso_urls":                         &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"iso_target_path":                  &hcldec.AttrSpec{Name: "iso_target_path", Type: cty.String, Required: false},
		"iso_target_extension":             &hcldec.AttrSpec{Name: "iso_target_extension", Type: cty.String, Required: false},
		"boot_keygroup_interval":           &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"boot_wait":                        &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"boot_command":                     &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"output_directory":                 &hcldec.AttrSpec{Name: "output_directory", Type: cty.String, Required: false},
		"communicator":                     &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":          &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                         &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                         &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                     &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                     &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":                 &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":          &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":        &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":             &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                          &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                      &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":                   &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding":     &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":           &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":                 &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":                 &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":           &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":             &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":             &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_private_key_file":     &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":         &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":                   &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":                   &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":               &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":               &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":          &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":           &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":               &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":                &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":                   &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":                  &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":                   &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":                   &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                       &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                       &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                    &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                    &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":                   &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":                   &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"floppy_files":                     &hcldec.AttrSpec{Name: "floppy_files", Type: cty.List(cty.String), Required: false},
		"floppy_dirs":                      &hcldec.AttrSpec{Name: "floppy_dirs", Type: cty.List(cty.String), Required: false},
		"floppy_label":                     &hcldec.AttrSpec{Name: "floppy_label", Type: cty.String, Required: false},
		"disk_block_size":                  &hcldec.AttrSpec{Name: "disk_block_size", Type: cty.Number, Required: false},
		"memory":                           &hcldec.AttrSpec{Name: "memory", Type: cty.Number, Required: false},
		"secondary_iso_images":             &hcldec.AttrSpec{Name: "secondary_iso_images", Type: cty.List(cty.String), Required: false},
		"disk_additional_size":             &hcldec.AttrSpec{Name: "disk_additional_size", Type: cty.List(cty.Number), Required: false},
		"guest_additions_mode":             &hcldec.AttrSpec{Name: "guest_additions_mode", Type: cty.String, Required: false},
		"guest_additions_path":             &hcldec.AttrSpec{Name: "guest_additions_path", Type: cty.String, Required: false},
		"vm_name":                          &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"switch_name":                      &hcldec.AttrSpec{Name: "switch_name", Type: cty.String, Required: false},
		"switch_vlan_id":                   &hcldec.AttrSpec{Name: "switch_vlan_id", Type: cty.String, Required: false},
		"mac_address":                      &hcldec.AttrSpec{Name: "mac_address", Type: cty.String, Required: false},
		"vlan_id":                          &hcldec.AttrSpec{Name: "vlan_id", Type: cty.String, Required: false},
		"cpus":                             &hcldec.AttrSpec{Name: "cpus", Type: cty.Number, Required: false},
		"generation":                       &hcldec.AttrSpec{Name: "generation", Type: cty.Number, Required: false},
		"enable_mac_spoofing":              &hcldec.AttrSpec{Name: "enable_mac_spoofing", Type: cty.Bool, Required: false},
		"enable_dynamic_memory":            &hcldec.AttrSpec{Name: "enable_dynamic_memory", Type: cty.Bool, Required: false},
		"enable_secure_boot":               &hcldec.AttrSpec{Name: "enable_secure_boot", Type: cty.Bool, Required: false},
		"secure_boot_template":             &hcldec.AttrSpec{Name: "secure_boot_template", Type: cty.String, Required: false},
		"enable_virtualization_extensions": &hcldec.AttrSpec{Name: "enable_virtualization_extensions", Type: cty.Bool, Required: false},
		"temp_path":                        &hcldec.AttrSpec{Name: "temp_path", Type: cty.String, Required: false},
		"configuration_version":            &hcldec.AttrSpec{Name: "configuration_version", Type: cty.String, Required: false},
		"keep_registered":                  &hcldec.AttrSpec{Name: "keep_registered", Type: cty.Bool, Required: false},
		"skip_compaction":                  &hcldec.AttrSpec{Name: "skip_compaction", Type: cty.Bool, Required: false},
		"skip_export":                      &hcldec.AttrSpec{Name: "skip_export", Type: cty.Bool, Required: false},
		"headless":                         &hcldec.AttrSpec{Name: "headless", Type: cty.Bool, Required: false},
		"shutdown_command":                 &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"shutdown_timeout":                 &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"clone_from_vmcx_path":             &hcldec.AttrSpec{Name: "clone_from_vmcx_path", Type: cty.String, Required: false},
		"clone_from_vm_name":               &hcldec.AttrSpec{Name: "clone_from_vm_name", Type: cty.String, Required: false},
		"clone_from_snapshot_name":         &hcldec.AttrSpec{Name: "clone_from_snapshot_name", Type: cty.String, Required: false},
		"clone_all_snapshots":              &hcldec.AttrSpec{Name: "clone_all_snapshots", Type: cty.Bool, Required: false},
		"differencing_disk":                &hcldec.AttrSpec{Name: "differencing_disk", Type: cty.Bool, Required: false},
		"copy_in_compare":                  &hcldec.AttrSpec{Name: "copy_in_compare", Type: cty.Bool, Required: false},
	}
	return s
}
