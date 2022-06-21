# SSI medical prescriptions-demo


## Run mock server
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