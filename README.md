# AppTrial
 Application trial expiration library for golang



Add this package with your go project

```
go get github.com/rafiulgits/apptrial
```



### How to use

```go
trail := NewAppTrial("MyApp", time.Minute*2, "_this_is_my_encrypt_key_")
trail.Start()
```

Here `time.Minute*2` indicate that this application will expired after 2minutes of fresh installtion.

**Note:** Encryption key should be `16/24/32 bytes` 

