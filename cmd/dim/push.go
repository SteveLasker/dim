package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/deislabs/oras/pkg/content"
	ctxo "github.com/deislabs/oras/pkg/context"
	"github.com/deislabs/oras/pkg/oras"

	"github.com/SteveLasker/dim/pkg/dim"
	"github.com/containerd/containerd/images"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type pushOptions struct {
	targetRef string
	fileRefs  []string

	verbose bool

	debug    bool
	username string
	password string
}

func pushCmd() *cobra.Command {
	var opts pushOptions
	cmd := &cobra.Command{
		Use:   "push name[:tag|@digest] file[:type] [file...]",
		Short: "Push files to remote registry",
		Long: `Push files to remote registry

Example - Push directory called docs with markdown and other content.
  dim push localhost:5000/hello:latest ./docs/
`,
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.targetRef = args[0]
			opts.fileRefs = args[1:]
			return runPush(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.verbose, "verbose", "v", false, "verbose output")
	cmd.Flags().BoolVarP(&opts.debug, "debug", "d", false, "debug mode")
	cmd.Flags().StringVarP(&opts.username, "username", "u", "", "registry username")
	cmd.Flags().StringVarP(&opts.password, "password", "p", "", "registry password")
	return cmd
}

func runPush(opts pushOptions) error {
	ctx := context.Background()
	if opts.debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else if !opts.verbose {
		ctx = ctxo.WithLoggerDiscarded(ctx)
	}

	// load files
	var (
		annotations map[string]map[string]string
		store       = content.NewFileStore("")
		pushOpts    []oras.PushOpt
	)
	defer store.Close()

	files, err := loadFiles(store, annotations, &opts)
	if err != nil {
		return err
	}

	// ready to push
	resolver := newResolver(opts.username, opts.password)
	pushOpts = append(pushOpts, oras.WithPushBaseHandler(pushStatusTrack()))
	desc, err := oras.Push(ctx, resolver, opts.targetRef, store, files, pushOpts...)
	if err != nil {
		return err
	}

	fmt.Println("Pushed", opts.targetRef)
	fmt.Println("Digest:", desc.Digest)

	return nil
}

func loadFiles(store *content.FileStore, annotations map[string]map[string]string, opts *pushOptions) ([]ocispec.Descriptor, error) {
	var files []ocispec.Descriptor
	for _, fileRef := range opts.fileRefs {
		filename, _ := parseFileRef(fileRef, "")
		name := filepath.Clean(filename)
		if !filepath.IsAbs(name) {
			// convert to slash-separated path unless it is absolute path
			name = filepath.ToSlash(name)
		}
		if opts.verbose {
			fmt.Println("Preparing", name)
		}
		fileInfo, err := os.Stat(filename)
		if err != nil {
			return nil, err
		}
		if fileInfo.IsDir() {
			return dim.LoadDirectory(store, filename)
		}
	}
	return files, nil
}

func pushStatusTrack() images.Handler {
	var printLock sync.Mutex
	return images.HandlerFunc(func(ctx context.Context, desc ocispec.Descriptor) ([]ocispec.Descriptor, error) {
		if name, ok := content.ResolveName(desc); ok {
			printLock.Lock()
			defer printLock.Unlock()
			fmt.Println("Uploading", desc.Digest.Encoded()[:12], name)
		}
		return nil, nil
	})
}
