* NCBI-replica with version history data storage platform. Phase 1 of creating a service for accessing old versions of NCBI data. For use in the Infectious Disease platform.

## Installation
1. Download the binary for your system:
    * `bin/darwin-amd64/ncbitool` for Mac
    * `bin/linux-amd64/ncbitool` for Linux
    * Or build from source: `go build -o ncbitool`
2. Move to `/usr/local` if desired. Or run with `./ncbitool command`
3. Run `chmod u+x ncbitool` if any permissions issues.

## Usage
* Reference API documentation [here](https://docs.google.com/document/d/1mRzOFqJvhAWb4954o1eV-DVSvm_RFukohnt5bvTch-4/edit).
* Usage:
    - `ncbitool command [command options] [arguments...]`
    - General file structure on the server is the same as on `ftp://ftp.ncbi.nlm.nih.gov/blast/db/`
* Commands:
    - `file` - file actions. Gets file info.
    - `directory` - directory actions. Gets directory info.
    - `help, h` - Shows a list of commands or help for one command
* Sub-commands:
  - `file at-time` - get a file version at or before a point in time
  - `file history` - get the version history of a file
  - `directory at-time` - get a directory state at or before a point in time
  - `directory compare` - compare a directory state across a start and end date
* Flags:
    - Make sure to put flags before the last file path argument.
    - `--download` - include to download the files to local disk
    - `--input-time [value]` - input time for 'at time' requests. Ex: `2017-07-07T00:06:12`
    - `--version-num [value]` - version number for file requests
    - `--dest [value]` - download destination on local disk
    - `--help, -h` - show help
    - `--start-date [value]` - start date for directory diff comparisons
    - `--end-date [value]` - end date for directory diff comparisons
* Example commands:
  - `ncbitool file /blast/db/README`
  - `ncbitool file --download --version-num 3 --dest ~/Desktop /blast/db/README`
  - `ncbitool file at-time --download --input-time "2017-07-12T08:38:46" /blast/db/nr.00.tar.gz`
  - `ncbitool file history /blast/db/nr.00.tar.gz`
  - `ncbitool directory /blast/db/FASTA`
  - `ncbitool directory --download --dest ~/Desktop /blast/db/FASTA`
  - `ncbitool directory at-time --download --input-time "2005-07-12T08:38:46" --dest ~/Desktop /blast/db/FASTA`
  - `ncbitool directory compare --start-date "2015-07-12T08:38:46" --end-date "2017-07-20T08:38:46" /blast/db/FASTA`

## Development info
1. `git clone https://github.com/chanzuckerberg/ncbi-tool-cliclient`
2. `go get github.com/urfave/cli`
3. `go build -o ncbitool` or `env GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/ncbitool` to cross-compile for a specific platform.

## Links
* Components:
  * Sync service: https://github.com/chanzuckerberg/ncbi-tool-sync
  * Server service: https://github.com/chanzuckerberg/ncbi-tool-server
  * Command line client: https://github.com/chanzuckerberg/ncbi-tool-cliclient

* Planning docs:
  * Part 1: https://docs.google.com/document/d/1y9Y6Q5HgPHT5CfIPCMtkK2gIINtzcTEhdNzEWwqIIw4/edit
  * Part 2 and API documentation: https://docs.google.com/document/d/1mRzOFqJvhAWb4954o1eV-DVSvm_RFukohnt5bvTch-4/edit