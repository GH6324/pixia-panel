package main

var (
	version = "0.3.4"
)

func normalizedVersion() string {
	v := version
	if v == "" {
		return "0.3.4"
	}
	return v
}
