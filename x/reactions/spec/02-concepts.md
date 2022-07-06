---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts 

## Reaction
Within each subspace, users can react to posts by using a _reaction_. This contains all the data referring to the post to which the reaction should be added, along with the reaction value itself. Reactions can be used to represent likes as well as any other type of response the subspace owners have decided to support. 

### Subspace ID
The subspace id represents the id of the subspace containing the post to which this reaction is associated.

### Post ID
The post id represents the id of post to which the reaction is associated.

### ID
A reaction id is a unique id within a post that can be used along with the post id itself to uniquely identify a reaction.

### Value
The value of the reaction. It can be a registered reaction or a text reaction.

### Author
The address of the author of the reaction.

## Registered Reaction Value
The registered reaction value contains the details of a reaction value that references
a reaction registered within the subspace.

### Registered reaction ID
The ID of the registered reaction the value refers to.

## Free Text Value
The free text value contains the details of a reaction value that is made of free text.

### Text
The actual value of the reaction.

## Registered reaction
The registered reaction structure contains the details of a user's registered reaction within
a subspace.

### Subspace ID
The ID of the subspaces where the reaction has been registered.

### ID 
The unique ID that identifies the registered reaction itself. This ID is automatically assigned to the reaction at the moment of its
creation in a sequential way.

### Shorthand code
The unique shorthand code associated to this reaction. (i.e :smile:)
[Here](https://emojipedia.org/shortcodes/) you can read the Emoji's shorthands.

### Display value
The value that should be displayed when using the reaction.

## Subspace Reactions Params
The params contains all the reactions details for a specific subspace, such as registered reactions,
free text reactions params.

### Subspace ID
The ID of the subspace for which the parameters are valid.

### Registered Reaction
The parameters associated with the registered reaction in the subspace.

### Free Text 
The parameters associated with the free text value reactions in the subspace.

## Registered Reactions Params
The parameters of registered reactions within the subspace.

### Enabled
This parameter tells if registered reactions are enabled or not.

## Free Text Value Params
The parameters of free text value based reactions.

### Enabled
This parameter tells if free text values are enabled or not.

### Max Length
The max length that a free text value reaction can have.

### RegEx
The regular expression that should be used to validate the free text value reaction.
It can be useful to limit characters to a certain group.

