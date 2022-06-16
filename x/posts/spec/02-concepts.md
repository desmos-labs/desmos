---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts 

## Post
A Post is a structure representing any kind of content in a social network. It can contain some text and be enriched with:
- Attachments such as media (pictures, clips, videos, GIFs) and polls;
- Links;
- Mentions of users;
- Hashtags values;
- Quotes of other posts.

### Subspace ID
The [subspace] ID indicates the ID of the Dapp where the post is hosted and lives. It is represented by an unsigned
64 bytes integer.

### Section ID 
The [section] ID indicates the ID of the subspace's section where the post lives. It is represented by an unsigned 32
bytes integer.

### ID
The unique ID that identifies the post itself. This ID is automatically assigned to the post at the moment of its 
creation in a sequential way (e.g. if there's 4 posts in the chain, the one we are creating will have id equal to 5).
It is represented by an unsigned 64 bytes integer. 

### External ID (Optional)
External ID indicates and external ID attached to the post. It is represented by a string. //TODO add some more info 

### Text (Optional)
The text is the actual textual content of the post. It has a fixed max length that is determined by an on-chain governance parameter.

## Entities
Entities are particular parts of the text that can be parsed out of it in order to be displayed in custom ways.
Entities are divided in 3 different categories:
- Hashtags (i.e. #desmos)
- Mentions (i.e. @desmos1xcfui...., @Forbole)
- Urls (i.e. https://desmos.network, ftp://user:password@example.com/image.png)

### Tag
Both hashtags and mentions are represented by a tag. The tag structure contains the necessary fields that ease the process of 
text's parsing.
#### Start

#### End

#### Tag