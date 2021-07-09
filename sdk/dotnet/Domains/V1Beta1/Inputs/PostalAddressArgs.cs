// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1Beta1.Inputs
{

    /// <summary>
    /// Represents a postal address, e.g. for postal delivery or payments addresses. Given a postal address, a postal service can deliver items to a premise, P.O. Box or similar. It is not intended to model geographical locations (roads, towns, mountains). In typical usage an address would be created via user input or from importing existing data, depending on the type of process. Advice on address input / editing: - Use an i18n-ready address widget such as https://github.com/google/libaddressinput) - Users should not be presented with UI elements for input or editing of fields outside countries where that field is used. For more guidance on how to use this schema, please see: https://support.google.com/business/answer/6397478
    /// </summary>
    public sealed class PostalAddressArgs : Pulumi.ResourceArgs
    {
        [Input("addressLines")]
        private InputList<string>? _addressLines;

        /// <summary>
        /// Unstructured address lines describing the lower levels of an address. Because values in address_lines do not have type information and may sometimes contain multiple values in a single field (e.g. "Austin, TX"), it is important that the line order is clear. The order of address lines should be "envelope order" for the country/region of the address. In places where this can vary (e.g. Japan), address_language is used to make it explicit (e.g. "ja" for large-to-small ordering and "ja-Latn" or "en" for small-to-large). This way, the most specific line of an address can be selected based on the language. The minimum permitted structural representation of an address consists of a region_code with all remaining information placed in the address_lines. It would be possible to format such an address very approximately without geocoding, but no semantic reasoning could be made about any of the address components until it was at least partially resolved. Creating an address only containing a region_code and address_lines, and then geocoding is the recommended way to handle completely unstructured addresses (as opposed to guessing which parts of the address should be localities or administrative areas).
        /// </summary>
        public InputList<string> AddressLines
        {
            get => _addressLines ?? (_addressLines = new InputList<string>());
            set => _addressLines = value;
        }

        /// <summary>
        /// Optional. Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state, a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community (e.g. "Barcelona" and not "Catalonia"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland this should be left unpopulated.
        /// </summary>
        [Input("administrativeArea")]
        public Input<string>? AdministrativeArea { get; set; }

        /// <summary>
        /// Optional. BCP-47 language code of the contents of this address (if known). This is often the UI language of the input form or is expected to match one of the languages used in the address' country/region, or their transliterated equivalents. This can affect formatting in certain countries, but is not critical to the correctness of the data and will never affect any validation or other non-formatting related operations. If this value is not known, it should be omitted (rather than specifying a possibly incorrect default). Examples: "zh-Hant", "ja", "ja-Latn", "en".
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use address_lines.
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// Optional. The name of the organization at the address.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// Optional. Postal code of the address. Not all countries use or require postal codes to be present, but where they are used, they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        [Input("recipients")]
        private InputList<string>? _recipients;

        /// <summary>
        /// Optional. The recipient at the address. This field may, under certain circumstances, contain multiline information. For example, it might contain "care of" information.
        /// </summary>
        public InputList<string> Recipients
        {
            get => _recipients ?? (_recipients = new InputList<string>());
            set => _recipients = value;
        }

        /// <summary>
        /// CLDR region code of the country/region of the address. This is never inferred and it is up to the user to ensure the value is correct. See http://cldr.unicode.org/ and http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: "CH" for Switzerland.
        /// </summary>
        [Input("regionCode", required: true)]
        public Input<string> RegionCode { get; set; } = null!;

        /// <summary>
        /// The schema revision of the `PostalAddress`. This must be set to 0, which is the latest revision. All new revisions **must** be backward compatible with old revisions.
        /// </summary>
        [Input("revision")]
        public Input<int>? Revision { get; set; }

        /// <summary>
        /// Optional. Additional, country-specific, sorting code. This is not used in most regions. Where it is used, the value is either a string like "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number alone, representing the "sector code" (Jamaica), "delivery area indicator" (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
        /// </summary>
        [Input("sortingCode")]
        public Input<string>? SortingCode { get; set; }

        /// <summary>
        /// Optional. Sublocality of the address. For example, this can be neighborhoods, boroughs, districts.
        /// </summary>
        [Input("sublocality")]
        public Input<string>? Sublocality { get; set; }

        public PostalAddressArgs()
        {
        }
    }
}
