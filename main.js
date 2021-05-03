const defaultCity = "city 1"

const restaurants = {
  "city 1": [
    {
      name: "ABC Pizza",
      weight: 1
    },
    "Arby's",
    "Subway"
  ],

  "City 2": [
    {
      name: "Taco Bell",
      weight: -2
    },
    "Hacienda"
  ]
};

const dayAbbrvs = abbv => {
  switch (abbv.toLowerCase()) {
    case "s":
    case "su":
      return 0;
    case "m":
      return 1;
    case "t":
      return 2;
    case "w":
      return 3;
    case "th":
    case "r":
      return 4;
    case "f":
      return 5;
    case "sa":
      return 6;
  }
};

let weightedChoices = [];

const day = new Date().getDay();
const userLocation = () => {
  const path = window.location.search.replace("%20", " ").replace("?", "");
  if (path === "" || !(path in restaurants)) {
    return defaultCity;
  }
  return path;
};

const randVal = max => Math.floor(Math.random() * max);
const openToday = schStr => schStr.split(",").some(dayAbbv => dayAbbrvs(dayAbbv) == day)
const makeChoice = () => document.getElementById("eat").textContent = weightedChoices[randVal(weightedChoices.length)].name;

restaurants[userLocation()].forEach(val => {
  if (typeof val === "string") {
    val = { name: val };
  }

  if (!openToday(val.sch ?? "s,m,t,w,r,f,sa")) {
    return;
  }

  const weight = val.weight ? val.weight < -4 ? -4 : val.weight : 0;
  for (let i = 0; i < (5 + weight); i++) {
    weightedChoices.push(val);
  }
});

window.addEventListener("click", () => makeChoice());
window.addEventListener("mousedown", () => document.body.style.backgroundColor = "rgb(162, 255, 73)");
window.addEventListener("mouseup", () => document.body.style.backgroundColor = "lawngreen");

makeChoice();
