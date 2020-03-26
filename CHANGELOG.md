# Version 0.4.0
## Changes
- Improved `alias.go` files (#103)
- Added the support for posting empty-message posts with medias (#110)
- Implemented the support for hashtags in posts (#96)
- Updated the post create CLI command in posts (#117)
- Implemented the support for registering new reactions (#94)

## Bug fixes
- Fixed a bug inside the migration procedure of the `magpie` module (#106)

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
