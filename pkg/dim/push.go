package dim

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/deislabs/oras/pkg/content"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

//LoadDirectory loads the content store with files in the directory
func LoadDirectory(store *content.FileStore, dir string) ([]ocispec.Descriptor, error) {
	var files []ocispec.Descriptor

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fileRef := range fileInfos {
		filename := path.Join("./", dir, fileRef.Name())
		name := filepath.ToSlash(filename)
		var mediaType string

		if strings.ToLower(filepath.Ext(filename)) == ".md" {
			mediaType = "vnd/application.stevelasker.docsinmarkdown.layer.v1+md"
		}

		fmt.Printf("Adding -> %s\n", filename)

		file, err := store.Add(name, mediaType, filename)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}
	return files, nil
}
