# Box

![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/FrostyTheSouthernSnowman/Box?include_prereleases)
![GitHub last commit](https://img.shields.io/github/last-commit/FrostyTheSouthernSnowman/Box)
![GitHub issues](https://img.shields.io/github/issues-raw/FrostyTheSouthernSnowman/Box)
![GitHub pull requests](https://img.shields.io/github/issues-pr/FrostyTheSouthernSnowman/Box)
![GitHub](https://img.shields.io/github/license/FrostyTheSouthernSnowman/Box)

Box is a tool for building scalable microservices from predefined templates.
Box is currently in Beta so if you find any issues or have some ideas for features, please open an issue.
If you are using Box, please star us on Github. Any support is highly appreciated.

# Table of contents
- [What is Box?](#Box)
- [TOC](#table-of-contents)
- [Installation](#installation-and-usage)
- [Development](#development)
- [Contribute](#contribute)
    - [Adding new features or fixing bugs (see Contributing.md)](#adding-new-features-or-fixing-bugs)
- [License](#license)

# Installation and Usage
[(Back to top)](#table-of-contents)


To install Box, clone the repo and run `go build .` to build a binaray or `go run .` to run without genarating binaries. Which ever way you chose, when you run the package, you must supply a path for Box to check for a Box project.

# Development
[(Back to top)](#table-of-contents)

To work on Box, clone the project or fork it. Once you have the source code on your machine, run `go mod tidy` to install all dependencies locally.
If you want to test your work, you can do `cd <directory-with-code-you-want-to-test>` and then `go test` to run the tests. You can also call `go run . example` after deleting the Box-genarated python file to see if it compiles the example configurations correctly.

# Contribute
[(Back to top)](#table-of-contents)

To contribute to Box, clone or fork this repo and follow the instructions above in [Development](#development). Once you have a working version of your changes, open a pull request and your code will be review, edited if any bugs are found, and merged into the main code base.

# License
This project is licenced under MIT Open Source License.

If you like this project, please star it on github and recomend it to people you know.
