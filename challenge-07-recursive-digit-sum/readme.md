# Recursive Digit Sum

We define super digit of an integer $x$ using the following rules:

- If $x$ has only $1$ digit, then its super digit is $x$.
- Otherwise, the super digit of $x$ is equal to the super digit of the digit-sum of $x$.

For example, the super digit of $9875$ will be calculated as:

```bash
super_digit(9875)   9+8+7+5 = 29
super_digit(29)     2 + 9 = 11
super_digit(11)     1 + 1 = 2
super_digit(2)      = 2
```

## Example

$n = 9875$

$k = 4$

The number $p$ is created by concatenating the string $n$ $k$ times so the value of $p$ is $9875987598759875$.

```bash
super_digit(P) = super_digit(9875987598759875)
               = super_digit(9+8+7+5+9+8+7+5+9+8+7+5+9+8+7+5)
               = super_digit(116)
               = super_digit(1+1+6)
               = super_digit(8)
               = 8.
```

All of the digits of $p$ sum to $116$. The digits of $116$ sum to $8$. $8$ is only one digit, so it's the super digit.

## Function Description

Complete the function `superDigit` in the editor below. It must return the calculated super digit as an integer.

`superDigit` has the following parameter(s):

- `n`: a string representation of an integer.
- `k`: an integer, the times to concatenate $n$ to make $p$.

## Returns

- `int`: the super digit of $p$.

## Input Format

The first line contains two space separated integers, $n$ and $k$.

## Constraints

- $1 \le n \le 10^{100000}$
- $1 \le k \le 10^{5}$
