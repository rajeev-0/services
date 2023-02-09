This package defines [`IEvidinceHandler`](ievidecehandler.go) and
[`IEndorsementHandler`](iendorsementhandler.go) [pluggable](../plugin/README.md)
interfaces and associated RPC channels. These are used to add new attestation
scheme sort to Veraison services. Additionally, the package defines a [couple
of wrappers](plugin.go) around `plugin.RegisterImplementation` for registering
implementations of these two interfaces.