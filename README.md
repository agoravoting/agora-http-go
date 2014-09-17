## agora-http-go [![Build Status][1]][2] [![Coverage Status](https://coveralls.io/repos/agoravoting/agora-http-go/badge.png)](https://coveralls.io/r/agoravoting/agora-http-go)

[1]: https://travis-ci.org/agoravoting/agora-http-go.png
[2]: https://travis-ci.org/agoravoting/agora-http-go

The agora-http-go is a server-side middleware written in go language that
provides an http server and some additional common services:

* Http server (codegangsta/negroni)
* Http routing (imdario/medeina and julienschmidt/httprouter)
* Http header based authentication and permissions based on HMAC's
* Raven error reporting (kisielk/raven-go/raven)
* Db connectivity (jmoiron/sqlx)
* Db migrations (liamstask / goose)