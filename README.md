# Hungry

A small webapp to choose a place to eat.

## Setup

Build using typical Go tools.

The restaurant file is line-based where each line has the following format:

`Name:Days Open:Score` E.g. `Subway:m,t,w,th,f,sa,su:0`

Shortcuts can be taken. If the score is omitted, it will default to 0.
If the days open is omitted, it will default to everyday.

These are the same:

`Subway:m,t,w,th,f,sa,su:0` and `Subway::`

The colons are still required.

- **Name**: The name can contain any letter or punctuation except a newline or colon.
- **Days Open**: Days open is a comma separated list of day Abbreviations.
  Letter case doesn't matter. Abbreviations are:
  - Monday: m
  - Tuesday: t
  - Wednesday: w
  - Thursday: th,r
  - Friday: f
  - Saturday: sa,a
  - Sunday: su,u
- **Score**: Score weights the restaurant against all the others. The score can range
  between -4 to 5 with -4 meaning "I don't like it as much" and 5 meaning "I like it
  a lot". The default is 0.

## Example file

```
# Lines starting with a pound sign are comments
Dairy Queen::0
McDonald's::-4
Arby's::0
Random Mexican Place:m,t,w,r,f:0
Subway::0
Grocery Store::-2
```
