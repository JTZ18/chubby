# chubby
A (very simplified) implementation of [Chubby](https://static.googleusercontent.com/media/research.google.com/en//archive/chubby-osdi06.pdf), Google's distributed lock service, written for [COS 518, Spring 2019](http://www.cs.princeton.edu/courses/archive/spr19/cos518/index.html).

## Instructions
<a href="https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/JTZ18/chubby"><img src="https://img.shields.io/static/v1?label=Dev%20Containers&message=Open&color=blue&logo=visualstudiocode" alt="Open in Dev Containers"></a>

Open the repo in vscode and choose to open in dev contaier or alternatively, click the button above to set up dev container locally on your machine

To bring up five Chubby nodes in individual Docker containers, run `docker-compose up`. Superuser privileges may be required.

Chubby nodes can also be run locally as individual processes. To build, `cd` into the `chubby` subdirectory and run `make chubby`. We can bring up three Chubby nodes as follows:

1. First node: `./chubby -id "node1" -raftdir ./node1 -listen ":5379" -raftbind ":15379"`
2. Second node: `./chubby -id "node2" -raftdir ./node2 -listen ":6379" -raftbind ":16379" -join "127.0.0.1:5379"`
3. Third node: `./chubby -id "node3" -raftdir ./node3 -listen ":7379" -raftbind ":17379" -join "127.0.0.1:5379"`

Example Chubby clients can be found in the `cmd` folder. To run, build using `make [CLIENT NAME]`, then run the resulting executable (e.g., `make simple_client; ./simple_client`).
