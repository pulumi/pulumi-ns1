// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "ns1:index/aPIKey:APIKey":
		r = &APIKey{}
	case "ns1:index/accountWhitelist:AccountWhitelist":
		r = &AccountWhitelist{}
	case "ns1:index/application:Application":
		r = &Application{}
	case "ns1:index/dataFeed:DataFeed":
		r = &DataFeed{}
	case "ns1:index/dataSource:DataSource":
		r = &DataSource{}
	case "ns1:index/dnsview:Dnsview":
		r = &Dnsview{}
	case "ns1:index/monitoringJob:MonitoringJob":
		r = &MonitoringJob{}
	case "ns1:index/notifyList:NotifyList":
		r = &NotifyList{}
	case "ns1:index/pulsarJob:PulsarJob":
		r = &PulsarJob{}
	case "ns1:index/record:Record":
		r = &Record{}
	case "ns1:index/subnet:Subnet":
		r = &Subnet{}
	case "ns1:index/team:Team":
		r = &Team{}
	case "ns1:index/tsigkey:Tsigkey":
		r = &Tsigkey{}
	case "ns1:index/user:User":
		r = &User{}
	case "ns1:index/zone:Zone":
		r = &Zone{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:ns1" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"ns1",
		"index/aPIKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/accountWhitelist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/dataFeed",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/dataSource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/dnsview",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/monitoringJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/notifyList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/pulsarJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/record",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/subnet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/team",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/tsigkey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ns1",
		"index/zone",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"ns1",
		&pkg{version},
	)
}
