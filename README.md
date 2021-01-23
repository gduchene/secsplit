`secsplit` is a simple wrapper around the
`github.com/hashicorp/vault/shamir` package.

`secsplit` lets you split a secret into several parts using [Shamir’s
Secret Sharing][shamir] algorithm. `secsplit` also encodes those parts
using Base64, and expects parts to be Base64-encoded when it reads them.

[shamir]: https://en.wikipedia.org/wiki/Shamir's_Secret_Sharing
