// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

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
	case "google-native:metastore/v1alpha:Backup":
		r = &Backup{}
	case "google-native:metastore/v1alpha:MetadataImport":
		r = &MetadataImport{}
	case "google-native:metastore/v1alpha:Service":
		r = &Service{}
	case "google-native:metastore/v1alpha:ServiceBackupIamPolicy":
		r = &ServiceBackupIamPolicy{}
	case "google-native:metastore/v1alpha:ServiceDatabaseIamPolicy":
		r = &ServiceDatabaseIamPolicy{}
	case "google-native:metastore/v1alpha:ServiceDatabaseTableIamPolicy":
		r = &ServiceDatabaseTableIamPolicy{}
	case "google-native:metastore/v1alpha:ServiceIamPolicy":
		r = &ServiceIamPolicy{}
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
		"metastore/v1alpha",
		&module{version},
	)
}
