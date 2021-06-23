# Reaction
Inside Desmos, the reactions is the fastest way users can react to posts.

Everyday each one of us use reactions inside the most popular social networks.
Reactions let you immediately express what you think about the post/photo/video you've just watched.

Each new reaction that will be registered inside Desmos will become part of the set of reactions supported by the system.  
Once you have registered your favorite GIF/image/emoji, you will be allowed to use them to react to every post inside Desmos.
Remember that a reaction can be registered only once per `subspace`, so if you ever try to register a previously 
registered reaction, your transaction will not be valid. 

## Contained data
A reaction is made of different parts. Following you will find out what are those and how they can be used.

### `ShortCode`
The `ShortCode` identifies the actual reaction short code.  
Short codes are codes used on various websites to speed up reaction insertion using a keyboard.
These begin and end with a colon, and contain the literal name of the reaction itself. 
For example, it can look something like `:emoji-shortcode:`.
When registering a new reaction, the shot code must be validated by the following regEx: `:[a-z0-9+-]([a-z0-9\d_-])*:`. 
[Here](https://www.webfx.com/tools/emoji-cheat-sheet/) the list of all available short codes.

### `Value`
The `Value` of a reaction identifies whether the reaction is a GIF an image or an emoji.  
Value can be a `URL` with the path of the GIF/image your using as a reaction or it can be a `UNICODE` 
that identifies a specific emoji.  
`URL`'s will be validated by the following regEx: `^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$`.
`Unicode` must be one of the [supported ones](https://github.com/desmos-labs/Go-Emoji-Utils/blob/master/data/emoji.json).

### `Subspace`
The `Subspace` field identifies the application inside which the reaction has been registered.  
Currently the subspace must be a SHA256 hash of the previously plain-text value.

### `Creator`
The `Creator` field is used to specify the Bech32 address of the creator of the reaction.  
In order for a creator address to be valid, it must begin with the `desmos` Bech32 human-readable part.
