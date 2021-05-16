# bbs_api


## generate openapi code

[OpenAPITools/openapi-generator](https://github.com/OpenAPITools/openapi-generator)

```sh
$ docker run --rm -v $PWD:/local openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g go-server -o /local/out/go

$ cp out/go/go/* openapi/
```
