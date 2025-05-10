# Changelog

## [1.2.0](https://github.com/zisuu/golang-playground/compare/v1.1.0...v1.2.0) (2025-05-10)


### Features

* implement dictionary package with search, add, update, and delete functionalities; add tests and update workflows ([#30](https://github.com/zisuu/golang-playground/issues/30)) ([2a0157f](https://github.com/zisuu/golang-playground/commit/2a0157f3b20beec97fea1990a998695f64a96ae4))

## [1.1.0](https://github.com/zisuu/golang-playground/compare/v1.0.0...v1.1.0) (2025-05-09)


### Features

* implement dictionary package with search, add, update, and dele functionalities; add tests for dictionary operations ([#25](https://github.com/zisuu/golang-playground/issues/25)) ([712da1c](https://github.com/zisuu/golang-playground/commit/712da1c18b6396fe21e75295509018f3bb9910dc))

## 1.0.0 (2025-05-09)


### Features

* add 'context' module ([#6](https://github.com/zisuu/golang-playground/issues/6)) ([d637211](https://github.com/zisuu/golang-playground/commit/d637211de09efada8f11d15a8965009d99a04aa4))
* add 'countdown' module ([#3](https://github.com/zisuu/golang-playground/issues/3)) ([72f6638](https://github.com/zisuu/golang-playground/commit/72f66385b5b96b22ee874211582808be116bbb0b))
* add 'di' module ([#2](https://github.com/zisuu/golang-playground/issues/2)) ([79f45f2](https://github.com/zisuu/golang-playground/commit/79f45f2ca037a84724c515cf4036701cb09e58cb))
* add 'numeral' module ([#7](https://github.com/zisuu/golang-playground/issues/7)) ([1bbb89c](https://github.com/zisuu/golang-playground/commit/1bbb89cb00f6e94f76d00b0813fb31a7eea87e25))
* add dictionary package with search, add, update, and delete functionalities ([b9f4533](https://github.com/zisuu/golang-playground/commit/b9f4533f9ed24e0fc04e611e7bf0c82acdad34cb))
* add GoReleaser configuration for building and releasing binaries ([7c415a2](https://github.com/zisuu/golang-playground/commit/7c415a24b4d2e27bdec8f77445dcfd43f001da46))
* add initial GoReleaser configuration for building and releasing binaries ([de02626](https://github.com/zisuu/golang-playground/commit/de026261811d3552b87ba480273220c195e597f2))
* add some more modules ([#4](https://github.com/zisuu/golang-playground/issues/4)) ([5009f2d](https://github.com/zisuu/golang-playground/commit/5009f2d8e86f3c5116682ac4391990e8c298d1df))
* add some more modules ([#8](https://github.com/zisuu/golang-playground/issues/8)) ([c80db03](https://github.com/zisuu/golang-playground/commit/c80db037ae72dc08577a52d93c77d4d1f41ca92e))
* add sync module ([#5](https://github.com/zisuu/golang-playground/issues/5)) ([163a56e](https://github.com/zisuu/golang-playground/commit/163a56ea44574d73c4e9b9e39830a665b10444db))
* implement Racer functionality with tests for server speed comparison ([ed556b1](https://github.com/zisuu/golang-playground/commit/ed556b10c8498069ca1bac8c6a20a2b3a8bf80b0))


### Bug Fixes

* add CHANGELOG.md to markdownlintignore ([c490f7c](https://github.com/zisuu/golang-playground/commit/c490f7c446eb5fbcb761057f87d600a2cef77b28))
* add GoReleaser configuration and update GitHub Actions workflow to use app token for releases ([62ea109](https://github.com/zisuu/golang-playground/commit/62ea109b3114533aac231e84d9da9e4e4e9d7950))
* correct output file naming and source file for Go binaries in release workflow ([58810b6](https://github.com/zisuu/golang-playground/commit/58810b60929e73ebe5bf7029424a35279ad93529))
* **deps:** update module github.com/approvals/go-approval-tests to v1.4.0 ([#9](https://github.com/zisuu/golang-playground/issues/9)) ([6a65a25](https://github.com/zisuu/golang-playground/commit/6a65a25df772b9216ad39f09c584e4a8537a839e))
* **deps:** update module github.com/approvals/go-approval-tests to v1.5.0 ([#11](https://github.com/zisuu/golang-playground/issues/11)) ([014c8dc](https://github.com/zisuu/golang-playground/commit/014c8dcb8e0bceced3c088dc99ea119c4f528131))
* remove unnecessary command from release action configuration ([cf5f65f](https://github.com/zisuu/golang-playground/commit/cf5f65fb2d986e12a9d560a3e214bc6e34352c60))
* remove version file and related versioning logic from main application ([e99c259](https://github.com/zisuu/golang-playground/commit/e99c25914023da6102f0e84d59de8559064db2c9))
* restructure push-release workflow to separate versioning and release jobs ([07f8a60](https://github.com/zisuu/golang-playground/commit/07f8a60a9c20f9f56773f1411d5da2df5cbe2944))
* update archive format to zip in GoReleaser configuration ([a89ce24](https://github.com/zisuu/golang-playground/commit/a89ce24ce56b0d9b56e4ea586be617ba3cc552e0))
* update build command to use correct Go source file for binaries ([ebed78a](https://github.com/zisuu/golang-playground/commit/ebed78aeff62cdbb4cda1421885294a1f1be0a20))
* update GoReleaser version to v2.9.0 in push-release workflow ([126646f](https://github.com/zisuu/golang-playground/commit/126646fca29e62d3209938b8abcee20ea949024f))
* update module path in go.mod and streamline GitHub Actions workflow ([9b9dcd6](https://github.com/zisuu/golang-playground/commit/9b9dcd6372b359f44bbde72b84f7f2c5c513ba53))
* update permissions and add Go module validation steps in push-release workflow ([76d9b86](https://github.com/zisuu/golang-playground/commit/76d9b8629f6b911ec9c00e1720c0573cbc816414))
* update release workflow and add version information to main application ([1fbb66d](https://github.com/zisuu/golang-playground/commit/1fbb66dc50127c470a4717826db216e8124ac1be))
* update release workflow to correctly handle new release outputs and conditions ([603b04e](https://github.com/zisuu/golang-playground/commit/603b04ed386f79e172505150d225641367dbbb69))
* update versioning job condition to exclude workflow_dispatch and adjust release job condition ([0fa46ec](https://github.com/zisuu/golang-playground/commit/0fa46ecfddaac13c5a6177edbf38ddae6b4dc169))
* update versioning job condition to include workflow_dispatch and specific commit messages ([8d6fc20](https://github.com/zisuu/golang-playground/commit/8d6fc20812fbb729bd9336e775bebabd81bd9c26))
* update versioning job outputs and streamline release process ([f0cb7c7](https://github.com/zisuu/golang-playground/commit/f0cb7c779abaaedd415da05a81891d52a16bde0f))
* update workflow to use the correct directory for Go module validation and GoReleaser ([91b0f99](https://github.com/zisuu/golang-playground/commit/91b0f9995dd0c3106a32ad40ca81f353468475d8))
