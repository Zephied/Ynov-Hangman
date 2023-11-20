# Hangman Classic

Hangman Classic is a text-based version of the classic word guessing game. The game reads from a file to get a list of words, and then randomly selects one for the player to guess.

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/hangman_classic.git
```

## Usage
Navigate to the project directory and run the game with the following command:
    
 ```bash
go run main.go [wordlist.txt] 
```

Replace [wordlist.txt] with the path to your word list file and you can add an [ascii.txt] to use the ascii art version of the game.

to load a saved game, run the following command:

```bash
go run main.go [wordlist.txt] --startWith [savefile.json]
```

replace [savefile.json] with the path to your save file.

## Features
Text-based hangman game
Reads words from a file
Saves and loads game state