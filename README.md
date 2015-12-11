Admin
---

Admin is meant to help administer a multiuser linux system

Building
---

```bash
go build -o admin_linux-amd64 -tags netgo *.go
# Or
./script/build
```
> `-tags netgo` will help you achieve static binaries :)

Running
---

```bash
./admin_linux-amd64
docker run --rm -ti pdxjohnny/admin
```
