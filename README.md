# SSI medical prescriptions-demo

## Demo server

### Build demo server
```bash
make demo-server
```

### Run demo server and swagger ui using docker
1. Run demo server
    ```
    make run-demo-server
    ```
    - Demo server api will be available at http://localhost:8888
    - Openapi will be available at http://localhost:8889/openapi

2. Stop demo server
    ```bash
    make stop-demo-server
    ```

## Mock server

### Generate mock stubs for openapi specs
1. Install oapi-codegen:
    ```bash
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    ```
    ```bash
    export PATH=$PATH:$HOME/go/bin
    ```
2. Generate mock stubs for openapi specs
    ```bash
    oapi-codegen -package mock ./api/openapi-spec/openapi.yml > internal/controller/mock/ssimp_mock.gen.go
    ```
### Build mock server
```bash
make mock-server
```

### Run mock server and swagger ui using docker
1. Run mock server container
    ```
    make run-mock-server
    ```
    - Mock server api will be available at http://localhost:8989
    - Openapi will be available at http://localhost:8889/openapi

2. Stop mock server
    ```bash
    make stop-mock-server
    ```