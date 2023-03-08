#!/bin/bash

kill -9 $(cat ./srv.pid)
rm srv.pid srv
