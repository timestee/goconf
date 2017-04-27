package goconf

// Values are resolved with the following priorities (lowest to highest):
// 1. Options struct default value
// 2. Flags default value
// 3. Config file value, TOML or JSON file
// 4. Command line flag
