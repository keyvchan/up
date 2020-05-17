package config

// Template for darwin
var Template = `
[
	{
		"Kind": "PackageManager",
		"Name": "HomeBrew",
		"Command": "brew",
		"Args": [
			"upgrade"
		]
	},
	{
		"Kind": "PackageManager",
		"Name": "HomeBrew Cask",
		"Command": "brew",
		"Args": [
			"cask",
			"upgrade"
		]
	},
	{
		"Kind": "PackageManager",
		"Name": "Yarn",
		"Command": "yarn",
		"Args": [
			"global",
			"upgrade"
		]
	},
	{
		"Kind": "Toolchain",
		"Name": "Rust",
		"Command": "rustup",
		"Args": [
			"update"
		]
	},
	{
		"Kind": "System",
		"Name": "MacOS",
		"Command": "softwareupdate",
		"Args": [
			"--install"
		]
	}
]
`
