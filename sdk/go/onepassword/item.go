// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Item struct {
	pulumi.CustomResourceState

	Attachments OutFieldMapOutput   `pulumi:"attachments"`
	Category    pulumi.StringOutput `pulumi:"category"`
	Fields      OutFieldMapOutput   `pulumi:"fields"`
	References  OutFieldMapOutput   `pulumi:"references"`
	Sections    OutSectionMapOutput `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringOutput `pulumi:"vault"`
}

// NewItem registers a new resource with the given unique name, arguments, and options.
func NewItem(ctx *pulumi.Context,
	name string, args *ItemArgs, opts ...pulumi.ResourceOption) (*Item, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	if isZero(args.Category) {
		args.Category = pulumi.StringPtr("Item")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"references",
		"sections",
	})
	opts = append(opts, secrets)
	var resource Item
	err := ctx.RegisterResource("onepassword:index:Item", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetItem gets an existing Item resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ItemState, opts ...pulumi.ResourceOption) (*Item, error) {
	var resource Item
	err := ctx.ReadResource("onepassword:index:Item", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Item resources.
type itemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type ItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (ItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*itemState)(nil)).Elem()
}

type itemArgs struct {
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	Category    *string                          `pulumi:"category"`
	Fields      map[string]Field                 `pulumi:"fields"`
	Sections    map[string]Section               `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a Item resource.
type ItemArgs struct {
	Attachments pulumi.AssetOrArchiveMapInput
	Category    pulumi.StringPtrInput
	Fields      FieldMapInput
	Sections    SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (ItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*itemArgs)(nil)).Elem()
}

func (r *Item) Attachment(ctx *pulumi.Context, args *ItemAttachmentArgs) (ItemAttachmentResultOutput, error) {
	out, err := ctx.Call("onepassword:index:Item/attachment", args, ItemAttachmentResultOutput{}, r)
	if err != nil {
		return ItemAttachmentResultOutput{}, err
	}
	return out.(ItemAttachmentResultOutput), nil
}

type itemAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name string `pulumi:"name"`
}

// The set of arguments for the Attachment method of the Item resource.
type ItemAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name pulumi.StringInput
}

func (ItemAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*itemAttachmentArgs)(nil)).Elem()
}

// The resolved reference value
type ItemAttachmentResult struct {
	// the value of the attachment
	Value string `pulumi:"value"`
}

type ItemAttachmentResultOutput struct{ *pulumi.OutputState }

func (ItemAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemAttachmentResult)(nil)).Elem()
}

// the value of the attachment
func (o ItemAttachmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ItemAttachmentResult) string { return v.Value }).(pulumi.StringOutput)
}

type ItemInput interface {
	pulumi.Input

	ToItemOutput() ItemOutput
	ToItemOutputWithContext(ctx context.Context) ItemOutput
}

func (*Item) ElementType() reflect.Type {
	return reflect.TypeOf((**Item)(nil)).Elem()
}

func (i *Item) ToItemOutput() ItemOutput {
	return i.ToItemOutputWithContext(context.Background())
}

func (i *Item) ToItemOutputWithContext(ctx context.Context) ItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemOutput)
}

// ItemArrayInput is an input type that accepts ItemArray and ItemArrayOutput values.
// You can construct a concrete instance of `ItemArrayInput` via:
//
//	ItemArray{ ItemArgs{...} }
type ItemArrayInput interface {
	pulumi.Input

	ToItemArrayOutput() ItemArrayOutput
	ToItemArrayOutputWithContext(context.Context) ItemArrayOutput
}

type ItemArray []ItemInput

func (ItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Item)(nil)).Elem()
}

func (i ItemArray) ToItemArrayOutput() ItemArrayOutput {
	return i.ToItemArrayOutputWithContext(context.Background())
}

func (i ItemArray) ToItemArrayOutputWithContext(ctx context.Context) ItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemArrayOutput)
}

// ItemMapInput is an input type that accepts ItemMap and ItemMapOutput values.
// You can construct a concrete instance of `ItemMapInput` via:
//
//	ItemMap{ "key": ItemArgs{...} }
type ItemMapInput interface {
	pulumi.Input

	ToItemMapOutput() ItemMapOutput
	ToItemMapOutputWithContext(context.Context) ItemMapOutput
}

type ItemMap map[string]ItemInput

func (ItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Item)(nil)).Elem()
}

func (i ItemMap) ToItemMapOutput() ItemMapOutput {
	return i.ToItemMapOutputWithContext(context.Background())
}

func (i ItemMap) ToItemMapOutputWithContext(ctx context.Context) ItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemMapOutput)
}

type ItemOutput struct{ *pulumi.OutputState }

func (ItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Item)(nil)).Elem()
}

func (o ItemOutput) ToItemOutput() ItemOutput {
	return o
}

func (o ItemOutput) ToItemOutputWithContext(ctx context.Context) ItemOutput {
	return o
}

type ItemArrayOutput struct{ *pulumi.OutputState }

func (ItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Item)(nil)).Elem()
}

func (o ItemArrayOutput) ToItemArrayOutput() ItemArrayOutput {
	return o
}

func (o ItemArrayOutput) ToItemArrayOutputWithContext(ctx context.Context) ItemArrayOutput {
	return o
}

func (o ItemArrayOutput) Index(i pulumi.IntInput) ItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Item {
		return vs[0].([]*Item)[vs[1].(int)]
	}).(ItemOutput)
}

type ItemMapOutput struct{ *pulumi.OutputState }

func (ItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Item)(nil)).Elem()
}

func (o ItemMapOutput) ToItemMapOutput() ItemMapOutput {
	return o
}

func (o ItemMapOutput) ToItemMapOutputWithContext(ctx context.Context) ItemMapOutput {
	return o
}

func (o ItemMapOutput) MapIndex(k pulumi.StringInput) ItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Item {
		return vs[0].(map[string]*Item)[vs[1].(string)]
	}).(ItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ItemInput)(nil)).Elem(), &Item{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemArrayInput)(nil)).Elem(), ItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemMapInput)(nil)).Elem(), ItemMap{})
	pulumi.RegisterOutputType(ItemOutput{})
	pulumi.RegisterOutputType(ItemAttachmentResultOutput{})
	pulumi.RegisterOutputType(ItemArrayOutput{})
	pulumi.RegisterOutputType(ItemMapOutput{})
}
