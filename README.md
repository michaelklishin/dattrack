# What is DatTrack

DatTrack is a small command line tool for [Digitally Imported](http://di.fm) fans.

It displays recent tracks from a particular channel in the shell.

DatTrack was developed as a small code kata excercise and should not be taken seriously.


## Installation

Assuming you have `~/bin` in `PATH`, from the repository root:

    go build
    cp dattrack ~/bin/


## Usage

    dattrack [--channel NAME --limit LIMIT]

Channel names are currently not parsed in any intelligent way so you have to know them:

 * Epic Trance: `et`, `epictrance`
 * Liquid DnB: `ldnb`, `liquiddnb`
 * Club Dubstep: `clubdubstep`, `clubds`
 * Electro House: `electrohouse`
 * Hands Up: `handsup`, `hu`
 * Glitch Hop: `glitchhup`, `gh`

This will be improved and expanded in future versions.


## License

Released under the BSD license.

Copyright Michael S. Klishin, 2012
