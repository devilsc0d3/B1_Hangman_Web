# Hangman Web

****

**Rules :**

A word is randomly selected by the website. The guessing player is given only 10 chances to correctly guess the word.
The guessing player guesses letters, one at a time, by typing them into a designated field.
Incorrect guesses are indicated by displaying the letters already used.

For each incorrect letter, a part of a stick figure is drawn and -1 life will be subtracted to player and For each incorrect word, -2 life will be subtracted to player.
If the stick figure is fully drawn before the word is guessed, the guessing player loses and the game is over.
If the word is guessed before the stick figure is fully drawn, the guessing player wins and get the point corresponding to the remaining chance.

****

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

****

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

****

**Update coming...**

- Add news languages
- translated word
- Add a functionalities :
- Create a save
- Page more responsive, adapt of mobile
- Add sound

****

**Credit :**

Fauré Léo