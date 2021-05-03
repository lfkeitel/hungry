# Hungry

A small web page to choose a place to eat.

## Setup

Throw the html and js files on a web server and you're good to go.

Restaurants are stored in the .js file.

```js
const defaultCity = "city"
const restaurants = {
  "city": [
    {
      name: "Mom and Pop Pizza",
      sch: "t,w,r,f,sa", // Schedule, see below
      weight: 2
    },
    "Taco Bell"
  ],

  "other city": [
    "Deli"
  ]
};
```

Shortcuts can be taken. If `weight` is omitted, it will default to 0.
If `sch` is omitted, it will default to everyday. If `sch` and `weight` are
both omitted, the array element can be a simple string of the name instead
of an object.

```js
[
  {
    name: "Mom and Pop Pizza"
  }
]
// or
[
  "Mom and Pop Pizza"
]
```

Schedule format:

Comma separated list of day abbreviations.
Letter case doesn't matter. Abbreviations are:

- Monday: m
- Tuesday: t
- Wednesday: w
- Thursday: th,r
- Friday: f
- Saturday: sa
- Sunday: s,su
