package main

var (
	version = "0.3.5"
)

func normalizedVersion() string {
	v := version
	if v == "" {
		return "0.3.5"
	}
	return v
}
