local utils = require "kong.tools.utils"

local AutoRequestIdPlugin = {
  VERSION = "0.0.1",
  PRIORITY = 1000,
}

function AutoRequestIdPlugin:access(conf)
  -- Get the key name from the plugin configuration
  local keyName = conf.key_name

  -- Check if the header exists
  local headerValue = kong.request.get_header(keyName)
  if not headerValue then
    -- Generate a UUID and add it to the request header
    headerValue = utils.uuid()
    kong.service.request.set_header(keyName, headerValue)
    kong.log.info("Generated UUID and added to request header: ", headerValue)
  end

  kong.response.set_header(keyName, headerValue)
  kong.log.info("Added header to response: ", headerValue)
end

return AutoRequestIdPlugin