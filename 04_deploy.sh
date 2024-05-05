#!/bin/bash
docker service create server --name server
docker service create client --name client