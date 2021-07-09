// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a study.
func LookupStudy(ctx *pulumi.Context, args *LookupStudyArgs, opts ...pulumi.InvokeOption) (*LookupStudyResult, error) {
	var rv LookupStudyResult
	err := ctx.Invoke("google-native:ml/v1:getStudy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStudyArgs struct {
	Location string `pulumi:"location"`
	Project  string `pulumi:"project"`
	StudyId  string `pulumi:"studyId"`
}

type LookupStudyResult struct {
	// Time at which the study was created.
	CreateTime string `pulumi:"createTime"`
	// A human readable reason why the Study is inactive. This should be empty if a study is ACTIVE or COMPLETED.
	InactiveReason string `pulumi:"inactiveReason"`
	// The name of a study.
	Name string `pulumi:"name"`
	// The detailed state of a study.
	State string `pulumi:"state"`
	// Configuration of the study.
	StudyConfig GoogleCloudMlV1__StudyConfigResponse `pulumi:"studyConfig"`
}
