module github.com/SteveLasker/dim

require (
	github.com/containerd/containerd v1.2.6
	github.com/deislabs/oras v0.5.0
	github.com/docker/docker v1.14.0-0.20190131205458-8a43b7bb99cd
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/opencontainers/image-spec v1.0.1
	github.com/sirupsen/logrus v1.3.0
	github.com/spf13/cobra v0.0.3
)

replace (
	github.com/docker/docker => github.com/docker/docker v0.0.0-20190131205458-8a43b7bb99cd
	rsc.io/letsencrypt => github.com/dmcgowan/letsencrypt v0.0.0-20160928181947-1847a81d2087
)
