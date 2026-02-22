package main

var (
	version = "0.3.3"
)

func normalizedVersion() string {
	v := version
	if v == "" {
		return "0.3.3"
	}
	return v
}
