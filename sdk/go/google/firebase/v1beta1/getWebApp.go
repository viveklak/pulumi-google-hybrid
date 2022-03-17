// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified WebApp.
func LookupWebApp(ctx *pulumi.Context, args *LookupWebAppArgs, opts ...pulumi.InvokeOption) (*LookupWebAppResult, error) {
	var rv LookupWebAppResult
	err := ctx.Invoke("google-native:firebase/v1beta1:getWebApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppArgs struct {
	Project  *string `pulumi:"project"`
	WebAppId string  `pulumi:"webAppId"`
}

type LookupWebAppResult struct {
	// The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
	ApiKeyId string `pulumi:"apiKeyId"`
	// Immutable. The globally unique, Firebase-assigned identifier for the `WebApp`. This identifier should be treated as an opaque token, as the data format is not specified.
	AppId string `pulumi:"appId"`
	// The URLs where the `WebApp` is hosted.
	AppUrls []string `pulumi:"appUrls"`
	// The user-assigned display name for the `WebApp`.
	DisplayName string `pulumi:"displayName"`
	// The resource name of the WebApp, in the format: projects/PROJECT_IDENTIFIER /webApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.webApps#WebApp.FIELDS.app_id)).
	Name string `pulumi:"name"`
	// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `WebApp`.
	Project string `pulumi:"project"`
	// Immutable. A unique, Firebase-assigned identifier for the `WebApp`. This identifier is only used to populate the `namespace` value for the `WebApp`. For most use cases, use `appId` to identify or reference the App. The `webId` value is only unique within a `FirebaseProject` and its associated Apps.
	WebId string `pulumi:"webId"`
}

func LookupWebAppOutput(ctx *pulumi.Context, args LookupWebAppOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppResult, error) {
			args := v.(LookupWebAppArgs)
			r, err := LookupWebApp(ctx, &args, opts...)
			return *r, err
		}).(LookupWebAppResultOutput)
}

type LookupWebAppOutputArgs struct {
	Project  pulumi.StringPtrInput `pulumi:"project"`
	WebAppId pulumi.StringInput    `pulumi:"webAppId"`
}

func (LookupWebAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppArgs)(nil)).Elem()
}

type LookupWebAppResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppResult)(nil)).Elem()
}

func (o LookupWebAppResultOutput) ToLookupWebAppResultOutput() LookupWebAppResultOutput {
	return o
}

func (o LookupWebAppResultOutput) ToLookupWebAppResultOutputWithContext(ctx context.Context) LookupWebAppResultOutput {
	return o
}

// The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
func (o LookupWebAppResultOutput) ApiKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.ApiKeyId }).(pulumi.StringOutput)
}

// Immutable. The globally unique, Firebase-assigned identifier for the `WebApp`. This identifier should be treated as an opaque token, as the data format is not specified.
func (o LookupWebAppResultOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.AppId }).(pulumi.StringOutput)
}

// The URLs where the `WebApp` is hosted.
func (o LookupWebAppResultOutput) AppUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppResult) []string { return v.AppUrls }).(pulumi.StringArrayOutput)
}

// The user-assigned display name for the `WebApp`.
func (o LookupWebAppResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The resource name of the WebApp, in the format: projects/PROJECT_IDENTIFIER /webApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.webApps#WebApp.FIELDS.app_id)).
func (o LookupWebAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Name }).(pulumi.StringOutput)
}

// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `WebApp`.
func (o LookupWebAppResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.Project }).(pulumi.StringOutput)
}

// Immutable. A unique, Firebase-assigned identifier for the `WebApp`. This identifier is only used to populate the `namespace` value for the `WebApp`. For most use cases, use `appId` to identify or reference the App. The `webId` value is only unique within a `FirebaseProject` and its associated Apps.
func (o LookupWebAppResultOutput) WebId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppResult) string { return v.WebId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppResultOutput{})
}
