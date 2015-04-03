# gocube

**gocube** will be a pure Go library for solving the 3x3x3 Rubik's cube and various parts of it.

# Design considerations

I want this code to be efficient. However, I don't claim to know the best way to solve the cube computationally. Thus, I will make this project as modular as possible so that it can be improved easily (unlike [some people's code](https://github.com/lgarron/shuang-chen-projects/blob/dee5de0485d20b6f7759e11b5aef248d9e0f2dda/min2phase-java/src/CoordCube.java#L225)).

# TODO

I am currently working towards a functional two-phase solver which uses Kociemba's algorithm. Here is what must be done:

 * Implement phase-2 solver
   * Phase2Cube which stores corner+edge+slice permutations.
   * Conversion from CubieCube to Phase2Cube (in a given axis)
     * Make sure encoding permutations works
   * Phase2Heuristic
   * Phase2Solver
 * Finish two-phase solver as TwoPhaseSolver class.
   * (Possibly) attempt pre-scramble and inverse solutions.

# License

**gocube** is licensed under the BSD 2-clause license. See [LICENSE](LICENSE).

```
Copyright (c) 2015, Alex Nichol.
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer. 
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
