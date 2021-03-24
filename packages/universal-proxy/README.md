## Lab Multi-Proxy

A Caddy web server that can be used to reverse proxy multiple GPA Digital Lab projects at once.

> One Proxy to rule them all, One Proxy to find them,\
> One Proxy to bring them all, and in the Dockers bind them,
>
> ~ J.R.R Tolkien

Specifically, it maps the following:

- Content Commons (localhost:3000) -> commons.dev.local
- CDP Public API (localhost:8080) -> api.dev.local
- Content Dev Site (content_web Docker container) -> content.dev.local
- Lab Dev Site (lab_web Docker container) -> lab.dev.local
