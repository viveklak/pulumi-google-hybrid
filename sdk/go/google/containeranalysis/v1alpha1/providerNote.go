// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new `Note`.
type ProviderNote struct {
	pulumi.CustomResourceState

	// A note describing an attestation role.
	AttestationAuthority AttestationAuthorityResponseOutput `pulumi:"attestationAuthority"`
	// A note describing a base image.
	BaseImage BasisResponseOutput `pulumi:"baseImage"`
	// Build provenance type for a verifiable build.
	BuildType BuildTypeResponseOutput `pulumi:"buildType"`
	// The time this note was created. This field can be used as a filter in list requests.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A note describing something that can be deployed.
	Deployable DeployableResponseOutput `pulumi:"deployable"`
	// A note describing a provider/analysis type.
	Discovery DiscoveryResponseOutput `pulumi:"discovery"`
	// Time of expiration for this note, null if note does not expire.
	ExpirationTime pulumi.StringOutput `pulumi:"expirationTime"`
	// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A detailed description of this `Note`.
	LongDescription pulumi.StringOutput `pulumi:"longDescription"`
	// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
	Name pulumi.StringOutput `pulumi:"name"`
	// A note describing a package hosted by various package managers.
	Package PackageResponseOutput `pulumi:"package"`
	// URLs associated with this note
	RelatedUrl RelatedUrlResponseArrayOutput `pulumi:"relatedUrl"`
	// A one sentence description of this `Note`.
	ShortDescription pulumi.StringOutput `pulumi:"shortDescription"`
	// The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// A note describing an upgrade.
	Upgrade UpgradeNoteResponseOutput `pulumi:"upgrade"`
	// A package vulnerability type of note.
	VulnerabilityType VulnerabilityTypeResponseOutput `pulumi:"vulnerabilityType"`
}

// NewProviderNote registers a new resource with the given unique name, arguments, and options.
func NewProviderNote(ctx *pulumi.Context,
	name string, args *ProviderNoteArgs, opts ...pulumi.ResourceOption) (*ProviderNote, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NotesId == nil {
		return nil, errors.New("invalid value for required argument 'NotesId'")
	}
	if args.ProvidersId == nil {
		return nil, errors.New("invalid value for required argument 'ProvidersId'")
	}
	var resource ProviderNote
	err := ctx.RegisterResource("google-cloud:containeranalysis/v1alpha1:ProviderNote", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderNote gets an existing ProviderNote resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderNote(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderNoteState, opts ...pulumi.ResourceOption) (*ProviderNote, error) {
	var resource ProviderNote
	err := ctx.ReadResource("google-cloud:containeranalysis/v1alpha1:ProviderNote", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderNote resources.
type providerNoteState struct {
	// A note describing an attestation role.
	AttestationAuthority *AttestationAuthorityResponse `pulumi:"attestationAuthority"`
	// A note describing a base image.
	BaseImage *BasisResponse `pulumi:"baseImage"`
	// Build provenance type for a verifiable build.
	BuildType *BuildTypeResponse `pulumi:"buildType"`
	// The time this note was created. This field can be used as a filter in list requests.
	CreateTime *string `pulumi:"createTime"`
	// A note describing something that can be deployed.
	Deployable *DeployableResponse `pulumi:"deployable"`
	// A note describing a provider/analysis type.
	Discovery *DiscoveryResponse `pulumi:"discovery"`
	// Time of expiration for this note, null if note does not expire.
	ExpirationTime *string `pulumi:"expirationTime"`
	// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
	Kind *string `pulumi:"kind"`
	// A detailed description of this `Note`.
	LongDescription *string `pulumi:"longDescription"`
	// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
	Name *string `pulumi:"name"`
	// A note describing a package hosted by various package managers.
	Package *PackageResponse `pulumi:"package"`
	// URLs associated with this note
	RelatedUrl []RelatedUrlResponse `pulumi:"relatedUrl"`
	// A one sentence description of this `Note`.
	ShortDescription *string `pulumi:"shortDescription"`
	// The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime *string `pulumi:"updateTime"`
	// A note describing an upgrade.
	Upgrade *UpgradeNoteResponse `pulumi:"upgrade"`
	// A package vulnerability type of note.
	VulnerabilityType *VulnerabilityTypeResponse `pulumi:"vulnerabilityType"`
}

type ProviderNoteState struct {
	// A note describing an attestation role.
	AttestationAuthority AttestationAuthorityResponsePtrInput
	// A note describing a base image.
	BaseImage BasisResponsePtrInput
	// Build provenance type for a verifiable build.
	BuildType BuildTypeResponsePtrInput
	// The time this note was created. This field can be used as a filter in list requests.
	CreateTime pulumi.StringPtrInput
	// A note describing something that can be deployed.
	Deployable DeployableResponsePtrInput
	// A note describing a provider/analysis type.
	Discovery DiscoveryResponsePtrInput
	// Time of expiration for this note, null if note does not expire.
	ExpirationTime pulumi.StringPtrInput
	// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
	Kind pulumi.StringPtrInput
	// A detailed description of this `Note`.
	LongDescription pulumi.StringPtrInput
	// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
	Name pulumi.StringPtrInput
	// A note describing a package hosted by various package managers.
	Package PackageResponsePtrInput
	// URLs associated with this note
	RelatedUrl RelatedUrlResponseArrayInput
	// A one sentence description of this `Note`.
	ShortDescription pulumi.StringPtrInput
	// The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime pulumi.StringPtrInput
	// A note describing an upgrade.
	Upgrade UpgradeNoteResponsePtrInput
	// A package vulnerability type of note.
	VulnerabilityType VulnerabilityTypeResponsePtrInput
}

func (ProviderNoteState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerNoteState)(nil)).Elem()
}

type providerNoteArgs struct {
	// A note describing an attestation role.
	AttestationAuthority *AttestationAuthority `pulumi:"attestationAuthority"`
	// A note describing a base image.
	BaseImage *Basis `pulumi:"baseImage"`
	// Build provenance type for a verifiable build.
	BuildType *BuildType `pulumi:"buildType"`
	// The time this note was created. This field can be used as a filter in list requests.
	CreateTime *string `pulumi:"createTime"`
	// A note describing something that can be deployed.
	Deployable *Deployable `pulumi:"deployable"`
	// A note describing a provider/analysis type.
	Discovery *Discovery `pulumi:"discovery"`
	// Time of expiration for this note, null if note does not expire.
	ExpirationTime *string `pulumi:"expirationTime"`
	// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
	Kind *string `pulumi:"kind"`
	// A detailed description of this `Note`.
	LongDescription *string `pulumi:"longDescription"`
	// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
	Name    *string `pulumi:"name"`
	NotesId string  `pulumi:"notesId"`
	// A note describing a package hosted by various package managers.
	Package     *Package `pulumi:"package"`
	ProvidersId string   `pulumi:"providersId"`
	// URLs associated with this note
	RelatedUrl []RelatedUrl `pulumi:"relatedUrl"`
	// A one sentence description of this `Note`.
	ShortDescription *string `pulumi:"shortDescription"`
	// The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime *string `pulumi:"updateTime"`
	// A note describing an upgrade.
	Upgrade *UpgradeNote `pulumi:"upgrade"`
	// A package vulnerability type of note.
	VulnerabilityType *VulnerabilityType `pulumi:"vulnerabilityType"`
}

// The set of arguments for constructing a ProviderNote resource.
type ProviderNoteArgs struct {
	// A note describing an attestation role.
	AttestationAuthority AttestationAuthorityPtrInput
	// A note describing a base image.
	BaseImage BasisPtrInput
	// Build provenance type for a verifiable build.
	BuildType BuildTypePtrInput
	// The time this note was created. This field can be used as a filter in list requests.
	CreateTime pulumi.StringPtrInput
	// A note describing something that can be deployed.
	Deployable DeployablePtrInput
	// A note describing a provider/analysis type.
	Discovery DiscoveryPtrInput
	// Time of expiration for this note, null if note does not expire.
	ExpirationTime pulumi.StringPtrInput
	// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
	Kind pulumi.StringPtrInput
	// A detailed description of this `Note`.
	LongDescription pulumi.StringPtrInput
	// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
	Name    pulumi.StringPtrInput
	NotesId pulumi.StringInput
	// A note describing a package hosted by various package managers.
	Package     PackagePtrInput
	ProvidersId pulumi.StringInput
	// URLs associated with this note
	RelatedUrl RelatedUrlArrayInput
	// A one sentence description of this `Note`.
	ShortDescription pulumi.StringPtrInput
	// The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime pulumi.StringPtrInput
	// A note describing an upgrade.
	Upgrade UpgradeNotePtrInput
	// A package vulnerability type of note.
	VulnerabilityType VulnerabilityTypePtrInput
}

func (ProviderNoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerNoteArgs)(nil)).Elem()
}

type ProviderNoteInput interface {
	pulumi.Input

	ToProviderNoteOutput() ProviderNoteOutput
	ToProviderNoteOutputWithContext(ctx context.Context) ProviderNoteOutput
}

func (*ProviderNote) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderNote)(nil))
}

func (i *ProviderNote) ToProviderNoteOutput() ProviderNoteOutput {
	return i.ToProviderNoteOutputWithContext(context.Background())
}

func (i *ProviderNote) ToProviderNoteOutputWithContext(ctx context.Context) ProviderNoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderNoteOutput)
}

type ProviderNoteOutput struct {
	*pulumi.OutputState
}

func (ProviderNoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderNote)(nil))
}

func (o ProviderNoteOutput) ToProviderNoteOutput() ProviderNoteOutput {
	return o
}

func (o ProviderNoteOutput) ToProviderNoteOutputWithContext(ctx context.Context) ProviderNoteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderNoteOutput{})
}