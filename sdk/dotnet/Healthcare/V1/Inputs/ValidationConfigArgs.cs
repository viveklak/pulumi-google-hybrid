// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Inputs
{

    /// <summary>
    /// Contains the configuration for FHIR profiles and validation.
    /// </summary>
    public sealed class ValidationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to disable FHIRPath validation for incoming resources. Set this to true to disable checking incoming resources for conformance against FHIRPath requirement defined in the FHIR specification. This property only affects resource types that do not have profiles configured for them, any rules in enabled implementation guides will still be enforced.
        /// </summary>
        [Input("disableFhirpathValidation")]
        public Input<bool>? DisableFhirpathValidation { get; set; }

        /// <summary>
        /// Whether to disable profile validation for this FHIR store. Set this to true to disable checking incoming resources for conformance against structure definitions in this FHIR store.
        /// </summary>
        [Input("disableProfileValidation")]
        public Input<bool>? DisableProfileValidation { get; set; }

        /// <summary>
        /// Whether to disable reference type validation for incoming resources. Set this to true to disable checking incoming resources for conformance against reference type requirement defined in the FHIR specification. This property only affects resource types that do not have profiles configured for them, any rules in enabled implementation guides will still be enforced.
        /// </summary>
        [Input("disableReferenceTypeValidation")]
        public Input<bool>? DisableReferenceTypeValidation { get; set; }

        /// <summary>
        /// Whether to disable required fields validation for incoming resources. Set this to true to disable checking incoming resources for conformance against required fields requirement defined in the FHIR specification. This property only affects resource types that do not have profiles configured for them, any rules in enabled implementation guides will still be enforced.
        /// </summary>
        [Input("disableRequiredFieldValidation")]
        public Input<bool>? DisableRequiredFieldValidation { get; set; }

        [Input("enabledImplementationGuides")]
        private InputList<string>? _enabledImplementationGuides;

        /// <summary>
        /// A list of implementation guide URLs in this FHIR store that are used to configure the profiles to use for validation. For example, to use the US Core profiles for validation, set `enabled_implementation_guides` to `["http://hl7.org/fhir/us/core/ImplementationGuide/ig"]`. If `enabled_implementation_guides` is empty or omitted, then incoming resources are only required to conform to the base FHIR profiles. Otherwise, a resource must conform to at least one profile listed in the `global` property of one of the enabled ImplementationGuides. The Cloud Healthcare API does not currently enforce all of the rules in a StructureDefinition. The following rules are supported: - min/max - minValue/maxValue - maxLength - type - fixed[x] - pattern[x] on simple types - slicing, when using "value" as the discriminator type When a URL cannot be resolved (for example, in a type assertion), the server does not return an error.
        /// </summary>
        public InputList<string> EnabledImplementationGuides
        {
            get => _enabledImplementationGuides ?? (_enabledImplementationGuides = new InputList<string>());
            set => _enabledImplementationGuides = value;
        }

        public ValidationConfigArgs()
        {
        }
    }
}
