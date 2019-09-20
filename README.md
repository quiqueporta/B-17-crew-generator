# B-17 Crew Generator
![](https://cf.geekdo-images.com/itemrep/img/jbIHDfAO6sCRE8QDzpxlg4M3KsA=/fit-in/246x300/pic44389.jpg)

https://boardgamegeek.com/boardgame/1032/b-17-queen-skies

A `go` script to generate the B-17 Queen of the skies crew, a great solo board game.

The crew names are randomly generated from this page [randomuser.me](https://randomuser.me)

You can download the `print a play` version from the [Felipe Santamaría webpage](https://sites.google.com/site/felisan88/b-17queenoftheskies).

# Usage
```go
go run main.go
```
This command generates a `B-17_crew.pdf` file with the B-17 crew.

Also you can run a HTTP server to generate the PDF.
```go
go run http/main.go
```
Then you can access to http://localhost:8080 to get the PDF.



# Result example

https://raw.githubusercontent.com/quiqueporta/B-17-crew-generator/master/B-17_crew.pdf

![](https://raw.githubusercontent.com/quiqueporta/B-17-crew-generator/master/example.jpg)

