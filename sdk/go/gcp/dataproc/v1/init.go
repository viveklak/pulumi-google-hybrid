// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp-native/sdk/go/gcp"
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
	case "gcp-native:dataproc/v1:AutoscalingPolicy":
		r = &AutoscalingPolicy{}
	case "gcp-native:dataproc/v1:AutoscalingPolicyIamPolicy":
		r = &AutoscalingPolicyIamPolicy{}
	case "gcp-native:dataproc/v1:RegionAutoscalingPolicy":
		r = &RegionAutoscalingPolicy{}
	case "gcp-native:dataproc/v1:RegionAutoscalingPolicyIamPolicy":
		r = &RegionAutoscalingPolicyIamPolicy{}
	case "gcp-native:dataproc/v1:RegionCluster":
		r = &RegionCluster{}
	case "gcp-native:dataproc/v1:RegionClusterIamPolicy":
		r = &RegionClusterIamPolicy{}
	case "gcp-native:dataproc/v1:RegionJob":
		r = &RegionJob{}
	case "gcp-native:dataproc/v1:RegionJobIamPolicy":
		r = &RegionJobIamPolicy{}
	case "gcp-native:dataproc/v1:RegionOperationIamPolicy":
		r = &RegionOperationIamPolicy{}
	case "gcp-native:dataproc/v1:RegionWorkflowTemplate":
		r = &RegionWorkflowTemplate{}
	case "gcp-native:dataproc/v1:RegionWorkflowTemplateIamPolicy":
		r = &RegionWorkflowTemplateIamPolicy{}
	case "gcp-native:dataproc/v1:WorkflowTemplate":
		r = &WorkflowTemplate{}
	case "gcp-native:dataproc/v1:WorkflowTemplateIamPolicy":
		r = &WorkflowTemplateIamPolicy{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"gcp-native",
		"dataproc/v1",
		&module{version},
	)
}
