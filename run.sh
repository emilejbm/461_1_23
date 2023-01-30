#!/bin/sh

#include check for num of arguments
case $1 in
    "install")
        echo -n "will install dependencies tbd";; 
    "build")
        echo -n "will compile tbd";;
    "URL_FILE") # i have no idea how to actually do this maybe regex ?
        echo -n "will process URLs tbd";;
    "test") # logging and testing stuff
        echo -n "will test and log tbd";;
esac