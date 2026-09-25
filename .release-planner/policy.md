# Release policy

Release Planner's agent reads this file before every release.

## Breaking changes

A change is breaking when it changes any of these in a way that forces users to change how they run the tool:

- Command names, flags, and exit codes
- Output that other tools parse
- Supported platforms

## Choosing a version

Versions follow [SemVer 2.0.0](https://semver.org/). The first release is v0.1.0.

Before 1.0.0:

- Minor: any breaking change or new feature
- Patch: bug fixes, documentation, and internal changes

## Who reads the release notes

Readers:

- People who run `greet` from the command line

## Order of the release notes

Start with the first section that has something to say. The notes present changes in this order, leaving out any with nothing to say:

1. New features
2. Improvements
3. Bug fixes
4. Breaking changes

## Always and never

- Always credit external contributors by GitHub handle.
- Never mention dependency updates unless they fix a security issue.
