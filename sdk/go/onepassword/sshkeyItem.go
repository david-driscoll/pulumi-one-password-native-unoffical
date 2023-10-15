// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SSHKeyItem struct {
	pulumi.CustomResourceState

	Category   pulumi.StringOutput    `pulumi:"category"`
	Fields     GetFieldArrayOutput    `pulumi:"fields"`
	Id         pulumi.StringOutput    `pulumi:"id"`
	Notes      pulumi.StringPtrOutput `pulumi:"notes"`
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	Sections   GetSectionArrayOutput  `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringOutput `pulumi:"vault"`
}

// NewSSHKeyItem registers a new resource with the given unique name, arguments, and options.
func NewSSHKeyItem(ctx *pulumi.Context,
	name string, args *SSHKeyItemArgs, opts ...pulumi.ResourceOption) (*SSHKeyItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	var resource SSHKeyItem
	err := ctx.RegisterResource("onepassword:index:SSHKeyItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSSHKeyItem gets an existing SSHKeyItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSSHKeyItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SSHKeyItemState, opts ...pulumi.ResourceOption) (*SSHKeyItem, error) {
	var resource SSHKeyItem
	err := ctx.ReadResource("onepassword:index:SSHKeyItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SSHKeyItem resources.
type sshkeyItemState struct {
}

type SSHKeyItemState struct {
}

func (SSHKeyItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshkeyItemState)(nil)).Elem()
}

type sshkeyItemArgs struct {
	Fields     []Field   `pulumi:"fields"`
	Notes      *string   `pulumi:"notes"`
	PrivateKey *string   `pulumi:"privateKey"`
	Sections   []Section `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a SSHKeyItem resource.
type SSHKeyItemArgs struct {
	Fields     FieldArrayInput
	Notes      pulumi.StringPtrInput
	PrivateKey pulumi.StringPtrInput
	Sections   SectionArrayInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (SSHKeyItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshkeyItemArgs)(nil)).Elem()
}

type SSHKeyItemInput interface {
	pulumi.Input

	ToSSHKeyItemOutput() SSHKeyItemOutput
	ToSSHKeyItemOutputWithContext(ctx context.Context) SSHKeyItemOutput
}

func (*SSHKeyItem) ElementType() reflect.Type {
	return reflect.TypeOf((**SSHKeyItem)(nil)).Elem()
}

func (i *SSHKeyItem) ToSSHKeyItemOutput() SSHKeyItemOutput {
	return i.ToSSHKeyItemOutputWithContext(context.Background())
}

func (i *SSHKeyItem) ToSSHKeyItemOutputWithContext(ctx context.Context) SSHKeyItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeyItemOutput)
}

// SSHKeyItemArrayInput is an input type that accepts SSHKeyItemArray and SSHKeyItemArrayOutput values.
// You can construct a concrete instance of `SSHKeyItemArrayInput` via:
//
//	SSHKeyItemArray{ SSHKeyItemArgs{...} }
type SSHKeyItemArrayInput interface {
	pulumi.Input

	ToSSHKeyItemArrayOutput() SSHKeyItemArrayOutput
	ToSSHKeyItemArrayOutputWithContext(context.Context) SSHKeyItemArrayOutput
}

type SSHKeyItemArray []SSHKeyItemInput

func (SSHKeyItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSHKeyItem)(nil)).Elem()
}

func (i SSHKeyItemArray) ToSSHKeyItemArrayOutput() SSHKeyItemArrayOutput {
	return i.ToSSHKeyItemArrayOutputWithContext(context.Background())
}

func (i SSHKeyItemArray) ToSSHKeyItemArrayOutputWithContext(ctx context.Context) SSHKeyItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeyItemArrayOutput)
}

// SSHKeyItemMapInput is an input type that accepts SSHKeyItemMap and SSHKeyItemMapOutput values.
// You can construct a concrete instance of `SSHKeyItemMapInput` via:
//
//	SSHKeyItemMap{ "key": SSHKeyItemArgs{...} }
type SSHKeyItemMapInput interface {
	pulumi.Input

	ToSSHKeyItemMapOutput() SSHKeyItemMapOutput
	ToSSHKeyItemMapOutputWithContext(context.Context) SSHKeyItemMapOutput
}

type SSHKeyItemMap map[string]SSHKeyItemInput

func (SSHKeyItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSHKeyItem)(nil)).Elem()
}

func (i SSHKeyItemMap) ToSSHKeyItemMapOutput() SSHKeyItemMapOutput {
	return i.ToSSHKeyItemMapOutputWithContext(context.Background())
}

func (i SSHKeyItemMap) ToSSHKeyItemMapOutputWithContext(ctx context.Context) SSHKeyItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeyItemMapOutput)
}

type SSHKeyItemOutput struct{ *pulumi.OutputState }

func (SSHKeyItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SSHKeyItem)(nil)).Elem()
}

func (o SSHKeyItemOutput) ToSSHKeyItemOutput() SSHKeyItemOutput {
	return o
}

func (o SSHKeyItemOutput) ToSSHKeyItemOutputWithContext(ctx context.Context) SSHKeyItemOutput {
	return o
}

type SSHKeyItemArrayOutput struct{ *pulumi.OutputState }

func (SSHKeyItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSHKeyItem)(nil)).Elem()
}

func (o SSHKeyItemArrayOutput) ToSSHKeyItemArrayOutput() SSHKeyItemArrayOutput {
	return o
}

func (o SSHKeyItemArrayOutput) ToSSHKeyItemArrayOutputWithContext(ctx context.Context) SSHKeyItemArrayOutput {
	return o
}

func (o SSHKeyItemArrayOutput) Index(i pulumi.IntInput) SSHKeyItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SSHKeyItem {
		return vs[0].([]*SSHKeyItem)[vs[1].(int)]
	}).(SSHKeyItemOutput)
}

type SSHKeyItemMapOutput struct{ *pulumi.OutputState }

func (SSHKeyItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSHKeyItem)(nil)).Elem()
}

func (o SSHKeyItemMapOutput) ToSSHKeyItemMapOutput() SSHKeyItemMapOutput {
	return o
}

func (o SSHKeyItemMapOutput) ToSSHKeyItemMapOutputWithContext(ctx context.Context) SSHKeyItemMapOutput {
	return o
}

func (o SSHKeyItemMapOutput) MapIndex(k pulumi.StringInput) SSHKeyItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SSHKeyItem {
		return vs[0].(map[string]*SSHKeyItem)[vs[1].(string)]
	}).(SSHKeyItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeyItemInput)(nil)).Elem(), &SSHKeyItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeyItemArrayInput)(nil)).Elem(), SSHKeyItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeyItemMapInput)(nil)).Elem(), SSHKeyItemMap{})
	pulumi.RegisterOutputType(SSHKeyItemOutput{})
	pulumi.RegisterOutputType(SSHKeyItemArrayOutput{})
	pulumi.RegisterOutputType(SSHKeyItemMapOutput{})
}
