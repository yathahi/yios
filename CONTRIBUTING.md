# Contributing to Yathahi

Thank you for your interest in contributing to **Yathahi**!

Contributions of all kinds are welcome, including bug fixes, feature implementations, documentation improvements, testing, technical discussions, and ideas.

Yathahi is an evolving open-source project. Whether you're making a small correction or proposing a significant improvement, your contribution is appreciated.

## 1. Before Contributing

Before making changes, take some time to understand the project and its current direction.

We recommend reviewing the following documents:

- [Project Vision](docs/PROJECT_VISION_v0.0.1-alpha.md) — Understand the project's goals and intended direction.
- [Architecture](docs/ARCHITECTURE_v0.0.1-alpha.md) — Understand the architectural components and their responsibilities.
- [Specification](docs/SPECIFICATION_v0.0.1-alpha.md) — Understand the intended behaviour and requirements.
- [Yathahi Engineering Standards (YES)](docs/YES.v0.0.1-alpha.md) — Understand the engineering standards and development conventions.

Contributions should respect the project's current architecture, specifications, and engineering standards.

For significant changes to architecture, component responsibilities, or project direction, please open a discussion or issue before beginning implementation. This helps avoid unnecessary work and keeps development coherent.

## 2. Licensing and Contributor Recognition

Yathahi is distributed under the **Apache License 2.0**. Contributors are not required to sign a Contributor License Agreement (CLA) or transfer ownership of their copyrights to the project.

- **Contribution Licensing:** By submitting a contribution, you agree that it will be made available under the project's Apache License 2.0, subject to its applicable terms.
- **Copyright Ownership:** You retain ownership of the copyright in your original contributions. No copyright assignment is required.
- **Patent Rights:** Patent rights and grants relating to contributions are governed by the Apache License 2.0.
- **Contributor Recognition:** Contributors are credited through Git commit history. Additional recognition may be provided through release notes and the project's `CONTRIBUTORS.md` file.

For the complete licensing terms, refer to the [LICENSE](LICENSE) file.

## 3. How to Submit a Contribution

### 3.1. Fork the Repository

Create your own fork of the Yathahi repository using the **Fork** button on GitHub.

### 3.2. Clone Your Fork

Clone your fork to your local machine:

```bash
git clone https://github.com/<your-username>/yathahi.git
cd yathahi
```

### 3.3. Create a Branch

Create a dedicated branch for your contribution. Avoid making changes directly on `main`.

For example:

```bash
git checkout -b feature/your-feature-name
```

Use a descriptive branch name that reflects the purpose of your changes.

### 3.4. Make Your Changes

When implementing changes:

- Follow the current Yathahi Engineering Standards (YES).
- Respect the responsibilities and boundaries of the existing architectural components.
- Keep changes focused and avoid unrelated modifications.
- Include appropriate tests where applicable.
- Update relevant documentation when behaviour, requirements, or interfaces change.
- Avoid introducing unnecessary dependencies.

### 3.5. Format and Test

Before committing, format your Go code and run the available tests:

```bash
go fmt ./...
go test ./...
```

Where supported, you may also run the race detector:

```bash
go test -race ./...
```

If a test cannot be executed in your environment, mention that when submitting your pull request.

### 3.6. Commit Your Changes

Stage and commit your changes:

```bash
git add .
git commit -m "Describe your change clearly"
```

Use a Git email address associated with your GitHub account if you want your commits to be attributed to your profile.

### 3.7. Push Your Branch

Push your contribution branch to your fork:

```bash
git push origin feature/your-feature-name
```

### 3.8. Open a Pull Request

Open your fork on GitHub and select **Contribute → Open pull request**.

Submit the pull request against the `main` branch of the original `yathahi/yios` repository.

Your pull request should include:

- A clear description of the changes.
- The reason for making those changes.
- Relevant testing information.
- References to related issues or discussions, if applicable.

Pull requests will be reviewed by project maintainers. Changes may be discussed or revised before they are accepted and merged.

## 4. Keeping Your Fork Up to Date

Before starting new work, make sure your fork is synchronized with the latest changes from the main repository.

You can use GitHub's **Sync fork → Update branch** option on your fork's repository page.

After synchronizing, update your local branch:

```bash
git checkout main
git pull origin main
```

If you are working on an existing contribution branch, update it against the latest `main` branch before continuing.

## 5. Reporting Bugs and Suggesting Features

Bug reports, feature requests, and technical suggestions are welcome.

When reporting a bug, provide as much relevant information as possible, including:

- A clear description of the issue.
- Steps to reproduce it.
- Expected behaviour.
- Actual behaviour.
- Relevant logs, error messages, or system information.

For feature suggestions, explain the problem you're trying to solve and how the proposed feature could benefit Yathahi.

Please check existing issues and discussions before creating a new one to avoid duplicates.

## 6. Security Vulnerabilities

Please do not disclose security vulnerabilities through public issues or discussions.

Refer to [SECURITY.md](SECURITY.md) for instructions on responsibly reporting security-related issues.

## 7. Project Development Status

Yathahi is currently in its early alpha stage.

Its architecture, specifications, engineering standards, interfaces, and implementation are actively evolving. Existing designs and requirements may be revised, expanded, replaced, or removed as development progresses.

Contributors should refer to the latest project documentation before starting work and understand that changes may introduce breaking modifications.

This flexibility is intentional and part of the project's current development stage.

## 8. Final Notes

Yathahi is built through open collaboration, technical exploration, and continuous improvement.

Contributions are not limited to code. Documentation, testing, discussions, constructive criticism, and new ideas are equally valuable to the project's development.

Thank you for helping build Yathahi.

**Unity through interoperability.**
