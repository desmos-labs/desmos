# Version 0.14.0
## Changes
- Implemented blocked users checks (#298)
- Implemented the possibility to set a minimum messages fee (#230) 

## Bug fixes
- Fixed height not working in REST queries (#299)

# Version 0.13.0
## Changes
- Removed the relationship event attribute key prefix (#300)
- Removed the user_block event attribute key prefix (#291)
- Changed posts' optional data representation (#272)
- Fixed bugs inside the DTag transfer process (#295, #296)
- Implemented the possibility to refuse and cancel DTag requests from both receiver and sender side (#297)

## Bug fixes
- Fixed the possibility of requesting a transfer of an empty DTag (#292)
- Fixed the impossibility  of querying all the relationships on chain (#306)

# Version 0.12.3
## Changes
- Renamed the `accept-dtag-transfer` CLI command to remove the `.md` suffix (#282)

# Version 0.12.2
## Changes
- Added the migration command from v0.10.0 to v0.12.0

# Version 0.12.1
## Bug fixes
- Fixed an upgrade migration bug

# Version 0.12.0
## Changes
- Changed `relationships`' implementation adding a `subspace` field to identify in which app users make relationships (#266)
- Implemented the possibility to (un)block a specific user (#169)
- Allow users to edit their DTag (#226)
- Allow users to give their DTag away (#225)

# Version 0.11.0
## Changes
- Allowed the possibility to edit a post's attachments and poll data using the `MsgEditPost` type (#202)
- Removed the `Open` field from within the `PollData` object. Now you should rely on the `CloseDate` field to determine whether a poll is close or open. (#252)
- Implemented users `Relationships` (#168)

# Version 0.10.0
## Changes
- Changed application errors to make them more clear (#237)
- Implemented the `x/upgrade` module (#174)
- Removed user specified post's `CreationDate` and `LastEdited` in favor of block time usage (#215)
- Renamed posts' parameters queries endpoints (#245)
- Renamed `PostMedia` into `Attachment` to fix incorrect singular and plural forms of variables (#203)
- Updated Cosmos to `v.0.39.1` (#257)

## Bug fixes
- Fixed a bug that caused valid URIs to be considered invalid (#233)
- Fixed a bug that didn't allow querying `x/profile` parameters from REST API (#244) 

# Version 0.9.0
## Changes
- Removed the JSON-style `String` methods where it was possible, changed the others to not rely on JSON for `String` representation (#199)
- Replaced `SetupTestInput()` with testify test suite (#198)
- Removed all the `internal` folders (#197)

# Version 0.8.2
## Changes
- Updated Cosmos to v0.38.5

# Version 0.8.1
## Changes
- Added the removal of invalid registered reactions during `v0.8.0` migration

## Bug fixes
- Fixed a bug inside the post validation method that did not consider poll-only posts valid
- Added  the registration of new modules when migrating towards `v0.8.0`

# Version 0.8.0
## Changes
- Changed the data stored inside the profile objects (#193)
- Renamed the `profile_cov` field of `MsgSaveProfile` to `cover_picture`
- Renamed the `profile_pic` field of `MsgSaveProfile` to `profile_picture`
- Renamed the `profile` module to `profiles` (#200)
- Moved `profiles` module constants to chain parameters (#171)
- Moved `posts` module constants to chain parameters (#172)
- Added the creation date inside the profile saving event (#210)
- Changed the way times are serialized inside event attributes (#211)
- Updated Cosmos to `v0.38.5`

## Bug fixes
- Fixed a bug inside the `Equals` method of the `Pictures` object 
- Changed the `tx profiles save` flags names (fixes #207)

# Version 0.7.0
## Changes
- Implemented benchmarks tests (#126)
- Implemented posts' reports (#50)
- Re-introduced the on-chain government module (#173)
- Fixed reactions registration bug (#187)

# Version 0.6.3
## Changes
- Restored evidence module (#189)

# Version 0.6.2
## Changes
- Updated Cosmos to v0.38.4 (#177)

# Version 0.6.1
## Changes
- Updated the way with which the profiles are created and edited (#170)

## Bug fixes
- Fixed the on-chain events usage (#175)

# Version 0.6.0
## Changes
- Added the option to use [RocksDB](https://github.com/facebook/rocksdb) as both Tendermint and/or Cosmos database backend (#111)
- Implemented tags in post medias (#118)
- Edited PostReaction struct to allow a better integration with middle layer applications (#157)

## Bug fixes
- Fixed the account query CLI command (#155)
- Fixed the profile deletion CLI command (#166)

# Version 0.5.3
## Changes
- Updated Cosmos to [v0.38.5](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.38.5)

# Version 0.5.2
## Bug fixes
- Fixed a bug that caused the state export to fail due to [cosmos/cosmos-sdk#6280](https://github.com/cosmos/cosmos-sdk/issues/6280)

# Version 0.5.1
## Bug fixes
- Fixed a bug that caused users to be unable to add more than one reaction to the same post

# Version 0.5.0
## Changes
- Implemented invariants for posts and profile modules (#90)
- Added YAML support for types (#124)
- Improved reactions events (#144)
- Removed automatic registration of emoji reactions (#145)
- Improved reaction registration error message (#147)
- Allow for empty message posts when they contain a poll (#148)

# Version 0.4.0
## Changes
- Improved the generation of post ids (#131)
- Improved `alias.go` files (#103)
- Added the support for posting empty-message posts with medias (#110)
- Implemented the support for hashtags in posts (#96)
- Updated the post create CLI command in posts (#117)
- Implemented the support for registering new reactions (#94)
- Implemented the support for decentralized profiles (#56)
- Improved the storage usage to reduce gas usage (#125)
- Removed the `gov` and `upgrade` modules as they are currently not used (#142)

## Bug fixes
- Fixed a bug inside the migration procedure of the `magpie` module (#106)

# Version 0.3.2
- Fixed a bug that should allow to properly export the state of the chain

# Version 0.3.1
- Updated Cosmos SDK to `v0.38.3` and Tendermint to `v0.33.3` to solve security issues.

# Version 0.3.0
## Changes
- Implemented the support for media posts (#36)
- Implemented the support for poll posts  (#14) 
- Added the support for posts sorting (#78)
- Added the support for magpie default session length inside genesis (#38)
- Posts now only supports `subspace` values in form of hex-encoded SHA-256 hashes (#82)
- Bumped Cosmos to `v0.38.0` (#10)

## Bug fixes
- Fixed the posts REST endpoint not working properly (#77)
- Fixed a bug that allowed to create multiple posts with the exact same contents (#92) 

## Migration
In order to migrate the chain state from version `v0.2.0` to `v0.3.0`, please run the following command:

```bash
desmosd migrate v0.3.0 <path-to-genesis-file> 
```

# Version 0.2.0
## Changes
- Implemented the support for arbitrary data inside a post (#52, #66)
- Implemented the support for posts reactions (#47)
- Implemented the support for posts subspaces (#46)
- Automated the default bond denom change to `desmos` (#25)
- Replaced the block height with timestamps inside posts' creation dates and edit dates (#62)
- Capped the post message length to 500 characters (#67)

## Migration
In order to migrate the chain state from version `v0.1.0` or `v0.1.1` to `v0.2.0`, please run the following command:

```bash
desmosd migrate v0.2.0 <path-to-genesis-file> 
```

# Version 0.1.1
## Bug fixes
- Fixed double children IDs insertion upon post edit (#63)
- Fixed a bug that made impossible to create a new post upon a post edit due to the `Post with ID X already exists` (#64)

# Version 0.1.0
## Features
- Create a session to associate an external chain address to a Desmos address. 
- Create a post using a `MsgCreatePost` and providing a message. You can also decide whether other users can comment on such post or not. 
- Like a post using a `MsgLikePost` and specifying a post id. 
- Unlike a post using a `MsgUnlikePost` and specifying a post id.

## Notes
- When generating Desmos accounts, the path to use is `m'/852'/0'/0/0`
- The stake token denomination is `desmos`  
