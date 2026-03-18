A simple in-memory TTL (Time-To-Live) cache implemented in Go.
Supports automatic key expiration and concurrent-safe operations.
Built as a learning project to explore cache eviction strategies and memory management.

## Features

- In-memory key-value store
- TTL-based expiration
- Concurrent-safe (sync.Mutex / RWMutex)
- Background cleanup worker
- Simple CLI interface


## Installation

```bash
git clone https://github.com/yourusername/ttl-cache
cd ttl-cache
go build -o ttl-cache

## Run
./ttl-cache


---

## 5️⃣ CLI Usage

```md
## Usage

SET key value ttl
GET key
DELETE key
EXIT


SET user1 dhirendra 10
GET user1


## Architecture

The cache stores entries in a Go map with metadata:

- value
- expiration timestamp

A background goroutine periodically scans and removes expired keys.

Concurrency is handled using sync.RWMutex to ensure thread-safe access.


## Example

> SET user1 username 5
OK

> GET user1
username

(after 5 seconds)

> GET user1
Key expired




## Roadmap

- Add LRU eviction policy
- Add persistence (disk snapshot)
- Add distributed support
- Replace map scanning with min-heap for efficient expiration


Client (CLI)
    ↓
Cache Layer
    ↓
In-Memory Map
    ↓
Background Cleaner