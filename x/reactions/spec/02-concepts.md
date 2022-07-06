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
The value of the reaction can be either a [registered reaction value](#registered-reaction-value) or a [text reaction value](#text-reaction-value).

### Author
The author of a reaction is the user that has created this reaction.

## Registered Reaction Value
The registered reaction value contains the details of a reaction value that references
a reaction registered within the subspace.

### Registered reaction ID
The id of the registered reaction that should be used as the post's reaction. 

## Free Text Value
The free text value contains the details of a reaction value that is made of free text. This is particularly useful to react to posts using emojis or other text inside a subspace that has not registered any supported reaction.

### Text
The actual value of the reaction.

## Registered reaction
In some cases, subspace owners and admins might want to allow users to only react to posts with a defined set of reactions. This might be the case of dApps that act similarly to Facebook, where you can only use a small set of emojis as reaction. In this case, subspace owners will have to create one _registered reaction_ for each emoji that can be used as a reaction.
At the same time, registered reactions can also be used to customize the set of emojis that can be used within a subspace. For example, you might want to register a reaction with a custom shorthand code that is visualized as a GIF. This is the case for dApps that act like Discord, allowing admins to register custom reactions associated to custom shorthand codes.

### Subspace ID
The id of the subspace inside which the reaction has been registered.

### ID 
Each registered reaction has a unique id within a subspace. This, along with the subspace id itself, is used to uniquely reference a registered reaction while adding a post reaction through the [registered reaction value](#registered-reaction-value) type.

### Shorthand code
A registered reaction shorthand code should be used by users to reference the reaction itself within a text. For this reason, each registered reaction should have a unique shorthand code within a subspace. 

Usually shorthand codes are in the form of `:<code>:` (e.g. the code `:rocket:` is associated to the :rocket: emoji). 

### Display value
The display value of a registered reaction represents the image, emoji, GIF or video that should be visualized instead of the reaction shorthand code. This can be a simple text value (like an emoji) or an URL pointing to the image/GIF/video wanted. 

## Subspace Reactions Params
Each subspace owner can decide what kind of reactions are supported inside their own subspace. The _subspace reactions params_ contains all the related configuration about it. 

### Subspace ID
The id of the subspace for which the parameters are valid.

### Registered Reaction
The parameters related to reactions using a [registered reaction value](#registered-reaction-value).

### Free Text 
The parameters related to reactions using a [free text reaction value](#free-text-reaction-value).

## Registered Reaction Value Params
The _registered reaction value params_ type contains all the parameters related to reactions that use a [registered reaction value](#registered-reaction-value).

### Enabled
Tells whether [registered reaction value](#registered-reaction-value) reactions are supported within the subspace.

## Free Text Value Params
The _free text value params_ type contains all the parameters related to reactions that use a [free text value value](#free-text-value).

### Enabled
Tells whether [free text value](#free-text-value) reactions are supported within the subspace.

### Max Length
The max length that a free text value reaction can have.

### RegEx
The regular expression that should be used to validate the free text value reaction.
It can be useful to limit characters to a certain group.

