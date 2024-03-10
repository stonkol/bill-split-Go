# bill-split-Go

A bill split app that supports multi live currency,
will write the output in the terminal and in the text/md file itself.
Will write in the top of the file.

## 1. Currencies

Will have two types of currency written ways:

1. Batch Write: all the list below will have the same currency.
a "\n\n" will mean change of currency following by a three digits
currecy, it only detecs the last 3 digits, so you can have a `=>`,
`-` or nothing, it will not affect the detection

  ```md
  => EUR
  4.3 item 1 
  55 item 2
  ```

1. Multi Currency: each line can have different currency.

  ```md
  4.3e item1
  55j item2
  ```

- Currency supported:

`USD(d), EUR(e), JPY(j), TWD(t), KOW(w), GBP(p), CHF(f)`

## 2. Categories

Have different categories:

1. [g]roceries
1. [f]un/concerts/events
1. [h]ealth
1. [t]ransport
1. [e]at/dine
1. [r]ent/home
1. [c]lothing
1. [b]ills

Will be written in the last part of each item with a "[acronym]": `44.4e lipsticks [h]`

If no categories are written will be categorize as [o]thers.

## Todo

- [ ] exchange currency
- [ ] live exchange currency
- [ ] split for 2 ppl
- [ ] split for >=3 ppl
- [ ] have categories: groceries(g), fun/events(f), health(h)

## Example file

```markdown
### 2020 Japan Travel

- from 3rd Jul to 6th Jul

# Susan

3e burguer
4.4j groceries

# Niklas

4d chicken
12.4d gas
f6.6d cloth #when is 'f' at the beginning
```

