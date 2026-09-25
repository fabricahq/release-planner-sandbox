---
name: release
description: Prepare a release pull request with drafted release notes. Use when asked to release, cut or issue a release, draft, revise, or correct release notes, or retry a failed release.
---
<!-- release-planner:generated 1b1356c58b4f162f43bcfe9b4d0c2ee589a9ced6 sha256:8fe0c720b9e4eaef. Do not edit; change .release-planner/config.yml and run release-planner install. -->

# Prepare a release

This repository publishes releases with [Release Planner](https://github.com/fabricahq/release-planner) 1b1356c58b4f162f43bcfe9b4d0c2ee589a9ced6. Print its release procedure and follow it:

```sh
release-planner guide
```

First check that `release-planner version` prints `1b1356c58b4f162f43bcfe9b4d0c2ee589a9ced6`. If it doesn't, or `release-planner` isn't installed, install that version:

```sh
go install github.com/fabricahq/release-planner/cmd/release-planner@1b1356c58b4f162f43bcfe9b4d0c2ee589a9ced6
```

Read `.release-planner/policy.md` first for this repository's release policy. You prepare the release pull request; the maintainer approves the release by merging it. Never tag, publish, or merge.
