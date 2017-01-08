tag = $(shell git describe --tags $(git rev-list --tags --max-count=1))
dir = $(CURDIR)/build

$(shell rm -rf $(dir); mkdir -p $(dir))

all:
	go build

release: windows_amd64 darwin_amd64 freebsd_amd64 linux_arm	linux_amd64

windows_amd64:
	GOOS=windows GOARCH=amd64 go build -o pwcalc.exe && zip $(dir)/pwcalc_$(tag)_$@.zip pwcalc.exe

darwin_amd64:
	GOOS=darwin  GOARCH=amd64 go build -o pwcalc && zip $(dir)/pwcalc_$(tag)_$@.zip pwcalc

freebsd_amd64:
	GOOS=freebsd GOARCH=amd64 go build -o pwcalc && tar cvzf $(dir)/pwcalc_$(tag)_$@.tgz pwcalc

linux_arm:
	GOOS=linux   GOARCH=arm   go build -o pwcalc && tar cvzf $(dir)/pwcalc_$(tag)_$@.tgz pwcalc

linux_amd64:
	GOOS=linux   GOARCH=amd64 go build -o pwcalc && tar cvzf $(dir)/pwcalc_$(tag)_$@.tgz pwcalc

clean:
	rm -rf $(dir) pwcalc pwcalc.exe

