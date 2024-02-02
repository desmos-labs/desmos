# ADR 024: Decentralized Identifiers resolver

## Changelog

- Jan 16th, 2024: First draft;
- Jan 23th, 2024: First review;

## Status

ACCEPTED Not Implemented

## Abstract

This ADR proposes the integration of a new functionality to incorporate Decentralized Identifiers (DIDs) into Desmos by resolving Desmos Profile into DID document.

## Context

Decentralized identifiers (DIDs) are a new type of identifier that enables verifiable, decentralized digital identity. Via verifiable credential, DID enables individuals to assert control over their personal information, aligning with principles of privacy, and user empowerment. Furthermore, as DIDs emerge as a universal standard, they are poised to play a pivotal role in identity authentication across various applications. Desmos serves as a decentralized social platform infrastructure, DID would contribute to an improved user experience in identity management.

## Decision

We will implement a query method `DidDoc` to resolve Desmos Profile into DID document. In addition, A DID that uses this method MUST begin with the following prefix: `did:desmos`.

The example of the resolver's response would be like as follows:

```json
{
    "context": [
        "https://www.w3.org/ns/did/v1",
        "https://w3id.org/security/suites/secp256k1-2019/v1"
    ],
    "id": "did:desmos:<desmos-address>",
    "alsoKnownAs": [
        "dtag:<DTag>", // Desmos Dtag
        "application:<application-name>:<id-in-application>", // application link
        "blockchain:<chain-name>:<chain-address>", // chain link
    ],
    "verificationMethod": [
        {
            "id": "did:desmos:<desmos-address>#DESMOS-KEY-1",
            "type": "EcdsaSecp256k1VerificationKey2019",
            "publicKeyMultibase": "<multibase-encoded-public-key>"
        }
    ],
    "authentication": [
        "did:desmos:<desmos-address>#DESMOS-KEY-1"
    ],
    "assertionMethod": [
        "did:desmos:<desmos-address>#DESMOS-KEY-1"
    ],
}
```

### `Query` service

```protobuf
service Query{
    // DidDoc queries for a single DID document. 
    rpc DidDoc(QueryDidDocRequest) returns(QueryDidDocResponse) {
        option (google.http.get) = "/desmos/profiles/v3/did/{id}";
    };
}

// QueryDidDocRequest is the request type for the Query/DidDoc RPC method
message QueryDidDocRequest {
    string id = 1;
}

// QueryDidDocResponse is the response type for the Query/DidDoc RPC method
message QueryDidDocResponse {
    // URIs used to identify the context of DID document.
    // Default: ["https://www.w3.org/ns/did/v1", "https://w3id.org/security/suites/secp256k1-2019/v1"]
    repeated string context = 1;

    // The unique identifier in DID syntax of DID document.
    // Example: did:desmos:<desmos-address> 
    string id = 2;
    
    // The assertion that resources refer to the DID.
    // In Desmos, it shows chain links and application links linked to profile.
    // Documentation: https://www.w3.org/TR/did-core/#also-known-as
    string also_known_as = 3;

    // Keys of verification methods for verifying digital signature.
    // In Desmos, it must be the public key(s) that associated to the profile owner.
    repeated VerificationMethod verification_methods = 4;

    // Id of keys for authentication within verification methods.
    // Documentation: https://www.w3.org/TR/did-core/#authentication
    repeated string authentication = 5;

    // Id of keys for assertion method within verification methods.
    // Documentation: https://www.w3.org/TR/did-core/#assertion
    repeated string assertion_method = 6;
}

// VerificationMethod represents the cryptographic public keys, which can be used to authenticate interaction.
// Documentation: https://www.w3.org/TR/did-core/#verification-methods
message VerificationMethod {
    // Unique identifier in DID URL syntax.
    // Example: did:desmos:<desmos-address>#DESMOS-KEY-1
    string id = 1;

    // Type of the verification method.
    // Example: "EcdsaSecp256k1VerificationKey2019"
    string type = 2;
    
    // Hex-encoded of the public key in the multibase format.
    // Documentation: https://w3c-ccg.github.io/multibase
    string public_key_multibase = 3;
}
```

### Limitation

Due to the necessity of public key(s) being directly controlled by the profile owner, any profile owner lacking public keys, such as those associated with a contract, cannot be resolved into a DID.

## Consequences

### Backwards Compatibility

The solution outlined above is fully backward compatible since we are just adding a new query method.

### Positive

- Enable the usage of Desmos in applications that support DID.

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

To be compatible to DID universal resolver, we SHOULD implement resolver driver for Desmos in the future.

## References

- [W3C Decentralized Identifiers (DIDs) v1.0](https://www.w3.org/TR/did-core/)
- [DID universal resolver](https://github.com/decentralized-identity/universal-resolver)