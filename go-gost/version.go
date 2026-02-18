package main

var (
	version = "0.3.2"
)

func normalizedVersion() string {
	v := version
	if v == "" {
		return "0.3.2"
	}
	return v
}
