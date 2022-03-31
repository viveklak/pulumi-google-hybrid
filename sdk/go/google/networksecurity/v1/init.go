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
	case "google-native:networksecurity/v1:AuthorizationPolicy":
		r = &AuthorizationPolicy{}
	case "google-native:networksecurity/v1:AuthorizationPolicyIamPolicy":
		r = &AuthorizationPolicyIamPolicy{}
	case "google-native:networksecurity/v1:ClientTlsPolicy":
		r = &ClientTlsPolicy{}
	case "google-native:networksecurity/v1:ClientTlsPolicyIamPolicy":
		r = &ClientTlsPolicyIamPolicy{}
	case "google-native:networksecurity/v1:ServerTlsPolicy":
		r = &ServerTlsPolicy{}
	case "google-native:networksecurity/v1:ServerTlsPolicyIamPolicy":
		r = &ServerTlsPolicyIamPolicy{}
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
		"networksecurity/v1",
		&module{version},
	)
}
