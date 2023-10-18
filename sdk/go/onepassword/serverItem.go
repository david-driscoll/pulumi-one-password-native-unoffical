// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-onepassword/sdk/go/onepassword/server"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerItem struct {
	pulumi.CustomResourceState

	AdminConsole    server.AdminConsoleSectionPtrOutput    `pulumi:"adminConsole"`
	Attachments     OutFieldMapOutput                      `pulumi:"attachments"`
	Category        pulumi.StringOutput                    `pulumi:"category"`
	Fields          OutFieldMapOutput                      `pulumi:"fields"`
	HostingProvider server.HostingProviderSectionPtrOutput `pulumi:"hostingProvider"`
	Notes           pulumi.StringPtrOutput                 `pulumi:"notes"`
	Password        pulumi.StringPtrOutput                 `pulumi:"password"`
	References      OutFieldMapOutput                      `pulumi:"references"`
	Sections        OutSectionMapOutput                    `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title    pulumi.StringOutput    `pulumi:"title"`
	Url      pulumi.StringPtrOutput `pulumi:"url"`
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringOutput `pulumi:"vault"`
}

// NewServerItem registers a new resource with the given unique name, arguments, and options.
func NewServerItem(ctx *pulumi.Context,
	name string, args *ServerItemArgs, opts ...pulumi.ResourceOption) (*ServerItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	args.Category = pulumi.StringPtr("Server")
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"adminConsole",
		"attachments",
		"fields",
		"password",
		"references",
		"sections",
	})
	opts = append(opts, secrets)
	var resource ServerItem
	err := ctx.RegisterResource("one-password-native:index:ServerItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerItem gets an existing ServerItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerItemState, opts ...pulumi.ResourceOption) (*ServerItem, error) {
	var resource ServerItem
	err := ctx.ReadResource("one-password-native:index:ServerItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerItem resources.
type serverItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type ServerItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (ServerItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverItemState)(nil)).Elem()
}

type serverItemArgs struct {
	AdminConsole *server.AdminConsoleSection      `pulumi:"adminConsole"`
	Attachments  map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category        *string                        `pulumi:"category"`
	Fields          map[string]Field               `pulumi:"fields"`
	HostingProvider *server.HostingProviderSection `pulumi:"hostingProvider"`
	Notes           *string                        `pulumi:"notes"`
	Password        *string                        `pulumi:"password"`
	Sections        map[string]Section             `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title    *string `pulumi:"title"`
	Url      *string `pulumi:"url"`
	Username *string `pulumi:"username"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a ServerItem resource.
type ServerItemArgs struct {
	AdminConsole server.AdminConsoleSectionPtrInput
	Attachments  pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category        pulumi.StringPtrInput
	Fields          FieldMapInput
	HostingProvider server.HostingProviderSectionPtrInput
	Notes           pulumi.StringPtrInput
	Password        pulumi.StringPtrInput
	Sections        SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title    pulumi.StringPtrInput
	Url      pulumi.StringPtrInput
	Username pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (ServerItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverItemArgs)(nil)).Elem()
}

func (r *ServerItem) GetAttachment(ctx *pulumi.Context, args *ServerItemGetAttachmentArgs) (ServerItemGetAttachmentResultOutput, error) {
	out, err := ctx.Call("one-password-native:index:ServerItem/attachment", args, ServerItemGetAttachmentResultOutput{}, r)
	if err != nil {
		return ServerItemGetAttachmentResultOutput{}, err
	}
	return out.(ServerItemGetAttachmentResultOutput), nil
}

type serverItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name string `pulumi:"name"`
}

// The set of arguments for the GetAttachment method of the ServerItem resource.
type ServerItemGetAttachmentArgs struct {
	// The name or uuid of the attachment to get
	Name pulumi.StringInput
}

func (ServerItemGetAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverItemGetAttachmentArgs)(nil)).Elem()
}

// The resolved reference value
type ServerItemGetAttachmentResult struct {
	// the value of the attachment
	Value string `pulumi:"value"`
}

type ServerItemGetAttachmentResultOutput struct{ *pulumi.OutputState }

func (ServerItemGetAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerItemGetAttachmentResult)(nil)).Elem()
}

// the value of the attachment
func (o ServerItemGetAttachmentResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ServerItemGetAttachmentResult) string { return v.Value }).(pulumi.StringOutput)
}

type ServerItemInput interface {
	pulumi.Input

	ToServerItemOutput() ServerItemOutput
	ToServerItemOutputWithContext(ctx context.Context) ServerItemOutput
}

func (*ServerItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerItem)(nil)).Elem()
}

func (i *ServerItem) ToServerItemOutput() ServerItemOutput {
	return i.ToServerItemOutputWithContext(context.Background())
}

func (i *ServerItem) ToServerItemOutputWithContext(ctx context.Context) ServerItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerItemOutput)
}

// ServerItemArrayInput is an input type that accepts ServerItemArray and ServerItemArrayOutput values.
// You can construct a concrete instance of `ServerItemArrayInput` via:
//
//	ServerItemArray{ ServerItemArgs{...} }
type ServerItemArrayInput interface {
	pulumi.Input

	ToServerItemArrayOutput() ServerItemArrayOutput
	ToServerItemArrayOutputWithContext(context.Context) ServerItemArrayOutput
}

type ServerItemArray []ServerItemInput

func (ServerItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerItem)(nil)).Elem()
}

func (i ServerItemArray) ToServerItemArrayOutput() ServerItemArrayOutput {
	return i.ToServerItemArrayOutputWithContext(context.Background())
}

func (i ServerItemArray) ToServerItemArrayOutputWithContext(ctx context.Context) ServerItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerItemArrayOutput)
}

// ServerItemMapInput is an input type that accepts ServerItemMap and ServerItemMapOutput values.
// You can construct a concrete instance of `ServerItemMapInput` via:
//
//	ServerItemMap{ "key": ServerItemArgs{...} }
type ServerItemMapInput interface {
	pulumi.Input

	ToServerItemMapOutput() ServerItemMapOutput
	ToServerItemMapOutputWithContext(context.Context) ServerItemMapOutput
}

type ServerItemMap map[string]ServerItemInput

func (ServerItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerItem)(nil)).Elem()
}

func (i ServerItemMap) ToServerItemMapOutput() ServerItemMapOutput {
	return i.ToServerItemMapOutputWithContext(context.Background())
}

func (i ServerItemMap) ToServerItemMapOutputWithContext(ctx context.Context) ServerItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerItemMapOutput)
}

type ServerItemOutput struct{ *pulumi.OutputState }

func (ServerItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerItem)(nil)).Elem()
}

func (o ServerItemOutput) ToServerItemOutput() ServerItemOutput {
	return o
}

func (o ServerItemOutput) ToServerItemOutputWithContext(ctx context.Context) ServerItemOutput {
	return o
}

type ServerItemArrayOutput struct{ *pulumi.OutputState }

func (ServerItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerItem)(nil)).Elem()
}

func (o ServerItemArrayOutput) ToServerItemArrayOutput() ServerItemArrayOutput {
	return o
}

func (o ServerItemArrayOutput) ToServerItemArrayOutputWithContext(ctx context.Context) ServerItemArrayOutput {
	return o
}

func (o ServerItemArrayOutput) Index(i pulumi.IntInput) ServerItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerItem {
		return vs[0].([]*ServerItem)[vs[1].(int)]
	}).(ServerItemOutput)
}

type ServerItemMapOutput struct{ *pulumi.OutputState }

func (ServerItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerItem)(nil)).Elem()
}

func (o ServerItemMapOutput) ToServerItemMapOutput() ServerItemMapOutput {
	return o
}

func (o ServerItemMapOutput) ToServerItemMapOutputWithContext(ctx context.Context) ServerItemMapOutput {
	return o
}

func (o ServerItemMapOutput) MapIndex(k pulumi.StringInput) ServerItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerItem {
		return vs[0].(map[string]*ServerItem)[vs[1].(string)]
	}).(ServerItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerItemInput)(nil)).Elem(), &ServerItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerItemArrayInput)(nil)).Elem(), ServerItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerItemMapInput)(nil)).Elem(), ServerItemMap{})
	pulumi.RegisterOutputType(ServerItemOutput{})
	pulumi.RegisterOutputType(ServerItemGetAttachmentResultOutput{})
	pulumi.RegisterOutputType(ServerItemArrayOutput{})
	pulumi.RegisterOutputType(ServerItemMapOutput{})
}
