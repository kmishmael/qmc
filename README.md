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

```

+-------+---------+--------+---------+      +------------------------+
|               START                |      |     1ST REDUCTION      |
+-------+---------+--------+---------+      +------------+-----------+
| index | minterm | binary |  check  |      |  minterm   |   binary  |
+-------+---------+--------+---------+      +------------+-----------+
| Grp 0 |  0      | 000    |  True   |      |    0, 1    |   00x     |
+------------------------------------+      |    0, 4    |   x00     |
| Grp 1 |  1      | 001    |  True   |      +------------------------+
|       |  4      | 100    |  True   |      |    1, 3    |   0x1     |
+------------------------------------+      |    1, 5    |   x01     |
| Grp 2 |  3      | 011    |  True   |      |    4, 5    |   10x     |
|       |  5      | 101    |  True   |      |    4, 6    |   1x0     |
|       |  6      | 110    |  True   |      +------------+-----------+
+-------+---------+--------+----+----+

```

### c. Repeat the process using the reduced table

```
+------------------------+---------+            +-----------------------+
|          1ST REDUCTION           |            |      2ND REDUCTION    |
+------------+-----------+---------+            +------------+----------+
|  minterm   |   binary  |  check  |            |  minterm   |  binary  |
+------------+-----------+---------+            +------------+----------+
|    0, 1    |   00x     |  True   |            | 0, 1, 4, 5 |          |
|    0, 4    |   x00     |  True   |            |            |    x0x   |
+------------------------+---------+            | 0, 4, 1, 5 |          |
|    1, 3    |   0x1     |  False  |            +------------+----------+
|    1, 5    |   x01     |  True   |
|    4, 5    |   10x     |  True   |
|    4, 6    |   1x0     |  False  |
+------------+-----------+---------+
```

### d.

At this second reduction, we should note two things:

1. Not all the minterm groups in the 1st reduction could form new bigger pairs ie. 0x1, 1x0
2. In the 2nd reduction table only one term exist i.e x0x and we cannot reduce further.

All these terms from 1 and 2, are our **prime implicants**

Further, you can now proceed to determine the essential, and if any, non-essential
prime implicants for our simplified expression.

## How it works


```bash
kmishmael@KIBET-PC:~$ go run main.go
Enter maxterms separated by commas: 0,1,3,4,5,6

INITIAL BOOLEAN TABLE
|-----|-------|---------|
| KEY | VALUE | MATCHED |
|-----|-------|---------|
| 0   | 000   | true    |
| 1   | 001   | true    |
| 4   | 100   | true    |
| 3   | 011   | true    |
| 5   | 101   | true    |
| 6   | 110   | true    |
|-----|-------|---------|

REDUCTION 1
|------|-------|---------|
| KEY  | VALUE | MATCHED |
|------|-------|---------|
| 0, 1 | 00x   | true    |
| 0, 4 | x00   | true    |
| 1, 3 | 0x1   | false   |
| 1, 5 | x01   | true    |
| 4, 5 | 10x   | true    |
| 4, 6 | 1x0   | false   |
|------|-------|---------|

FINAL REDUCTION 2
|------------|-------|---------|
|    KEY     | VALUE | MATCHED |
|------------|-------|---------|
| 0, 1, 4, 5 | x0x   | false   |
| 0, 4, 1, 5 | x0x   | false   |
|------------|-------|---------|

Prime implicants:
=> 0x1
=> 1x0
=> x0x

```

Since QMC methods handles minterms and maxterms similary, it can there work for both terms. The interpretation of the
prime implicants, which would be different for both scenarios, it left up to the user.
