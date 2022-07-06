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
This address represents the creator of the subspace. Note that, once a subspace is created, the creator address can never change. However, if you want to transfer the ownership to another user you can use the [owner](#owner) field. 

### Creation time
The creation time of a subspace represents the block time at which the subspace was created. Note that this cannot be set externally nor edited, and is assigned automatically when a `MsgCreateSubspace` is handled.

## Section
A section can be seen as a folder within a subspace. It can be useful to represent a category or a topic, or to build
forum-like social networks, allowing a better way to manage and categorize contents. Just like folders, you can also create nested sections.
By default, each subspace has a root section with id `0`.

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