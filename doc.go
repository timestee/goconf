/*
Package gonconf provides a simple conf loader support multiple sources.

Read configuration automatically based on the given struct's field name.
Load configuration from multiple sources
multiple file inherit

Values are resolved with the following priorities (lowest to highest):
1. Options struct default value
2. Flags default value
3. Config file value, TOML or JSON file
4. Command line flag

*/

package goconf
