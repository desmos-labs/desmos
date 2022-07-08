---
id: permissions
title: Permissions
sidebar_label: Permissions
slug: permissions
---

# Permissions

Based on which `x/subspaces` related action your users want to perform, they need to be granted with one or more
of the following permissions.

| **Permission Value** | **Permission Description**                                          | 
|:---------------------|:--------------------------------------------------------------------|
| `EDIT_SUBSPACE`      | Allows to change the subspace's information                         |
| `DELETE_SUBSPACE`    | Allows to delete a subspace                                         |
| `MANAGE_SECTIONS`    | Allows to manage the subspace's sections                            |
| `MANAGE_GROUPS`      | Allows to manage the subspace's groups                              |
| `SET_PERMISSIONS`    | Allows to set other users' permissions except for `SET_PERMISSIONS` |
| `EVERYTHING`         | Allows to do everything                                             |

> **Warning**
> Note that when setting permission `EVERYTHING` to a user, that user will de facto be the same as the subspace owner, having control over everything and being able to do everything within that subspace. Use this with caution.