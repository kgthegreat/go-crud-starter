#! /bin/bash
entity=$1
rm controllers/$entity.go
rm models/$entity.go
rm repositories/$entity.go
rm -rf templates/entity1
