// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Requests the creation of a new IosApp in the specified FirebaseProject. The result of this call is an `Operation` which can be used to track the provisioning process. The `Operation` is automatically deleted after completion, so there is no need to call `DeleteOperation`.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type IosApp struct {
	pulumi.CustomResourceState

	// The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
	ApiKeyId pulumi.StringOutput `pulumi:"apiKeyId"`
	// Immutable. The globally unique, Firebase-assigned identifier for the `IosApp`. This identifier should be treated as an opaque token, as the data format is not specified.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The automatically generated Apple ID assigned to the iOS app by Apple in the iOS App Store.
	AppStoreId pulumi.StringOutput `pulumi:"appStoreId"`
	// Immutable. The canonical bundle ID of the iOS app as it would appear in the iOS AppStore.
	BundleId pulumi.StringOutput `pulumi:"bundleId"`
	// The user-assigned display name for the `IosApp`.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The resource name of the IosApp, in the format: projects/PROJECT_IDENTIFIER /iosApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.iosApps#IosApp.FIELDS.app_id)).
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `IosApp`.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Apple Developer Team ID associated with the App in the App Store.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewIosApp registers a new resource with the given unique name, arguments, and options.
func NewIosApp(ctx *pulumi.Context,
	name string, args *IosAppArgs, opts ...pulumi.ResourceOption) (*IosApp, error) {
	if args == nil {
		args = &IosAppArgs{}
	}

	var resource IosApp
	err := ctx.RegisterResource("google-native:firebase/v1beta1:IosApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIosApp gets an existing IosApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIosApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IosAppState, opts ...pulumi.ResourceOption) (*IosApp, error) {
	var resource IosApp
	err := ctx.ReadResource("google-native:firebase/v1beta1:IosApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IosApp resources.
type iosAppState struct {
}

type IosAppState struct {
}

func (IosAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*iosAppState)(nil)).Elem()
}

type iosAppArgs struct {
	// The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
	ApiKeyId *string `pulumi:"apiKeyId"`
	// Immutable. The globally unique, Firebase-assigned identifier for the `IosApp`. This identifier should be treated as an opaque token, as the data format is not specified.
	AppId *string `pulumi:"appId"`
	// The automatically generated Apple ID assigned to the iOS app by Apple in the iOS App Store.
	AppStoreId *string `pulumi:"appStoreId"`
	// Immutable. The canonical bundle ID of the iOS app as it would appear in the iOS AppStore.
	BundleId *string `pulumi:"bundleId"`
	// The user-assigned display name for the `IosApp`.
	DisplayName *string `pulumi:"displayName"`
	// The resource name of the IosApp, in the format: projects/PROJECT_IDENTIFIER /iosApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.iosApps#IosApp.FIELDS.app_id)).
	Name *string `pulumi:"name"`
	// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `IosApp`.
	Project *string `pulumi:"project"`
	// The Apple Developer Team ID associated with the App in the App Store.
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a IosApp resource.
type IosAppArgs struct {
	// The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
	ApiKeyId pulumi.StringPtrInput
	// Immutable. The globally unique, Firebase-assigned identifier for the `IosApp`. This identifier should be treated as an opaque token, as the data format is not specified.
	AppId pulumi.StringPtrInput
	// The automatically generated Apple ID assigned to the iOS app by Apple in the iOS App Store.
	AppStoreId pulumi.StringPtrInput
	// Immutable. The canonical bundle ID of the iOS app as it would appear in the iOS AppStore.
	BundleId pulumi.StringPtrInput
	// The user-assigned display name for the `IosApp`.
	DisplayName pulumi.StringPtrInput
	// The resource name of the IosApp, in the format: projects/PROJECT_IDENTIFIER /iosApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.iosApps#IosApp.FIELDS.app_id)).
	Name pulumi.StringPtrInput
	// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `IosApp`.
	Project pulumi.StringPtrInput
	// The Apple Developer Team ID associated with the App in the App Store.
	TeamId pulumi.StringPtrInput
}

func (IosAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iosAppArgs)(nil)).Elem()
}

type IosAppInput interface {
	pulumi.Input

	ToIosAppOutput() IosAppOutput
	ToIosAppOutputWithContext(ctx context.Context) IosAppOutput
}

func (*IosApp) ElementType() reflect.Type {
	return reflect.TypeOf((**IosApp)(nil)).Elem()
}

func (i *IosApp) ToIosAppOutput() IosAppOutput {
	return i.ToIosAppOutputWithContext(context.Background())
}

func (i *IosApp) ToIosAppOutputWithContext(ctx context.Context) IosAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IosAppOutput)
}

type IosAppOutput struct{ *pulumi.OutputState }

func (IosAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IosApp)(nil)).Elem()
}

func (o IosAppOutput) ToIosAppOutput() IosAppOutput {
	return o
}

func (o IosAppOutput) ToIosAppOutputWithContext(ctx context.Context) IosAppOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IosAppInput)(nil)).Elem(), &IosApp{})
	pulumi.RegisterOutputType(IosAppOutput{})
}
