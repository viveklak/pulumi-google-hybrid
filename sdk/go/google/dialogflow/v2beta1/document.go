// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new document. Note: The `projects.agent.knowledgeBases.documents` resource is deprecated; only use `projects.knowledgeBases.documents`.
type Document struct {
	pulumi.CustomResourceState

	// The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types. Note: This field is in the process of being deprecated, please use raw_content instead.
	Content pulumi.StringOutput `pulumi:"content"`
	// The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
	ContentUri pulumi.StringOutput `pulumi:"contentUri"`
	// Required. The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
	EnableAutoReload pulumi.BoolOutput `pulumi:"enableAutoReload"`
	// Required. The knowledge type of document content.
	KnowledgeTypes pulumi.StringArrayOutput `pulumi:"knowledgeTypes"`
	// The time and status of the latest reload. This reload may have been triggered automatically or manually and may not have succeeded.
	LatestReloadStatus GoogleCloudDialogflowV2beta1DocumentReloadStatusResponseOutput `pulumi:"latestReloadStatus"`
	// Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Required. The MIME type of this document.
	MimeType pulumi.StringOutput `pulumi:"mimeType"`
	// Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
	RawContent pulumi.StringOutput `pulumi:"rawContent"`
}

// NewDocument registers a new resource with the given unique name, arguments, and options.
func NewDocument(ctx *pulumi.Context,
	name string, args *DocumentArgs, opts ...pulumi.ResourceOption) (*Document, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KnowledgeBaseId == nil {
		return nil, errors.New("invalid value for required argument 'KnowledgeBaseId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Document
	err := ctx.RegisterResource("google-native:dialogflow/v2beta1:Document", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocument gets an existing Document resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocument(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentState, opts ...pulumi.ResourceOption) (*Document, error) {
	var resource Document
	err := ctx.ReadResource("google-native:dialogflow/v2beta1:Document", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Document resources.
type documentState struct {
}

type DocumentState struct {
}

func (DocumentState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentState)(nil)).Elem()
}

type documentArgs struct {
	// The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types. Note: This field is in the process of being deprecated, please use raw_content instead.
	Content *string `pulumi:"content"`
	// The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
	ContentUri *string `pulumi:"contentUri"`
	// Required. The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
	DisplayName *string `pulumi:"displayName"`
	// Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
	EnableAutoReload        *bool   `pulumi:"enableAutoReload"`
	ImportGcsCustomMetadata *string `pulumi:"importGcsCustomMetadata"`
	KnowledgeBaseId         string  `pulumi:"knowledgeBaseId"`
	// Required. The knowledge type of document content.
	KnowledgeTypes []string `pulumi:"knowledgeTypes"`
	Location       string   `pulumi:"location"`
	// Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
	Metadata map[string]string `pulumi:"metadata"`
	// Required. The MIME type of this document.
	MimeType *string `pulumi:"mimeType"`
	// Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
	RawContent *string `pulumi:"rawContent"`
}

// The set of arguments for constructing a Document resource.
type DocumentArgs struct {
	// The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types. Note: This field is in the process of being deprecated, please use raw_content instead.
	Content pulumi.StringPtrInput
	// The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
	ContentUri pulumi.StringPtrInput
	// Required. The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
	DisplayName pulumi.StringPtrInput
	// Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
	EnableAutoReload        pulumi.BoolPtrInput
	ImportGcsCustomMetadata pulumi.StringPtrInput
	KnowledgeBaseId         pulumi.StringInput
	// Required. The knowledge type of document content.
	KnowledgeTypes DocumentKnowledgeTypesItemArrayInput
	Location       pulumi.StringInput
	// Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
	Metadata pulumi.StringMapInput
	// Required. The MIME type of this document.
	MimeType pulumi.StringPtrInput
	// Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
	RawContent pulumi.StringPtrInput
}

func (DocumentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentArgs)(nil)).Elem()
}

type DocumentInput interface {
	pulumi.Input

	ToDocumentOutput() DocumentOutput
	ToDocumentOutputWithContext(ctx context.Context) DocumentOutput
}

func (*Document) ElementType() reflect.Type {
	return reflect.TypeOf((*Document)(nil))
}

func (i *Document) ToDocumentOutput() DocumentOutput {
	return i.ToDocumentOutputWithContext(context.Background())
}

func (i *Document) ToDocumentOutputWithContext(ctx context.Context) DocumentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentOutput)
}

type DocumentOutput struct {
	*pulumi.OutputState
}

func (DocumentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Document)(nil))
}

func (o DocumentOutput) ToDocumentOutput() DocumentOutput {
	return o
}

func (o DocumentOutput) ToDocumentOutputWithContext(ctx context.Context) DocumentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DocumentOutput{})
}
