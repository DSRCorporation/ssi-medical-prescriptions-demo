# SSI medical prescriptions-demo

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
1. Build mock server docker image:
    ```bash
    make mock-server-docker
    ```
2. Run mock server container
    ```
    make run-mock-server
    ```
    - Mock server api will be available at http://localhost:8989
    - Openapi will be available at http://localhost:8889/openapi

3. Stop mock server
    ```bash
    make stop-mock-server
    ```
