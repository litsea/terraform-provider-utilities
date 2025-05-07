# terraform-provider-filedownloader

A simple Terraform provider to download files from a given URL with optional HTTP headers.

## Example Usage

```hcl
terraform {
  required_providers {
    filedownloader = {
      source  = "litsea/filedownloader"
      version = "0.1.1"
    }
  }
}

provider "filedownloader" {}

resource "filedownloader_file" "example" {
  url     = "https://example.com/file.zip"
  path    = "${path.module}/file.zip"
  headers = {
    Authorization = "Bearer <token>"
  }
}
```
