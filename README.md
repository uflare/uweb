uWeb
====
> a uflarians micro framework based on go plugins, it is all about just a plugin

How
===
> just write a go plugin that has a `Run(*sync.Map)` function, sync.Map has `*echo.Echo` with key `echo`,  `go build -buildmode=plugin` and add it to the `.env` and start the `uweb` server

Credits
=======
uFlare (c) 2018
