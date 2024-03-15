# Bill Split App in Go

A bill split app that supports multi live currency,will write
the output in the terminal and in the text/md file itself.Will
write in the top of the file.

## 0. Why I made this app

1. As the famous Splitwise only allow 6 bills added in one day,
have annoying adds blocking too often while going to add a bill,
I want to resolve it by make my own.
1. For more privacy, compared to the normal bill split apps
which stored it in a server.

## 1. Currencies

Will have two types of currency written ways:

1. Batch Write: all the list below will have the same currency.
A "\n\n" will mean change of currency following by a three digits
currency, it only detects the last 3 digits, so you can have a `=>`,
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

- Currency supported: `USD(d), EUR(e), JPY(j), TWD(t), KOW(w), GBP(p), CHF(f)`

## 2. Categories

These are the different categories for the items:

1. [g]roceries
1. [f]un/concerts/events/museums
1. [h]ealth
1. [t]ransport/oil
1. [e]at/dine
1. [r]ent/home
1. [c]lothing
1. [b]ills
1. [l]earn/studies
1. [s]tationery/books/magazine
1. ta[x]es
1. [o]thers (default if not written)

Will be written in the last part of each item with a "[acronym]"
-> `44.4e lipsticks [h]`
Also can be written like "!h", which will be more easy for mobile users
-> `44.4e lipsticks !h`

If no categories are written will be categorize as [o]thers.

## Todo

- [x] retrieve from a markdown
- detect and process
  - [x] names
  - [x] currency
  - [x] item.price
  - [x] item.no_split
  - [ ] item.category  
- exchange
  - [x] currency
  - [ ] live exchange
  - [ ] update live and auto save it locally
so can exchange when the computer do not have internet
- [x] split for 2 ppl
- [ ] split for >=3 ppl
- [ ] implement categories: groceries(g), fun/events(f), health(h)
- visualise
  - [x] in the terminal
  - [ ] in the browser locally
  - [ ] a graph
- [ ] self host it
- [ ] others can add bill using the browser with the phone/pc

## Example file

```markdown
# 2020 Japan Travel

- from 3rd Jul to 6th Jul
- default currency: eur

## Susan

3e burguer
4.4j groceries

## Niklas

4d chicken
12.4d gas
f6.6d cloth #when is 'f' at the beginning

## Amanda

-> EUR
5.5 Susan
6.3 taxi

-> JPY
3.32 chesse
```
