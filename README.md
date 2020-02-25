# getThebackup

Scans backup folders/files on target sites. Searches archived files in the folders it finds.

1 - config/extensions.txt - This adds new extensions to the file, for example: by adding in the form of .example allows you to retry all the possibilities tried in the new extensions.

2 - config/files.txt - It can scan these folders according to the extensions you added, by giving them new file names.

3 - config/folders.txt - Recursively scans the specified folders. You can add to this list yourself.


## Installation & Run

`go run getthebackup.go --hostname host.test`

or 

`go build getthebackup.go`

`./getthebackup --hostname host.test`


