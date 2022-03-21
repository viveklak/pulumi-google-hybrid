// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new CertificateTemplate in a given Project and Location.
// Auto-naming is currently not supported for this resource.
type CertificateTemplate struct {
	pulumi.CustomResourceState

	// The time at which this CertificateTemplate was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A human-readable description of scenarios this template is intended for.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints CertificateIdentityConstraintsResponseOutput `pulumi:"identityConstraints"`
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions CertificateExtensionConstraintsResponseOutput `pulumi:"passthroughExtensions"`
	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
	PredefinedValues X509ParametersResponseOutput `pulumi:"predefinedValues"`
	// The time at which this CertificateTemplate was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCertificateTemplate registers a new resource with the given unique name, arguments, and options.
func NewCertificateTemplate(ctx *pulumi.Context,
	name string, args *CertificateTemplateArgs, opts ...pulumi.ResourceOption) (*CertificateTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateTemplateId'")
	}
	var resource CertificateTemplate
	err := ctx.RegisterResource("google-native:privateca/v1:CertificateTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateTemplate gets an existing CertificateTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateTemplateState, opts ...pulumi.ResourceOption) (*CertificateTemplate, error) {
	var resource CertificateTemplate
	err := ctx.ReadResource("google-native:privateca/v1:CertificateTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateTemplate resources.
type certificateTemplateState struct {
}

type CertificateTemplateState struct {
}

func (CertificateTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateTemplateState)(nil)).Elem()
}

type certificateTemplateArgs struct {
	// Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	CertificateTemplateId string `pulumi:"certificateTemplateId"`
	// Optional. A human-readable description of scenarios this template is intended for.
	Description *string `pulumi:"description"`
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints *CertificateIdentityConstraints `pulumi:"identityConstraints"`
	// Optional. Labels with user-defined metadata.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions *CertificateExtensionConstraints `pulumi:"passthroughExtensions"`
	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
	PredefinedValues *X509Parameters `pulumi:"predefinedValues"`
	Project          *string         `pulumi:"project"`
	// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a CertificateTemplate resource.
type CertificateTemplateArgs struct {
	// Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	CertificateTemplateId pulumi.StringInput
	// Optional. A human-readable description of scenarios this template is intended for.
	Description pulumi.StringPtrInput
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints CertificateIdentityConstraintsPtrInput
	// Optional. Labels with user-defined metadata.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions CertificateExtensionConstraintsPtrInput
	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
	PredefinedValues X509ParametersPtrInput
	Project          pulumi.StringPtrInput
	// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
}

func (CertificateTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateTemplateArgs)(nil)).Elem()
}

type CertificateTemplateInput interface {
	pulumi.Input

	ToCertificateTemplateOutput() CertificateTemplateOutput
	ToCertificateTemplateOutputWithContext(ctx context.Context) CertificateTemplateOutput
}

func (*CertificateTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateTemplate)(nil)).Elem()
}

func (i *CertificateTemplate) ToCertificateTemplateOutput() CertificateTemplateOutput {
	return i.ToCertificateTemplateOutputWithContext(context.Background())
}

func (i *CertificateTemplate) ToCertificateTemplateOutputWithContext(ctx context.Context) CertificateTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateTemplateOutput)
}

type CertificateTemplateOutput struct{ *pulumi.OutputState }

func (CertificateTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateTemplate)(nil)).Elem()
}

func (o CertificateTemplateOutput) ToCertificateTemplateOutput() CertificateTemplateOutput {
	return o
}

func (o CertificateTemplateOutput) ToCertificateTemplateOutputWithContext(ctx context.Context) CertificateTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateTemplateInput)(nil)).Elem(), &CertificateTemplate{})
	pulumi.RegisterOutputType(CertificateTemplateOutput{})
}
