# Pizzabot

## Robot pizza delivery

Section                                              | Description
-----------------------------------------------------|------------
[Overview](#overview)                                | Overview of problem and solution
[Challenge](#challenge)                              | Description of task and solution
[Implementation](#implementation)                    | Implementation of solution
[Building](#building)                                | Building the program
[Execution Instructions](#execution-instructions)    | Instructions on how to execute program

### Overview

This is a simple program to plot the course of a pizza delivery bot using sets of [x,y] co-ordinates,
where `x` is a lateral position along the x-axis of a grid, corresponding to East & West, and `y` 
is the vertical position on the y-axis, corresponding to North & South.

___

### Challenge

#### Task
Your task is to instruct Pizzabot on how to deliver pizzas to all the houses in a neighborhood.
In more specific terms, given a grid (where each point on the grid is one house) and a list of
points representing houses in need of pizza delivery, return a list of instructions for getting
Pizzabot to those locations and delivering. An instruction is one of:

- N: Move north 
- S: Move south 
- E: Move east 
- W: Move west 
- D: Drop pizza

Pizzabot always starts at the origin point, (0, 0). As with a Cartesian plane, this point lies 
at the most south- westerly point of the grid. Therefore, given the following input string: 5x5 (1, 3) (4, 4)
one correct solution would be: `ENNNDEEEND`  In other words: move east once and north thrice; drop a pizza; move east thrice and north once; 
drop a final pizza.

#### Solution

Your solution should have a `pizzabot` command, executable from the command line, 
i.e. `./pizzabot "5x5 (1, 3) (4, 4)"` 

There are multiple correct ways to navigate between locations. We do not take optimality of route 
into account when grading: all correct solutions are good solutions.

To complete the challenge, please solve for the following exact input string:
 
&nbsp; &nbsp; `5x5 (0, 0) (1, 3) (4,4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)`

#### Notes

1. Please submit the solution to your challenge as a tarball, with clear instructions on how to 
execute it.

2. When the test is anonymously reviewed by a panel of your potential peers, here's whatwe're looking for:
   - Correctness:
     - Does the code fulfill all the requirements of the challenge?
   - Production Readiness:
     - Is the code well-structured by the standards of the host language?
     - Is the solution maintainable and easy to make changes to?
     - Is the code clean, readable and easy for other team members to understand? 
     - Is there appropriate test coverage?
   - Fit and polish:
     - Is there a README? A build script?
     - Are there spelling errors or extraneous comments? • How does it handle unspecified input?

___

### Implementation

#### Assumptions

1. One main assumption I have made is that the co-ordinates must be 'all-or-nothing'.  So if there is
an error with any co-ordinates, or any are invalid, the program will exit.  The assumption here is
that we want the same number of deliveries as co-ordinates input to the program.  If we chose to skip
invalid ones, the number of deliveries would not equal the number of co-ordinates input to the program.

2. Nothing was mentioned in the requirements as to whether or not the delivery co-ordinates should be
allowed to extend beyond the size of the grid.  However, as the grid is explicitly defined in the argument
string, I am working under the assumption that we should not go outside this grid.  One option in
the case of a co-ordinate lying outside the grid would be to ignore that location.  However, as
mentioned above, this would cause a mismatch between the number of co-ordinates entered, and the number
of deliveries.  I have therefore chosen to flag the input as invalid.

3. Another assumption is that all co-ordinates from the origin co-ordinates, `(0,0)`, will be positive.
So if we pass a negative co-ordinate, for example `(-1, -3)`, the program will exit with an error
message. 

#### Code

The code is implemented and tested in Go.  There is a main `cmd` package, along with a `utils`
package to handle parsing of the input arguments, and sorting to optimise the route.  As efficiency
and optimisation are not requirements of this test, I have only implemented a simple sort for the
`x` co-ordinates.

Each method should encapsulate only the functionality required of it, allowing for a level of 
abstraction between the caller and the implementation.  This will make testing and maintenance
easier.

#### Unit Tests

I have added unit tests to all files and methods.  See below on how to run tests and coverage.

---

### Building

Use `make` to clean, build and test the project.

To create the binary, run `make build`.  The Makefile will place the binary in the `./bin` folder.
To run tests or produce a test coverage report, run the appropriate `make` commands and open the 
resulting files.

> clean &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; Clean build files  
> build &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; Build the executable  
> test &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; Run all tests and output to test.out  
> cover &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; Generate test coverage report coverage.html  
> run &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; Run the executable  

---

### Execution Instructions

To execute the `pizzabot` program from the command line, change directory to the `bin` folder on
either Linux or Mac and run the following command:

`./pizzabot "args"`

where _args_ are the arguments to pass to the program in the form:
`grid_size (co-ordingate_1) ... (co-ordingate_n)`.  For example:

`./pizzabot "5x5 (1, 3) (4, 4)"`

`./pizzabot "5x5 (0, 0) (1, 3) (4,4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"`

**Notes:** 
- there must be a space between the grid size and the subsequent co-ordinates
- no co-ordinates should be greater than the corresponding x or y grid size
- if the program is run using `build run`, the arguments must be passed using `ARGS=`, and spaces
and parenthesis must be escaped.  For example: `make ARGS="5x5\ \(1, 3\)\ \(4, 4\)" run`
