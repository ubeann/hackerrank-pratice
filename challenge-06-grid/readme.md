# Grid Challenge

Given a square grid of characters in the range $\text{ascii[a-z]}$, rearrange elements of each row alphabetically, ascending. Determine if the columns are also in ascending alphabetical order, top to bottom. Return `YES` if they are or `NO` if they are not.

## Example

$grid = [\text{abc, ade, efg}]$

The grid is illustrated below:

```bash
a b c
a d e
e f g
```

The rows are already in alphabetical order. The columns `a a e`, `b d f` and `c e g` are also in alphabetical order, so the answer would be `YES`. Only elements within the same row can be rearranged. They cannot be moved to a different row.

## Function Description

Complete the function `gridChallenge` in the editor below.

`gridChallenge` has the following parameter(s):

- `string grid[n]`: an array of strings

## Returns

- `string`: either `YES` or `NO`

## Input Format

The first line contains $t$, the number of testcases.

Each of the next $t$ sets of lines are described as follows:

- The first line contains $n$, the number of rows in the grid.
- The next $n$ lines each contain a string of length $n$.

## Constraints

- $1 \le t \le 10$
- $1 \le n \le 100$
- Each string consists of lowercase letters in the range $\text{ascii[a-z]}$

## Output Format

For each test case, on a separate line print `YES` if it is possible to rearrange the grid alphabetically ascending in both its rows and columns, or `NO` otherwise.

## Sample Input

```bash
STDIN   Function
-----   --------
1       t = 1
5       n = 5
ebacd   grid = ['ebacd', 'fghij', 'olmkn', 'trpqs', 'xywuv']
fghij
olmkn
trpqs
xywuv
```

## Sample Output

```bash
YES
```

## Explanation

The $5 \times 5$ grid in the  test case can be reordered to

```bash
a b c d e
f g h i j
k l m n o
p q r s t
u v w x y
```

This fulfills the condition since the rows 1, 2, ..., 5 and the columns 1, 2, ..., 5 are all alphabetically sorted.
