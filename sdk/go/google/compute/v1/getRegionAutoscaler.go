// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified autoscaler.
func LookupRegionAutoscaler(ctx *pulumi.Context, args *LookupRegionAutoscalerArgs, opts ...pulumi.InvokeOption) (*LookupRegionAutoscalerResult, error) {
	var rv LookupRegionAutoscalerResult
	err := ctx.Invoke("google-native:compute/v1:getRegionAutoscaler", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionAutoscalerArgs struct {
	Autoscaler string `pulumi:"autoscaler"`
	Project    string `pulumi:"project"`
	Region     string `pulumi:"region"`
}

type LookupRegionAutoscalerResult struct {
	// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
	//
	// If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
	AutoscalingPolicy AutoscalingPolicyResponse `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Type of the resource. Always compute#autoscaler for autoscalers.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
	RecommendedSize int `pulumi:"recommendedSize"`
	// URL of the region where the instance group resides (for autoscalers living in regional scope).
	Region string `pulumi:"region"`
	// Status information of existing scaling schedules.
	ScalingScheduleStatus map[string]string `pulumi:"scalingScheduleStatus"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// The status of the autoscaler configuration. Current set of possible values:
	// - PENDING: Autoscaler backend hasn't read new/updated configuration.
	// - DELETING: Configuration is being deleted.
	// - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field.
	// - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
	Status string `pulumi:"status"`
	// Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
	StatusDetails []AutoscalerStatusDetailsResponse `pulumi:"statusDetails"`
	// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
	Target string `pulumi:"target"`
	// URL of the zone where the instance group resides (for autoscalers living in zonal scope).
	Zone string `pulumi:"zone"`
}
