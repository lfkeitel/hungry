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

let currentLoc = userLocation();

const randVal = max => Math.floor(Math.random() * max);
const openToday = schStr => schStr.split(",").some(dayAbbv => dayAbbrvs(dayAbbv) == day)
const makeChoice = () => document.getElementById("eat").textContent = weightedChoices[randVal(weightedChoices.length)].name;

const loadChoices = (currentLoc) => {
  weightedChoices = [];

  restaurants[currentLoc].forEach(val => {
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
};

document.getElementById("result").addEventListener("click", () => makeChoice());
document.getElementById("result").addEventListener("mousedown", () => document.body.style.backgroundColor = "rgb(162, 255, 73)");
document.getElementById("result").addEventListener("touchstart", () => document.body.style.backgroundColor = "rgb(162, 255, 73)");
document.getElementById("result").addEventListener("mouseup", () => document.body.style.backgroundColor = "lawngreen");
document.getElementById("result").addEventListener("touchend", () => document.body.style.backgroundColor = "lawngreen");
window.addEventListener("popstate", (event) => {
  const r = event.state;
  if (r !== currentLoc) {
    currentLoc = r;
  }
  loadChoices(r);
  makeChoice();
})

for (const r in restaurants) {
  const b = document.createElement("button")
  b.textContent = r;
  b.classList.add("restButton");
  b.addEventListener("click", () => {
    if (r !== currentLoc) {
      window.history.pushState(r, "", `?${r}`);
      currentLoc = r;
    }
    loadChoices(r);
    makeChoice();
  });

  document.getElementById("cities").appendChild(b);
}

loadChoices(currentLoc);
makeChoice();
