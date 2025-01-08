package = "auto-request-id"
version = "0.0.1"

dependencies = {
  "kong >= 3.9",
  "lua-resty-uuid",
}

build = {
  type = "builtin",
  modules = {
    ["kong.plugins.auto-request-id.handler"] = "handler.lua",
    ["kong.plugins.auto-request-id.schema"] = "schema.lua",
  },
}