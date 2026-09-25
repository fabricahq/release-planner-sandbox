<!-- release-planner:begin 153dac5742a4a67d2700a1e3564b370d0b19e35f sha256:2c9a78aaea430ab2 -->
## Releases

This repository publishes releases with [Release Planner](https://github.com/fabricahq/release-planner) 153dac5742a4a67d2700a1e3564b370d0b19e35f. When asked to make a release, draft, revise, or correct release notes, or retry a failed release, print the release procedure and follow it:

```sh
release-planner guide
```

First check that `release-planner version` prints `153dac5742a4a67d2700a1e3564b370d0b19e35f`. If it doesn't, or `release-planner` isn't installed, install that version:

```sh
go install github.com/fabricahq/release-planner/cmd/release-planner@153dac5742a4a67d2700a1e3564b370d0b19e35f
```

Read `.release-planner/policy.md` first for this repository's release policy. You prepare the release pull request; the maintainer approves the release by merging it. Never tag, publish, or merge.
<!-- release-planner:end -->
