// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new DICOM store within the parent dataset.
type DicomStore struct {
	pulumi.CustomResourceState

	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification destination for new DICOM instances. Supplied by the client.
	NotificationConfig NotificationConfigResponseOutput `pulumi:"notificationConfig"`
	// A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
	StreamConfigs GoogleCloudHealthcareV1beta1DicomStreamConfigResponseArrayOutput `pulumi:"streamConfigs"`
}

// NewDicomStore registers a new resource with the given unique name, arguments, and options.
func NewDicomStore(ctx *pulumi.Context,
	name string, args *DicomStoreArgs, opts ...pulumi.ResourceOption) (*DicomStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	var resource DicomStore
	err := ctx.RegisterResource("google-native:healthcare/v1beta1:DicomStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDicomStore gets an existing DicomStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDicomStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DicomStoreState, opts ...pulumi.ResourceOption) (*DicomStore, error) {
	var resource DicomStore
	err := ctx.ReadResource("google-native:healthcare/v1beta1:DicomStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DicomStore resources.
type dicomStoreState struct {
}

type DicomStoreState struct {
}

func (DicomStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreState)(nil)).Elem()
}

type dicomStoreArgs struct {
	DatasetId string `pulumi:"datasetId"`
	// The ID of the DICOM store that is being created. Any string value up to 256 characters in length.
	DicomStoreId *string `pulumi:"dicomStoreId"`
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
	Name *string `pulumi:"name"`
	// Notification destination for new DICOM instances. Supplied by the client.
	NotificationConfig *NotificationConfig `pulumi:"notificationConfig"`
	Project            *string             `pulumi:"project"`
	// A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
	StreamConfigs []GoogleCloudHealthcareV1beta1DicomStreamConfig `pulumi:"streamConfigs"`
}

// The set of arguments for constructing a DicomStore resource.
type DicomStoreArgs struct {
	DatasetId pulumi.StringInput
	// The ID of the DICOM store that is being created. Any string value up to 256 characters in length.
	DicomStoreId pulumi.StringPtrInput
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
	Name pulumi.StringPtrInput
	// Notification destination for new DICOM instances. Supplied by the client.
	NotificationConfig NotificationConfigPtrInput
	Project            pulumi.StringPtrInput
	// A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
	StreamConfigs GoogleCloudHealthcareV1beta1DicomStreamConfigArrayInput
}

func (DicomStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreArgs)(nil)).Elem()
}

type DicomStoreInput interface {
	pulumi.Input

	ToDicomStoreOutput() DicomStoreOutput
	ToDicomStoreOutputWithContext(ctx context.Context) DicomStoreOutput
}

func (*DicomStore) ElementType() reflect.Type {
	return reflect.TypeOf((**DicomStore)(nil)).Elem()
}

func (i *DicomStore) ToDicomStoreOutput() DicomStoreOutput {
	return i.ToDicomStoreOutputWithContext(context.Background())
}

func (i *DicomStore) ToDicomStoreOutputWithContext(ctx context.Context) DicomStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DicomStoreOutput)
}

type DicomStoreOutput struct{ *pulumi.OutputState }

func (DicomStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DicomStore)(nil)).Elem()
}

func (o DicomStoreOutput) ToDicomStoreOutput() DicomStoreOutput {
	return o
}

func (o DicomStoreOutput) ToDicomStoreOutputWithContext(ctx context.Context) DicomStoreOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DicomStoreInput)(nil)).Elem(), &DicomStore{})
	pulumi.RegisterOutputType(DicomStoreOutput{})
}
