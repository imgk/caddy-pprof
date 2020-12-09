# Caddy-PProf

## Build with [xcaddy](https://github.com/caddyserver/xcaddy)

```
$ xcaddy build \
    --with github.com/imgk/caddy-pprof
```

## Config

```
{
    "apps": {
        "http": {
            "servers": {
                "": {
                    "routes": [
                        {
                            "handle": [
                                {
                                    "handler": "pprof"
                                }
                            ]
                        }
                    ]
                }
            }
        }
    }
}

```
