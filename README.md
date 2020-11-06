# Pizzabot

## Robot pizza delivery

Section                           | Description
----------------------------------|------------
[Overview](#overview)             | Overview of problem and solution
[Challenge](#challenge)           | Description of task and solution
[Implementation](#implementation) | Implementation of solution
[Execution](#execution)           | Instructions on how to execute program

### Overview

This is a simple program to plot the course of a pizza delivery bot using sets of [x,y] co-ordinates,
where `x` is a lateral position along the x-axis of a grid, and `y` is the vertical position on the
y-axis.

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
 
- `5x5 (0, 0) (1, 3) (4,4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)`

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

#### Code

The code is implemented and tested in Go.  There is a main `pizzabot` package, along with a `utils`
package to handle parsing of the input arguments, and sorting to optimise the route.  As efficiency
and optimisation are not requirements of this test, I have only implemented a simple sort for the
`x` co-ordinates.  I have added a method that could order the y-axis, but not implemented the code.
Each method should encapsulate only the functionality required of it, and could be replaced with a
more efficient implementation if required.  But using the method signatures and return values allows
the code to stay decoupled from the caller.

#### Unit Tests

I have added unit tests to all the utils methods, covering extreme cases.

### Execution

To execute the `pizzabot` program, place the executable in your current directory and run the 
following command from a terminal on either Linux or Mac:

`./pizzabot "args"`

where _args_ are the arguments to pass to the program in the form:
`grid_size (co-ordingate_1) (co-ordingate_n)`.  For example:

`./pizzabot "5x5 (1, 3) (4, 4)"`

`./pizzabot "5x5 (0, 0) (1, 3) (4,4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"`
