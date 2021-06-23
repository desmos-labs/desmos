# `MsgUnlinkApplication`
This message allows you to remove a previously linked application from your Desmos profile.

## Structure

```json
{
  "@type": "/desmos.profiles.v1beta1.MsgUnlinkApplication",
  "application": "<Name of the application to unlink>",
  "username": "<Name of the account inside the application that should be unlinked>",
  "signer": "<Desmos address of the profile that should remove the link>"
}
```

### Attributes

| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `application`  | `string` | Name of the application to unlink |
| `username`| `string` | Name of the account inside the application that should be unlinked |
| `signer` | String | Desmos address of the profile that should remove the link |

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgUnlinkApplication",
  "application": "twitter",
  "username": "RicMontagnin",
  "signer": "desmos1qchdngxk8zkl4c4mheqdlpgcegkdrtucmwllpx"
} 
````

## Message action
The action associated to this message is the following:

```
unlink_application
```
