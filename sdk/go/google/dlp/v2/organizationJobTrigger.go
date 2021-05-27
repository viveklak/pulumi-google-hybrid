// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a job trigger to run DLP actions such as scanning storage for sensitive information on a set schedule. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.
type OrganizationJobTrigger struct {
	pulumi.CustomResourceState

	// The creation timestamp of a triggeredJob.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// User provided description (max 256 chars)
	Description pulumi.StringOutput `pulumi:"description"`
	// Display name (max 100 chars)
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
	Errors GooglePrivacyDlpV2ErrorResponseArrayOutput `pulumi:"errors"`
	// For inspect jobs, a snapshot of the configuration.
	InspectJob GooglePrivacyDlpV2InspectJobConfigResponseOutput `pulumi:"inspectJob"`
	// The timestamp of the last time this trigger executed.
	LastRunTime pulumi.StringOutput `pulumi:"lastRunTime"`
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. A status for this trigger.
	Status pulumi.StringOutput `pulumi:"status"`
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers GooglePrivacyDlpV2TriggerResponseArrayOutput `pulumi:"triggers"`
	// The last update timestamp of a triggeredJob.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewOrganizationJobTrigger registers a new resource with the given unique name, arguments, and options.
func NewOrganizationJobTrigger(ctx *pulumi.Context,
	name string, args *OrganizationJobTriggerArgs, opts ...pulumi.ResourceOption) (*OrganizationJobTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource OrganizationJobTrigger
	err := ctx.RegisterResource("google-native:dlp/v2:OrganizationJobTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationJobTrigger gets an existing OrganizationJobTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationJobTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationJobTriggerState, opts ...pulumi.ResourceOption) (*OrganizationJobTrigger, error) {
	var resource OrganizationJobTrigger
	err := ctx.ReadResource("google-native:dlp/v2:OrganizationJobTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationJobTrigger resources.
type organizationJobTriggerState struct {
	// The creation timestamp of a triggeredJob.
	CreateTime *string `pulumi:"createTime"`
	// User provided description (max 256 chars)
	Description *string `pulumi:"description"`
	// Display name (max 100 chars)
	DisplayName *string `pulumi:"displayName"`
	// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
	Errors []GooglePrivacyDlpV2ErrorResponse `pulumi:"errors"`
	// For inspect jobs, a snapshot of the configuration.
	InspectJob *GooglePrivacyDlpV2InspectJobConfigResponse `pulumi:"inspectJob"`
	// The timestamp of the last time this trigger executed.
	LastRunTime *string `pulumi:"lastRunTime"`
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name *string `pulumi:"name"`
	// Required. A status for this trigger.
	Status *string `pulumi:"status"`
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers []GooglePrivacyDlpV2TriggerResponse `pulumi:"triggers"`
	// The last update timestamp of a triggeredJob.
	UpdateTime *string `pulumi:"updateTime"`
}

type OrganizationJobTriggerState struct {
	// The creation timestamp of a triggeredJob.
	CreateTime pulumi.StringPtrInput
	// User provided description (max 256 chars)
	Description pulumi.StringPtrInput
	// Display name (max 100 chars)
	DisplayName pulumi.StringPtrInput
	// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
	Errors GooglePrivacyDlpV2ErrorResponseArrayInput
	// For inspect jobs, a snapshot of the configuration.
	InspectJob GooglePrivacyDlpV2InspectJobConfigResponsePtrInput
	// The timestamp of the last time this trigger executed.
	LastRunTime pulumi.StringPtrInput
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name pulumi.StringPtrInput
	// Required. A status for this trigger.
	Status pulumi.StringPtrInput
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers GooglePrivacyDlpV2TriggerResponseArrayInput
	// The last update timestamp of a triggeredJob.
	UpdateTime pulumi.StringPtrInput
}

func (OrganizationJobTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationJobTriggerState)(nil)).Elem()
}

type organizationJobTriggerArgs struct {
	// User provided description (max 256 chars)
	Description *string `pulumi:"description"`
	// Display name (max 100 chars)
	DisplayName *string `pulumi:"displayName"`
	// For inspect jobs, a snapshot of the configuration.
	InspectJob *GooglePrivacyDlpV2InspectJobConfig `pulumi:"inspectJob"`
	Location   string                              `pulumi:"location"`
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Required. A status for this trigger.
	Status *string `pulumi:"status"`
	// The trigger id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
	TriggerId *string `pulumi:"triggerId"`
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers []GooglePrivacyDlpV2Trigger `pulumi:"triggers"`
}

// The set of arguments for constructing a OrganizationJobTrigger resource.
type OrganizationJobTriggerArgs struct {
	// User provided description (max 256 chars)
	Description pulumi.StringPtrInput
	// Display name (max 100 chars)
	DisplayName pulumi.StringPtrInput
	// For inspect jobs, a snapshot of the configuration.
	InspectJob GooglePrivacyDlpV2InspectJobConfigPtrInput
	Location   pulumi.StringInput
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Required. A status for this trigger.
	Status pulumi.StringPtrInput
	// The trigger id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
	TriggerId pulumi.StringPtrInput
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers GooglePrivacyDlpV2TriggerArrayInput
}

func (OrganizationJobTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationJobTriggerArgs)(nil)).Elem()
}

type OrganizationJobTriggerInput interface {
	pulumi.Input

	ToOrganizationJobTriggerOutput() OrganizationJobTriggerOutput
	ToOrganizationJobTriggerOutputWithContext(ctx context.Context) OrganizationJobTriggerOutput
}

func (*OrganizationJobTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationJobTrigger)(nil))
}

func (i *OrganizationJobTrigger) ToOrganizationJobTriggerOutput() OrganizationJobTriggerOutput {
	return i.ToOrganizationJobTriggerOutputWithContext(context.Background())
}

func (i *OrganizationJobTrigger) ToOrganizationJobTriggerOutputWithContext(ctx context.Context) OrganizationJobTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationJobTriggerOutput)
}

type OrganizationJobTriggerOutput struct {
	*pulumi.OutputState
}

func (OrganizationJobTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationJobTrigger)(nil))
}

func (o OrganizationJobTriggerOutput) ToOrganizationJobTriggerOutput() OrganizationJobTriggerOutput {
	return o
}

func (o OrganizationJobTriggerOutput) ToOrganizationJobTriggerOutputWithContext(ctx context.Context) OrganizationJobTriggerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationJobTriggerOutput{})
}
