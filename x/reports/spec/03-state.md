---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## Next Report ID
The next report ID is stored using the subspace ID where it lives as key. This allows to easily query the ID to be used next
for the newest report created:

* `0x01 | Subspace ID | -> bytes(NextReportID)`

## Report
A report is stored using the subspace ID and its ID combined as its key. This allows to easily query:
- All the reports of a subspace;
- A specific report in a subspace.

* `0x02 | Subspace ID | Report ID | -> ProtocolBuffer(Report)`

## Posts Report
A post report is stored using the combination of subspace, post IDs and reporter address as key. This allows to easily query
all the reports towards a specific post.

* `0x03 | Subspace ID | Post ID | Reporter | -> bytes(ReportID)`

## User Report
A user report is stored using the combination of subspace, post IDs and reporter address as key. This allows to easily query
all the reports towards a specific user.

* `0x04 | Subspace ID | User | Reporter | -> bytes(ReportID)`

## Next Reason ID
The next reason ID is stored using the subspace ID where it lives as key. This allows to easily query the ID to be used next
for the newest reason created:

* `0x10 | Subspace ID | -> bytes(NextReasonID)`

## Reason
A reason is stored using the subspace ID where it lives combined with its own ID as key. This allows to easily query:
- All the reasons of a subspace;
- A specific reason in a subspace.

* `Ox11 | Subspace ID | Reason ID | -> bytes(Reason)`