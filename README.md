# go-tls-benchmark

Minimal HTTPS server. Built for [a blogpost on certificate performance](https://blog.espe.tech/posts/rsa-may-slow-you-down/).

## Building

```bash
$ go build -o tlsbench
```

## Profiling

Use `go tool pprof profile-rsa` to dig into the cpu profile.

### RSA
```
$ ./tlsbench -cert-file rsa-cert.pem -key-file rsa-key.pem -cpu-profile profile-rsa
```

### ECDSA
```
$ ./tlsbench -cert-file ecdsa-cert.pem -key-file ecdsa-key.pem -cpu-profile profile-ecdsa
```

### ED25519
```
$ ./tlsbench -cert-file ed25519-cert.pem -key-file ed25519-key.pem -cpu-profile profile-ed25519
```

