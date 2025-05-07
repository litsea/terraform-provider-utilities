

resource "filedownloader_file" "example" {
  url  = "https://example.com/file.zip"
  path = "${path.module}/file.zip"

  headers = {
    Authorization = "Bearer token"
  }
}
