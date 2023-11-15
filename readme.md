# Semver test

I want to document the functionality of the semver library.
This is a simple script that compares a series of version strings to help document how it behaves.

Here is the output of the script:

Identical strings are in fact considered equal:
`4.0.1-fast == 4.0.1-fast`

A larger minor release with the same `-fast` suffix is considered bigger:
`4.0.1-fast < 4.1.0-fast`

A larger micro number with the same `-fast` suffix is considered bigger:
`4.0.1-fast < 4.0.2-fast`

Major, minor, and micro values supersede the suffix:
`4.0.1-fast < 4.0.2-aaron`

When major, micro, and minor are the same, the suffixes determine order:
`4.0.1-fast > 4.0.1-aaron`

Strings with no suffix are sorted *after* strings with a suffix:
`4.0.1-fast < 4.0.1`

Confirming major, minor and micro versions all supersede the suffix:
`4.0.1-fast < 4.1.0`

Switching from fast stream to stable stream of the same minor version won't be considered an upgrade:
`4.1.123-fast > 4.1.0`

Just confirming that the micro version is compared numerically not lexicographically:
`4.1.123-fast > 4.1.9`
