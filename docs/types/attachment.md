# Attachment
The `Attachment` type contains the details of a single attachment file object that can be associated within a [`Post`](post.md) object. With attachment, you can add images and multimedia files (vocals, documents, videos, etc.) to posts.

Following you will find a description of all the fields it is composed of. 

## `URI`
The first field of an `Attachment` is the `URI` field. This field should contain the URI of the attachment file that is represented. 

When creating a `Post` on the chain, a regular expression checks this `URI`. If the check does not pass, the post will not be stored and an error will be thrown instead: 

```phpregexp
^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$
```

## `MimeType`
The second field of an `Attachment` is the `MimeType` field. This one allows you to specify the [MIME type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types) of the included media file. 

No check is ever performed on this field's values, and any string is accepted as valid as long as it is not empty. That being said, please make sure to use a valid MIME type each time you specify it as it will make it easier for other apps to read your data. 

## `Tags`
`Tags` is the third field of a `Attachment`. This field allows you to tag any user on a particular attachment.
This field can be omitted and the system will check that every tag inside the array `Tags` is a valid `bech32` 
encoded `address` of desmos.
E.g.
`desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr`