#!/bin/bash
docker service update --image client:1.0.1 client
docker service update --image server:1.0.1 server
