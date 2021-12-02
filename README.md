<div id="top"></div>

<!-- PROJECT LOGO -->
<br />
<div align="center">

<h1 align="center">Tai - Eng easy dictionary API</h3>

  <p align="center">
    **API** for Tai to English and English to Tai in easy words.
    <br />
    <br />
    <a href="https://tai-eng-dictionaryapi.herokuapp.com/api/v1/api_key=NO8p3FC4qMrTzx1RUjRXNXWrqlLa8DkDjmRgt7s9rDE=/eng/apple">View Demo</a>
    ·
    <a href="https://github.com/NoerNova/tai-eng-dictionaryAPI/issues">Report Bug</a>
    ·
    <a href="https://github.com/NoerNova/tai-eng-dictionaryAPI/issues">Request Feature</a>
  </p>
</div>


<!-- ABOUT THE PROJECT -->
## About The Project
```json
// 20211202233858
// https://tai-eng-dictionaryapi.herokuapp.com/api/v1/api_key=NO8p3FC4qMrTzx1RUjRXNXWrqlLa8DkDjmRgt7s9rDE=/eng/apple

{
  "data": [
    {
      "id": 1687,
      "english": "apple",
      "shan": "မၢၵ်ႇၵႅမ်ႈၶွင်ႇ",
      "Antonym": {
        "word_id": 1687,
        "english": "",
        "shan": ""
      },
      "Definition": {
        "word_id": 1687,
        "english": "a round fruit with firm white flesh and a green, red or yellow skin ",
        "shan": "မၢၵ်ႇၵႅမ်ႈၶွင်ႇ ၊ မၢၵ်ႇၶႃႈၶုင်ႉ ။",
        "pos": "Noun",
        "uncount": "( prefix )"
      }
    }
  ]
}
```

<br />
I started this project on my personal purpose, I used the Tai dictionary a lot, to learn and write in Shan words, previously I use taidictionary.com but for some reason, they are out of service now so I start this project.
I also develop a website using this api [check it out](https://taidictionary.noernova.com).


This is a very easy Dictionary with a small database ,thanks to [https://github.com/saitawngpha](https://github.com/saitawngpha) for dictionary database, it's may not cover all words and some of the words may not quite well be translated. if you have another big or good database its could be very please, provide your own API, folk this project or direct them to me, anyway might be helpful for everyone.

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Golang](https://go.dev/)
* [Gin web framework](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io/index.html)

<p align="right">(<a href="#top">back to top</a>)</p>


## Dictionary website

* Project source Chack [https://github.com/NoerNova/tai-eng-dictionary](https://github.com/NoerNova/tai-eng-dictionary) for web project build with Nextjs.
* Website [https://taidictionary.noernova.com](https://taidictionary.noernova.com)



<!-- GETTING STARTED -->
## Getting Started

**This project currently does not include the Database file due to a third party license.**

### Prerequisites

Check out [Golang](https://go.dev/learn/) website first in case you are not familiar with Golang yet.

### Installation

1. Clone this project 
    ```sh
    go get https://github.com/NoerNova/tai-eng-dictionaryAPI.git
    ```
2. ```sh
    cd /tai-eng-dictionaryAPI
	
    # API Endpoint http://127.0.0.1:8080/api/v1/api_key=${API_KEY}/eng/:text
    # API Endpoint  http://127.0.0.1:8080/api/v1/api_key=${API_KEY}/shn/:text
    ```



<p align="right">(<a href="#top">back to top</a>)</p>


## Structure
```sh
.
├── _test
│   ├── database
│   │   └── tai-eng.db
│   └── endpoints_test.go
├── controllers
│   └── translate.go
├── database
│   └── tai-eng.db
├── favicon.ico
├── go.mod
├── go.sum
├── main.go
└── models
    ├── setup.go
    └── word.go
```

<!-- USAGE EXAMPLES -->
## Usage

This dictionary can translate in Tai to English and English to Shan
```sh
# API Endpoint http://127.0.0.1:8080/api/v1/api_key=${API_KEY}/eng/:text
# API Endpoint  http://127.0.0.1:8080/api/v1/api_key=${API_KEY}/shn/:text
```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/NoerNova/tai-eng-dictionaryAPI/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



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

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the [GNU](https://www.gnu.org/licenses/gpl-3.0.txt) License.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CREDIT -->
## Credit

* Database by: Sai TawngPha [https://github.com/saitawngpha]()

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

NorHsangPha - [@noer_nova](https://twitter.com/noer_nova) - noernova666@gmail.com

Project Link: [https://github.com/NoerNova/tai-eng-dictionaryAPI](https://github.com/NoerNova/tai-eng-dictionaryAPI)

<p align="right">(<a href="#top">back to top</a>)</p>
