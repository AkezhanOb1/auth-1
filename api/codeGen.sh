#!/bin/bash
protoc auth.proto  --go_out=plugins=grpc:.