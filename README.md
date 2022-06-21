# SSI medical prescriptions-demo

## Mock server

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
    - Openapi will be available at http://localhost:8889

3. Stop mock server
    ```bash
    make stop-mock-server
    ```