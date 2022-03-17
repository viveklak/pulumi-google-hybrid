// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-google-native/sdk/go/google"
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
	case "google-native:apigee/v1:Alias":
		r = &Alias{}
	case "google-native:apigee/v1:Api":
		r = &Api{}
	case "google-native:apigee/v1:ApiProduct":
		r = &ApiProduct{}
	case "google-native:apigee/v1:App":
		r = &App{}
	case "google-native:apigee/v1:ArchiveDeployment":
		r = &ArchiveDeployment{}
	case "google-native:apigee/v1:CanaryEvaluation":
		r = &CanaryEvaluation{}
	case "google-native:apigee/v1:DataCollector":
		r = &DataCollector{}
	case "google-native:apigee/v1:Datastore":
		r = &Datastore{}
	case "google-native:apigee/v1:DebugSession":
		r = &DebugSession{}
	case "google-native:apigee/v1:Developer":
		r = &Developer{}
	case "google-native:apigee/v1:EndpointAttachment":
		r = &EndpointAttachment{}
	case "google-native:apigee/v1:Envgroup":
		r = &Envgroup{}
	case "google-native:apigee/v1:EnvgroupAttachment":
		r = &EnvgroupAttachment{}
	case "google-native:apigee/v1:Environment":
		r = &Environment{}
	case "google-native:apigee/v1:Export":
		r = &Export{}
	case "google-native:apigee/v1:HostQuery":
		r = &HostQuery{}
	case "google-native:apigee/v1:Instance":
		r = &Instance{}
	case "google-native:apigee/v1:InstanceAttachment":
		r = &InstanceAttachment{}
	case "google-native:apigee/v1:Keystore":
		r = &Keystore{}
	case "google-native:apigee/v1:NatAddress":
		r = &NatAddress{}
	case "google-native:apigee/v1:Organization":
		r = &Organization{}
	case "google-native:apigee/v1:OrganizationEnvironmentIamPolicy":
		r = &OrganizationEnvironmentIamPolicy{}
	case "google-native:apigee/v1:Override":
		r = &Override{}
	case "google-native:apigee/v1:Query":
		r = &Query{}
	case "google-native:apigee/v1:RatePlan":
		r = &RatePlan{}
	case "google-native:apigee/v1:Reference":
		r = &Reference{}
	case "google-native:apigee/v1:Report":
		r = &Report{}
	case "google-native:apigee/v1:Resourcefile":
		r = &Resourcefile{}
	case "google-native:apigee/v1:Sharedflow":
		r = &Sharedflow{}
	case "google-native:apigee/v1:Subscription":
		r = &Subscription{}
	case "google-native:apigee/v1:TargetServer":
		r = &TargetServer{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := google.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"google-native",
		"apigee/v1",
		&module{version},
	)
}
