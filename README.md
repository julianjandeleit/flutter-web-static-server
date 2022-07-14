# README
## WHY another server
In my usecase python simple http server was not able to serve a simple flutter static app. However, a simple go server works.

## Deployment
make install
it builds the executable _static-server_ and then installs the file
or build executable only with go build main.go

### Windows
There exists a Makefile for Windows, used as ´make install -f MakefileWin´.
Assumes ´/bin/´ exists and is in path.

Dependencies like go and sudo can be installed using [scoop](http://scoop.sh).

## Development
Currently, the directory that is served is the current directory. It is served on 0.0.0.0:9000. This configuration is currently hardcoded.

The functionality may vary and made more modular when deemed appropriate. 
