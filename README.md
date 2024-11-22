# netcat

This project consists on recreating the netcat in a server-client architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

## Prerequisites

- [Go](https://go.dev/doc/install) programming language version 1.22 and above.

- Basic understanding of [TCP](https://en.wikipedia.org/wiki/Transmission_Control_Protocol).

## Installation

- Clone the repository from the remote location to your local environment and switch to the project's directory using the below terminal command.
    ```bash
    git clone https://github.com/johneliud/netcat.git
    cd netcat
    ```

## Usage
- Run the program from the `root` using either of the below terminal commands.
    ```bash
    go run .
    ```

    ```bash
    go run . [SPECIFY-PORT]
    ```

**NOTE:** The program uses port `8989` as its default port using the first run command. A success message is displayed if the program succeeds to run.

- Open a new terminal session and paste the below command to access the program using netcat's server.
```bash
nc localhost [SPECIFIED-PORT]
```

## Example
1. Run the program's server

![](/readme-img/run-program.png)


2. Switch to new terminal sessions and access the program using netcat's server.

![](/readme-img/access-netcat-server.png)

## Contributions

Contributions towards improving the project are welcome. Feel free to create a new branch and submit a pull request or create issues on how to improve the project.

### Creating a branch

-
    ```bash
    git checkout -b feat-improvement-your_name
    ```

-
    ```bash
    git add modified_file
    ```

-
    ```bash
    git commit -m "description_of_your_changes"
    ```

-
    ```bash
    git push --set-upstream origin feat-improvement-your_name
    ```

## Contact

Incase of further enquiries, do not hesitate to reach [out](johneliud4@gmail.com).