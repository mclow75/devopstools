package main

// не устанавливать вручную, использовать go build -ldflags "-X 'main.version=1.0.0'"
var (
	version         = "0.0.0" // major.minor.patch
	longCommitHash  = "N/A"
	shortCommitHash = "N/A"
)

func LongCommitHash() string {
	return longCommitHash
}

func ShortCommitHash() string {
	return shortCommitHash
}

func Version() string {
	return version
}
