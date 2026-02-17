package main

var (
	version = "0.3.0"
)

func normalizedVersion() string {
	v := version
	if v == "" {
		return "0.3.0"
	}
	return v
}
