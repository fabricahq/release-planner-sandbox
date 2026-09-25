<!-- release-planner:begin 1b1356c58b4f162f43bcfe9b4d0c2ee589a9ced6 sha256:73a422b16f4b53b2 -->
## Releases

This repository publishes releases with [Release Planner](https://github.com/fabricahq/release-planner) 1b1356c58b4f162f43bcfe9b4d0c2ee589a9ced6. When asked to make a release, draft, revise, or correct release notes, or retry a failed release, print the release procedure and follow it:

```sh
release-planner guide
```

First check that `release-planner version` prints `1b1356c58b4f162f43bcfe9b4d0c2ee589a9ced6`. If it doesn't, or `release-planner` isn't installed, install that version:

```sh
go install github.com/fabricahq/release-planner/cmd/release-planner@1b1356c58b4f162f43bcfe9b4d0c2ee589a9ced6
```

Read `.release-planner/policy.md` first for this repository's release policy. You prepare the release pull request; the maintainer approves the release by merging it. Never tag, publish, or merge.
<!-- release-planner:end -->
