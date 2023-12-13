﻿using System.Collections.Immutable;
using pulumi_resource_one_password_native_unofficial.OnePasswordCli;
using Pulumi.Experimental.Provider;

namespace pulumi_resource_one_password_native_unofficial.Domain;

public partial record FunctionType(
    string Urn,
    string InputCategory,
    string OutputCategory,
    TemplateMetadata.TransformOutputs TransformItemToOutputs
) : IPulumiItemType
{
    public ImmutableDictionary<string, PropertyValue> TransformOutputs(Item.Response item)
    {
        return TransformItemToOutputs(this, item, null);
    }
    
    public static FunctionType GetVault { get; } = new(
        "one-password-native-unofficial:index:GetVault",
        "GetVault",
        "",
        (_, _, _) => ImmutableDictionary<string, PropertyValue>.Empty);
    public static FunctionType GetSecretReference { get; } = new(
        "one-password-native-unofficial:index:GetSecretReference",
        "GetSecretReference",
        "",
        (_, _, _) => ImmutableDictionary<string, PropertyValue>.Empty);
    public static FunctionType Read { get; } = new(
        "one-password-native-unofficial:index:Read",
        "Read",
        "",
        (_, _, _) => ImmutableDictionary<string, PropertyValue>.Empty);
    public static FunctionType Inject { get; } = new(
        "one-password-native-unofficial:index:Inject",
        "Inject",
        "",
        (_, _, _) => ImmutableDictionary<string, PropertyValue>.Empty);
    public static FunctionType GetAttachment { get; } = new(
        "one-password-native-unofficial:index:GetAttachment",
        "GetAttachment",
        "",
        (_, _, _) => ImmutableDictionary<string, PropertyValue>.Empty);

    public virtual bool Equals(FunctionType? other)
    {
        if (ReferenceEquals(null, other)) return false;
        if (ReferenceEquals(this, other)) return true;
        return string.Equals(Urn, other.Urn, StringComparison.OrdinalIgnoreCase);
    }

    public override int GetHashCode()
    {
        return StringComparer.OrdinalIgnoreCase.GetHashCode(Urn);
    }
}
