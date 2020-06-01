# wup

Making your system always up to date.

## Installation

1. Clone the repo.

```
git clone https://github.com/keyvchan/up
```

2. Install it into your `$GOBIN`.

```
cd up/
go install ./cmd/up
```

3. If you have set `$GOBIN` to your `$PATH`, you can simply use `up` to trigger the update process.

## Congiguration

`up` is shipped with a default configuration. The configuration is made up by an array of json. It looks like below:

```
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
```

### Config

#### Kind

The `Kind` in the configuration above have three types:

- PackageManager
- Toolchain
- System

It will appare in the indicator of update items.

#### Name

Same as `Kind`, shown right after `Kind` in the indicator.

#### Command

The command of upgrade items. It should place in `$PATH`. `up` will lookup your `PATH`.

#### Args

The `Args` is an array of string, it will pass to `Command`.
