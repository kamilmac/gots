Template for Go server with Typescript client

1. Install deps with "npm install" in client folder.
2. Run "go run main.go" for dev
3. For prod:
    - "go run build.go" for bundling client files and server compilation.
    - "main -prod" + optional "-port PORT" to run production server.
4. Have fun