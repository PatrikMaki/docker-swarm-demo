#!/bin/bash
(cd client && docker build -t client:1.0.1 .)
(cd server && docker build -t server:1.0.1 .)

