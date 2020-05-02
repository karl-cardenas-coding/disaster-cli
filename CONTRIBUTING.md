# Contributing to Disaster-CLI

1. Open an issue if one does not exist.
2. Simply fork the repo and get started!
3. Have a working Go [environment](https://golang.org/doc/install)
```
git clone https://github.com/karl-cardenas-coding/disaster-cli.git
cd disaster-cli
go build
```
4. Submit merge request. Reference the issue number (if applicable)


## CI/CD 

All CI/CD is powered by Github Actions. See the `.github/workflows/` to view the templates. 

* Linting: [golangci-lint](https://github.com/golangci/golangci-lint)
* SAST: [LGTM](https://lgtm.com/help/lgtm/about-lgtm)
* Dependecy Scan: [dependabot](https://dependabot.com/)