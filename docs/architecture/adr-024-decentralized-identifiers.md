# ADR 024: Decentralized Identifiers registry

## Changelog

- Jan 16th, 2024: First draft;

## Status

PROPOSED

## Abstract

This ADR proposes the integration of a new functionality to incorporate Decentralized Identifiers (DIDs) into Desmos.

## Context

Decentralized identifiers (DIDs) are a new type of identifier that enables verifiable, decentralized digital identity. DIDs enable individuals to assert control over their personal information, aligning with principles of privacy, and user empowerment. Furthermore, as DIDs emerge as a universal standard, they are poised to play a pivotal role in identity authentication across various applications. Desmos serves as a decentralized social platform infrastructure, becoming DIDs registry would contribute to an improved user experience in identity management.

## Decision

We will implement a new module named `x/did` that allows users to create/manage their DID documents.

In addition, the name that identifies our DID method is: `desmos`, in other word, a DID that uses this method MUST begin with following prefix: `did:desmos`. 

### Types

#### DID document

`DidDoc` represents the DID document which includes full context with the standard published by W3C DID core.

The store keys of DID document inside KVStore will be as follows:

* 0x1 | {id} | {version_id} -> ProtocolBuffer(DidDoc)

> NOTE  
> To be compatible to DID resolver queries `did:example:123?versionId={version_id}` and `did:example:123?versionTime=2016-10-17T02:41:00Z`, we SHOULD store all the versions of DID document.

##### Protobuf definition

```protobuf
// DidDoc represents a set of data describing DID subject.
// Documentation: https://www.w3.org/TR/did-core
message DidDoc {
    // URIs used to identify the context of DID document.
    // Example: [https://www.w3.org/ns/did/v1]
    repeated string context = 1;

    // The unique identifier in DID syntax of DID document.
    // Example: did:desmos:<unique-identifier> 
    string id = 2;
    
    // External controller in DID syntax who can manage the verification method (optional).
    // Documentation: https://w3id.org/security#controller
    string controller = 3;

    // The assertion that two or more DIDs or other types of UIR refer to the same DID.
    // Documentation: https://www.w3.org/TR/did-core/#also-known-as
    string also_known_as = 4;

    // Keys of verification methods for verifying digital signature, must have at least one associated to Desmos address.
    repeated VerificationMethod verification_methods = 5;

    // Id of keys for authentication within verification methods (optional).
    // Documentation: https://w3id.org/security#authenticationMethod
    repeated string authentication = 6;

    // Id of keys for assertion method within verification methods (optional).
    // Documentation: https://w3id.org/security#assertionMethod
    repeated string assertion_method = 7;

    // Id of keys for capability invocation that can manage the DID documentation within verification methods, must have at least one associated to Desmos address.
    // Documentation: https://w3id.org/security#capabilityInvocationMethod
    repeated string capability_invocation = 8;

    // Id of keys for capability delegation within verification methods (optional).
    // Documentation: https://w3id.org/security#keyAgreementMethod
    repeated string capability_delegation = 9;

    // Id of keys for kee agreement within verification methods (optional).
    // Documentation: https://w3id.org/security#keyAgreementMethod
    repeated string key_agreement = 10;

    // A set of services endpoint maps (optional).
    repeated Service service = 11;

    // Version ID of the current DID document for resolver queries.
    // Documentation: https://www.w3.org/TR/did-core/#did-parameters
    uint64 version_id = 12;

    // Timestamp of the current DID document update for resolver queries.
    // Documentation: https://www.w3.org/TR/did-core/#did-parameters
    google.protobuf.timestamp version_time = 13;

}

// VerificationMethod express the cryptographic public keys, which can be used to authenticate or authorize interaction.
// Documentation: https://www.w3.org/TR/did-core/#verification-methods
message VerificationMethod {
    // Unique identifier in DID URL syntax.
    // Example: did:desmos:<unique-identifier>#<key-id>
    string id = 1;

    // Type of the verification method.
    string type = 2;
    
    // External controller in DID syntax who can manage the verification method (optional).
    // Documentation: https://w3id.org/security#controller
    string controller = 3;

    // Material of the public key in the verification method.
    // Documentation: https://www.w3.org/TR/did-core/#verification-material
    oneof verification_material {
        // JSON web key format public key in JSON.
        string public_key_jwk = 4;
        
        // Base58 encoded public key.
        string public_key_multibase = 5;
    };
}

// Service represents the ways of communicating with DID subjects or associated entities.
// Documentation: https://www.w3.org/TR/did-core/#services
message Service {
    // Unique identifier for the service.
    // Example: did:desmos:<unique-identifier>#<service-name>
    string id = 1;

    // Type of the service.
    string type = 2;

    // Endpoints of service.
    repeated string service_endpoint = 3;
}
```

#### DidMetadata

`DidMetadata` represents the metadata to show the latest status of DID document.

The store keys of `DIDMetadata` for KVStore will be:

* 0x2 | {id} -> ProtocolBuffer(DidMetadata)

##### Protobuf definition

```protobuf
// DidMetadata represents metadata of the DID document.
// Documentation: https://www.w3.org/TR/did-core/#did-document-metadata
message DidMetadata {
    // The unique identifier in DID syntax of DID document.
    string id = 1;

    // Timestamp of the DID document creation.
    google.protobuf.Timestamp created = 2;

    // Timestamp of the last DID document update.
    google.protobuf.Timestamp updated = 3;

    // Flag that indicates DID document is deactivated or not.
    bool deactivated = 4;

    // Last version id of the DID document.
    uint64 version_id = 5;
}
```

### `Msg` Service

We will allow the following operations to be performed:

- Create a new DID document
- Update an existing DID document
- Deactivate DID document
- Edit verification inside DID document

> NOTE  
> When updating DID document, sender's public key MUST be its controller or one of keys inside capability invocation field.

#### Protobuf definition

```protobuf
service Msg{
    // CreateDidDoc allows to create a newly DID document.
    rpc CreateDidDoc(MsgCreateDidDoc) returns(MsgCreateDidDocResponse);

    // UpdateDidDoc allows to update an existing DID document.
    rpc UpdateDidDoc(MsgUpdateDidDoc) returns(MsgUpdateDidDocResponse);

    // DeactivateDidDoc allows to deactivate DID document.
    rpc DeactivateDidDoc(MsgsDeactivateDidDoc) returns(MsgsDeactivateDidDocResponse);

    // EditVerificationMethod allows to edit an existing verification method inside DID document.
    rpc EditVerificationMethod(MsgEditVerificationMethod) returns(MsgEditVerificationMethodResponse);
}

// MsgCreateDidDoc represents the message to be used to create a new DID document.
message MsgCreateDidDoc {
    // URIs used to identify the context of DID document to be set.
    repeated string context = 1;

    // ID of the newly DID document to be set. 
    string id = 2;

    // ID of the newly DID document controller to be set.
    string controller = 3;

    // DIDs or URIs referring to the newly DID document to be set.
    string also_known_as = 4;

    // Keys of the newly DID document for verifying digital signature, must have at least one associated to Desmos address.
    repeated VerificationMethod verification_methods = 5;

    // Verification method for authentication to be set.
    repeated string authentication = 6;

    // Verification method for assertion method to be set.
    repeated string assertion_method = 7;

    // Verification method for capability invocation to be set, must have at least one associated to Desmos address.
    repeated string capability_invocation = 8;

    // Verification method for capability delegation to be set.
    repeated string capability_delegation = 9;

    // Verification method for key agreement to be set.
    repeated string key_agreement = 10;

    // Endpoints of services to be set.
    repeated Service service = 11;

    // Address of the message sender.
    string sender = 12;
}

// MsgCreateDidDocResponse defines the Msg/CreateDidDoc response type.
message MsgCreateDidDocResponse {
    // The metadata of the DID document.
    DidMetadata metadata = 1;
}

// MsgUpdateDidDoc represents the message to be used to update an existing DID document.
message MsgUpdateDidDoc {
    // ID of the DID document to update. 
    string id = 1;

    // URIs used to identify the context of DID document to be set.
    repeated string context = 2;

    // ID of the controller to be set.
    string controller = 3;

    // DIDs or URIs referring to the newly DID document to be set.
    string also_known_as = 4;

    // Keys of the newly DID document for verifying digital signature, must have at least one associated to Desmos address.
    repeated VerificationMethod verification_methods = 5;

    // Verification method for authentication to be set.
    repeated string authentication = 6;

    // Verification method for assertion method to be set.
    repeated string assertion_method = 7;

    // Verification method for capability invocation to be set, must have at least one associated to Desmos address.
    repeated string capability_invocation = 8;

    // Verification method for capability delegation to be set.
    repeated string capability_delegation = 9;

    // Verification method for key agreement to be set.
    repeated string key_agreement = 10;

    // Endpoints of services to be set.
    repeated Service service = 11;

    // Address of the DID document editor.
    string sender = 12;
}

// MsgUpdateDidDocResponse defines the Msg/UpdateDidDoc response type.
message MsgUpdateDidDocResponse {
    // The updated metadata of the DID document.
    DidMetadata metadata = 1;
}

// MsgDeactivateDidDoc represents the message to be used to deactivate a DID document.
message MsgDeactivateDidDoc {
    // ID of the DID document to deactivate. 
    string id = 1;

    // Address of the DID document editor.
    string sender = 2;
}

// MsgDeactivateDidDocResponse defines the Msg/DeactivateDidDoc response type.
message MsgDeactivateDidDocResponse {
    // The updated metadata of the DID document.
    DidMetadata metadata = 1;
}

// MsgEditVerificationMethod represents the message to be used to edit an existing verification inside DID document.
message MsgEditVerificationMethod {
    // Identifier of verification method to edit.
    string id = 1;

    // Key type to be set.
    string type = 2;

    // ID of controller to be set.
    string controller = 3;

    // Material of public key to be set.
    oneof verification_material {
        // JSON web key format public key in JSON.
        string public_key_jwk = 4;
        
        // Base58 encoded public key.
        string public_key_multibase = 5;
    };

    // Address of the DID document editor.
    string sender = 6;
}

// MsgEditVerificationMethodResponse defines the Msg/EditVerificationMethod response type.
message MsgEditVerificationMethodResponse {
    // The updated metadata of the DID document.
    DidMetadata metadata = 1;
}
```

### `Query` service

```protobuf
service Query{
    // DidDoc queries for a single DID document with the specified version. 
    rpc DidDoc(QueryDidDocRequest) returns(QueryDidDocResponse) {
        option (google.http.get) = "/desmos/did/v1/did/{id}";
    };

    // DidMetadata queries for the metadata of the DID document with its ID.
    rpc DidMetadata(QueryDidMetadataRequest) returns(DidMetadataResponse) {
        option (google.http.get) = "/desmos/did/v1/did/{id}/metadata";
    };
}

// QueryDidDocRequest is the request type for the Query/DidDoc RPC method
message QueryDidDocRequest {
    string id = 1;

    // Version ID of DID document to query, returns latest version if version is not provided or equals to 0.
    uint64 version = 2;
}

// QueryDidDocResponse is the response type for the Query/DidDoc RPC method
message QueryDidDocResponse {
    DidDoc did_doc = 1;
}

// QueryDidMetadataRequest is the request type for the Query/DidMetadata RPC method
message QueryDidMetadataRequest {
    string id = 1;
}

// QueryDidMetadataResponse is the response type for the Query/DidMetadata RPC method
message QueryDidMetadataResponse {
    DidMetadata metadata = 1;
}
```

## Consequences

### Backwards Compatibility

The solution outlined above is fully backward compatible since we are just adding a new module.

### Positive

- Improve the experience of identity management inside Desmos ecosystem.

### Negative

- Having DID implementation increases the complexity and storage usage of Desmos core.

### Neutral

(none known)

## Further Discussions

To fully compatible to DID universal resolver, we SHOULD implement resolver driver for Desmos in the future.

## References

- [W3C Decentralized Identifiers (DIDs) v1.0](https://www.w3.org/TR/did-core/)
- [DID universal resolver](https://github.com/decentralized-identity/universal-resolver)