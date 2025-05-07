// Copyright (c) Litsea
// SPDX-License-Identifier: MIT

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func New() provider.Provider {
	return &fileDownloaderProvider{}
}

var _ provider.Provider = (*fileDownloaderProvider)(nil)

type fileDownloaderProvider struct{}

func (p *fileDownloaderProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "filedownloader"
}

func (p *fileDownloaderProvider) Schema(_ context.Context, _ provider.SchemaRequest, _ *provider.SchemaResponse) {
}

func (p *fileDownloaderProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (p *fileDownloaderProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewFileResource,
	}
}

func (p *fileDownloaderProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}
