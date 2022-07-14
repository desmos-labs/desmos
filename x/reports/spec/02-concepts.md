---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts 

## Report
A report represents the way that users have to report misbehaving users or bad contents inside Desmos. 
Each report contains the necessary information to identify the target and reason that lead to its creation.

### Subspace ID
The subspace id represents the id of the subspace where the report has been created.

### ID
The most important thing about a report is its id. This is a unique identifier across the subspace that is used to uniquely reference the report itself along with the subspace id. Report ids are assigned automatically during the handling of a `MsgCreateReport`. 

### Reasons IDs
The reasons ids represent the array of the reasons that this report has been created for. Each id references a specific reason registered within the subspace of this report.

### Message (Optional)
A report message can be optionally used to further describe why a user or content has been reported. This can be useful if the reporter wants to leave a message to other users or administrators that will deal with the report itself.

### Reporter
The address of the user that has created the report.

### Target
A report target represents the content that has been reported. This can be either a [UserTarget](#UserTarget) or a [PostTarget](#PostTarget).

### Creation Date
The creation date of a report represents the block time at which the report has been stored on the chain. This cannot be specified externally and is assigned automatically when handling a `MsgCreateReport`.

## User Target
A user target object should be used when reporting a specific user.

### User
The address of the reported user.

## Post Target
A post target should be used when reporting a specific post within the same subspace where the report has been created.

### Post ID
The ID of the reported post.

## Reason
A reason is the structure representing the motivation behind a report.

### Subspace ID
The subspace id of a reason represents the subspace inside which this reason is valid. Since subspaces can have very different Term of Services from one another, each of them should register their reasons independently so that users are limited in why a report can be created. 

### ID
A reason id represents the unique id within the subspace that can be used to uniquely reference the registered reason. This is assigned automatically when handling either a `MsgSupportStandardReason` or `MsgAddReason` message.

### Title
The title of a reason should be used to give users a quick idea about why they might want to select this reason during the report creation process. Good titles should be short and easy to understand (e.g. __Spam__, __Explicit content__, etc). 

### Description (Optional)
A reason description can be optionally used to allow the users to further understand when they should select this reason during the report creation process.
