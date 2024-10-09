# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- New `ctllm sync` command to synchronize the configuration with the current project structure.
- Functionality to add new project types and ignore patterns during sync.
- Option to re-initialize the project with `ctllm init --force`.
- Interactive prompts to add the output directory to `.gitignore`.

### Changed

- Moved original scripts to a dedicated `scripts` folder.
- Updated help text to include information about the new `sync` command.
- Replaced deprecated `ioutil` functions with `os` package equivalents.

### Fixed

- Improved error handling and user feedback during initialization and configuration.
- Resolved duplicate function declarations causing build errors.

## [0.1.0] - 2024-10-09

### Added

- Initial release of `ctllm`.
- Core functionality for processing code projects for LLMs.
- Recursive file collection with custom ignore patterns.
- Generation of project tree structure and file chunks.
- Interactive project initialization with `ctllm init`.

[Unreleased]: https://github.com/baudevs/code-to-llm/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/baudevs/code-to-llm/releases/tag/v0.1.0
