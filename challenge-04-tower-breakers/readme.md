# Tower Breakers

Two players are playing a game of Tower Breakers! Player $1$ always moves first, and both players always play optimally.The rules of the game are as follows:

- Initially there are $n$ towers.
- Each tower is of height $m$.
- The players move in alternating turns.
- In each turn, a player can choose a tower of height $x$ and reduce its height to $y$, where $1 \le y < x$ and $y$ [evenly divides](https://en.wiktionary.org/wiki/evenly_divisible) $x$.
- If the current player is unable to make a move, they lose the game.

Given the values of $n$ and $m$, determine which player will win. If the first player wins, return `1`. Otherwise, return `2`.

## Example

$$n = 2$$
$$m = 6$$

There are $n = 2$ towers, each $m = 6$ high. Player $1$ has two options:

- Remove $3$ pieces from tower $1$ to leave $3$ as $6 \mod 3 = 0$.
- Remove $5$ pieces from tower $2$ to leave $1$ as $6 \mod 5 = 1$.

Let Player $1$ remove $3$. Now the towers are $[3, 6]$.

Player $2$ matches the move. Now the towers are $[3, 3]$.

Now Player $1$ has only one move:

Player $1$ removes $2$ pieces from tower $1$ to leave $1$. Now the towers are $[1, 3]$.

Player $2$ matches again. Now the towers are $[1, 1]$.

Player $1$ has no move and loses. Return $2$.

## Function Description

Complete the `towerBreakers` function in the editor below.

`towerBreakers` has the following parameter(s):

- `int n`: the number of towers
- `int m`: the height of each tower

## Returns

- `int`: the winner of the game

## Input Format

The first line contains a single integer, $t$, the number of test cases.

Each of the next $t$ lines describes a test case in the form of $2$ space-separated integers, $n$ and $m$.

## Constraints

- $1 \le t \le 100$
- $1 \le n, m \le 10^6$

## Sample Input

```bash
STDIN   Function
-----   --------
2       t = 2
2 2     n = 2, m = 2
1 4     n = 1, m = 4
```

## Sample Output

```bash
2
1
```

## Explanation

We'll refer to player $1$ as $P_1$ and player $2$ as $P_2$.

In the first test case, $P_1$ chooses one of the two towers and reduces it to $1$. Then $P_2$ reduces the remaining tower to $1$. As both towers now have height $1$, $P_1$ is unable to make a move so $P_2$ wins and we return $2$.

In the second test case, there is only one tower of height $4$. $P_1$ can reduce it to a height of either $1$ or $2$. $P_1$ chooses $1$ as both players always move optimally. Because $P_2$ has no possible move, $P_1$ wins and we return $1$.
