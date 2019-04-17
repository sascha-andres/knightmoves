# knight moves

A small program to get moves for a knight from one position to another.

Rather brute force and not optimized as it traverses through a tree of moves using a recursive function.

## Install

    go get -u livingit.de/code/knightmoves/cmd/knightmoves

Make sure go modules are activated.

## Usage

    knightmoves
    
      Flags: 
           --version  Displays the program version string.
        -h --help  Displays help with available flag, subcommand, and positional value parameters.
        -i --initial  initial position
        -t --target  target position
        -d --depth  maximum depth to analyse
        -v --verbose  change log level

Example:

    knightmoves --initial A1 --target B1 --depth 10

    INFO[0000] initial := [A1] target := [B1], maxDepth := 10, verbose := false  func=NewGame package=knightmoves
    INFO[0000] received possible solution: A1,C2,E3,G4,E5,G6,E7,D5,C3,B1  func=SetSolution package=knightmoves
    INFO[0000] received possible solution: A1,C2,E3,G4,E5,C4,A3,B1  func=SetSolution package=knightmoves
    INFO[0000] received possible solution: A1,C2,E3,C4,A3,B1  func=SetSolution package=knightmoves
    INFO[0000] received possible solution: A1,C2,A3,B1       func=SetSolution package=knightmoves
    4 moves:
        1: A1
        2: C2
        3: A3
        4: B1