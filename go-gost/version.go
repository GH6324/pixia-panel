package main

var (
	version = "0.3.8"
)

func normalizedVersion() string {
	v := version
	if v == "" {
		return "0.3.8"
	}
	return v
}
