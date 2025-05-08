schema_version = 1

project {
  license          = "MIT"
  copyright_holder = "Litsea"
  copyright_year   = 2025

  header_ignore = [
    # examples used within documentation (prose)
    "examples/**",

    # golangci-lint tooling configuration
    ".golangci.yml",

    # goreleaser configuration
    ".goreleaser.yml",
  ]
}