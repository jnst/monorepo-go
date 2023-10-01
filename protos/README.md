# protos

## Notes

Generate `buf.yaml`

```shell
buf mod int
```

Add proto files

```shell
touch auth/v1/auth.proto
touch user/v1/user.proto
```

Lint and Generate

```shell
buf lint
buf generate
```
