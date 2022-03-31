// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
	case "google-native:networkservices/v1beta1:EndpointPolicy":
		r = &EndpointPolicy{}
	case "google-native:networkservices/v1beta1:EndpointPolicyIamPolicy":
		r = &EndpointPolicyIamPolicy{}
	case "google-native:networkservices/v1beta1:Gateway":
		r = &Gateway{}
	case "google-native:networkservices/v1beta1:GatewayIamPolicy":
		r = &GatewayIamPolicy{}
	case "google-native:networkservices/v1beta1:GrpcRoute":
		r = &GrpcRoute{}
	case "google-native:networkservices/v1beta1:HttpRoute":
		r = &HttpRoute{}
	case "google-native:networkservices/v1beta1:Mesh":
		r = &Mesh{}
	case "google-native:networkservices/v1beta1:MeshIamPolicy":
		r = &MeshIamPolicy{}
	case "google-native:networkservices/v1beta1:ServiceBinding":
		r = &ServiceBinding{}
	case "google-native:networkservices/v1beta1:ServiceBindingIamPolicy":
		r = &ServiceBindingIamPolicy{}
	case "google-native:networkservices/v1beta1:TcpRoute":
		r = &TcpRoute{}
	case "google-native:networkservices/v1beta1:TlsRoute":
		r = &TlsRoute{}
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
		"networkservices/v1beta1",
		&module{version},
	)
}
