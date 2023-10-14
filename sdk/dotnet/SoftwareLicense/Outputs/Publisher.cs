// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword.SoftwareLicense.Outputs
{

    [OutputType]
    public sealed class Publisher
    {
        public readonly string? DownloadPage;
        public readonly string? Publisher;
        public readonly string? RetailPrice;
        public readonly string? SupportEmail;
        public readonly string? Website;

        [OutputConstructor]
        private Publisher(
            string? downloadPage,

            string? publisher,

            string? retailPrice,

            string? supportEmail,

            string? website)
        {
            DownloadPage = downloadPage;
            Publisher = publisher;
            RetailPrice = retailPrice;
            SupportEmail = supportEmail;
            Website = website;
        }
    }
}
