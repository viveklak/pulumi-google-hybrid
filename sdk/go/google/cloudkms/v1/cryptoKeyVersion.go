// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new CryptoKeyVersion in a CryptoKey. The server will assign the next sequential id. If unset, state will be set to ENABLED.
type CryptoKeyVersion struct {
	pulumi.CustomResourceState

	// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	// Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only provided for key versions with protection_level HSM.
	Attestation KeyOperationAttestationResponseOutput `pulumi:"attestation"`
	// The time at which this CryptoKeyVersion was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time this CryptoKeyVersion's key material was destroyed. Only present if state is DESTROYED.
	DestroyEventTime pulumi.StringOutput `pulumi:"destroyEventTime"`
	// The time this CryptoKeyVersion's key material is scheduled for destruction. Only present if state is DESTROY_SCHEDULED.
	DestroyTime pulumi.StringOutput `pulumi:"destroyTime"`
	// ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
	ExternalProtectionLevelOptions ExternalProtectionLevelOptionsResponseOutput `pulumi:"externalProtectionLevelOptions"`
	// The time this CryptoKeyVersion's key material was generated.
	GenerateTime pulumi.StringOutput `pulumi:"generateTime"`
	// The root cause of an import failure. Only present if state is IMPORT_FAILED.
	ImportFailureReason pulumi.StringOutput `pulumi:"importFailureReason"`
	// The name of the ImportJob used to import this CryptoKeyVersion. Only present if the underlying key material was imported.
	ImportJob pulumi.StringOutput `pulumi:"importJob"`
	// The time at which this CryptoKeyVersion's key material was imported.
	ImportTime pulumi.StringOutput `pulumi:"importTime"`
	// The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
	ProtectionLevel pulumi.StringOutput `pulumi:"protectionLevel"`
	// The current state of the CryptoKeyVersion.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewCryptoKeyVersion registers a new resource with the given unique name, arguments, and options.
func NewCryptoKeyVersion(ctx *pulumi.Context,
	name string, args *CryptoKeyVersionArgs, opts ...pulumi.ResourceOption) (*CryptoKeyVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CryptoKeyId == nil {
		return nil, errors.New("invalid value for required argument 'CryptoKeyId'")
	}
	if args.KeyRingId == nil {
		return nil, errors.New("invalid value for required argument 'KeyRingId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource CryptoKeyVersion
	err := ctx.RegisterResource("google-native:cloudkms/v1:CryptoKeyVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCryptoKeyVersion gets an existing CryptoKeyVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCryptoKeyVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CryptoKeyVersionState, opts ...pulumi.ResourceOption) (*CryptoKeyVersion, error) {
	var resource CryptoKeyVersion
	err := ctx.ReadResource("google-native:cloudkms/v1:CryptoKeyVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CryptoKeyVersion resources.
type cryptoKeyVersionState struct {
}

type CryptoKeyVersionState struct {
}

func (CryptoKeyVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyVersionState)(nil)).Elem()
}

type cryptoKeyVersionArgs struct {
	CryptoKeyId string `pulumi:"cryptoKeyId"`
	// ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
	ExternalProtectionLevelOptions *ExternalProtectionLevelOptions `pulumi:"externalProtectionLevelOptions"`
	KeyRingId                      string                          `pulumi:"keyRingId"`
	Location                       string                          `pulumi:"location"`
	Project                        string                          `pulumi:"project"`
	// The current state of the CryptoKeyVersion.
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a CryptoKeyVersion resource.
type CryptoKeyVersionArgs struct {
	CryptoKeyId pulumi.StringInput
	// ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.
	ExternalProtectionLevelOptions ExternalProtectionLevelOptionsPtrInput
	KeyRingId                      pulumi.StringInput
	Location                       pulumi.StringInput
	Project                        pulumi.StringInput
	// The current state of the CryptoKeyVersion.
	State *CryptoKeyVersionStateEnum
}

func (CryptoKeyVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyVersionArgs)(nil)).Elem()
}

type CryptoKeyVersionInput interface {
	pulumi.Input

	ToCryptoKeyVersionOutput() CryptoKeyVersionOutput
	ToCryptoKeyVersionOutputWithContext(ctx context.Context) CryptoKeyVersionOutput
}

func (*CryptoKeyVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKeyVersion)(nil))
}

func (i *CryptoKeyVersion) ToCryptoKeyVersionOutput() CryptoKeyVersionOutput {
	return i.ToCryptoKeyVersionOutputWithContext(context.Background())
}

func (i *CryptoKeyVersion) ToCryptoKeyVersionOutputWithContext(ctx context.Context) CryptoKeyVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyVersionOutput)
}

type CryptoKeyVersionOutput struct {
	*pulumi.OutputState
}

func (CryptoKeyVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKeyVersion)(nil))
}

func (o CryptoKeyVersionOutput) ToCryptoKeyVersionOutput() CryptoKeyVersionOutput {
	return o
}

func (o CryptoKeyVersionOutput) ToCryptoKeyVersionOutputWithContext(ctx context.Context) CryptoKeyVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CryptoKeyVersionOutput{})
}
