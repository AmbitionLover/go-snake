GO_VERSION=`go version`

convey:
	goconvey -excludedDirs internal,cmd,config,deployments,docs,init,log,scripts,tmp

run:
	go run main.go