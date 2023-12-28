#!/bin/bash
fuser -k 8080/tcp && echo "Stop Server" || echo "Not Running"