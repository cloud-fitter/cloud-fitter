#！/bin/bash

rm -rf gen/*
buf beta mod update
buf generate