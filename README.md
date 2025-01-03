<br />
<div align="center">
  <a href="https://github.com/lhoogenraad/shelfd">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>
<a id="readme-top"></a>

<h3 align="center">Letterbookd</h3>

  <p align="center">
    A book review system.
    <br />
    <a href="https://github.com/lhoogenraad/letterbookd"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/lhoogenraad/letterbookd">View Demo</a>
    ·
    <a href="https://github.com/lhoogenraad/letterbookd/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/lhoogenraad/letterbookd/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>


### Built With

* [![Next][Next.js]][Next-url]
* [![React][React.js]][React-url]
* [![Golang][Golang.com]][Golang-url]
* [![MySQL][MySQL.com]][MySQL-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Letterbookd

### Description

This application serves as a book review system. I built this as a passion project, and a
way to learn new programming langs/frameworks/dbms' listed below.
I've used [openlibrary](https://openlibrary.org/) to populate the database with all of the book information.


### Feature List

- Login/Signup.
- Search books within Letterbookd.
- Add books to readlist.
- Set status of readlist items (unread, reading, read)
- Create rating/review of a book.
- Add comments to reviews.
- Like reviews.
- View other user's reviews.
- View other user's review comments.
- Featured books.
- Popular reviews (in the past week)


## Getting Started

### Prerequisites

This project uses:
- npm
- next.js
- Golang
- MySQL
- Docker


* To run the frontend server:
  ```sh
  npm install; npm run dev;
  ```

* To run the backend sever:
    ```sh
    go get .; go run server/cmd/api/main.go
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Leon Hoogenraad - le.o.n@outlook.com

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/lhoogenraad/letterbookd.svg?style=for-the-badge
[contributors-url]: https://github.com/lhoogenraad/letterbookd/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/lhoogenraad/letterbookd.svg?style=for-the-badge
[forks-url]: https://github.com/lhoogenraad/letterbookd/network/members
[stars-shield]: https://img.shields.io/github/stars/lhoogenraad/letterbookd.svg?style=for-the-badge
[stars-url]: https://github.com/lhoogenraad/letterbookd/stargazers
[issues-shield]: https://img.shields.io/github/issues/lhoogenraad/letterbookd.svg?style=for-the-badge
[issues-url]: https://github.com/lhoogenraad/letterbookd/issues
[license-shield]: https://img.shields.io/github/license/lhoogenraad/letterbookd.svg?style=for-the-badge
[license-url]: https://github.com/lhoogenraad/letterbookd/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 
[Golang.com]: https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge
[Golang-url]: https://go.dev
[MySQL.com]: https://shields.io/badge/MySQL-lightgrey?logo=mysql&style=plastic&logoColor=white&labelColor=blue
[MySQL-url]: https://www.mysql.com/



# Extras

## Server deployment commands

- ssh shelfd@ssh-shelfd.alwaysdata.net
- scp -s ./main shelfd@ssh-shelfd.alwaysdata.net:/home/shelfd/www/main
