tag = $(shell git describe --tags $(git rev-list --tags --max-count=1))
dir = $(CURDIR)/build

$(shell rm -rf $(dir); mkdir -p $(dir))

all:
	go build

release: win32 win64 osx freebsd arm linux-x64

win64:
	GOOS=windows GOARCH=amd64 go build -o pwcalc.exe && zip $(dir)/pwcalc_$(tag)_$@.zip pwcalc.exe

win32:
	GOOS=windows GOARCH=386   go build -o pwcalc.exe && zip $(dir)/pwcalc_$(tag)_$@.zip pwcalc.exe

osx:
	GOOS=darwin  GOARCH=amd64 go build -o pwcalc && zip $(dir)/pwcalc_$(tag)_$@.zip pwcalc

freebsd:
	GOOS=freebsd GOARCH=amd64 go build -o pwcalc && tar cvzf $(dir)/pwcalc_$(tag)_$@.tgz pwcalc

arm:
	GOOS=linux   GOARCH=arm   go build -o pwcalc && tar cvzf $(dir)/pwcalc_$(tag)_$@.tgz pwcalc

linux-x64:
	GOOS=linux   GOARCH=amd64 go build -o pwcalc && tar cvzf $(dir)/pwcalc_$(tag)_$@.tgz pwcalc

clean:
	rm -rf $(dir) pwcalc pwcalc.exe

