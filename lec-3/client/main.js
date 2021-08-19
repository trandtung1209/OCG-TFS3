const display = document.querySelector(".display");
const numbers = document.querySelectorAll(".number");
const operators = document.querySelectorAll(".operator");
const equal = document.querySelector(".equal");
const clear = document.querySelector(".clear");
const del = document.querySelector(".del");

let num1 = "";
let num2 = "";
let opr = "";
let temp = "";
let haveDot = false;

numbers.forEach((number) => {
  number.addEventListener("click", (e) => {
    if (e.target.innerText === "." && !haveDot) {
      haveDot = true;
    } else if (e.target.innerText === "." && haveDot) {
      return;
    }
    temp += e.target.innerText;
    display.innerText = temp;
  });
});

operators.forEach((operator) => {
  operator.addEventListener("click", (e) => {
    num1 = temp;
    temp = "";
    opr = e.target.innerText;
  });
});

clear.addEventListener("click", () => {
  num1 = "";
  haveDot = false;
  opr = "";
  temp = "";
  num2 = "";
  display.innerText = "0";
});

del.addEventListener("click", () => {
  temp = temp.substring(0, temp.length - 1);
  display.innerText = temp !== "" ? temp : "0";
});

equal.addEventListener("click", () => {
  num2 = temp;
  console.log(num1, opr, num2);
});
