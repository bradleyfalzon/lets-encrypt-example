# Overview

This example shows how to run Let's Encrypt's `certbot` with Go's `ListenAndServeTLS`.

Note, this method requires `certbot` to be installed with its dependencies such as Python.
It also requires the daemon the be restarted at least monthly after running `certbot`.

# Running Example

Build or Install your binary

```
go install
```

Run as root via `sudo`, but ideally you would use runit/upstart/systemd to
automatically restart daemons and use `setcap` to start as a non privileged
user (required as this example binds to port 80 and 443)

```
sudo ./lets-encrypt-example -domain example.com -webroot /home/user/go/src/github.com/bradleyfalzon/lets-encrypt-example
```

Upon running this for the first time, the `ListenAndServeTLS` will fail due to
non existent certificates, this is OK and expected on the first run.

Next step is the run `certbot`, which will communicate with Let's Encrypt, and
then place a randomly generated key inside `.well-known/acme-challenge/random-key.txt`
which Let's Encrypt will then fetch by accessing <http://example.com/.well-known/acme-challenge/random-key.txt>.

```
certbot-auto --renew-by-default certonly --webroot -w /home/user/go/src/github.com/bradleyfalzon/lets-encrypt-example -d example.com
```

If this is the first time you've ran `certbot` you will be asked for your email
address and must agree to the Terms of Service.

Once a certificate has been successfully obtained, restart the `lets-encrypt-example` to use the new certificates.


# Other Approaches

## Certbot Alternatives

`certbot` isn't the only tool that can obtain certificates for you, and with its Python dependencies it's not light
weight either.

- Go: <https://github.com/hlandau/acme> (has many package manager repositories available)
- Go: <https://github.com/xenolf/lego>
- Bash: <https://github.com/lukas2511/letsencrypt.sh>

## Libraries

Other libraries available, such as <https://github.com/dkumor/acmewrapper> and <https://github.com/rsc/letsencrypt>
fetch and renew Let's Encrypt certificates automatically from within your application,
this removes the `certbot` requirements (and its dependencies) as well as removes the
requirement to constantly restart your application when `certbot` updates certificates.

## Web Servers

Some web servers, such as [Caddy](https://caddyserver.com/) provide standard web server
capabilities, as well as automatic HTTPS via Let's Encrypt.

## Configuration Management

Many approaches focus on a single server model, if you require the same certificate
distributed to multiple servers, the web servers themselves usually should not generate
the certificates themselves. Another approach is to have a single server running `certbot`
and the web servers or load balancers proxy all requires to `.well-known/acme-challenge`
to this dedicated server. This server can then generate the certificates once and
you can use your existing configuration management tools to push these certificates
to your web servers or load balancers and reload the relevant daemons (if required).
