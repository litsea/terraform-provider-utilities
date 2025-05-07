// Copyright (c) Litsea
// SPDX-License-Identifier: MIT

package provider

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type fileResource struct{}

func NewFileResource() resource.Resource {
	return &fileResource{}
}

func (r *fileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "filedownloader_file"
}

func (r *fileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Resource to download a remote file via HTTP(S) using GET or POST, optionally with custom headers.",
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				Required:    true,
				Description: "URL to download the file from.",
			},
			"path": schema.StringAttribute{
				Required:    true,
				Description: "Local path to save the downloaded file.",
			},
			"method": schema.StringAttribute{
				Optional:    true,
				Description: "HTTP method to use (default: GET). Only GET and POST are allowed.",
				Validators: []validator.String{
					stringvalidator.OneOf("GET", "POST"),
				},
			},
			"headers": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom headers to include in the request.",
			},
		},
	}
}

func (r *fileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan fileResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	method := "GET"
	if !plan.Method.IsNull() && plan.Method.ValueString() != "" {
		method = strings.ToUpper(plan.Method.ValueString())
	}

	headers := make(map[string]string)
	for k, v := range plan.Headers.Elements() {
		if strVal, ok := v.(types.String); ok {
			headers[k] = strVal.ValueString()
		}
	}

	err := downloadFile(method, plan.URL.ValueString(), plan.Path.ValueString(), headers)
	if err != nil {
		resp.Diagnostics.AddError("Download Failed", err.Error())
		return
	}
}

func (r *fileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No-op: not reading remote status
}

func (r *fileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan fileResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	method := "GET"
	if !plan.Method.IsNull() && plan.Method.ValueString() != "" {
		method = strings.ToUpper(plan.Method.ValueString())
	}

	headers := make(map[string]string)
	for k, v := range plan.Headers.Elements() {
		if strVal, ok := v.(types.String); ok {
			headers[k] = strVal.ValueString()
		}
	}

	err := downloadFile(method, plan.URL.ValueString(), plan.Path.ValueString(), headers)
	if err != nil {
		resp.Diagnostics.AddError("Download Failed", err.Error())
		return
	}
}

func (r *fileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Optional: delete the downloaded file if needed
}

type fileResourceModel struct {
	URL     types.String `tfsdk:"url"`
	Path    types.String `tfsdk:"path"`
	Method  types.String `tfsdk:"method"`
	Headers types.Map    `tfsdk:"headers"`
}

func downloadFile(method, url, filepath string, headers map[string]string) error {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("failed to download file: " + resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
