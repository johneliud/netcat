# netcat

This project consists on recreating the netcat in a server-client architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

## Prerequisites

- [Go](https://go.dev/doc/install) programming language version 1.22 and above.

- Basic understanding of what TCP is and how it works.

## Installation

- Clone the repository from the remote location to your local environment and switch to the project's directory using the below terminal command.
    ```bash
    git clone https://github.com/johneliud/netcat.git
    cd netcat
    ```

## Usage
- Run the program from the `root` using either of the below terminal commands.

    1.
    ```bash
    go run .
    ```

    2.
    ```bash
    go run . [SPECIFY-PORT]
    ```

**NOTE:** The program uses port `8989` as its default port using the first run command. A success message is displayed if the program succeeds to run.

- Open a new terminal session and paste the below command to access the program using netcat's server.
```bash
nc localhost [DISPLAYED-PORT]
```