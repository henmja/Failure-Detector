# Simple failure detector using heartbeat
Three nodes are used for this example.

## Background
Each node has a server and a client. Each client broadcasts heartbeats to all the servers (which means that each node also sends heartbeats to themselves). If a client does not receive a reply within 10 seconds, the client will broadcast to all the servers that the given node has disconnected.

## Compile/run
First, set ip address and port number of the three machines in the three clients & servers.
Start up all the servers, then the clients.
You might need to turn off your firewall.
