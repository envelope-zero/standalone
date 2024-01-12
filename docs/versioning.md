# Versioning

This project uses [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html). This means that the version has three parts: `major.minor.patch`.
Releases are fully automated and happen on every _feature_ and _bug fix_ that is merged into the `main` branch.

The versions increase as follows:

- The **major** version increases when there are breaking changes, meaning that the behavior of the software changes compared to an earlier version (unless that behavior was wrong, then it's a **patch** version increase). Please check the [upgrading documentation](docs/upgrading.md) before updating major versions!
- The **minor** version increases when there are new features
- The **patch** version increases when bugs are fixed.
- If the Envelope Zero backend or frontend are updated, the version bump is the same as for the dependency that is updated
- If a release with only dependency updates is made, it bumps the `PATCH` version.

Whenever a version increases, all numbers to the right of it are reset to 0.

The following things are looked at for versioning (called “public API” in Semantic Versioning):

- Location of the data on your computer
- Behavior of the application, e.g. how budget values are calculated, including new features
