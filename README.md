# AWS Lambda Go functions with Netlify

Read more about Go functions support on Netlify in [our documentation](https://www.netlify.com/docs/lambda-functions).


Example:


Curl requests isnt working due to rewrite-rule but things work on POSTMAN:

```
{
	"signingMessage": "SIGNING MESSAGE",
	"privKeyHex": "PRIVATE KEY HEX",
	"senderAddress": "Sender ADDRESS",
	"sigType": "ecdsa_recovery",
	"payloadSigType": "ecdsa_recovery"
}
```


```
curl -X POST \
  http://serverlessgo.mobeenashraf.com/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	"signingMessage": "SIGNING MESSAGE",
	"privKeyHex": "PRIVATE KEY HEX",
	"senderAddress": "Sender ADDRESS",
	"sigType": "ecdsa_recovery",
	"payloadSigType": "ecdsa_recovery"
}'
```
# Build

Function is built in functions directory. You can see details by reading make command in Makefile

```
build:
	mkdir -p functions
	go get ./...
	go build -o functions/serverless-go ./...
```

There is a redirect rule applied to direct .netlify/function/serverless-go to '/'

```
[[redirects]]
  from = "/"
  to = "/.netlify/functions/serverless-go"
  status = 200
  force = true
```

Netlify uses go 1.14, To update it, had to update Function's Environment variable:

https://app.netlify.com/sites/site_name/settings/deploys#environment
GO_VERSION 1.16
