# AWS-Cognito-Tutorials
Code from Cognito Youtube Tutorials (Youtube: Nikhil Rao)

SDK's Needed
------------
  - https://sdk.amazonaws.com/js/aws-sdk-2.7.16.min.js
  - amazon-cognito-identity.min.js
  
 Other Resources
---------------
The full How-to can be found at https://docs.aws.amazon.com/cognito/latest/developerguide/tutorial-integrating-user-pools-javascript.html#tutorial-integrating-user-pools-user-sign-in-javascript


## Upload images
Using UploadImage.html

*  Config Indentity Pool

```
{
   "Version": "2012-10-17",
   "Statement": [
      {
         "Effect": "Allow",
         "Action": [
            "s3:*"
         ],
         "Resource": [
            "arn:aws:s3:::BUCKET_NAME/*"
         ]
      }
   ]
}


````
> Need To Enable access to unauthenticated identities

* Config bucket CORS
```
<?xml version="1.0" encoding="UTF-8"?>
<CORSConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
    <CORSRule>
        <AllowedOrigin>*</AllowedOrigin>
        <AllowedMethod>POST</AllowedMethod>
        <AllowedMethod>GET</AllowedMethod>
        <AllowedMethod>PUT</AllowedMethod>
        <AllowedMethod>DELETE</AllowedMethod>
        <AllowedMethod>HEAD</AllowedMethod>
        <AllowedHeader>*</AllowedHeader>
    </CORSRule>
</CORSConfiguration>
```
* Config file js/config.js
```
s3: {
        BUCKET_NAME: '',
        REGION: 'ap-northeast-1',
        IDENTITY_POOL_ID: ''
    }
```









