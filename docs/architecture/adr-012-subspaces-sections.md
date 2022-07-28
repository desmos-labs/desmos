# ADR 012: Subspaces sections

## Changelog

- May 06th, 2022: Initial draft;
- May 09th, 2022: Fixed some typos and improved Proto file definitions;

## Status

ACCEPTED Implemented

## Abstract

This ADR contains the specific of the new subspaces section, which will allow users to better manage content inside their subspaces.

## Context

Currently, the `x/subspaces` module allows to create simple subspaces that are suited to represent applications that display contents inside a single section. An example of this are social applications such as Twitter, where each tweet is put in the same bucket and is not categorized. 

Although this allows to create a lot of applications, we are not allowing to easily create forum-like applications where contents are organized in sections, subsections, and so on. This currently requires the developers to create and maintain a very large number of subspaces.

Consider, as an example, a forum with 3 sections each one having 3 subsections. This would require the admins to create and manage 9 (one per each subsection) subspaces at the same time. Although technically possible, it would require a lot of duplicated work in some occasions: when a moderator gets elected, for example, it would require the original subspace owner to add it to all 9 subspaces. 

This obviously gets more and more complicated very quickly as this kind of applications scale: each time a section (or subsection) needs to be added, a new subspace must be created and managed, making it more complex to operate the whole system.

## Decision

In order to solve the above problems, we will implement a new concept inside the `x/subspaces` module: **sections**. 

Each section will represent a portion of a subspace that can have its own user groups and host its own contents. Also, it will be possible to nest sections into one another thanks to parent-children relationships.

```
Parent:                    Section P
                  /            |           \
Children:      Section C1  Section C2  Section C3
```

Permissions inside sections will be inherited from the parent section to all the children sections. Permissions defined in the parent section will have the same permissions also on children sections, but permissions defined in children sections will be valid only there.

All subspaces will have a default section having id `0` which identifies the subspace itself. This will be used to define user groups that should be present in all (sub)sections and to post content directly inside the subspace itself.

To manage different sections, we will also introduce a new permission inside the `x/subspaces` module: `PermissionManageSections` that will allow users to create, edit and delete a subspace sections as they please.

### Types

#### Section
```protobuf
syntax = "proto3";

// Section contains the data of a single subspace section
message Section {
  // Unique id of the section within the subspace
  uint32 id = 1;
  
  // (optional) Id of the parent section
  uint32 parent_id = 2;
  
  // Name of the section within the subspace
   string name = 3; 
  
  // (optional) Description of the section
  string description = 4;
}
```

### `Msg` Service
We will allow the following operations: 
- create a new section 
- edit an existing section
- move a section to another parent
- delete an existing section


```protobuf
syntax = "proto3";

service Msg {
  // CreateSection allows to create a new subspace section
  rpc CreateSection(MsgCreateSection) returns (MsgCreateSectionResponse);
  
  // EditSection allows to edit an existing section
  rpc EditSection(MsgEditSection) returns (MsgEditSectionResponse);
  
  // MoveSection allows to move an existing section to another parent
  rpc MoveSection(MsgMoveSection) returns (MsgMoveSectionResponse);
  
  // DeleteSection allows to delete an existing section
  rpc DeleteSection(MsgDeleteSection) returns (MsgDeleteSectionResponse);
}

// MsgCreateSection represents the message to be used when creating a subspace section
message MsgCreateSection {
  // Id of the subspace inside which the section will be placed
  uint64 subspace_id = 1;
  
  // Name of the section to be created
  string name = 2;
  
  // (optional) Description of the section 
  string description = 3;
  
  // (optional) Id of the parent section 
  uint32 parent_id = 4;
  
  // User creating the section
  string creator = 5;
}

// MsgCreateSectionResponse represents the Msg/CreateSection response type
message MsgCreateSectionResponse {
  // Id of the newly created section
  uint32 section_id = 1;
}

// MsgEditSection represents the message to be used when editing a subspace section
message MsgEditSection {
  // Id of the subspace inside which the section to be edited is
  uint64 subspace_id = 1;
  
  // Id of the section to be edited
  uint32 section_id = 2;
  
  // (optional) New name of the section
  string name = 3;
  
  // (optional) New description of the section
  string description = 4;
  
  // User editing the section
  string editor = 5;
}

// MsgEditSectionResponse represents the Msg/EditSection response type
message MsgEditSectionResponse {}

// MsgMoveSection represents the message to be used when moving a section to another parent
message MsgMoveSection { 
  // Id of the subspace inside which the section lies
  uint64 subspace_id = 1;
  
  // Id of the section to be moved 
  uint32 section_id = 2;
  
  // Id of the new parent 
  uint32 new_parent_id = 3;
  
  // Signer of the message
  string signer = 4;
}

// MsgMoveSectionResponse
message MsgMoveSectionResponse {}

// MsgDeleteSection represents the message to be used when deleting a section
message MsgDeleteSection {
  // Id of the subspace inside which the section to be deleted is
  uint64 subspace_id = 1;
  
  // Id of the section to delete
  uint32 section_id = 2;
  
  // User deleting the section
  string signer = 3;
}

// MsgDeleteSectionResponse represents the Msg/DeleteSection response type
message MsgDeleteSectionResponse {}
```

### `Query` Service
```protobuf
syntax = "proto3";

service Query {
  // Sections allows to query for the sections of a specific subspace
  rpc Sections(QuerySectionsRequest) returns (QuerySectionsResponse) {
    option (google.api.http).get = "/desmos/subspaces/v2/{subspace_id}/sections";
  }
}

// QuerySectionsRequest is the request type for Query/Sections RPC method 
message QuerySectionsRequest {
  // Id of the subspace to query the sections for
  uint64 subspace_id = 1;

  // pagination defines an optional pagination for the request.
  cosmos.base.query.v1beta1.PageRequest pagination = 2;
}

// QuerySectionsResponse is the response type for Query/Sections RPC method 
message QuerySectionsResponse {
  repeated Section sections = 1;
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}
```

## Consequences

### Backwards Compatibility

The changes described inside this ADR are **not** backward compatible. The introduction of subspace sections requires changes inside the `x/subspaces` module as well as inside the upcoming `x/posts` module. In particular, permission management and post creation need to be adapted accordingly.

User groups need to be assigned to a specific section instead of a subspace, and users need to be able to assign permissions within a specific section as well. 

Also, posts will need to be put inside a particular section (by default the `0` one), and not a generic subspace only. To do this, when creating a post users will have to specify the section id along with the subspace id.

### Positive

- DApp developers will be able to manage the contents of their apps more easily
- Subspace sections will allow to create forum-like applications, opening up Desmos to even more use cases

### Negative

### Neutral

- Changes are required to both the `x/subspaces` as well as the `x/posts` modules

## References

- Issue [#856](https://github.com/desmos-labs/desmos/issues/856).