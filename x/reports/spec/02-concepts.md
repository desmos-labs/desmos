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
A Reason is the structure representing the motivation behind a report.

### Subspace ID
The [subspace] ID indicates the ID of the Dapp where the reason lives.

### ID
The unique ID that identifies the reason itself. This ID is automatically assigned to the reason at the moment of its
creation in a sequential way, just like the report ID.

### Title
The title of the reason.

### Description (Optional)
AN optional extended description of the reason and why the report has been made.
