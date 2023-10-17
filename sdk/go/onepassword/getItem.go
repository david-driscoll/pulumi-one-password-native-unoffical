// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupItem(ctx *pulumi.Context, args *LookupItemArgs, opts ...pulumi.InvokeOption) (*LookupItemResult, error) {
	var rv LookupItemResult
	err := ctx.Invoke("onepassword:index:GetItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupItemArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type LookupItemResult struct {
	Attachments map[string]OutField   `pulumi:"attachments"`
	Category    string                `pulumi:"category"`
	Fields      map[string]OutField   `pulumi:"fields"`
	References  map[string]OutField   `pulumi:"references"`
	Sections    map[string]OutSection `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

func LookupItemOutput(ctx *pulumi.Context, args LookupItemOutputArgs, opts ...pulumi.InvokeOption) LookupItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupItemResult, error) {
			args := v.(LookupItemArgs)
			r, err := LookupItem(ctx, &args, opts...)
			var s LookupItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupItemResultOutput)
}

type LookupItemOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (LookupItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupItemArgs)(nil)).Elem()
}

type LookupItemResultOutput struct{ *pulumi.OutputState }

func (LookupItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupItemResult)(nil)).Elem()
}

func (o LookupItemResultOutput) ToLookupItemResultOutput() LookupItemResultOutput {
	return o
}

func (o LookupItemResultOutput) ToLookupItemResultOutputWithContext(ctx context.Context) LookupItemResultOutput {
	return o
}

func (o LookupItemResultOutput) Attachments() OutFieldMapOutput {
	return o.ApplyT(func(v LookupItemResult) map[string]OutField { return v.Attachments }).(OutFieldMapOutput)
}

func (o LookupItemResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupItemResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o LookupItemResultOutput) Fields() OutFieldMapOutput {
	return o.ApplyT(func(v LookupItemResult) map[string]OutField { return v.Fields }).(OutFieldMapOutput)
}

func (o LookupItemResultOutput) References() OutFieldMapOutput {
	return o.ApplyT(func(v LookupItemResult) map[string]OutField { return v.References }).(OutFieldMapOutput)
}

func (o LookupItemResultOutput) Sections() OutSectionMapOutput {
	return o.ApplyT(func(v LookupItemResult) map[string]OutSection { return v.Sections }).(OutSectionMapOutput)
}

// An array of strings of the tags assigned to the item.
func (o LookupItemResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupItemResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o LookupItemResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupItemResult) string { return v.Title }).(pulumi.StringOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o LookupItemResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupItemResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// The UUID of the vault the item is in.
func (o LookupItemResultOutput) Vault() pulumi.StringOutput {
	return o.ApplyT(func(v LookupItemResult) string { return v.Vault }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupItemResultOutput{})
}
