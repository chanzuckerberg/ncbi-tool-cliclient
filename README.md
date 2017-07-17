* NCBI-replica with version history data storage platform. Phase 1 of creating a service for accessing old versions of NCBI data. For use in the Infectious Disease platform.

* Components:
  * Sync service: https://github.com/chanzuckerberg/ncbi-tool-sync
  * Server service: https://github.com/chanzuckerberg/ncbi-tool-server
  * Command line client: https://github.com/chanzuckerberg/ncbi-tool-cliclient

* Planning docs:
  * Part 1: https://docs.google.com/document/d/1y9Y6Q5HgPHT5CfIPCMtkK2gIINtzcTEhdNzEWwqIIw4/edit
  * Part 2 and API documentation: https://docs.google.com/document/d/1mRzOFqJvhAWb4954o1eV-DVSvm_RFukohnt5bvTch-4/edit

* Testing:
  - To avoid running some of the acceptance tests, run go test with -short, e.g.
    - ```go test -short ./...```