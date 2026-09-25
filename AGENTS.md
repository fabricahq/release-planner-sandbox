<!-- release-planner:begin 30074f1d3b428bbca1cf8d905888ae7c3d305993 sha256:558c4a2b5b3a3d99 -->
## Releases

This repository publishes releases with [Release Planner](https://github.com/fabricahq/release-planner) 30074f1d3b428bbca1cf8d905888ae7c3d305993. When asked to make a release, draft, revise, or correct release notes, or retry a failed release, print the release procedure and follow it:

```sh
release-planner guide
```

First check that `release-planner version` prints `30074f1d3b428bbca1cf8d905888ae7c3d305993`. If it doesn't, or `release-planner` isn't installed, install that version:

```sh
go install github.com/fabricahq/release-planner/cmd/release-planner@30074f1d3b428bbca1cf8d905888ae7c3d305993
```

Read `.release-planner/policy.md` first for this repository's release policy. You prepare the release pull request; the maintainer approves the release by merging it. Never tag, publish, or merge.
<!-- release-planner:end -->
