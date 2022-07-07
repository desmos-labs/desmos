---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## Next Subspace ID
The next subspace id is stored on the chain as follows:

* `0x00 | -> bytes(NextSubspaceID)`

## Subspace
A subspace is stored on the chain by using its unique id as the key:

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
The next group id is stored using the subspace id to which it is associated as the key:

* Next Group ID: `0x02 | Subspace ID | -> bytes(NextUserGroupID)`

## User Group
A user group is stored on the chain with a combination of subspace id, section id and user group id as key. This make it easier to query:
- all the user groups of a subspace;
- all the user groups of a section;
- a specific user group.

* User Group: `0x03 | Subspace ID | Section ID | User Group ID -> ProtocolBuffer(UserGroup)`

## User Group Member
A user group member is stored on the chain with a combination of subspace id and user group id as key:

* User Group Member: `0x04 | Subspace ID | User Group ID | Address | -> 0x01`

## User Permission
A user permission is stored on the chain with a combination of subspace id, section id and user address as key. This make it easy to query:
- all the user permissions set within a subspace;
- all the user permissions set within a a section;
- all the permissions set to an address.

* User Permission: `0x05 | Subspace ID | Section ID | Address | -> ProtocolBuffer(UserPermission)`