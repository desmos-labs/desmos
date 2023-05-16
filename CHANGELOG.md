<!--
All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).
-->
## Version 5.0.0
### Features
#### Fees
- ([\#1135](https://github.com/desmos-labs/desmos/pull/1135)) Removed the `x/fees` module

#### Subspaces
- ([\#1067](https://github.com/desmos-labs/desmos/pull/1067)) Changed how subspace treasury addresses are generated and managed
- ([\#1071](https://github.com/desmos-labs/desmos/pull/1071)) Implemented the ability for subspace admins to grant authorizations on behalf of the subspace treasury

#### Other
- ([\#1059](https://github.com/desmos-labs/desmos/pull/1059)) Added proper support for Ledger devices inside the `sign` command
- ([\#1081](https://github.com/desmos-labs/desmos/pull/1081)) Added the IBC fee middleware
- ([\#1082](https://github.com/desmos-labs/desmos/pull/1082)) Added the InterChain Accounts module
- ([\#1083](https://github.com/desmos-labs/desmos/pull/1083)) Added missing stargate querier accepted list

### Bug Fixes
#### Subspaces
- ([\#1074](https://github.com/desmos-labs/desmos/pull/1074)) Added missing `MsgDeleteSubspace` Amino registration
- ([\#1139](https://github.com/desmos-labs/desmos/pull/1139)) Made it impossible to use the `DefaultUserGroup` to another section

#### Posts
- ([\#1075](https://github.com/desmos-labs/desmos/pull/1075)) Changed `Poll_ProvidedAnswer` attachments type from `Attachment` to `AttachmentContent`

#### Other
- ([\#1069](https://github.com/desmos-labs/desmos/pull/1069)) Registered missing Wasm snapshotter extension

### Dependencies
#### Other
- ([\#1079](https://github.com/desmos-labs/desmos/pull/1079)) Updated `ibc-go` to `v4.3.0`
- ([\#1124](https://github.com/desmos-labs/desmos/pull/1124)) Updated `cosmos-sdk` to `v0.47.2`
- ([\#1133](https://github.com/desmos-labs/desmos/pull/1133)) Updated `wasmd` to `v0.40.0-rc.1`

## Version 4.8.1
### Dependencies
#### Other
- Updated CosmWasm VM to 1.1.2

## Version 4.8.0
### Bug Fixes
#### Reactions
- ([\#1091](https://github.com/desmos-labs/desmos/pull/1091)) Fixed duplicated reaction storing

## Version 4.7.1
### Bug Fixes
#### Other
- ([\#1069](https://github.com/desmos-labs/desmos/pull/1069)) Register missing Wasm snapshotter extension

## Version 4.7.0
### Features
#### Bank
- ([\#1037](https://github.com/desmos-labs/desmos/pull/1037)) Added missing on-chain denom metadata

#### Other
- ([\#1006](https://github.com/desmos-labs/desmos/pull/1006)) Added `v4` upgrade handler

### Dependencies
- ([\#1041](https://github.com/desmos-labs/desmos/pull/1041)) Updated `cosmos-sdk` to `v0.45.11`
- ([\#1043](https://github.com/desmos-labs/desmos/pull/1043)) Updated `ibc-go` to `v3.4.0`

## Version 4.6.3
### Bug Fixes
#### Other
- ([\#1031](https://github.com/desmos-labs/desmos/pull/1031)) Add missing `BaseApp` options when creating a new `DesmosApp` instance

## Version 4.6.2
### Bug Fixes
#### Other
- ([\#1028](https://github.com/desmos-labs/desmos/pull/1028)) Fixed the export command not running properly in some cases

## Version 4.6.1
### Build
- ([\#1026](https://github.com/desmos-labs/desmos/pull/1026)) Updated `wasmvm` to `v1.1.1`

## Version 4.6.0
### Bug Fixes
#### Subspaces
- ([\#1014](https://github.com/desmos-labs/desmos/pull/1014)) Fixed the argument order when editing a subspace

### Dependencies
- ([\#1020](https://github.com/desmos-labs/desmos/pull/1020)) Updated Cosmos SDK to `v0.45.9`

## Version 4.5.0
### Bug Fixes
#### Relationships
- ([\#1010](https://github.com/desmos-labs/desmos/pull/1010)) Added full support for `MsgCreateRelationship` when used with `GenericSubspaceAuthorization`

#### Subspaces
- ([\#1008](https://github.com/desmos-labs/desmos/pull/1008)) Fixed outdated CLI examples

## Version 4.4.1
### Features
- ([\#1002](https://github.com/desmos-labs/desmos/pull/1002)) Added missing on-chain upgrade handler

## Version 4.4.0
### Features
#### Posts
- ([\#998](https://github.com/desmos-labs/desmos/pull/998)) Allow to set the permission to only comment content

### Bug Fixes
#### Profiles
- ([\#986](https://github.com/desmos-labs/desmos/pull/986)) Added missing WASM support for default external address messages

### Dependencies
- ([\#988](https://github.com/desmos-labs/desmos/pull/988)) Updated wasmd to `v0.28.0`
- ([\#1000](https://github.com/desmos-labs/desmos/pull/1000)) Updated Cosmos SDK to `v0.45.8`

## Version 4.3.0
### Bug Fixes
#### Profiles
- ([\#980](https://github.com/desmos-labs/desmos/pull/980)) Added missing params migration

#### Subspaces
- ([\#981](https://github.com/desmos-labs/desmos/pull/981)) Added missing subspaces authorizations migration

#### Posts
- ([\#978](https://github.com/desmos-labs/desmos/pull/978)) Fixed the REST endpoints version

## Version 4.2.0
### Features
#### Profiles
- ([\#723](https://github.com/desmos-labs/desmos/pull/723)) Added expiration time to application links
- ([\#887](https://github.com/desmos-labs/desmos/pull/887)) Added the ability to specify per-chain default external links
- ([\#968](https://github.com/desmos-labs/desmos/pull/968)) Added the support for EVM-specific chain link signatures

#### Subspaces
- ([\#974](https://github.com/desmos-labs/desmos/pull/974)) Allow setting initial group members when creating a user group

### Bug Fixes
#### Profiles
- ([\#971](https://github.com/desmos-labs/desmos/pull/971)) `QueryProfile` now returns an error instead of `nil` when a profile is not found

#### Posts
- ([\#964](https://github.com/desmos-labs/desmos/pull/964)) Do not allow to answer a poll after voting period ends

## Version 4.1.0
### Features
#### Posts
- ([\#941](https://github.com/desmos-labs/desmos/pull/941)) Allow to specify start and end indexes of tags and post references even when a post text is stored outside the chain
- ([\#942](https://github.com/desmos-labs/desmos/pull/942)) Allow to specify a post tags for content categorization

### Bug Fixes
#### Profiles
- ([\#954](https://github.com/desmos-labs/desmos/pull/954)) Added missing unlink messages WASM parsers

#### Subspaces
- ([\#945](https://github.com/desmos-labs/desmos/pull/945)) Allow to specify the section id inside user permissions query

#### Reactions
- ([\#940](https://github.com/desmos-labs/desmos/pull/940)) Fixed wrong event types and missing attributes

#### Reports
- ([\#955](https://github.com/desmos-labs/desmos/pull/955)) Fixed wrong CLI commands naming
- ([\#961](https://github.com/desmos-labs/desmos/pull/961)) Fixed the JSON fields used to deserialize report create and delete report messages received from CosmWASM

### Dependencies
- ([\#946](https://github.com/desmos-labs/desmos/pull/946)) Updated WASM parser to be compatible with `cosmwasm-std v1.0.0`

## Version 4.0.1
### Bug fixes
#### Other
- ([\#952](https://github.com/desmos-labs/desmos/pull/952)) Fixed how the `v.4.0.0` upgrade is handled

## Version 4.0.0
### Features
#### Subspaces
- ([\#866](https://github.com/desmos-labs/desmos/pull/866)) Added support for subspaces sections
- ([\#886](https://github.com/desmos-labs/desmos/pull/886)) Improved how permissions are managed
- ([\#919](https://github.com/desmos-labs/desmos/pull/919)) Added the ability to create a GenericSubspaceAuthorization to authorize operations only inside a specific subspace

#### Posts
- ([\#847](https://github.com/desmos-labs/desmos/pull/847)) Added the new `x/posts` module

#### Reactions
- ([\#898](https://github.com/desmos-labs/desmos/pull/898)) Added the new `x/reactions` module

#### Reports
- ([\#860](https://github.com/desmos-labs/desmos/pull/860)) Added the new `x/reports` module

#### Other
- ([\#822](https://github.com/desmos-labs/desmos/pull/822)) Added CosmWASM messages and queries parsers for all modules
- ([\#919](https://github.com/desmos-labs/desmos/pull/919)) Added proper support for `x/authz`

### Dependencies
- ([\#875](https://github.com/desmos-labs/desmos/pull/875)) Updated ibc-go to v3.0.0 and wasmd to v0.27.0

## Version 3.2.0
### Bug Fixes
#### Supply
- ([\#883](https://github.com/desmos-labs/desmos/pull/883)) Removed a wrongfully added supply store key

## Version 3.1.0
### Features
#### Profiles
- ([\#793](https://github.com/desmos-labs/desmos/pull/793)) Added the ability to reverse search application links and chain links

#### Fees
- ([\#796](https://github.com/desmos-labs/desmos/pull/796)) Added the new `x/fees` module

#### Supply
- ([\#782](https://github.com/desmos-labs/desmos/pull/782)) Added the new `x/supply` module

### Bug Fixes
#### Profiles
- ([\#831](https://github.com/desmos-labs/desmos/pull/831)) Renamed the `dtag_transfer_requests` REST endpoint to `dtag-transfer-requests`
- ([\#832](https://github.com/desmos-labs/desmos/pull/832)) Fixed the CLI commands expected number of arguments
- ([\#840](https://github.com/desmos-labs/desmos/pull/840)) Added missing profiles query commands

#### Subspaces
- ([\#801](https://github.com/desmos-labs/desmos/pull/801)) Added permissions sanitization
- ([\#802](https://github.com/desmos-labs/desmos/pull/802)) Made it not possible for users to set their own permissions
- ([\#805](https://github.com/desmos-labs/desmos/pull/805)) Added missing required permission when creating user groups
- ([\#815](https://github.com/desmos-labs/desmos/pull/815)) Fixed wrong Proto files notations and added missing CLI commands
- ([\#849](https://github.com/desmos-labs/desmos/pull/849)) Added correct Amino encoding support

#### Relationships
- ([\#838](https://github.com/desmos-labs/desmos/pull/838)) Replaced store keys to use less disk space

#### Other
- ([\#829](https://github.com/desmos-labs/desmos/pull/829)) Added missing events and events attributes

### Dependencies
- ([\#812](https://github.com/desmos-labs/desmos/pull/812)) Updated CosmWASM to v1.0.0-beta10
- ([\#844](https://github.com/desmos-labs/desmos/pull/844)) Updated Cosmos SDK to v0.45.5

## Version 3.0.1
### Bug Fixes
#### Other
- ([\#843](https://github.com/desmos-labs/desmos/pull/843)) Fixed the default home path inside the export command

## Version 3.0.0
### Notes
This version introduces breaking changes to the `x/profiles` Protobuf definitions. Particularly:
- all profiles types are now inside the `desmos/profiles/v2` package (it was `desmos/profiles/v1beta1`)
- relationships and user blocks models and queries have been moved to the `desmos/relationships/v1` package (it was `desmos/profiles/v1beta1`)

This requires all the clients to re-generate the Protobuf implementations in order to be compatible with the new version of the chain. Using previously generated models will result in errors when broadcasting messages and reading data from the chain itself.

### Features
#### Profiles
- ([\#688](https://github.com/desmos-labs/desmos/pull/688)) Added support for hex addresses inside chain links
- ([\#708](https://github.com/desmos-labs/desmos/pull/708)) Added support for multisig chain links
- ([\#785](https://github.com/desmos-labs/desmos/pull/785)) Added keeper hooks

#### Subspaces
- ([\#728](https://github.com/desmos-labs/desmos/pull/728)) Added the new `x/subspaces` module

#### Relationships
- ([\#750](https://github.com/desmos-labs/desmos/pull/750)) Split relationships and user blocks into the new `x/relationships` module

#### Other
- ([\#717](https://github.com/desmos-labs/desmos/pull/717)) Added IBC AnteHandler
- ([\#720](https://github.com/desmos-labs/desmos/pull/720)) Added CosmWasm module

### Bug Fixes
#### Profiles
- ([\#759](https://github.com/desmos-labs/desmos/pull/759)) Added the emission of missing events
- ([\#784](https://github.com/desmos-labs/desmos/pull/784)) Fixed how the profiles data are deleted (DTag transfer requests, chain links and application links)

#### Other
- ([\#693](https://github.com/desmos-labs/desmos/pull/693)) Added missing server version

### Dependencies
- ([\#716](https://github.com/desmos-labs/desmos/pull/716)) Updated IBC to v2.0.2
- ([\#769](https://github.com/desmos-labs/desmos/pull/769)) Updated Cosmos SDK to v0.45.1

## Version 2.3.1
### Bug Fixes
#### Profiles
- ([\#679](https://github.com/desmos-labs/desmos/pull/679)) Fixed the vesting accounts not working after the 2.3.0 upgrade
- ([\#680](https://github.com/desmos-labs/desmos/pull/680)) Fixed the wrong serialization of the AddressData interface when using Amino

## Version 2.2.1
### Dependencies
- ([\#668](https://github.com/desmos-labs/desmos/pull/668)) Updated Cosmos to v0.44.3 and fixed a bug inside the x/upgrade module

## Version 2.2.0
### Bug Fixes
#### Profiles
- ([\#662](https://github.com/desmos-labs/desmos/pull/662)) Fixed the application links from not being verified correctly

## Version 2.1.0
### Features
#### Profiles
- ([\#646](https://github.com/desmos-labs/desmos/pull/646)) Added the possibility to mutually exchange DTags between users
- ([\#649](https://github.com/desmos-labs/desmos/pull/649)) Removed custom json tags from Proto files
- ([\#651](https://github.com/desmos-labs/desmos/pull/651)) Made `MsgSaveProfile` DTag optional
- ([\#652](https://github.com/desmos-labs/desmos/pull/652)) Changed the plain text encoding of links to hex

## Version 2.0.2
### Dependencies
- ([\#647](https://github.com/desmos-labs/desmos/issues/647)) Updated Cosmos to v0.44.2

## Version 2.0.1
### Changes
- ([\#637](https://github.com/desmos-labs/desmos/pull/637)) Updated `go.mod` module to `v2`

## Version 2.0.0
### Features
#### Profiles
- ([\#539](https://github.com/desmos-labs/desmos/pull/539)) Made profiles query user parameter optional

### Bug Fixes
#### Profiles
- ([\#598](https://github.com/desmos-labs/desmos/pull/598)) Fixed the help of the `desmos tx profiles save` command

### Dependencies
- ([\#595](https://github.com/desmos-labs/desmos/pull/595)) Updated Cosmos to v0.44.0
- ([\#619](https://github.com/desmos-labs/desmos/pull/619)) Updated Band Protocol to v2.3.1

## Version 1.0.3
### Bug fixes
- Fixed the help of the `desmos tx profiles save` command

## Version 1.0.2
### Bug fixes
- Fixed the help of the `desmos tx profiles save` command

## Version 1.0.1
### Bug fixes
- Fixed a bug of the `init` command that prevented the creation of the `priv_validator_key.json` file when using the `--recover` flag.

## Version 1.0.0
### Changes
- Require chain name to be lowercase ([#533](https://github.com/desmos-labs/desmos/pull/533))
- Improved pagination ([#544](https://github.com/desmos-labs/desmos/pull/544))
- Improved the performance of profile validation checks ([\#557](https://github.com/desmos-labs/desmos/pull/557))
- Added `create-chain-link-json` command ([\#583](https://github.com/desmos-labs/desmos/pull/583))
- Changed Ledger app name from Cosmos to Desmos ([\#590](https://github.com/desmos-labs/desmos/pull/590))

## Version 0.17.7
### Changes
- Fixed tags not being fetched correctly inside `release` workflow

## Version 0.17.7
### Changes
- Fixed tags not being fetched correctly inside `release` workflow

## Version 0.17.6
### Bug fixes
- Fixed the Cosmos capability issue ([cosmos\#9835](https://github.com/cosmos/cosmos-sdk/pull/9835))

### Changes
- Updated Cosmos to `v0.42.8` ([\#578](https://github.com/desmos-labs/desmos/issues/578))

## Version 0.17.5
### Bug fixes
- Fixed the `--dry-run` flag not working properly

## Version 0.17.4
### Changes
- Added the on-chain upgrade handler

## Version 0.17.3
### Changes
- Removed the version check when creating an IBC channel

## Version 0.17.2
### Changes
- Renamed x/profiles messages response Proto definitions to match others

## Version 0.17.1
### Changes
- Improved the `x/profiles` params structure ([\#538](https://github.com/desmos-labs/desmos/issues/538))
- Added oracle-related data to `x/profiles` params ([\#537](https://github.com/desmos-labs/desmos/issues/537))
- Updated Cosmos to `v0.42.7` ([\#550](https://github.com/desmos-labs/desmos/issues/550))

## Version 0.17.0
### Changes
- Added the new `x/subspaces` module ([\#392](https://github.com/desmos-labs/desmos/issues/392))
- Made DTags case-insensitive ([\#492](https://github.com/desmos-labs/desmos/issues/492))
- Added the ability to paginate the relationships of a profile ([\#467](https://github.com/desmos-labs/desmos/issues/467))
- Added the ability to paginate user blocks ([\#495](https://github.com/desmos-labs/desmos/issues/495))
- Added the ability to paginate incoming DTag transfer requests ([\#519](https://github.com/desmos-labs/desmos/pull/519))
- Added the possibility to connect external chain accounts to a Desmos profile ([\#192](https://github.com/desmos-labs/desmos/issues/192))
- Added the possibility to verify a profile with an external application ([\#472](https://github.com/desmos-labs/desmos/issues/472))
- Added the ability to edit whether users can comment on a post or not ([\#446](https://github.com/desmos-labs/desmos/issues/446))
- Added the ability to paginate the registered reactions ([\#471](https://github.com/desmos-labs/desmos/issues/471))
- Added the ability to paginate the user poll answers ([\#478](https://github.com/desmos-labs/desmos/issues/478))
- Added the ability to paginate post reactions query ([\#515](https://github.com/desmos-labs/desmos/pull/515))
- Added the ability to query posts comments ([\#510](https://github.com/desmos-labs/desmos/pull/510))
- Improved the posts query ([\#499](https://github.com/desmos-labs/desmos/issues/499))
- Improved the registered reactions query ([\#515](https://github.com/desmos-labs/desmos/pull/515))
- Removed all the legacy REST endpoints ([\#447](https://github.com/desmos-labs/desmos/issues/447))
- Removed all the legacy queriers ([\#496](https://github.com/desmos-labs/desmos/issues/496))
- Added the logging of some interesting operations ([\#483](https://github.com/desmos-labs/desmos/issues/483))
- Updated Cosmos SDK to `v0.42.6` ([\#509](https://github.com/desmos-labs/desmos/issues/509))

#### Messages
- Renamed the following messages
  - `MsgCancelDTagTransfer` -> `MsgCancelDTagTransferRequest`
  - `MsgAcceptDTagTransfer` -> `MsgAcceptDTagTransferRequest`
  - `MsgRefuseDTagTransfer` -> `MsgRefuseDTagTransferRequest`

#### CLI
- Renamed the following CLI commands
  - `profiles tx request-dtag` -> `profiles tx request-dtag-transfer`
  - `profiles tx cancel-dtag-transfer` -> `profiles tx cancel-dtag-transfer-request`
  - `profiles tx accept-dtag-transfer` -> `profiles tx accept-dtag-transfer-request`
  - `profiles tx refuse-dtag-transfer` -> `profiles tx refuse-dtag-transfer-request`
  - `profiles query dtag-requests` -> `profiles query incoming-dtag-transfer-requests`
  - `profiles query blocklist` -> `profiles query blocks`

## Version 0.16.3
### Changes
- Renamed all `OptionalData` into `AdditionalAttributes` and `OptionalDataEntry` into `Attribute` ([\#422](https://github.com/desmos-labs/desmos/issues/422))
- Capped the `OptionalDataEntry.Key` length ([\#393](https://github.com/desmos-labs/desmos/issues/393)])
- Removed `--allows-comments` in favor of `--disable-comments` ([\#386](https://github.com/desmos-labs/desmos/issues/386))
- Standardized `PostID` naming convetions ([\#419](https://github.com/desmos-labs/desmos/issues/419))
- Renamed all events `Attribute` into `AttributeKey` ([\#423](https://github.com/desmos-labs/desmos/issues/423))
- Updated Cosmos to `v0.42.5` ([\#433](https://github.com/desmos-labs/desmos/issues/433))
- Merged `x/reports` inside `x/posts` ([\#429](https://github.com/desmos-labs/desmos/issues/429))

### Bug fixes
- Added missing `CommunitySpendProposal` handler ([\#421](https://github.com/desmos-labs/desmos/issues/421))

## Version 0.16.2
### Changes
- Renamed profile's `moniker` into `nickname` ([\#413](https://github.com/desmos-labs/desmos/issues/413))

### Bug fixes
- Fixed failing transactions performed by a `Profile` continuously depleting
  funds ([\#416](https://github.com/desmos-labs/desmos/issues/416))
- Fixed vesting accounts not working properly ([\#384](https://github.com/desmos-labs/desmos/issues/384))

## Version 0.16.1
### Changes
- Added the ability to query all accounts ([cosmos#8522](https://github.com/cosmos/cosmos-sdk/pull/8522))
- Refactored relationships related CLI commands to make them less
  implicits ([#408](https://github.com/desmos-labs/desmos/issues/408))
- Reintroduced deleted relationships CLI commands' tests ([#409](https://github.com/desmos-labs/desmos/issues/409))

### Bug fixes
- Removed duplicated `vesting` CLI commands
- Fixed the `--recover` flag of the `init` command not
  working ([cosmos#9201](https://github.com/cosmos/cosmos-sdk/pull/9201))
- Fixed `x/mint` queries not working ([\#403](https://github.com/desmos-labs/desmos/issues/403))
- Fixed the not-working `delete` CLI command for relationships inside `x/profiles` ([\#407](https://github.com/desmos-labs/desmos/issues/407))

## Version 0.16.0
### Changes
- Updated Cosmos to `v0.42.4`
- Disabled all the modules, except the `x/profiles` module
- Changed how data are stored inside the `x/profiles`
  modules ([\#261](https://github.com/desmos-labs/desmos/issues/261))
- Enabled the `x/ibc` transfer module ([\#382](https://github.com/desmos-labs/desmos/issues/382))

## Version 0.15.5
- Updated Cosmos to `v0.42.3` ([\#387](https://github.com/desmos-labs/desmos/issues/387))

## Version 0.15.4
### Changes
- Updated Cosmos to `v0.42.1` ([\#378](https://github.com/desmos-labs/desmos/issues/378))

### Bug fixes
- Fixed a security vulnerability identified in the `app.go` file.

## Version 0.15.3
### Bug fixes
- Added missing gRPC gateways ([\#370](https://github.com/desmos-labs/desmos/issues/370))

## Version 0.15.2
### Changes
- Updated Cosmos to v0.41.3 ([\#359](https://github.com/desmos-labs/desmos/issues/359))
- Added Tendermint LD flag ([\#360](https://github.com/desmos-labs/desmos/issues/360))
- Added ARM-32 support

## Version 0.15.1
### Changes
- Replaced `gogoproto.jsontag` usages to be coherent with Proto field names
- Removed usage of `json` inside genesis-related methods ([\#348](https://github.com/desmos-labs/desmos/issues/348))
- Updated Cosmos to `v0.40.1` ([\#345](https://github.com/desmos-labs/desmos/issues/345))

### Bug fixes
- Fixed some buys inside the `verify-genesis` command that did not detect some errors

## Version 0.15.0
### Changes
- Updated Cosmos to v0.40.0 ([\#313](https://github.com/desmos-labs/desmos/issues/313))
- Updated validators documentation ([\#333](https://github.com/desmos-labs/desmos/issues/333))
- Renamed binary from `desmosd` to `desmos` ([\#342](https://github.com/desmos-labs/desmos/issues/342))

## Version 0.14.0
### Changes
- Implemented blocked users checks ([\#298](https://github.com/desmos-labs/desmos/issues/298))
- Implemented the possibility to set a minimum messages fee ([\#230](https://github.com/desmos-labs/desmos/issues/230))

### Bug fixes
- Fixed height not working in REST queries ([\#299](https://github.com/desmos-labs/desmos/issues/299))

## Version 0.13.0
### Changes
- Removed the relationship event attribute key prefix ([\#300](https://github.com/desmos-labs/desmos/issues/300))
- Removed the user_block event attribute key prefix ([\#291](https://github.com/desmos-labs/desmos/issues/291))
- Changed posts' optional data representation ([\#272](https://github.com/desmos-labs/desmos/issues/272))
- Fixed bugs inside the DTag transfer process ([\#295](https://github.com/desmos-labs/desmos/issues/295)
  , [\#296](https://github.com/desmos-labs/desmos/issues/296))
- Implemented the possibility to refuse and cancel DTag requests from both receiver and sender side ([\#297](https://github.com/desmos-labs/desmos/issues/297))

### Bug fixes
- Fixed the possibility of requesting a transfer of an empty DTag ([\#292](https://github.com/desmos-labs/desmos/issues/292))
- Fixed the impossibility of querying all the relationships on chain ([\#306](https://github.com/desmos-labs/desmos/issues/306))

## Version 0.12.3
### Changes
- Renamed the `accept-dtag-transfer` CLI command to remove the `.md`
  suffix ([\#282](https://github.com/desmos-labs/desmos/issues/282))

## Version 0.12.2
### Changes
- Added the migration command from v0.10.0 to v0.12.0

## Version 0.12.1
### Bug fixes
- Fixed an upgrade migration bug

## Version 0.12.0
### Changes
- Changed `relationships`' implementation adding a `subspace` field to identify in which app users make relationships ([\#266](https://github.com/desmos-labs/desmos/issues/266))
- Implemented the possibility to (un)block a specific user ([\#169](https://github.com/desmos-labs/desmos/issues/169))
- Allow users to edit their DTag ([\#226](https://github.com/desmos-labs/desmos/issues/226))
- Allow users to give their DTag away ([\#225](https://github.com/desmos-labs/desmos/issues/225))

## Version 0.11.0
### Changes
- Allowed the possibility to edit a post's attachments and poll data using the `MsgEditPost`
  type ([\#202](https://github.com/desmos-labs/desmos/issues/202))
- Removed the `Open` field from within the `PollData` object. Now you should rely on the `CloseDate` field to determine
  whether a poll is close or open. ([\#252](https://github.com/desmos-labs/desmos/issues/252))
- Implemented users `Relationships` ([\#168](https://github.com/desmos-labs/desmos/issues/168))

## Version 0.10.0
### Changes
- Changed application errors to make them more clear ([\#237](https://github.com/desmos-labs/desmos/issues/237))
- Implemented the `x/upgrade` module ([\#174](https://github.com/desmos-labs/desmos/issues/174))
- Removed user specified post's `CreationDate` and `LastEdited` in favor of block time usage ([\#215](https://github.com/desmos-labs/desmos/issues/215))
- Renamed posts' parameters queries endpoints ([\#245](https://github.com/desmos-labs/desmos/issues/245))
- Renamed `PostMedia` into `Attachment` to fix incorrect singular and plural forms of variables ([\#203](https://github.com/desmos-labs/desmos/issues/203))
- Updated Cosmos to `v.0.39.1` ([\#257](https://github.com/desmos-labs/desmos/issues/257))

### Bug fixes
- Fixed a bug that caused valid URIs to be considered invalid ([\#233](https://github.com/desmos-labs/desmos/issues/233))
- Fixed a bug that didn't allow querying `x/profile` parameters from REST API ([\#244](https://github.com/desmos-labs/desmos/issues/244))

## Version 0.9.0
### Changes
- Removed the JSON-style `String` methods where it was possible, changed the others to not rely on JSON for `String`
  representation ([\#199](https://github.com/desmos-labs/desmos/issues/199))
- Replaced `SetupTestInput()` with testify test suite ([\#198](https://github.com/desmos-labs/desmos/issues/198))
- Removed all the `internal` folders ([\#197](https://github.com/desmos-labs/desmos/issues/197))

## Version 0.8.2
### Changes
- Updated Cosmos to v0.38.5

## Version 0.8.1
### Changes
- Added the removal of invalid registered reactions during `v0.8.0` migration

### Bug fixes
- Fixed a bug inside the post validation method that did not consider poll-only posts valid
- Added the registration of new modules when migrating towards `v0.8.0`

## Version 0.8.0
### Changes
- Changed the data stored inside the profile objects ([\#193](https://github.com/desmos-labs/desmos/issues/193))
- Renamed the `profile_cov` field of `MsgSaveProfile` to `cover_picture`
- Renamed the `profile_pic` field of `MsgSaveProfile` to `profile_picture`
- Renamed the `profile` module to `profiles` ([\#200](https://github.com/desmos-labs/desmos/issues/200))
- Moved `profiles` module constants to chain parameters ([\#171](https://github.com/desmos-labs/desmos/issues/171))
- Moved `posts` module constants to chain parameters ([\#172](https://github.com/desmos-labs/desmos/issues/172))
- Added the creation date inside the profile saving event ([\#210](https://github.com/desmos-labs/desmos/issues/210))
- Changed the way times are serialized inside event attributes ([\#211](https://github.com/desmos-labs/desmos/issues/211))
- Updated Cosmos to `v0.38.5`

### Bug fixes
- Fixed a bug inside the `Equals` method of the `Pictures` object
- Changed the `tx profiles save` flags names (fixes #207)

## Version 0.7.0
### Changes
- Implemented benchmarks tests ([\#126](https://github.com/desmos-labs/desmos/issues/126))
- Implemented posts' reports ([\#50](https://github.com/desmos-labs/desmos/issues/50))
- Re-introduced the on-chain government module ([\#173](https://github.com/desmos-labs/desmos/issues/173))
- Fixed reactions registration bug ([\#187](https://github.com/desmos-labs/desmos/issues/187))

## Version 0.6.3
### Changes
- Restored evidence module ([\#189](https://github.com/desmos-labs/desmos/issues/189))

## Version 0.6.2
### Changes
- Updated Cosmos to v0.38.4 ([\#177](https://github.com/desmos-labs/desmos/issues/177))

## Version 0.6.1
### Changes
- Updated the way with which the profiles are created and edited ([\#170](https://github.com/desmos-labs/desmos/issues/170))

### Bug fixes
- Fixed the on-chain events usage ([\#175](https://github.com/desmos-labs/desmos/issues/175))

## Version 0.6.0
### Changes
- Added the option to use [RocksDB](https://github.com/facebook/rocksdb) as both Tendermint and/or Cosmos database
  backend ([\#111](https://github.com/desmos-labs/desmos/issues/111))
- Implemented tags in post medias ([\#118](https://github.com/desmos-labs/desmos/issues/118))
- Edited PostReaction struct to allow a better integration with middle layer applications ([\#157](https://github.com/desmos-labs/desmos/issues/157))

### Bug fixes
- Fixed the account query CLI command ([\#155](https://github.com/desmos-labs/desmos/issues/155))
- Fixed the profile deletion CLI command ([\#166](https://github.com/desmos-labs/desmos/issues/166))

## Version 0.5.3
### Changes
- Updated Cosmos to [v0.38.5](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.38.5)

## Version 0.5.2
### Bug fixes

- Fixed a bug that caused the state export to fail due
  to [cosmos/cosmos-sdk#6280](https://github.com/cosmos/cosmos-sdk/issues/6280)

## Version 0.5.1
### Bug fixes
- Fixed a bug that caused users to be unable to add more than one reaction to the same post

## Version 0.5.0
### Changes
- Implemented invariants for posts and profile modules ([\#90](https://github.com/desmos-labs/desmos/issues/90))
- Added YAML support for types ([\#124](https://github.com/desmos-labs/desmos/issues/124))
- Improved reactions events ([\#144](https://github.com/desmos-labs/desmos/issues/144))
- Removed automatic registration of emoji reactions ([\#145](https://github.com/desmos-labs/desmos/issues/145))
- Improved reaction registration error message ([\#147](https://github.com/desmos-labs/desmos/issues/147))
- Allow for empty message posts when they contain a poll ([\#148](https://github.com/desmos-labs/desmos/issues/148))

## Version 0.4.0
### Changes
- Improved the generation of post ids ([\#131](https://github.com/desmos-labs/desmos/issues/131))
- Improved `alias.go` files ([\#103](https://github.com/desmos-labs/desmos/issues/103))
- Added the support for posting empty-message posts with medias ([\#110](https://github.com/desmos-labs/desmos/issues/110))
- Implemented the support for hashtags in posts ([\#96](https://github.com/desmos-labs/desmos/issues/96))
- Updated the post create CLI command in posts ([\#117](https://github.com/desmos-labs/desmos/issues/117))
- Implemented the support for registering new reactions ([\#94](https://github.com/desmos-labs/desmos/issues/94))
- Implemented the support for decentralized profiles ([\#56](https://github.com/desmos-labs/desmos/issues/56))
- Improved the storage usage to reduce gas usage ([\#125](https://github.com/desmos-labs/desmos/issues/125))
- Removed the `gov` and `upgrade` modules as they are currently not used ([\#142](https://github.com/desmos-labs/desmos/issues/142))

### Bug fixes
- Fixed a bug inside the migration procedure of the `magpie`
  module ([\#106](https://github.com/desmos-labs/desmos/issues/106))

## Version 0.3.2
- Fixed a bug that should allow to properly export the state of the chain

## Version 0.3.1
- Updated Cosmos SDK to `v0.38.3` and Tendermint to `v0.33.3` to solve security issues.

## Version 0.3.0
### Changes
- Implemented the support for media posts ([\#36](https://github.com/desmos-labs/desmos/issues/36))
- Implemented the support for poll posts  ([\#14](https://github.com/desmos-labs/desmos/issues/14))
- Added the support for posts sorting ([\#78](https://github.com/desmos-labs/desmos/issues/78))
- Added the support for magpie default session length inside genesis ([\#38](https://github.com/desmos-labs/desmos/issues/38))
- Posts now only supports `subspace` values in form of hex-encoded SHA-256 hashes ([\#82](https://github.com/desmos-labs/desmos/issues/82))
- Bumped Cosmos to `v0.38.0` ([\#10](https://github.com/desmos-labs/desmos/issues/10))

### Bug fixes
- Fixed the posts REST endpoint not working properly ([\#77](https://github.com/desmos-labs/desmos/issues/77))
- Fixed a bug that allowed to create multiple posts with the exact same contents ([\#92](https://github.com/desmos-labs/desmos/issues/92))

### Migration
In order to migrate the chain state from version `v0.2.0` to `v0.3.0`, please run the following command:

```bash
desmos migrate v0.3.0 <path-to-genesis-file>
```

## Version 0.2.0
### Changes
- Implemented the support for arbitrary data inside a post ([\#52](https://github.com/desmos-labs/desmos/issues/52)
  , [\#66](https://github.com/desmos-labs/desmos/issues/66))
- Implemented the support for posts reactions ([\#47](https://github.com/desmos-labs/desmos/issues/47))
- Implemented the support for posts subspaces ([\#46](https://github.com/desmos-labs/desmos/issues/46))
- Automated the default bond denom change to `desmos` ([\#25](https://github.com/desmos-labs/desmos/issues/25))
- Replaced the block height with timestamps inside posts' creation dates and edit dates ([\#62](https://github.com/desmos-labs/desmos/issues/62))
- Capped the post message length to 500 characters ([\#67](https://github.com/desmos-labs/desmos/issues/67))

### Migration
In order to migrate the chain state from version `v0.1.0` or `v0.1.1` to `v0.2.0`, please run the following command:

```bash
desmos migrate v0.2.0 <path-to-genesis-file>
```

## Version 0.1.1
### Bug fixes
- Fixed double children IDs insertion upon post edit ([\#63](https://github.com/desmos-labs/desmos/issues/63))
- Fixed a bug that made impossible to create a new post upon a post edit due to the `Post with ID X already exists` ([\#64](https://github.com/desmos-labs/desmos/issues/64))

## Version 0.1.0
### Changes
- Create a session to associate an external chain address to a Desmos address.
- Create a post using a `MsgCreatePost` and providing a message. You can also decide whether other users can comment on
  such post or not.
- Like a post using a `MsgLikePost` and specifying a post id.
- Unlike a post using a `MsgUnlikePost` and specifying a post id.

### Notes

- When generating Desmos accounts, the path to use is `m'/852'/0'/0/0`
- The stake token denomination is `desmos`