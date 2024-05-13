# Chatrooms

This project implements a simple RPC chat application with a server and client written in Go.

## Prerequisites

- Go 1.15 or later installed (<https://go.dev/doc/install>)

## Running the project

### 1. **Build the server and client:**

```bash
go build ./server/server
go build ./client/client
```

### 2. **Run the server:**

Open a terminal and run:

```bash
./server.exe
```

This will start the server listening on port 7422 by default.

### 3. **Run multiple clients:**

Open separate terminals for each client and run:

```bash
./client.exe
```

Each client will be prompted for a nickname and can then send messages to the chat. All messages will be broadcast to all connected clients.
