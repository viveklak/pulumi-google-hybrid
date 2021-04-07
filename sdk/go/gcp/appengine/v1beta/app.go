// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an App Engine application for a Google Cloud Platform project. Required fields: id - The ID of the target Cloud Platform project. location - The region (https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.For more information about App Engine applications, see Managing Projects, Applications, and Billing (https://cloud.google.com/appengine/docs/standard/python/console/).
type App struct {
	pulumi.CustomResourceState

	// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
	AuthDomain pulumi.StringOutput `pulumi:"authDomain"`
	// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
	CodeBucket pulumi.StringOutput `pulumi:"codeBucket"`
	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	DatabaseType pulumi.StringOutput `pulumi:"databaseType"`
	// Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
	DefaultBucket pulumi.StringOutput `pulumi:"defaultBucket"`
	// Cookie expiration policy for this application.
	DefaultCookieExpiration pulumi.StringOutput `pulumi:"defaultCookieExpiration"`
	// Hostname used to reach this application, as resolved by App Engine.@OutputOnly
	DefaultHostname pulumi.StringOutput `pulumi:"defaultHostname"`
	// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
	DispatchRules UrlDispatchRuleResponseArrayOutput `pulumi:"dispatchRules"`
	// The feature specific settings to be used in the application.
	FeatureSettings FeatureSettingsResponseOutput `pulumi:"featureSettings"`
	// The Google Container Registry domain used for storing managed build docker images for this application.
	GcrDomain pulumi.StringOutput              `pulumi:"gcrDomain"`
	Iap       IdentityAwareProxyResponseOutput `pulumi:"iap"`
	// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
	LocationId pulumi.StringOutput `pulumi:"locationId"`
	// Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
	Name pulumi.StringOutput `pulumi:"name"`
	// Serving status of this application.
	ServingStatus pulumi.StringOutput `pulumi:"servingStatus"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppsId == nil {
		return nil, errors.New("invalid value for required argument 'AppsId'")
	}
	var resource App
	err := ctx.RegisterResource("gcp-native:appengine/v1beta:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("gcp-native:appengine/v1beta:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
	// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
	AuthDomain *string `pulumi:"authDomain"`
	// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
	CodeBucket *string `pulumi:"codeBucket"`
	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	DatabaseType *string `pulumi:"databaseType"`
	// Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
	DefaultBucket *string `pulumi:"defaultBucket"`
	// Cookie expiration policy for this application.
	DefaultCookieExpiration *string `pulumi:"defaultCookieExpiration"`
	// Hostname used to reach this application, as resolved by App Engine.@OutputOnly
	DefaultHostname *string `pulumi:"defaultHostname"`
	// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
	DispatchRules []UrlDispatchRuleResponse `pulumi:"dispatchRules"`
	// The feature specific settings to be used in the application.
	FeatureSettings *FeatureSettingsResponse `pulumi:"featureSettings"`
	// The Google Container Registry domain used for storing managed build docker images for this application.
	GcrDomain *string                     `pulumi:"gcrDomain"`
	Iap       *IdentityAwareProxyResponse `pulumi:"iap"`
	// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
	LocationId *string `pulumi:"locationId"`
	// Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
	Name *string `pulumi:"name"`
	// Serving status of this application.
	ServingStatus *string `pulumi:"servingStatus"`
}

type AppState struct {
	// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
	AuthDomain pulumi.StringPtrInput
	// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
	CodeBucket pulumi.StringPtrInput
	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	DatabaseType pulumi.StringPtrInput
	// Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
	DefaultBucket pulumi.StringPtrInput
	// Cookie expiration policy for this application.
	DefaultCookieExpiration pulumi.StringPtrInput
	// Hostname used to reach this application, as resolved by App Engine.@OutputOnly
	DefaultHostname pulumi.StringPtrInput
	// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
	DispatchRules UrlDispatchRuleResponseArrayInput
	// The feature specific settings to be used in the application.
	FeatureSettings FeatureSettingsResponsePtrInput
	// The Google Container Registry domain used for storing managed build docker images for this application.
	GcrDomain pulumi.StringPtrInput
	Iap       IdentityAwareProxyResponsePtrInput
	// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
	LocationId pulumi.StringPtrInput
	// Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
	Name pulumi.StringPtrInput
	// Serving status of this application.
	ServingStatus pulumi.StringPtrInput
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	AppsId string `pulumi:"appsId"`
	// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
	AuthDomain *string `pulumi:"authDomain"`
	// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
	CodeBucket *string `pulumi:"codeBucket"`
	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	DatabaseType *string `pulumi:"databaseType"`
	// Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
	DefaultBucket *string `pulumi:"defaultBucket"`
	// Cookie expiration policy for this application.
	DefaultCookieExpiration *string `pulumi:"defaultCookieExpiration"`
	// Hostname used to reach this application, as resolved by App Engine.@OutputOnly
	DefaultHostname *string `pulumi:"defaultHostname"`
	// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
	DispatchRules []UrlDispatchRule `pulumi:"dispatchRules"`
	// The feature specific settings to be used in the application.
	FeatureSettings *FeatureSettings `pulumi:"featureSettings"`
	// The Google Container Registry domain used for storing managed build docker images for this application.
	GcrDomain *string             `pulumi:"gcrDomain"`
	Iap       *IdentityAwareProxy `pulumi:"iap"`
	// Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
	Id *string `pulumi:"id"`
	// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
	LocationId *string `pulumi:"locationId"`
	// Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
	Name *string `pulumi:"name"`
	// Serving status of this application.
	ServingStatus *string `pulumi:"servingStatus"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	AppsId pulumi.StringInput
	// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
	AuthDomain pulumi.StringPtrInput
	// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
	CodeBucket pulumi.StringPtrInput
	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	DatabaseType pulumi.StringPtrInput
	// Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
	DefaultBucket pulumi.StringPtrInput
	// Cookie expiration policy for this application.
	DefaultCookieExpiration pulumi.StringPtrInput
	// Hostname used to reach this application, as resolved by App Engine.@OutputOnly
	DefaultHostname pulumi.StringPtrInput
	// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
	DispatchRules UrlDispatchRuleArrayInput
	// The feature specific settings to be used in the application.
	FeatureSettings FeatureSettingsPtrInput
	// The Google Container Registry domain used for storing managed build docker images for this application.
	GcrDomain pulumi.StringPtrInput
	Iap       IdentityAwareProxyPtrInput
	// Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
	Id pulumi.StringPtrInput
	// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
	LocationId pulumi.StringPtrInput
	// Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
	Name pulumi.StringPtrInput
	// Serving status of this application.
	ServingStatus pulumi.StringPtrInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((*App)(nil))
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

type AppOutput struct {
	*pulumi.OutputState
}

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*App)(nil))
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppOutput{})
}