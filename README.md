# GoTofuEnv

This project have been merged with [tenv](https://github.com/tofuutils/tenv), and the work (improvements and bugfixes) will continue on the [tenv](https://github.com/tofuutils/tenv) repository.

[OpenTofu](https://opentofu.org) version manager (inspired by [tofuenv](https://github.com/tofuutils/tofuenv), written in Go)

Support [Terraform](https://www.terraform.io/) too (see [here](#terraform-support)).

Handle [Semver 2.0.0](https://semver.org/) with [go-version](https://github.com/hashicorp/go-version) and use the [HCL](https://github.com/hashicorp/hcl) parser to extract required version constraint from OpenTofu files.

GoTofuEnv can use [cosign](https://github.com/sigstore/cosign) (if present) to check OpenTofu signature or fallback to [PGP](https://www.openpgp.org/about) using [gopenpgp](https://github.com/ProtonMail/gopenpgp) implementation. However, unstable OpenTofu versions are signed only with cosign (in this case, if cosign is not found GoTofuEnv will display a warning).

## Installation

### Automatic

Install via [Homebrew](https://brew.sh/) (migration from dvaumoron/tap and renaming to [tenv](https://github.com/tofuutils/tenv) are in place)

```console
brew tap tofuutils/tap
brew install gotofuenv
```

You can enable cosign check with :

```console
brew install cosign
```

### Manual

Get the last packaged binaries (use .deb, .rpm, .apk or .zip) found [here](https://github.com/tofuutils/gotofuenv/releases).

For the .zip case, the unzipped folder must be added to your PATH.

## Usage

### tofu

This project version of `tofu` command is a proxy to OpenTofu `tofu` command  managed by `tenv`. The default resolution strategy is latest-allowed (without [TOFUENV_TOFU_VERSION](#tofuenv_tofu_version) environment variable or [`.opentofu-version`](#opentofu-version-file) file).

### terraform

This project version of `terraform` command is a proxy to Hashicorp `terraform` command  managed by `tenv`. The default resolution strategy is latest-allowed (without [TFENV_TERRAFORM_VERSION](#tfenv_terraform_version) environment variable or `.terraform-version` file).

### tenv install [version]

Install a requested version of OpenTofu (into TOFUENV_ROOT directory from TOFUENV_REMOTE url).

Without parameter the version to use is resolved automatically via TOFUENV_TOFU_VERSION or [`.opentofu-version`](#opentofu-version-file) files
(searched in working directory, user home directory and TOFUENV_ROOT directory).
Use "latest-stable" when none are found.

If a parameter is passed, available options:

- an exact [Semver 2.0.0](https://semver.org/) version string to install
- a [version constraint](https://opentofu.org/docs/language/expressions/version-constraints) string (checked against available at TOFUENV_REMOTE url)
- latest or latest-stable (checked against available at TOFUENV_REMOTE url)
- latest-allowed or min-required to scan your OpenTofu files to detect which version is maximally allowed or minimally required.

See [required_version](https://opentofu.org/docs/language/settings#specifying-a-required-opentofu-version) docs.

```console
tenv install 1.6.0-beta5
tenv install "~> 1.6.0"
tenv install latest
tenv install latest-stable
tenv install latest-allowed
tenv install min-required
```

### Environment Variables

GoTofuEnv commands support the following environment variables.

#### TOFUENV_AUTO_INSTALL (alias TFENV_AUTO_INSTALL)

String (Default: true)

If set to true gotofuenv will automatically install missing OpenTofu version needed (fallback to latest-allowed strategy when no [`.opentofu-version`](#opentofu-version-file) files are found).

`tenv` subcommands `detect` and `use` support a `--no-install`, `-n` disabling flag version.

Example: use 1.6.0-rc1 version that is not installed, and auto installation is disabled. (-v flag is equivalent to `TOFUENV_VERBOSE=true`)

```console
$ TOFUENV_AUTO_INSTALL=false tenv use -v 1.6.0-rc1
Write 1.6.0-rc1 in /home/dvaumoron/.gotofuenv/.opentofu-version
```

Example: use 1.6.0-rc1 version that is not installed, and auto installation stay enabled.

```console
$ tenv use -v 1.6.0-rc1
Installation of OpenTofu 1.6.0-rc1
Write 1.6.0-rc1 in /home/dvaumoron/.gotofuenv/.opentofu-version
```

#### TOFUENV_FORCE_REMOTE (alias TFENV_FORCE_REMOTE)

String (Default: false)

If set to true gotofuenv detection of needed version will skip local check and verify compatibiliy on remote list.

`tenv` subcommands `detect` and `use` support a `--force-remote`, `-f` flag version.

#### TOFUENV_GITHUB_TOKEN

String (Default: "")

Allow to specify a GitHub token to increase [GitHub Rate limits for the REST API](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api). Useful because OpenTofu binares are downloaded from the OpenTofu GitHub repository.

`tenv` subcommands `detect`, `install`, `list-remote` and `use` support a `--github-token`, `-t` flag version.

#### TOFUENV_OPENTOFU_PGP_KEY

String (Default: "")

Allow to specify a local file path to OpenTofu PGP public key, if not present download https://get.opentofu.org/opentofu.asc.

`tenv` subcommands `detect`, `ìnstall` and `use` support a `--key-file`, `-k` flag version.

#### TFENV_HASHICORP_PGP_KEY

String (Default: "")

Allow to specify a local file path to Hashicorp PGP public key, if not present download https://www.hashicorp.com/.well-known/pgp-key.txt.

`tenv tf` subcommands `detect`, `ìnstall` and `use` support a `--key-file`, `-k` flag version.

#### TOFUENV_REMOTE

String (Default: https://api.github.com/repos/opentofu/opentofu/releases)

To install OpenTofu from a remote other than the default (must comply with [Github REST API](https://docs.github.com/en/rest?apiVersion=2022-11-28))

`tenv` subcommands `detect`, `install`, `list-remote` and `use` support a `--remote-url`, `-u` flag version.

#### TFENV_REMOTE

String (Default: https://releases.hashicorp.com/terraform)

To install Terraform from a remote other than the default (must comply with [Hashicorp Release API](https://releases.hashicorp.com/docs/api/v1))

`tenv tf` subcommands `detect`, `install`, `list-remote` and `use` support a `--remote-url`, `-u` flag version.

#### TOFUENV_ROOT (alias TFENV_ROOT)

Path (Default: `$HOME/.gotofuenv`)

The path to a directory where the local OpenTofu versions, Terraform versions and GoTofuEnv configuration files exist.

`tenv` support a `--root-path`, `-r` flag version.

#### TOFUENV_TOFU_VERSION

String (Default: "")

If not empty string, this variable overrides OpenTofu version, specified in [`.opentofu-version`](#opentofu-version-file) files.
`tenv` subcommands `install` and `detect` also respects this variable.

e.g. with :

```console
$ tofu version
OpenTofu v1.6.0
on linux_amd64
```

then :

```console
$ TOFUENV_TOFU_VERSION=1.6.0-rc1 tofu version
OpenTofu v1.6.0-rc1
on linux_amd64
```

#### TFENV_TERRAFORM_VERSION

String (Default: "")

If not empty string, this variable overrides Terraform version, specified in `.terraform-version` files.

`tenv tf` subcommands `install` and `detect` also respects this variable.

#### TOFUENV_VERBOSE (alias TFENV_VERBOSE)

String (Default: false)

Active the verbose display of gotofuenv.

`tenv` support a `--verbose`, `-v` flag version.

### tenv use version

Switch the default OpenTofu version to use (set in [`.opentofu-version`](#opentofu-version-file) file in TOFUENV_ROOT).

`tenv use` has a `--working-dir`, `-w` flag to write [`.opentofu-version`](#opentofu-version-file) file in working directory.

Available parameter options:

- an exact [Semver 2.0.0](https://semver.org/) version string to use
- a [version constraint](https://opentofu.org/docs/language/expressions/version-constraints) string (checked against available in TOFUENV_ROOT directory)
- latest or latest-stable (checked against available in TOFUENV_ROOT directory)
- latest-allowed or min-required to scan your OpenTofu files to detect which version is maximally allowed or minimally required.

See [required_version](https://opentofu.org/docs/language/settings#specifying-a-required-opentofu-version) docs.

```console
tenv use min-required
tenv use v1.6.0-beta5
tenv use latest
tenv use latest-allowed
```

### tenv detect

Detect the used version of OpenTofu for the working directory.

```console
$ tenv detect
OpenTofu 1.6.0 will be run from this directory.
```

### tenv reset

Reset used version of OpenTofu (remove .opentofu-version file from TOFUENV_ROOT).

```console
tenv reset
```

### tenv uninstall version

Uninstall a specific version of OpenTofu (remove it from TOFUENV_ROOT directory without interpretation).

```console
tenv uninstall v1.6.0-alpha4
```

### tenv list

List installed OpenTofu versions (located in TOFUENV_ROOT directory), sorted in ascending version order.

`tenv list` has a `--descending`, `-d` flag to sort in descending order.

```console
$ tenv list
  1.6.0-rc1 
* 1.6.0 (set by /home/dvaumoron/.gotofuenv/.opentofu-version)
```

### tenv list-remote

List installable OpenTofu versions (from TOFUENV_REMOTE url), sorted in ascending version order.

`tenv list-remote` has a `--descending`, `-d` flag to sort in descending order.

`tenv list-remote` has a `--stable`, `-s` flag to display only stable version.

```console
$ tenv list-remote
1.6.0-alpha1
1.6.0-alpha2
1.6.0-alpha3
1.6.0-alpha4
1.6.0-alpha5
1.6.0-beta1
1.6.0-beta2
1.6.0-beta3
1.6.0-beta4
1.6.0-beta5
1.6.0-rc1 (installed)
1.6.0 (installed)
```

### tenv help [command]

Help about any command.

You can use `--help` `-h` flag instead.

```console
$ tenv help tf detect
Display Terraform current version.

Usage:
  tenv tf detect [flags]

Flags:
  -f, --force-remote        force search on versions available at TFENV_REMOTE url
  -h, --help                help for detect
  -k, --key-file string     local path to PGP public key file (replace check against remote one)
  -n, --no-install          disable installation of missing version
  -u, --remote-url string   remote url to install from (default "https://releases.hashicorp.com/terraform")

Global Flags:
  -r, --root-path string   local path to install versions of OpenTofu and Terraform (default "/home/dvaumoron/.gotofuenv")
  -v, --verbose            verbose output
```

```console
$ tenv use -h
Switch the default OpenTofu version to use (set in .opentofu-version file in TOFUENV_ROOT)

Available parameter options:
- an exact Semver 2.0.0 version string to use
- a version constraint string (checked against version available in TOFUENV_ROOT directory)
- latest or latest-stable (checked against version available in TOFUENV_ROOT directory)
- latest-allowed or min-required to scan your OpenTofu files to detect which version is maximally allowed or minimally required.

Usage:
  tenv use version [flags]

Flags:
  -f, --force-remote          force search on versions available at TOFUENV_REMOTE url
  -t, --github-token string   GitHub token (increases GitHub REST API rate limits)
  -h, --help                  help for use
  -k, --key-file string       local path to PGP public key file (replace check against remote one)
  -n, --no-install            disable installation of missing version
  -u, --remote-url string     remote url to install from (default "https://api.github.com/repos/opentofu/opentofu/releases")
  -w, --working-dir           create .opentofu-version file in working directory

Global Flags:
  -r, --root-path string   local path to install versions of OpenTofu and Terraform (default "/home/dvaumoron/.gotofuenv")
  -v, --verbose            verbose output
```

## .opentofu-version file

If you put a `.opentofu-version` file  in working directory, user home directory or TOFUENV_ROOT directory, gotofuenv detects it and uses the version written in it.
Note, that TOFUENV_TOFU_VERSION can be used to override version specified by `.opentofu-version` file.

Recognized value (same as `tenv use` command) :

- an exact [Semver 2.0.0](https://semver.org/) version string to use
- a [version constraint](https://opentofu.org/docs/language/expressions/version-constraints) string (checked against available in TOFUENV_ROOT directory)
- latest or latest-stable (checked against available in TOFUENV_ROOT directory)
- latest-allowed or min-required to scan your OpenTofu files to detect which version is maximally allowed or minimally required.

See [required_version](https://opentofu.org/docs/language/settings#specifying-a-required-opentofu-version) docs.

## Terraform support

GoTofuEnv rely on `.terraform-version` files, [TFENV_HASHICORP_PGP_KEY](#tfenv_hashicorp_pgp_key), [TFENV_REMOTE](#tfenv_remote) and [TFENV_TERRAFORM_VERSION](#tfenv_terraform_version) specifically to manage Terraform versions.

`tenv tf` have the same managing subcommands for Terraform versions (`detect`, `install`, `list`, `list-remote`, `reset`, `uninstall` and `use`).

GoTofuEnv check Terraform PGP signature (there is no cosign signature available).

## LICENSE

The GoTofuEnv project is released under the Apache 2.0 license. See [LICENSE](LICENSE).
