---
id: halting
title: Halting
sidebar_position: 4
---

# Halting Your Validator
When attempting to perform routine maintenance or planning for an upcoming coordinated upgrade, it can be useful to have
your validator systematically and gracefully halt. You can achieve this by either setting the `halt-height` to the
height at which you want your node to shutdown or by passing the `--halt-height` flag to `desmos`. The node will
shutdown with a zero exit code at that given height after committing the block.
