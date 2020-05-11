# PostMedia
The `PostMedia` type contains the details of a single media file object that can be associated within a [`Post`](post.md) object. Such type is usually used to add images to posts, but it can also be used to specify multimedia files such as vocals, documents, etc. 

Following you will find a description of all the fields it is composed of. 

## `URI`
The first field of a `PostMedia` is the `URI` field. This field should contain the URI of the media file that is represented. 

When creating a `Post` on the chain, this `URI` is checked against the following regular expression. If the check does not pass, the post will not be stored and an error will be thrown instead: 

```phpregexp
^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$
```

## `MimeType`
The second field of a `PostMedia` is the `MimeType` field. This one allows you to specify the [MIME type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types) of the included media file. 

No check is ever performed on this field's values, and any string is accepted as valid as long as it is not empty. That being said, please make sure to use a valid MIME type each time you specify it as it will make it easier for other apps to read your data. 

## `Tags`
`Tags` is the third field of a `PostMedia`. This field allows you to tag any user on a particular media.
This field can be omitted and the system will check that every tag inside the array `Tags` is a valid `bech32` 
encoded `address` of desmos.
E.g.
`desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr`