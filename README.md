<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
<h3 align="center">go-terminator</h3>
  <p align="center">
    Handle graceful termination of your program
    <br />
    <a href="https://github.com/github_username/repo_name/issues">Report Bug</a>
    ·
    <a href="https://github.com/github_username/repo_name/issues">Request Feature</a>
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About The Project

A simple object to manage graceful termination of your program. Works with context, and supports adding
termination tasks from multiple threads.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

### Installation

```bash
go get github.com/zrbecker/go-terminator
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

```go
func main() {
	ctx, t := terminator.WithTerminator(context.Background())
	defer t.Terminate()

	terminator.MustFromContext(ctx).AddTerminationTask(func() {
		fmt.Println("Performing cleanup...")
	})

	fmt.Println("Application active. Press Ctrl+C to stop or wait 10 seconds for automatic shutdown.")
	time.Sleep(10 * time.Second)
}
```

_For more examples, please refer to the [Examples](https://github.com/zrbecker/go-terminator/tree/main/examples)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Zachary Becker - [@zrbecker](https://twitter.com/zrbecker) - zrbecker@gmail.com

Project Link: [https://github.com/zrbecker/go-terminator](https://github.com/zrbecker/go-terminator)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/zrbecker/go-terminator.svg?style=for-the-badge
[contributors-url]: https://github.com/zrbecker/go-terminator/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/zrbecker/go-terminator.svg?style=for-the-badge
[forks-url]: https://github.com/zrbecker/go-terminator/network/members
[stars-shield]: https://img.shields.io/github/stars/zrbecker/go-terminator.svg?style=for-the-badge
[stars-url]: https://github.com/zrbecker/go-terminator/stargazers
[issues-shield]: https://img.shields.io/github/issues/zrbecker/go-terminator.svg?style=for-the-badge
[issues-url]: https://github.com/zrbecker/go-terminator/issues
[license-shield]: https://img.shields.io/github/license/zrbecker/go-terminator.svg?style=for-the-badge
[license-url]: https://github.com/zrbecker/go-terminator/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/zrbecker
