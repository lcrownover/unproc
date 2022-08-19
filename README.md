# unproc

`unproc` is a client/server tool used to query user process counts via gRPC.

For clarification, this project is definitely not done :)

## Installing

`unproc` builds into two binaries, `unproc_server` and `unproc_client`.

### Installing both tools

```bash
make
sudo make install
```

### Installing only server

```bash
make server
sudo make install_server
```

### Installing only client

```bash
make client
sudo make install_client
```

## Usage

Run the server with `unproc -host 0.0.0.0 -port 6661`.

TODO(lcrown): write systemd unit file

Run the client with `unproc -host localhost -port 6661 -username myuser`
