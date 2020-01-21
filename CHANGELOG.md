# Version 0.3.0
## Changes
- Implemented the support for media in posts (#36)
 
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

```shell
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
