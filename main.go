package main

import (
	semver "github.com/blang/semver/v4"
)

func main() {
	println("Hello, world!")
	v1, err := semver.Make("1.0.0-beta")
	v2, err := semver.Make("2.0.0-beta")

	if err != nil {
		panic(err)
	}

	println(v1.Compare(v2))
}
