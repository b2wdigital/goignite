package gifiber

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	"github.com/gofiber/fiber/v2"
)

const (
	root                      = "gi.fiber"
	port                      = root + ".port"
	configRoot                = root + ".config"
	prefork                   = configRoot + ".prefork"
	serverHeader              = configRoot + ".serverHeader"
	strictRouting             = configRoot + ".strictRouting"
	caseSensitive             = configRoot + ".caseSensitive"
	immutable                 = configRoot + ".immutable"
	unescapePath              = configRoot + ".unescapePath"
	ETag                      = configRoot + ".ETag"
	bodyLimit                 = configRoot + ".bodyLimit"
	concurrency               = configRoot + ".concurrency"
	readTimeout               = configRoot + ".readTimeout"
	writeTimeout              = configRoot + ".writeTimeout"
	idleTimeout               = configRoot + ".idleTimeout"
	readBufferSize            = configRoot + ".readBufferSize"
	writeBufferSize           = configRoot + ".writeBufferSize"
	compressedFileSuffix      = configRoot + ".compressedFileSuffix"
	proxyHeader               = configRoot + ".proxyHeader"
	GETOnly                   = configRoot + ".GETOnly"
	reduceMemoryUsage         = configRoot + ".reduceMemoryUsage"
	network                   = configRoot + ".network"
	disableKeepalive          = configRoot + ".disableKeepalive"
	disableDefaultDate        = configRoot + ".disableDefaultDate"
	disableDefaultContentType = configRoot + ".disableDefaultContentType"
	disableHeaderNormalizing  = configRoot + ".disableHeaderNormalizing"
	disableStartupMessage     = configRoot + ".disableStartupMessage"
	ExtRoot                   = root + ".ext"
)

func init() {
	giconfig.Add(port, 8082, "Server http port")
	giconfig.Add(prefork, false, "Enables use of the SO_REUSEPORT socket option. This will spawn multiple Go processes listening on the same port. learn more about socket sharding.")
	giconfig.Add(serverHeader, "", "Enables the Server HTTP header with the given value.")
	giconfig.Add(strictRouting, false, "When enabled, the router treats /foo and /foo/ as different. Otherwise, the router treats /foo and /foo/ as the same.")
	giconfig.Add(caseSensitive, false, "When enabled, /Foo and /foo are different routes. When disabled, /Fooand /foo are treated the same.")
	giconfig.Add(immutable, false, "When enabled, all values returned by context methods are immutable. By default, they are valid until you return from the handler; see issue #185.")
	giconfig.Add(unescapePath, false, "Converts all encoded characters in the route back before setting the path for the context, so that the routing can also work with URL encoded special characters")
	giconfig.Add(ETag, false, "Enable or disable ETag header generation, since both weak and strong etags are generated using the same hashing method (CRC-32). Weak ETags are the default when enabled.")
	giconfig.Add(bodyLimit, 4*1024*1024, "Sets the maximum allowed size for a request body, if the size exceeds the configured limit, it sends 413 - Request Entity Too Large response.")
	giconfig.Add(concurrency, 256*1024, "Maximum number of concurrent connections.")
	giconfig.Add(readTimeout, "0s", "The amount of time allowed to read the full request, including the body. The default timeout is unlimited.")
	giconfig.Add(writeTimeout, "0s", "The maximum duration before timing out writes of the response. The default timeout is unlimited.")
	giconfig.Add(idleTimeout, "0s", "The maximum amount of time to wait for the next request when keep-alive is enabled. If IdleTimeout is zero, the value of ReadTimeout is used.")
	giconfig.Add(readBufferSize, 4096, "per-connection buffer size for requests' reading. This also limits the maximum header size. Increase this buffer if your clients send multi-KB RequestURIs and/or multi-KB headers (for example, BIG cookies).")
	giconfig.Add(writeBufferSize, 4096, "Per-connection buffer size for responses' writing.")
	giconfig.Add(compressedFileSuffix, ".fiber.gz", "Adds a suffix to the original file name and tries saving the resulting compressed file under the new file name.")
	giconfig.Add(proxyHeader, "", "This will enable c.IP() to return the value of the given header key. By default c.IP()will return the Remote IP from the TCP connection, this property can be useful if you are behind a load balancer e.g. X-Forwarded-*.")
	giconfig.Add(GETOnly, false, "Rejects all non-GET requests if set to true. This option is useful as anti-DoS protection for servers accepting only GET requests. The request size is limited by ReadBufferSize if GETOnly is set.")
	giconfig.Add(reduceMemoryUsage, false, "Aggressively reduces memory usage at the cost of higher CPU usage if set to true")
	giconfig.Add(network, fiber.NetworkTCP4, "Known networks are \"tcp\", \"tcp4\" (IPv4-only), \"tcp6\" (IPv6-only)")
	giconfig.Add(disableKeepalive, false, "Disable keep-alive connections, the Server will close incoming connections after sending the first response to the client")
	giconfig.Add(disableDefaultDate, false, "When set to true causes the default date header to be excluded from the response.")
	giconfig.Add(disableDefaultContentType, false, "When set to true, causes the default Content-Type header to be excluded from the Response.")
	giconfig.Add(disableHeaderNormalizing, false, "By default all header names are normalized: conteNT-tYPE -> Content-Type")
	giconfig.Add(disableStartupMessage, false, "When set to true, it will not print out debug information")
}
