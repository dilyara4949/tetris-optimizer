
--



## The program must :

* Compile successfully\
* Assemble all of the tetrominoes in order to create the smallest square possible.\
* Identify each tetromino in the solution by printing them with uppercase latin letters (A for the first one, B for the second, etc)\
* Expect at least one tetromino in the text file\
* In case of bad format on the tetrominoes or bad file format it should print ERROR


#### Example of a text File

#...\
#...\
#...\
#...

....\
....\
..##\
..##

If it isn't possible to form a complete square, the program should leave spaces between the tetrominoes. 
#### For example:
ABB.\
ABB.\
A...\
A...

$ cat -e sample.txt\
...#$\
...#$\
...#$\
...#$\
$\
....$\
....$\
....$\
####$\
$\
.###$\
...#$\
....$\
....$\
$\
....$\
..##$\
.##.$\
....$\
$\
....$\
.##.$\
.##.$\
....$\
$\
....$\
....$\
##..$\
.##.$\
$\
##..$\
.#..$\
.#..$\
....$\
$\
....$\
###.$\
.#..$\
....$

$ go run . sample.txt | cat -e\
ABBBB.$\
ACCCEE$\
AFFCEE$\
A.FFGG$\
HHHDDG$\
.HDD.G$\
$