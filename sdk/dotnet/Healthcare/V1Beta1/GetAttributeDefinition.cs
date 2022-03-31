// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1
{
    public static class GetAttributeDefinition
    {
        /// <summary>
        /// Gets the specified Attribute definition.
        /// </summary>
        public static Task<GetAttributeDefinitionResult> InvokeAsync(GetAttributeDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAttributeDefinitionResult>("google-native:healthcare/v1beta1:getAttributeDefinition", args ?? new GetAttributeDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Attribute definition.
        /// </summary>
        public static Output<GetAttributeDefinitionResult> Invoke(GetAttributeDefinitionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAttributeDefinitionResult>("google-native:healthcare/v1beta1:getAttributeDefinition", args ?? new GetAttributeDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAttributeDefinitionArgs : Pulumi.InvokeArgs
    {
        [Input("attributeDefinitionId", required: true)]
        public string AttributeDefinitionId { get; set; } = null!;

        [Input("consentStoreId", required: true)]
        public string ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAttributeDefinitionArgs()
        {
        }
    }

    public sealed class GetAttributeDefinitionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("attributeDefinitionId", required: true)]
        public Input<string> AttributeDefinitionId { get; set; } = null!;

        [Input("consentStoreId", required: true)]
        public Input<string> ConsentStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAttributeDefinitionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAttributeDefinitionResult
    {
        /// <summary>
        /// Possible values for the attribute. The number of allowed values must not exceed 500. An empty list is invalid. The list can only be expanded after creation.
        /// </summary>
        public readonly ImmutableArray<string> AllowedValues;
        /// <summary>
        /// The category of the attribute. The value of this field cannot be changed after creation.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Optional. Default values of the attribute in Consents. If no default values are specified, it defaults to an empty value.
        /// </summary>
        public readonly ImmutableArray<string> ConsentDefaultValues;
        /// <summary>
        /// Optional. Default value of the attribute in User data mappings. If no default value is specified, it defaults to an empty value. This field is only applicable to attributes of the category `RESOURCE`.
        /// </summary>
        public readonly string DataMappingDefaultValue;
        /// <summary>
        /// Optional. A description of the attribute.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Resource name of the Attribute definition, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/attributeDefinitions/{attribute_definition_id}`. Cannot be changed after creation.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetAttributeDefinitionResult(
            ImmutableArray<string> allowedValues,

            string category,

            ImmutableArray<string> consentDefaultValues,

            string dataMappingDefaultValue,

            string description,

            string name)
        {
            AllowedValues = allowedValues;
            Category = category;
            ConsentDefaultValues = consentDefaultValues;
            DataMappingDefaultValue = dataMappingDefaultValue;
            Description = description;
            Name = name;
        }
    }
}
