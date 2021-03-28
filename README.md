# captcha

:smile: captcha

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/captcha)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/captcha/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/captcha.svg)](https://github.com/moul/captcha/releases)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/captcha.svg)](https://microbadger.com/images/moul/captcha)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

[![Go](https://github.com/moul/captcha/workflows/Go/badge.svg)](https://github.com/moul/captcha/actions?query=workflow%3AGo)
[![Release](https://github.com/moul/captcha/workflows/Release/badge.svg)](https://github.com/moul/captcha/actions?query=workflow%3ARelease)
[![PR](https://github.com/moul/captcha/workflows/PR/badge.svg)](https://github.com/moul/captcha/actions?query=workflow%3APR)
[![GolangCI](https://golangci.com/badges/github.com/moul/captcha.svg)](https://golangci.com/r/github.com/moul/captcha)
[![codecov](https://codecov.io/gh/moul/captcha/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/captcha)
[![Go Report Card](https://goreportcard.com/badge/moul.io/captcha)](https://goreportcard.com/report/moul.io/captcha)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/captcha/badge)](https://www.codefactor.io/repository/github/moul/captcha)

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/moul/captcha)

## Usage

[embedmd]:# (.tmp/usage.txt console)
```console
foo@bar:~$ captcha -engine=math
1 + 3
->
foo@bar:~$ captcha -engine=math
5 + 3
->
foo@bar:~$ captcha -engine=banner
 _                _
| |_   _ _  _  _ | |__ _ __
| ' \ | '_|| || || / /| '_ \
|_||_||_|   \_,_||_\_\| .__/
                      |_|
->
foo@bar:~$ captcha -engine=banner
                  _
 __  ___ __ __ __| |__  ___
/ _|/ _ \\ V  V /| '_ \/ -_)
\__|\___/ \_/\_/ |_.__/\___|
->
```

## Install

### Using go

```sh
go get moul.io/captcha
```

### Releases

See https://github.com/moul/captcha/releases

## Contribute

![Contribute <3](https://raw.githubusercontent.com/moul/moul/master/contribute.gif)

I really welcome contributions.
Your input is the most precious material.
I'm well aware of that and I thank you in advance.
Everyone is encouraged to look at what they can do on their own scale;
no effort is too small.

Everything on contribution is sum up here: [CONTRIBUTING.md](./CONTRIBUTING.md)

### Contributors ✨

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg)](#contributors)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://manfred.life"><img src="https://avatars1.githubusercontent.com/u/94029?v=4" width="100px;" alt=""/><br /><sub><b>Manfred Touron</b></sub></a><br /><a href="#maintenance-moul" title="Maintenance">🚧</a> <a href="https://github.com/moul/captcha/commits?author=moul" title="Documentation">📖</a> <a href="https://github.com/moul/captcha/commits?author=moul" title="Tests">⚠️</a> <a href="https://github.com/moul/captcha/commits?author=moul" title="Code">💻</a></td>
    <td align="center"><a href="https://manfred.life/moul-bot"><img src="https://avatars1.githubusercontent.com/u/41326314?v=4" width="100px;" alt=""/><br /><sub><b>moul-bot</b></sub></a><br /><a href="#maintenance-moul-bot" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors)
specification. Contributions of any kind welcome!

### Stargazers over time

[![Stargazers over time](https://starchart.cc/moul/captcha.svg)](https://starchart.cc/moul/captcha)

## License

© 2021   [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT)
([`LICENSE-MIT`](LICENSE-MIT)), at your option.
See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
