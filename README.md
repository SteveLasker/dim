# docs-in-markdown
An example for using OCI Artifact Registries for storing new artifact types

## Samples
```powershell
oras push demo42.azurecr.io/samples/docs-in-markdown:v1 `
     --manifest-config ./config.json:application/vnd.stevelasker.docsinmarkdown.config.v1+json `
    ./readme.md:application/vnd.stevelasker.docsinmarkdown.layer.v1+md `
    ./moredetail.md:application/vnd.stevelasker.docsinmarkdown.layer.v1+md
```