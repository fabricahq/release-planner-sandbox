---
name: release
description: Prepare a release pull request with drafted release notes. Use when asked to release, cut or issue a release, draft, revise, or correct release notes, or retry a failed release.
---
<!-- release-planner:generated a5f7cbffcb38b2b8162d81bb420c06e363f4bd53 sha256:a2db2bc8bb03e383. Do not edit; change .release-planner/config.yml and run release-planner install. -->

# Prepare a release

This repository publishes releases with [Release Planner](https://github.com/fabricahq/release-planner) a5f7cbffcb38b2b8162d81bb420c06e363f4bd53. Print its release procedure and follow it:

```sh
release-planner guide
```

First check that `release-planner version` prints `a5f7cbffcb38b2b8162d81bb420c06e363f4bd53`. If it doesn't, or `release-planner` isn't installed, install that version:

```sh
go install github.com/fabricahq/release-planner/cmd/release-planner@a5f7cbffcb38b2b8162d81bb420c06e363f4bd53
```

Read `.release-planner/policy.md` first for this repository's release policy. You prepare the release pull request; the maintainer approves the release by merging it. Never tag, publish, or merge.
