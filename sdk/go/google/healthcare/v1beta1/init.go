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
	case "google-native:healthcare/v1beta1:Dataset":
		r = &Dataset{}
	case "google-native:healthcare/v1beta1:DatasetAnnotationStore":
		r = &DatasetAnnotationStore{}
	case "google-native:healthcare/v1beta1:DatasetAnnotationStoreAnnotation":
		r = &DatasetAnnotationStoreAnnotation{}
	case "google-native:healthcare/v1beta1:DatasetAnnotationStoreIamPolicy":
		r = &DatasetAnnotationStoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetConsentStore":
		r = &DatasetConsentStore{}
	case "google-native:healthcare/v1beta1:DatasetConsentStoreAttributeDefinition":
		r = &DatasetConsentStoreAttributeDefinition{}
	case "google-native:healthcare/v1beta1:DatasetConsentStoreConsent":
		r = &DatasetConsentStoreConsent{}
	case "google-native:healthcare/v1beta1:DatasetConsentStoreConsentArtifact":
		r = &DatasetConsentStoreConsentArtifact{}
	case "google-native:healthcare/v1beta1:DatasetConsentStoreIamPolicy":
		r = &DatasetConsentStoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetConsentStoreUserDataMapping":
		r = &DatasetConsentStoreUserDataMapping{}
	case "google-native:healthcare/v1beta1:DatasetDicomStore":
		r = &DatasetDicomStore{}
	case "google-native:healthcare/v1beta1:DatasetDicomStoreIamPolicy":
		r = &DatasetDicomStoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetFhirStore":
		r = &DatasetFhirStore{}
	case "google-native:healthcare/v1beta1:DatasetFhirStoreIamPolicy":
		r = &DatasetFhirStoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetHl7V2Store":
		r = &DatasetHl7V2Store{}
	case "google-native:healthcare/v1beta1:DatasetHl7V2StoreIamPolicy":
		r = &DatasetHl7V2StoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetHl7V2StoreMessage":
		r = &DatasetHl7V2StoreMessage{}
	case "google-native:healthcare/v1beta1:DatasetIamPolicy":
		r = &DatasetIamPolicy{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := google.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"google-native",
		"healthcare/v1beta1",
		&module{version},
	)
}
