#!/bin/bash
(cd client && docker build -t client .)
(cd server && docker build -t server .)

