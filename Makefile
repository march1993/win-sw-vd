
all:
	GOOS=windows go build -ldflags "-H=windowsgui -s -w" .
