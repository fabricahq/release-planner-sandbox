<!-- release-planner:begin a5f7cbffcb38b2b8162d81bb420c06e363f4bd53 sha256:a7847edc53d960fc -->
## Releases

This repository publishes releases with [Release Planner](https://github.com/fabricahq/release-planner) a5f7cbffcb38b2b8162d81bb420c06e363f4bd53. When asked to make a release, draft, revise, or correct release notes, or retry a failed release, print the release procedure and follow it:

```sh
release-planner guide
```

First check that `release-planner version` prints `a5f7cbffcb38b2b8162d81bb420c06e363f4bd53`. If it doesn't, or `release-planner` isn't installed, install that version:

```sh
go install github.com/fabricahq/release-planner/cmd/release-planner@a5f7cbffcb38b2b8162d81bb420c06e363f4bd53
```

Read `.release-planner/policy.md` first for this repository's release policy. You prepare the release pull request; the maintainer approves the release by merging it. Never tag, publish, or merge.
<!-- release-planner:end -->
