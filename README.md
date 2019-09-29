# ma-zero-trust-prototype

Prototype of the zero trust model for a research network.

## Proxy

Routes every authenticated and authorized data traffic to the corresponding service.

Should ideally only get mtls connections. However, requests without a certificate will also be accepted if policy permits.

* Running on moodle.uni-ulm.de:10443
* Expects a running moodle service on moodle.uni-ulm.de:8443
* Expects a running sso authentication service on sso.uni-ulm.de:4435
* Expects a running policy engine on policy.uni-ulm.de:4438

### Environment Variables

For Test Environment:

* IGNORE_FAILED_AUTHORIZATION=1 
* MOODLE_DISABLED=1 (for testing with the dummy web service)
* DUMMY_WEB_ADDRESS="dummy-web.uni-ulm.de:4438"
* ...

Values can be changed locally in `moodle-proxy/.env` or in `./docker-compose.yml` like:
```
environment:
 − IGNORE_FAILED_AUTHORIZATION=0
```

### Moodle Service

* Example Docker File and source code in moodle_service
* Running on moodle.uni-ulm.de:8443
* Requires activated Authentication Plugin for JWT (rudimentary implemented in the moodle_data source code)

## SSO Auth

A Single Sign On Authentication Server, which is addressed exclusively by the proxies of the services. 
Provides the various login pages (basicAuth or 2FA).
Sends failed and succeeded authentication attempts to the logger.

Creates, signs and verifies Json Web Tokens for Authentication (RSA keys).

* Running on sso.uni-ulm.de:4435
* Expects a running logger service on logger.uni-ulm.de:4432
* Expects running ldap test server at localhost:1389 for basic authentication. https://github.com/rroemhild/docker-test-openldap

## Policy Engine

Handles the Authorization of (ideally) every Request.
Builds Network Agent and reveices Trust Score from Trust Engine.
Checks whether the client's request complies with the given policies.

:warning: Only accepts MTLS Connections from listed Proxies.

* Running on policy.uni-ulm.de:4438
* Expects a running trust engine on trust.uni-ulm.de:4439

## Trust Engine

Handles the Trust Management for the Authorization Decisions.
Calls logger, to get the subjects' latest auth attempts.

:warning: Only accepts MTLS Connection from Policy Engine.

* Running on trust.uni-ulm.de:4439
* Expects a running logger service on logger.uni-ulm.de:4432

## Logger

Used by SSO to log the user authentication attempts. 
Requested by Trust Engine to get the log entries of a particular subject (user or device) since given timestamp.

:warning: Only accepts MTLS Connection from SSO Service and Trust Engine.

* Running on logger.uni-ulm.de:4432

## Prerequisites

* Insert Project into $GOPATH/src/github.com/ma-zero-trust-prototype
* Required Tools
    * Go
    * Docker
    * Docker Compose
    * Step CLI
    * Step Certificates
* **Edit Host File**:
    ```
    # Required for the X.509 certificates
    127.0.0.1	localhost moodle.uni-ulm.de ca.uni-ulm.de sso.uni-ulm.de policy.uni-ulm.de trust.uni-ulm.de logger.uni-ulm.de
    ```

