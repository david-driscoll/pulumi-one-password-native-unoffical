// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Rocket.Surgery.OnePasswordNativeUnofficial
{
    public static class GetVault
    {
        /// <summary>
        /// Use this data source to get details of a vault by either its name or uuid.
        /// </summary>
        public static Task<GetVaultResult> InvokeAsync(GetVaultArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVaultResult>("one-password-native-unofficial:index:GetVault", args ?? new GetVaultArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get details of a vault by either its name or uuid.
        /// </summary>
        public static Output<GetVaultResult> Invoke(GetVaultInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVaultResult>("one-password-native-unofficial:index:GetVault", args ?? new GetVaultInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVaultArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The vault to get information of.  Can be either the name or the UUID.
        /// </summary>
        [Input("vault", required: true)]
        public string Vault { get; set; } = null!;

        public GetVaultArgs()
        {
        }
    }

    public sealed class GetVaultInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The vault to get information of.  Can be either the name or the UUID.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public GetVaultInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVaultResult
    {
        /// <summary>
        /// The UUID of the vault to retrieve. This field will be populated with the UUID of the vault if the vault it looked up by its name.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the vault to retrieve. This field will be populated with the name of the vault if the vault it looked up by its UUID.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetVaultResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
