<!--
order: 1
-->

# Concepts

## Subspace
A Subspace is a structure representing a specific zone inside Desmos where a social network live on.
Each subspace can have its own term of services, tokenomics, sections and user groups with different permissions.

### ID
The unique ID identifying a Subspace. This ID is automatically assigned to the subspace at the moment of its
creation in a sequential way (e.g. if there's 2 subspace in the chain, the one we are creating will have id equal to 3).

### Name
The human-readable name of the subspace. This name will most likely be the same of the social app built on the top of it.

### Description (Optional)
An optional description of what the subspace is about.

### Treasury
The treasury is an address owned by the subspaces that should be used to connect it to external applications
in order to verify it.

### Owner
The address of the user (or smart contract) that owns the subspace. 

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
The unique ID identifying a Section. This ID is automatically assigned to the Section at the moment of its
creation in a sequential way.

### Parent ID (Optional)
The ID that identifies the parent section of a section.

### Name
The human-readable name of a section.

### Description (Optional)
An optional description of the section topic.

## User Group
User group is a set of users who have the same access in a subspace/section.

## Permissions
Permission is defined as the right insinde a subspace/section.