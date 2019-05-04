# docs-in-markdown
An example for using OCI Artifact Registries for storing new artifact types

## Samples
## login
```sh
oras login demo42.azurecr.io -u $demo42user -p $demo42pwd
```

```powershell
oras login demo42.azurecr.io -u $env:demo42user -p $env:demo42pwd
```

### Using common medaiTypes
```powershell
oras push demo42.azurecr.io/samples/docs-in-markdown:1 `
    ./readme.md `
    ./moredetail.md
```
```sh
oras push demo42.azurecr.io/samples/docs-in-markdown:1 \
    ./readme.md \
    ./moredetail.md
```
### With Custom medaTypes
```powershell
oras push demo42.azurecr.io/samples/docs-in-markdown:1 `
     --manifest-config ./config.json:application/vnd.stevelasker.docsinmarkdown.config.v1+json `
    ./readme.md:application/vnd.stevelasker.docsinmarkdown.layer.v1+md `
    ./moredetail.md:application/vnd.stevelasker.docsinmarkdown.layer.v1+md
```
```sh
oras push demo42.azurecr.io/samples/docs-in-markdown:1 \
     --manifest-config ./config.json:application/vnd.stevelasker.docsinmarkdown.config.v1+json \
    ./readme.md:application/vnd.stevelasker.docsinmarkdown.layer.v1+md \
    ./moredetail.md:application/vnd.stevelasker.docsinmarkdown.layer.v1+md
```