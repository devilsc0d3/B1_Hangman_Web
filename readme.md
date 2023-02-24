# Hangman Web

---

### instruction : ###

Create a program that will take a file as parameter . Create a file which contains a bunch of words with which the program will play. Each word is separated with a newline.hangman-webwords.txtwords.txt

The behavior of the game is the same as the hangman project, refers to it for more details.

In this project you will need to implement at least the following endpoints:

GET : Sends HTML response - the main page, it will basically display your interface.
1.1. GET Tip: go templates to receive and display data from the server/

POST : that sends data to the Golang server (the letter you want to find)
2.1. POST Tip: use form and other types of tags to make the post request. The form must redirect to /hangman/hangman

The main page must have at least:

A text representing the word to reveal.
A text input

A button which sends a POST request to and outputs the result on page./hangman

### Bonus : level ###
Hangman-level consists in adding difficulties to your website : You need to implement a dedicated page at the start of your game that allows the user to choose between the following difficulties

EASY: is file words.txt.
NORMAL: is file words2 .txt.
HARD is: file words3.txt.

### Bonus : JOSE ###

Hangman-bring-to-death consists in making a stick-man, representing the attempts you have left. 

Just like in the real hangman game. Each time you suggest a wrong letter or a wrong word, the stick-man grows,
until the scheme is done, and the man completely hanging.

### Bonus : Gamify ###

Hangman-gamify consists in making your interface more game oriented.

You can implement other routes to help you have other interfaces representing the state of your game For example:

A startup page asking for you UserName and a button to start
A losing or/and winning screen
A button somewhere to allow the user to play again.
A scoreboard
A list of all letters suggested during the game.
Those are only examples, you are free to implement any features you like that can make the project look better, and feel like a real game.

### Rules : ###

A word is randomly selected by the website. The guessing player is given only 10 chances to correctly guess the word.
The guessing player guesses letters, one at a time, by typing them into a designated field.
Incorrect guesses are indicated by displaying the letters already used.

For each incorrect letter, a part of a stick figure is drawn and -1 life will be subtracted to player and For each incorrect word, -2 life will be subtracted to player.
If the stick figure is fully drawn before the word is guessed, the guessing player loses and the game is over.
If the word is guessed before the stick figure is fully drawn, the guessing player wins and get the point corresponding to the remaining chance.

---

**Features :**

- version 1.0 available !!!
- Page on the theme PAC-MAN : menu, win, lose, hangman, setting, 404, rules, scoreboard, ranking board
- Language :
- English
- French
- Spanish
- German
- you can choose a level & a username !
- A hangman (PacMan move step by step around the ghost)
- Scoreboard sort !
- Array of letter & word already entered

---

**Installation :**

- Clone the repo with :

```bash
git clone https://ytrack.learn.ynov.com/git/fleo/hangman-web/src/branch/devphav2
```

or download and extract the repo : [click here !](https://ytrack.learn.ynov.com/git/fleo/hangman-web/src/branch/devphav2)

- Start the Hangman-Web with :

```bash
cd go_file
```

and this command :

```bash
go run .
```

or run this command with "words.txt", "words2.txt" or "words3.txt" for play at the first version :

```bash
go run . words.txt
```

- Enjoy :)

---

**Update coming...**

- translated word ___(coming soon)___
- Create a save
- Page more responsive, adapt of mobile
- Add sound

---

**Credit :**

Fauré Léo
