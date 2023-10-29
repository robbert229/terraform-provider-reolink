package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &cameraResource{}
)

// NewCameraResource is a helper function to simplify the provider implementation.
func NewCameraResource() resource.Resource {
	return &cameraResource{}
}

// cameraResource is the resource implementation.
type cameraResource struct{}

// Metadata returns the resource type name.
func (r *cameraResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camera"
}

// Schema defines the schema for the resource.
func (r *cameraResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "a reolink camera",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"address": schema.StringAttribute{
				Description: "The hostname, or ip address of the camera",
				Required:    true,
			},
			"username": schema.StringAttribute{
				Description: "The username to use to login to the camera",
				Required:    true,
			},
			"password": schema.StringAttribute{
				Description: "The password to use to login to the camera",
				Required:    true,
				Sensitive:   true,
			},
		},
	}
}

type cameraResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Address  types.String `tfsdk:"address"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

// Create creates the resource and sets the initial Terraform state.
func (r *cameraResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan cameraResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *cameraResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *cameraResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *cameraResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
