---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## Next Subspace ID
The Next Subspace ID is stored on the chain as follows:

* `0x00 | -> bytes(NextSubspaceID)`

## Subspace
The Subspace is stored on the chain with its ID as key:

* Subspace: `0x01 | Subspace ID | -> ProtocolBuffer(Subspace)`

## Next Section ID
The Next Section ID is stored on the chain using its associated subspace ID as key.

* Next Section ID: `0x06 | Subspace ID | -> bytes(NextSectionID)`

## Section
The Section is stored using both the Subspace ID and its ID as keys. This make it easier to query:
- All the subspace related sections;
- A specific section inside a given subspace.

* Section: `0x07 | Subspace ID | Section ID | -> ProtocolBuffer(Section)`

## Next Group ID
The Next Group ID is stored on the chain using its associated subspace ID as key:

* Next Group ID: `0x02 | Subspace ID | -> bytes(NextUserGroupID)`

## User Group
The user group is stored on the chain with a combination of subspace, section and user group IDs as key. This make it easier to query:
- All the user groups of a subspace;
- All the user groups of a section;
- A specific user group.

* User Group: `0x03 | Subspace ID | Section ID | User Group ID -> ProtocolBuffer(UserGroup)`

## User Group Member
The user group member is stored on the chain with a combination of subspace and user group IDs as key.

* User Group Member: `0x04 | Subspace ID | User Group ID | Address | -> 0x01`

## User Permission
The user permission is stored on the chain with a combination of Subspace ID, Section ID and user address as key. This make it easier to query:
- All the User Permissions for a subspace;
- All the User Permissions for a section;
- All the User Permissions for an address.

* User Permission: `0x05 | Subspace ID | Section ID | Address | -> ProtocolBuffer(UserPermission)`