---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts 

## Reaction
The reaction structure contains the data that identify a post's reaction.
Reactions can be registered by users inside their subspaces in order to later use
them to react contents. They can be text values or more classic kind of reactions (i.e. emojis).

### Subspace ID
The [subspace] ID indicates the ID of the Dapp where the reaction has been made.

### Post ID
The [post] ID indicates the ID of the posts to which the reaction is associated.

### ID
The unique ID that identifies the reaction itself. This ID is automatically assigned to the reaction at the moment of its
creation in a sequential way (e.g. if there's 5 reactions in the chain, the one we are creating will have id equal to 6).

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

