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
A section id represents the unique id within the subspace that the section has. This is used everywhere else along with the subspace id to uniquely identify this section.

### Parent ID (Optional)
The ID that identifies the parent of a section. This is set by default to `0`, which represents the root section of a subspace, but it can set to any other section's id to create a tree-like sections structure.

### Name
A section name represents the human-readable name of a section. This can be useful for developers to quickly see what sections are about, so that they can easily understand where to put a content if there are multiple ones.

### Description (Optional)
A section description allows to describe more in detail what a section is about. This can be useful to expand on the topic or motivation of a section.

## User Group
A User group is a set of users who share the same permissions inside a particular subspace or subspace section.

### Subspace ID
The ID of the subspace where the group exists.

### Section ID
The ID of the section where the group exists.

### ID
A user group id uniquely identifies this group within the subspace itself. This, along with the subspace id itself, is used to uniquely reference this group within Desmos.

### Name
The user group name represents the human-readable name of the group. In most cases, this is going to be a short name that makes it possible to easily understand who the users within the group are (e.g. __admins__, __moderators__, etc).

### Description (Optional)
An optional description of the group.

### Permissions
A user group permissions represent the set of permissions that are granted to all the users that are part of the group itself.

## User Permission
A user permission can be used to set an individual user's permission withing a subspace or subspace section. This is particularly useful when you want to set some kind of permission to a single wallet without having to create a user group to do so (e.g. you have a smart contract that needs to have a custom permission set). 

### Subspace ID
The ID of the subspace where the permission has been set.

### Section ID
The ID of the section where the permission has been set. Note that permissions set within parent sections are automatically applied to children sections as well. This means that if you set a user to have a specific permission inside the root section, such permission will be valid throught the entire subspace.

### User
The address of the user to which the permissions have been set.

### Permissions
The array of permissions granted to the user.