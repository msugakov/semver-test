package main

import (
	"fmt"
	semver "github.com/blang/semver/v4"
)

func sv(vs string) semver.Version {
	v, err := semver.Make(vs)

	if err != nil {
		panic(err)
	}

	return v
}

func compare(a, b string) {
	val := sv(a).Compare(sv(b))

	switch val {
	case -1:
		fmt.Printf("%s < %s\n", a, b)
	case 0:
		fmt.Printf("%s == %s\n", a, b)
	case 1:
		fmt.Printf("%s > %s\n", a, b)
	}
}

func describe(s string) {
	fmt.Printf("\n%s:\n", s)
}

func main() {
	describe("Identical strings are in fact considered equal")
	compare("4.0.1-fast", "4.0.1-fast")

	describe("A larger minor release with the same `-fast` suffix is considered bigger")
	compare("4.0.1-fast", "4.1.0-fast")

	describe("A larger micro number with the same `-fast` suffix is considered bigger")
	compare("4.0.1-fast", "4.0.2-fast")

	describe("Major, minor, and micro values supersede the suffix")
	compare("4.0.1-fast", "4.0.2-aaron")

	describe("When major, micro, and minor are the same, the suffixes determine order")
	compare("4.0.1-fast", "4.0.1-aaron")

	describe("Strings with no suffix are sorted *after* strings with a suffix")
	compare("4.0.1-fast", "4.0.1")

	describe("Confirming major, minor and micro versions all supersede the suffix")
	compare("4.0.1-fast", "4.1.0")

	describe("Switching from fast stream to stable stream of the same minor version won't be considered an upgrade")
	compare("4.1.123-fast", "4.1.0")

	describe("Just confirming that the micro version is compared numerically not lexicographically")
	compare("4.1.123-fast", "4.1.9")

	fmt.Println("----------")

	describe("Fast for the next release is greater than the current stable")
	compare("4.5.0-fast.1", "4.4.0")

	describe("Fast for the next release is less than the next stable")
	compare("4.5.0-fast.1", "4.5.0")

	describe("Curious what the result would be")
	compare("4.5.0-fast", "4.5.0-fast.1")

	describe("Numeric suffixes participate in comparison")
	compare("4.5.0-fast.1", "4.5.0-fast.2")

	describe("Numeric suffixes get compared as numbers but not lexicographically")
	compare("4.5.0-fast.16", "4.5.0-fast.2")
}
