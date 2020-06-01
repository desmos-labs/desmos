# Version 0.7.0
- Implemented benchmarks tests (#126)
- Implemented posts' reports (#50)
- Re-introduced the on-chain government module (#173)

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
