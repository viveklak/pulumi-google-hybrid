// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1Alpha2.Outputs
{

    [OutputType]
    public sealed class PostalAddressResponse
    {
        /// <summary>
        /// Unstructured address lines describing the lower levels of an address. Because values in address_lines do not have type information and may sometimes contain multiple values in a single field (e.g. "Austin, TX"), it is important that the line order is clear. The order of address lines should be "envelope order" for the country/region of the address. In places where this can vary (e.g. Japan), address_language is used to make it explicit (e.g. "ja" for large-to-small ordering and "ja-Latn" or "en" for small-to-large). This way, the most specific line of an address can be selected based on the language. The minimum permitted structural representation of an address consists of a region_code with all remaining information placed in the address_lines. It would be possible to format such an address very approximately without geocoding, but no semantic reasoning could be made about any of the address components until it was at least partially resolved. Creating an address only containing a region_code and address_lines, and then geocoding is the recommended way to handle completely unstructured addresses (as opposed to guessing which parts of the address should be localities or administrative areas).
        /// </summary>
        public readonly ImmutableArray<string> AddressLines;
        /// <summary>
        /// Optional. Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state, a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community (e.g. "Barcelona" and not "Catalonia"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland this should be left unpopulated.
        /// </summary>
        public readonly string AdministrativeArea;
        /// <summary>
        /// Optional. BCP-47 language code of the contents of this address (if known). This is often the UI language of the input form or is expected to match one of the languages used in the address' country/region, or their transliterated equivalents. This can affect formatting in certain countries, but is not critical to the correctness of the data and will never affect any validation or other non-formatting related operations. If this value is not known, it should be omitted (rather than specifying a possibly incorrect default). Examples: "zh-Hant", "ja", "ja-Latn", "en".
        /// </summary>
        public readonly string LanguageCode;
        /// <summary>
        /// Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use address_lines.
        /// </summary>
        public readonly string Locality;
        /// <summary>
        /// Optional. The name of the organization at the address.
        /// </summary>
        public readonly string Organization;
        /// <summary>
        /// Optional. Postal code of the address. Not all countries use or require postal codes to be present, but where they are used, they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
        /// </summary>
        public readonly string PostalCode;
        /// <summary>
        /// Optional. The recipient at the address. This field may, under certain circumstances, contain multiline information. For example, it might contain "care of" information.
        /// </summary>
        public readonly ImmutableArray<string> Recipients;
        /// <summary>
        /// CLDR region code of the country/region of the address. This is never inferred and it is up to the user to ensure the value is correct. See http://cldr.unicode.org/ and http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: "CH" for Switzerland.
        /// </summary>
        public readonly string RegionCode;
        /// <summary>
        /// The schema revision of the `PostalAddress`. This must be set to 0, which is the latest revision. All new revisions **must** be backward compatible with old revisions.
        /// </summary>
        public readonly int Revision;
        /// <summary>
        /// Optional. Additional, country-specific, sorting code. This is not used in most regions. Where it is used, the value is either a string like "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number alone, representing the "sector code" (Jamaica), "delivery area indicator" (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
        /// </summary>
        public readonly string SortingCode;
        /// <summary>
        /// Optional. Sublocality of the address. For example, this can be neighborhoods, boroughs, districts.
        /// </summary>
        public readonly string Sublocality;

        [OutputConstructor]
        private PostalAddressResponse(
            ImmutableArray<string> addressLines,

            string administrativeArea,

            string languageCode,

            string locality,

            string organization,

            string postalCode,

            ImmutableArray<string> recipients,

            string regionCode,

            int revision,

            string sortingCode,

            string sublocality)
        {
            AddressLines = addressLines;
            AdministrativeArea = administrativeArea;
            LanguageCode = languageCode;
            Locality = locality;
            Organization = organization;
            PostalCode = postalCode;
            Recipients = recipients;
            RegionCode = regionCode;
            Revision = revision;
            SortingCode = sortingCode;
            Sublocality = sublocality;
        }
    }
}
