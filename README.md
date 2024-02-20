# Quine-McCluskey method of minimisation

The Quine–McCluskey algorithm (QMC), also known as the method of prime implicants, is a method used for minimization of Boolean functions that was developed by [Willard V. Quine](https://en.wikipedia.org/wiki/Willard_Van_Orman_Quine) in 1952 and enhanced by [Edward J. McCluskey](https://en.wikipedia.org/wiki/Edward_J._McCluskey) in 1956.

The Quine–McCluskey algorithm is functionally identical to [Karnaugh mapping](https://en.wikipedia.org/wiki/Karnaugh_map), but the tabular form makes it more efficient for use in computer algorithms, and it also gives a deterministic way to check that the minimal form of a Boolean function has been reached. It is sometimes referred to as the tabulation method.

The steps involved:

1. Expand the minterms to binary equivalent.
2. Index the minterms into groups of equal numbered 1s.
3. Go through a progressive reduction process to assimilate minterms to generate prime implicants
4. Tabulate our findings in Step 3 to form our minimised function

Let's look at an example

Consider simplifying the expression

```
F = Σ m(0,1,3,4,5,6)
```

## Step 1: expand our minterms to binary

```
m₀ = 000
m₁ = 001
m₂ = 011
m₃ = 100
m₄ = 101
m₅ = 110
```

## Step 2: index our expansions in groups of equal numbered ones

```
Group 0 → 000

Group 1 → 001, 100

Group 2 → 011, 101, 110
```

## Step 3: progressive reduction to assimilate minterms to prime implicants

this step requires a bit of explanation:

### a. First, express the groups in order of increasing ones

```
+-------+---------+--------+
| index | minterm | binary |
+-------+---------+--------+
| Grp 0 |  0      | 000    |
+--------------------------+
| Grp 1 |  1      | 001    |
|       |  4      | 100    |
+--------------------------+
| Grp 2 |  3      | 011    |
|       |  5      | 101    |
|       |  6      | 110    |
+-------+---------+--------+
```

### b. Reduction begins by looking for commonality between terms in neigbour index groups. Commonality sought is a shift by only logic 1.

We systematically check each binary equivalent with members in neigbour index group and **_tag & pair_** those that shift by only 1 value.



+-------+---------+--------+---------+
|               START                |
+-------+---------+--------+---------+
| index | minterm | binary |  check  |
+-------+---------+--------+----+----+
| Grp 0 |  0      | 000    | ✓✓ |    |
+------------------------------------+
| Grp 1 |  1      | 001    | ✓  | ✓✓ |
|       |  4      | 100    | ✓  | ✓✓ |
+------------------------------------+
| Grp 2 |  3      | 011    |    | ✓  |
|       |  5      | 101    |    | ✓✓ |
|       |  6      | 110    |    | ✓  |
+-------+---------+--------+----+----+

|     1ST REDUCTION      |
|:----------:|:---------:|
|  minterm   |   binary  |

|    0, 1    |   00x     |
|    0, 4    |   x00     |
|    1, 3    |   0x1     |
|    1, 5    |   x01     |
|    4, 5    |   10x     |
|    4, 6    |   1x0     |
