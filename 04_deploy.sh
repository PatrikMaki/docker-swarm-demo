#!/bin/bash
docker service create --network my-network --name server server
docker service create --network my-network --name client client
