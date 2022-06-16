<!--
order: 2
-->

# State

## Subspace

* Next Subspace ID: `0x00 -> Next Subspace ID`
* Subspace: `0x01 | ID | -> ProtocolBuffer(Subspace)`

## Section

* Next Section ID: `0x06 | Subspace ID -> Next Section ID`
* Section: `0x07 | Subspace ID | Section ID -> ProtocolBuffer(Section)`

## User Group

* Next Group ID: `0x02 | Subspace ID -> Next User Group ID`
* User Group: `0x03 | Subspace ID | Section ID | User Group ID -> ProtocolBuffer(User Group)`
* User Group Member: `0x04 | Subspace ID | User Group ID | Address -> 0x01`

## User Permission

* User Permission: `0x05 | Subspace ID | Section ID | Address -> ProtocolBuffer(User Permission)`