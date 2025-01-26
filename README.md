# gRPC CSV Reader

This is a gRPC CSV reader system that streams data from a server to a client. The server reads rows from a CSV file and streams them to the client via gRPC.

---

## How to Run
### 1. Install Dependencies
Run the following command to install all required dependencies:
```bash
go mod tidy
```

### 2. Generate `.pb.go` Files (if needed)
If the files `csv_reader.pb.go` and `csv_reader_grpc.pb.go` are not already present in the `proto/` folder, generate them using the this command:
```bash
protoc --go_out=. --go-grpc_out=. proto/csv_reader.proto
```

### 3. Install Protocol Buffers Compiler and Plugins
Ensure the following tools are installed:
- **Protocol Buffers Compiler (`protoc`)**:
  Install using Homebrew:
  ```bash
  brew install protobuf
  ```

- **Protobuf Go Plugin (`protoc-gen-go`)**:
  Install with:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  ```

- **gRPC Go Plugin (`protoc-gen-go-grpc`)**:
  Install with:
  ```bash
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

### 4. Add Go Binaries to PATH
Make sure the Go binaries are in your `PATH` environment variable. Run the this command:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
Add this line to your shell configuration file (ie, `.bashrc` or `.zshrc`) to make it permanent:
```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
```
Reload your shell configuration:
```bash
source ~/.zshrc
```

### 5. Run the Server
Start the server by running:
```bash
go run main.go
```
You should see this message:
```
Server is running on port 50051...
```

### 6. Run the Client
Open a new terminal and run the client to fetch the CSV data:
```bash
go run client/client.go
```

You should now see the streamed rows from the server in the client’s output.
