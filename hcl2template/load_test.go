package hcl2template

import (
	"fmt"
	"testing"

	"github.com/hashicorp/packer/packer"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"

	amazonebs "github.com/hashicorp/packer/builder/amazon/ebs"
	"github.com/hashicorp/packer/builder/virtualbox/iso"

	"github.com/hashicorp/packer/provisioner/file"
	"github.com/hashicorp/packer/provisioner/shell"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func getBasicParser() *Parser {
	return &Parser{
		Parser: hclparse.NewParser(),
		ProvisionersSchemas: mapOfProvisioner(map[string]packer.Provisioner{
			"shell": &shell.Provisioner{},
			"file":  &file.Provisioner{},
		}),
		// PostProvisionersSchemas: mapOfProvisioner(map[string]packer.PostProcessor{
		// 	"amazon-import": &amazon_import.PostProcessor{},
		// }),
		// CommunicatorSchemas: mapOfDecodable(map[string]Decodable{
		// 	"ssh":   &communicator.SSH{},
		// 	"winrm": &communicator.WinRM{},
		// }).Get,
		BuilderSchemas: mapOfBuilder(map[string]packer.Builder{
			"amazon-ebs":     &amazonebs.Builder{},
			"virtualbox-iso": &iso.Builder{},
		}).Get,
	}
}

type mapOfBuilder map[string]packer.Builder

func (mob mapOfBuilder) Get(builder string) (packer.Builder, error) {
	d, found := mob[builder]
	var err error
	if !found {
		err = fmt.Errorf("Unknown entry %s", builder)
	}
	return d, err
}

type mapOfProvisioner map[string]packer.Provisioner

func (mop mapOfProvisioner) Get(provisioner string) (packer.Provisioner, error) {
	p, found := mop[provisioner]
	var err error
	if !found {
		err = fmt.Errorf("Unknown provisioner %s", provisioner)
	}
	return p, err
}

func (mod mapOfProvisioner) List() []string {
	res := []string{}
	for k := range mod {
		res = append(res, k)
	}
	return res
}

func TestParser_ParseFile(t *testing.T) {
	defaultParser := getBasicParser()

	type fields struct {
		Parser *hclparse.Parser
	}
	type args struct {
		filename string
		cfg      *PackerConfig
	}
	tests := []struct {
		name             string
		parser           *Parser
		args             args
		wantPackerConfig *PackerConfig
		wantDiags        bool
	}{
		{
			"valid " + sourceLabel + " load",
			defaultParser,
			args{"testdata/sources/basic.pkr.hcl", new(PackerConfig)},
			&PackerConfig{
				Sources: map[SourceRef]*Source{
					SourceRef{
						Type: "virtualbox-iso",
						Name: "ubuntu-1204",
					}: {
						Type:    "virtualbox-iso",
						Name:    "ubuntu-1204",
						Builder: &iso.Builder{
							// iso.Config{
							// 	HTTPDir:         "xxx",
							// 	ISOChecksum:     "769474248a3897f4865817446f9a4a53",
							// 	ISOChecksumType: "md5",
							// 	RawSingleISOUrl: "http://releases.ubuntu.com/12.04/ubuntu-12.04.5-server-amd64.iso",
							// 	BootCommand:     []string{"..."},
							// 	ShutdownCommand: "echo 'vagrant' | sudo -S shutdown -P now",
							// 	BootWait:        "10s",
							// 	VBoxManage:      [][]string{},
							// 	VBoxManagePost:  [][]string{},
							// },
						},
					},
					SourceRef{
						Type: "amazon-ebs",
						Name: "ubuntu-1604",
					}: {
						Type: "amazon-ebs",
						Name: "ubuntu-1604",
						// Cfg: &amazonebs.FlatConfig{
						// 	RawRegion:            "eu-west-3",
						// 	AMIEncryptBootVolume: boolPtr(true),
						// 	InstanceType:         "t2.micro",
						// 	SourceAmiFilter: &awscommon.FlatAmiFilterOptions{
						// 		Filters: map[string]string{
						// 			"name":                "ubuntu/images/*ubuntu-xenial-{16.04}-amd64-server-*",
						// 			"root-device-type":    "ebs",
						// 			"virtualization-type": "hvm",
						// 		},
						// 		Owners: []string{"099720109477"},
						// 	},
						// 	AMIMappings:    []awscommon.FlatBlockDevice{},
						// 	LaunchMappings: []awscommon.FlatBlockDevice{},
						// },
					},
					SourceRef{
						Type: "amazon-ebs",
						Name: "that-ubuntu-1.0",
					}: {
						Type: "amazon-ebs",
						Name: "that-ubuntu-1.0",
						// Cfg: &amazonebs.FlatConfig{
						// 	RawRegion:            "eu-west-3",
						// 	AMIEncryptBootVolume: boolPtr(true),
						// 	InstanceType:         "t2.micro",
						// 	SourceAmiFilter: &awscommon.FlatAmiFilterOptions{
						// 		MostRecent: boolPtr(true),
						// 	},
						// 	AMIMappings:    []awscommon.FlatBlockDevice{},
						// 	LaunchMappings: []awscommon.FlatBlockDevice{},
						// },
					},
				},
			},
			false,
		},

		{
			"valid " + communicatorLabel + " load",
			defaultParser,
			args{"testdata/communicator/basic.pkr.hcl", new(PackerConfig)},
			&PackerConfig{
				Communicators: map[CommunicatorRef]*Communicator{
					{Type: "ssh", Name: "vagrant"}: {
						Type: "ssh", Name: "vagrant",
						// Cfg: &communicator.FlatSSH{
						// 	SSHUsername:               "vagrant",
						// 	SSHPassword:               "s3cr4t",
						// 	SSHClearAuthorizedKeys:    boolPtr(true),
						// 	SSHHost:                   "sssssh.hashicorp.io",
						// 	SSHHandshakeAttempts:      intPtr(32),
						// 	SSHPort:                   intPtr(42),
						// 	SSHFileTransferMethod:     "scp",
						// 	SSHPrivateKeyFile:         "file.pem",
						// 	SSHPty:                    boolPtr(false),
						// 	SSHTimeout:                "5m",
						// 	SSHAgentAuth:              boolPtr(false),
						// 	SSHDisableAgentForwarding: boolPtr(true),
						// 	SSHBastionHost:            "",
						// 	SSHBastionPort:            intPtr(0),
						// 	SSHBastionAgentAuth:       boolPtr(true),
						// 	SSHBastionUsername:        "",
						// 	SSHBastionPassword:        "",
						// 	SSHBastionPrivateKeyFile:  "",
						// 	SSHProxyHost:              "ninja-potatoes.com",
						// 	SSHProxyPort:              intPtr(42),
						// 	SSHProxyUsername:          "dark-father",
						// 	SSHProxyPassword:          "pickle-rick",
						// 	SSHKeepAliveInterval:      "10s",
						// 	SSHReadWriteTimeout:       "5m",
						// },
					},
				},
			},
			false,
		},

		{
			"duplicate " + sourceLabel, defaultParser,
			args{"testdata/sources/basic.pkr.hcl", &PackerConfig{
				Sources: map[SourceRef]*Source{
					SourceRef{
						Type: "virtualbox-iso",
						Name: "ubuntu-1204",
					}: {
						Type: "virtualbox-iso",
						Name: "ubuntu-1204",
						// Cfg: &iso.FlatConfig{
						// 	HTTPDir: "xxx",
						// },
					},
				},
			},
			},
			&PackerConfig{
				Sources: map[SourceRef]*Source{
					SourceRef{
						Type: "virtualbox-iso",
						Name: "ubuntu-1204",
					}: {
						Type: "virtualbox-iso",
						Name: "ubuntu-1204",
						// Cfg: &iso.FlatConfig{
						// 	HTTPDir: "xxx",
						// },
					},
					SourceRef{
						Type: "amazon-ebs",
						Name: "ubuntu-1604",
					}: {
						Type: "amazon-ebs",
						Name: "ubuntu-1604",
						// Cfg: &amazonebs.FlatConfig{
						// 	RawRegion:            "eu-west-3",
						// 	AMIEncryptBootVolume: boolPtr(true),
						// 	InstanceType:         "t2.micro",
						// 	SourceAmiFilter: &awscommon.FlatAmiFilterOptions{
						// 		Filters: map[string]string{
						// 			"name":                "ubuntu/images/*ubuntu-xenial-{16.04}-amd64-server-*",
						// 			"root-device-type":    "ebs",
						// 			"virtualization-type": "hvm",
						// 		},
						// 		Owners: []string{"099720109477"},
						// 	},
						// 	AMIMappings:    []awscommon.FlatBlockDevice{},
						// 	LaunchMappings: []awscommon.FlatBlockDevice{},
						// },
					},
					SourceRef{
						Type: "amazon-ebs",
						Name: "that-ubuntu-1.0",
					}: {
						Type: "amazon-ebs",
						Name: "that-ubuntu-1.0",
						// Cfg: &amazonebs.FlatConfig{
						// 	RawRegion:            "eu-west-3",
						// 	AMIEncryptBootVolume: boolPtr(true),
						// 	InstanceType:         "t2.micro",
						// 	SourceAmiFilter: &awscommon.FlatAmiFilterOptions{
						// 		MostRecent: boolPtr(true),
						// 	},
						// 	AMIMappings:    []awscommon.FlatBlockDevice{},
						// 	LaunchMappings: []awscommon.FlatBlockDevice{},
						// },
					},
				},
			},
			true,
		},

		{"valid variables load", defaultParser,
			args{"testdata/variables/basic.pkr.hcl", new(PackerConfig)},
			&PackerConfig{
				Variables: PackerV1Variables{
					"image_name": "foo-image-{{user `my_secret`}}",
					"key":        "value",
					"my_secret":  "foo",
				},
			},
			false,
		},

		{"valid " + buildLabel + " load", defaultParser,
			args{"testdata/build/basic.pkr.hcl", new(PackerConfig)},
			&PackerConfig{
				Builds: Builds{
					{
						Froms: BuildFromList{
							{
								Src: SourceRef{"amazon-ebs", "ubuntu-1604"},
							},
							{
								Src: SourceRef{"virtualbox-iso", "ubuntu-1204"},
							},
						},
						ProvisionerGroups: ProvisionerGroups{
							&ProvisionerGroup{
								CommunicatorRef: CommunicatorRef{"ssh", "vagrant"},
								Provisioners: []Provisioner{
									{
										// Cfg: &shell.FlatConfig{
										// 	Inline: []string{"echo '{{user `my_secret`}}' :D"},
										// },
									},
									{
										// Cfg: &shell.FlatConfig{
										// 	Scripts:        []string{"script-1.sh", "script-2.sh"},
										// 	ValidExitCodes: []int{0, 42},
										// },
									},
									{
										// Cfg: &file.FlatConfig{
										// 	Source:      "app.tar.gz",
										// 	Destination: "/tmp/app.tar.gz",
										// },
									},
								},
							},
						},
						PostProvisionerGroups: ProvisionerGroups{
							&ProvisionerGroup{
								Provisioners: []Provisioner{
									{
										// Cfg: &amazon_import.FlatConfig{
										// 	Name: "that-ubuntu-1.0",
										// },
									},
								},
							},
						},
					},
					&Build{
						Froms: BuildFromList{
							{
								Src: SourceRef{"amazon", "that-ubuntu-1"},
							},
						},
						ProvisionerGroups: ProvisionerGroups{
							&ProvisionerGroup{
								Provisioners: []Provisioner{
									{
										// Cfg: &shell.FlatConfig{
										// 	Inline: []string{"echo HOLY GUACAMOLE !"},
										// },
									},
								},
							},
						},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.parser
			f, moreDiags := p.ParseHCLFile(tt.args.filename)
			if moreDiags != nil {
				t.Fatalf("diags: %s", moreDiags)
			}
			diags := p.ParseFile(f, tt.args.cfg)
			if tt.wantDiags == (diags == nil) {
				for _, diag := range diags {
					t.Errorf("PackerConfig.Load() unexpected diagnostics. %v", diag)
				}
				t.Error("")
			}
			if diff := cmp.Diff(tt.wantPackerConfig, tt.args.cfg,
				cmpopts.IgnoreUnexported(cty.Value{}),
				cmpopts.IgnoreUnexported(iso.Builder{}),
				cmpopts.IgnoreTypes(HCL2Ref{}),
				cmpopts.IgnoreTypes([]hcl.Range{}),
				cmpopts.IgnoreTypes(hcl.Range{}),
				cmpopts.IgnoreInterfaces(struct{ hcl.Expression }{}),
				cmpopts.IgnoreInterfaces(struct{ hcl.Body }{}),
			); diff != "" {
				t.Errorf("PackerConfig.Load() wrong packer config. %s", diff)
			}
			if t.Failed() {
				t.Fatal()
			}
		})
	}
}
