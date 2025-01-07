local typedefs = require "kong.db.schema.typedefs"

return {
  name = "auto-request-id-plugin",
  fields = {
    { consumer = typedefs.no_consumer },
    { protocols = typedefs.protocols_http },
    { config = {
        type = "record",
        fields = {
          { key_name = { type = "string", required = true } },
        },
      },
    },
  },
}