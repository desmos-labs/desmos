---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts

## Subspace
A subspace is a structure representing a specific dApp within the Desmos ecosystem where content can be created.
Each subspace can have its own [sections](#section), [user groups](#user-group) and [permissions](#user-permission).

### ID
The most important part of a subspace it's its own id. This represents the **unique identifier** that is assigned to the subspace during its creation. 

### Name
The subspace name allows the subspace creator and owner to set a human-readable name so that users and developers can easily understand what a subspace is used for. In most cases, when you create a subspace you want to set its name to the same name of the application that will later store its content there. 

### Description (Optional)
An optional description of what the subspace is about.

### Treasury
The treasury address represents the wallet of the subspace itself. This can be used for different reasons such as verifying a subspace with external applications (to prove its authenticity), or paying for fees when executing some smart contracts. 

### Owner
The subspace owner represents the wallet that owns the subspace. As the owner, this wallet will have the `EVERYTHING` permission always set, which allows it to perform any kind of operation within this subspace.

### Creator
The address of the creator of the subspace.

### Creation time
The creation time of the subspace.

## Section
A Section is a zone within a subspace. Each section can represent a category or a topic, and they can be useful to build
forum-like social networks, allowing a deeper way to manage and categorize contents. Sections can also be nested.
Furthermore, each subspace must have a root section (identified by the `0` ID) which is the highest section. 

### Subspace ID
The ID of the subspace where the section exists.

### ID
The unique ID identifying the section. This ID is automatically assigned to the section at the moment of its
creation in a sequential way.

### Parent ID (Optional)
The ID that identifies the parent section of a section.

### Name
The human-readable name of a section.

### Description (Optional)
An optional description of the section topic.

## User Group
A User group is a sub set of users who share the same permissions level inside a particular subspace or section.

### Subspace ID
The ID of the subspace where the group exists.

### Section ID
The ID of the section where the group exists.

### ID
The unique ID identifying the group. This ID is automatically assigned to the section at the moment of its creation in a
sequential way.

### Name
The human-readable name of the group.

### Description (Optional)
An optional description of the group.

### Permissions
The array of permissions granted to all the users part of the group.

## User Permission
The user permission represent a user's permissions inside a specific subspace or section.

### Subspace ID
The ID of the subspace where the user has the permissions.

### Section ID
The ID of the section the user has the permissions.

### User
The address of the user who has the permissions.

### Permissions
The array of permissions granted to the user.