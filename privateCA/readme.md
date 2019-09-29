# Private CA

Created with [Smallstep Certificates](https://github.com/smallstep/certificates) (version Smallstep CLI/0.8.6-dev)

## Configuration

Create new Certificates Example: (requires a running step ca)

used:`step ca certificate "moodle.uni-ulm.de" moodle_proxy.crt moodle_proxy.key` (CLI/0.8.6-dev)

alternative from [certificate create Docs](https://smallstep.com/docs/cli/certificate/create/): `step certificate create moodle.uni-ulm.de moodle_proxy.crt moodle_proxy.key` 

* Name: uniulm
* Domain: ca.uni-ulm.de
* Port: 4443
* Subject: CN=uniulm Root CA
* Password: uniulmca

## Let's get started

Detailed explanation in Smallsteps [Let's get started](https://github.com/smallstep/certificates#lets-get-started)

1.  Init  `step ca init`
2.  Run `step-ca $(step path)/config/ca.json`

You can find these artifacts in `$STEPPATH` (or `~/.step` by default)

## Smallstep Certificates

An online certificate authority and related tools for secure automated
certificate management, so you can use TLS everywhere.

For more information and docs see [the Step website](https://smallstep.com/cli/)
and the [blog post](https://smallstep.com/blog/step-certificates.html)
announcing Step Certificate Authority.

![Animated terminal showing step certificates in practice](https://github.com/smallstep/certificates/raw/master/docs/images/step-ca-2-legged.gif)

