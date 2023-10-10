# ADR 023: Subspace relationships export

## Changelog

- June 16th, 2023: First draft;
- October 10th, 2023: Status update.

## Abstract

This ADR introduces a new feature that enables users export relationships inside a subspace into another subspace.

## Context

Inside Desmos, users can follow each other by creating the so-called ___relationships___. Each relationship represents a user following another user, and two users can be considered ___friends___ if they are following each other.

Currently, relationships are managed at the subspace level. This is meaning that Alice can follow Bob within subspace A, and choose whether to follow Bob or not within subspace B. This approach allows for more freedom and avoids privacy issues that may arise from making all followings public across all subspaces.

However, this system does not grant users full ownership of their follower base. If a user is banned from a subspace, they cannot easily transfer all their relationships to another subspace. This is because subspaces are isolated and relationships cannot be moved between them.

## Decision

To address the issue mentioned above, we propose implementing a new operation to enable users to export their relationships from a subspace to another.
In addition, we will also enables subspace admins manage the relationships import/export from users.

### Permission

We will add new two permissions to allow subspace admins manage relationships import/export from users:

- PermissionExportRelationships: user having this permission can export relationships from this subspace to another.
- PermissionsImportRelationships: user having this permission can import relationships from other subspace to this subspace.

### `Msg` Service

We will implement a new handler that allows users to export their relationships from one subspace to another.

```proto
service Msg {
  ...

  // ExportRelationships allows users to export their relationships from one subspace to another
  rpc ExportRelationships(MsgExportRelationships) returns (MsgExportRelationshipsResponse);
}

// MsgExportRelationships represents a message used to export relationships from one subspace to another
message MsgExportRelationships {
    // Id of the subspace from which the relationships are being exported
    uint64 subspace_id = 1;

    // Id of the subspace to which the relationships are being exported
    uint64 target_subspace_id = 2;

    // Address of user exporting the relationships.
    string exporter = 3;
}

// MsgExportRelationshipsResponse defines the Msg/ExportRelationships response type
message MsgExportRelationshipsResponse {}
```

## Consequences

### Backwards Compatibility

The solution outlined above is fully completely backward compatible since it introduces new a new message handler.

### Positive

- Enable users to export the relationships from one subspace to another

### Negative

(none known)

### Neutral

(none known)

## References
