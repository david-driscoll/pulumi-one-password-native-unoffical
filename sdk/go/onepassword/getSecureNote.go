// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSecureNote(ctx *pulumi.Context, args *GetSecureNoteArgs, opts ...pulumi.InvokeOption) (*GetSecureNoteResult, error) {
	var rv GetSecureNoteResult
	err := ctx.Invoke("onepassword:index:GetSecureNote", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSecureNoteArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetSecureNoteResult struct {
	Category *string               `pulumi:"category"`
	Fields   map[string]GetField   `pulumi:"fields"`
	Id       *string               `pulumi:"id"`
	Notes    *string               `pulumi:"notes"`
	Sections map[string]GetSection `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

func GetSecureNoteOutput(ctx *pulumi.Context, args GetSecureNoteOutputArgs, opts ...pulumi.InvokeOption) GetSecureNoteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecureNoteResult, error) {
			args := v.(GetSecureNoteArgs)
			r, err := GetSecureNote(ctx, &args, opts...)
			var s GetSecureNoteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecureNoteResultOutput)
}

type GetSecureNoteOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetSecureNoteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecureNoteArgs)(nil)).Elem()
}

type GetSecureNoteResultOutput struct{ *pulumi.OutputState }

func (GetSecureNoteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecureNoteResult)(nil)).Elem()
}

func (o GetSecureNoteResultOutput) ToGetSecureNoteResultOutput() GetSecureNoteResultOutput {
	return o
}

func (o GetSecureNoteResultOutput) ToGetSecureNoteResultOutputWithContext(ctx context.Context) GetSecureNoteResultOutput {
	return o
}

func (o GetSecureNoteResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecureNoteResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o GetSecureNoteResultOutput) Fields() GetFieldMapOutput {
	return o.ApplyT(func(v GetSecureNoteResult) map[string]GetField { return v.Fields }).(GetFieldMapOutput)
}

func (o GetSecureNoteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecureNoteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GetSecureNoteResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecureNoteResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetSecureNoteResultOutput) Sections() GetSectionMapOutput {
	return o.ApplyT(func(v GetSecureNoteResult) map[string]GetSection { return v.Sections }).(GetSectionMapOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetSecureNoteResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecureNoteResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetSecureNoteResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecureNoteResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetSecureNoteResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecureNoteResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

// The UUID of the vault the item is in.
func (o GetSecureNoteResultOutput) Vault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecureNoteResult) *string { return v.Vault }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecureNoteResultOutput{})
}
