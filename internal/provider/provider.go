// Copyright (c) Litsea
// SPDX-License-Identifier: MIT

package provider

import (
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"

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

type fileChecksums struct {
	sha1Hex   string
	sha256Hex string
}

func genFileChecksums(data []byte) *fileChecksums {
	sha1Sum := sha1.Sum(data)
	sha256Sum := sha256.Sum256(data)

	return &fileChecksums{
		sha1Hex:   hex.EncodeToString(sha1Sum[:]),
		sha256Hex: hex.EncodeToString(sha256Sum[:]),
	}
}
