<!-- release-planner:begin d62a73b31ac4dde988dcedc7b96de8a9ec610c9a sha256:edf586bc8f92a11c -->
## Releases

This repository publishes releases with [Release Planner](https://github.com/fabricahq/release-planner) d62a73b31ac4dde988dcedc7b96de8a9ec610c9a. When asked to make a release, draft, revise, or correct release notes, or retry a failed release, print the release procedure and follow it:

```sh
release-planner guide
```

First check that `release-planner version` prints `d62a73b31ac4dde988dcedc7b96de8a9ec610c9a`. If it doesn't, or `release-planner` isn't installed, install that version:

```sh
go install github.com/fabricahq/release-planner/cmd/release-planner@d62a73b31ac4dde988dcedc7b96de8a9ec610c9a
```

Read `.release-planner/policy.md` first for this repository's release policy. You prepare the release pull request; the maintainer approves the release by merging it. Never tag, publish, or merge.
<!-- release-planner:end -->
