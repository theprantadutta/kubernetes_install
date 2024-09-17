### Description
This script installs kubernetes inside a ubuntu machine

### First Clone the Repo
```
git clone https://github.com/theprantadutta/kubernetes_install
```

### Then Build the binary
```
go build -o ./bin/install-kube ./src/main.go
```

### Finally, Run it
```
./bin/install-kube
```

### Or maybe combine them all and run it as
```
git clone https://github.com/theprantadutta/kubernetes_install && cd kubernetes_install && go build -o ./bin/install-kube ./src/main.go && ./bin/install-kube
```