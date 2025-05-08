

resource "filedownloader_file" "example" {
  url      = "https://example.com/file.zip"
  filename = "${path.module}/file.zip"

  headers = {
    Authorization = "Bearer token"
  }
}
