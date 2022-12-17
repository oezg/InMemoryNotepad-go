## In-Memory Notepad
A simple in-memory database that supports CRUD operations. I worked with slices, loops, string formatting, standard IO and the bufio scanner.
### About
Every day we face a shower of information that pours on us from everywhere. It's easy to forget the most important things in that flow. This simple notepad isn't able to store the data on the hard drive; it is one of the simplest, yet a practical program.
### Stage 1/4: Basic CLI program
Create a CLI program that can read commands, print them, and exit.
### Stage 2/4: Core functionality
Implement basic commands that allow you to create, list and clear the notes.
### Stage 3/4: Dynamic storage
Implement the dynamic storage option for notes and few other improvements.
### Stage 4/4: Working with the notes
Implement the commands to update and delete particular notes.
## Objectives
When the command is `create`:

* The program should interpret the data as a type string and make a new note that contains this text.

When the command is `clear`:

* The program should delete all records.

When the command is `list`:

* The program should print all entries.

When the command is `update`:

* The program should use a space (`" "`) as a separator to split the data into two pieces – the position (an integer) and the text (a string). Update the text of a note at the given position.

When the command is `delete`:

* The program should interpret the data as a position, delete a note at the given position, and clear this position. Be careful to properly update the remaining positions: the position of a note should represent the order of the notes.

When the command is `exit`:

* The program should run in an infinite loop and terminate only when the `exit` command is received.
