---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## Next Report ID
The next report id is stored using the subspace id that it references as the key:

* `0x01 | Subspace ID | -> bytes(NextReportID)`

## Report
A report is stored using the subspace id and its id combined as the key. This allows to easily query:
- all the reports of a subspace;
- a specific report in a subspace.

* `0x02 | Subspace ID | Report ID | -> ProtocolBuffer(Report)`

## Posts Report
A post report is stored using the combination of subspace id, post id and reporter address as the key. This allows to easily query all the reports towards a specific post.

* `0x03 | Subspace ID | Post ID | Reporter | -> bytes(ReportID)`

## User Report
A user report is stored using the combination of subspace id, post id and reporter address as the key. This allows to easily query all the reports towards a specific user.

* `0x04 | Subspace ID | User | Reporter | -> bytes(ReportID)`

## Next Reason ID
The next reason id is stored using the subspace id that it references as the key:

* `0x10 | Subspace ID | -> bytes(NextReasonID)`

## Reason
A reporting reason is stored using the subspace id and its own id as the key. This allows to easily query:
- all the reasons of a subspace;
- a specific reason in a subspace.

* `Ox11 | Subspace ID | Reason ID | -> bytes(Reason)`