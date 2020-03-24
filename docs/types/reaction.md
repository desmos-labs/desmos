# Reaction
Inside Desmos, the reactions is the fastest way users can react to posts.

Everyday each one of us use reactions inside the most popular social networks.
Reactions let you immediately express what you think about the post/photo/video you've just watched.

Each new reaction that will be registered inside Desmos will become part of the set of reactions supported by the system.  
Once you have registered your favorites GIFs/images/emojis, you will be allowed to use them to react to every post inside Desmos.
Remember that a reaction can be registered only once per `subspace`, so if you ever try to register a previously registered
reaction, your transaction will not be valid. 

## Contained data
A reaction is made of different parts. Following you will find out what are those and how they can be used.

### `ShortCode`
The `ShortCode` identifies the actual emoji's short code.  
Emoji's Short codes are codes used on various websites to speed up emoji insertion using a keyboard. 
These begin with a colon and include a shorter version of an emoji name.  
For example It must look like this `:emoji-shortcode:` and it will be validated
by the following regEx: `:[a-z]([a-z\d_])*:`.  
[Here](https://www.webfx.com/tools/emoji-cheat-sheet/) the list of all available short codes.

### `Value`
The `Value` of a reaction identifies whether the reaction is a GIF an image or an emoji.  
Value can be a `URL` with the path of the GIF/image your using as a reaction or it can be a `UNICODE` 
that identifies a specific emoji.  
`URL`'s will be validated by the following regEx: `^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$`.
`Unicode` must be one of the following [list](https://unicode.org/emoji/charts/full-emoji-list.html).

### `Subspace`
The `Subspace` field identifies the application inside which the reaction has been registered.  
Currently the subspace must be a SHA256 hash of the previously plain-text value.

### `Creator`
The `Creator` field is used to specify the Bech32 address of the creator of the reaction.  
In order for a creator address to be valid, it must begin with the `desmos` Bech32 human-readable part.

 


