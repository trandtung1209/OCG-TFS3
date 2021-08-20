const display = document.querySelector(".display");
const numbers = document.querySelectorAll(".number");
const operators = document.querySelectorAll(".operator");
const equal = document.querySelector(".equal");
const clear = document.querySelector(".clear");
const del = document.querySelector(".del");

let temp = "";
let haveDot = false;
let haveOpr = false;
let data = {
  num1: "",
  num2: "",
  opr: "",
};

numbers.forEach((number) => {
  number.addEventListener("click", (e) => {
    // if (data.opr !== "" && display.innerText !== "") {
    //   temp = "";
    //   haveOpr = false;
    // }
    if (data.opr !== "")
      operators.forEach((operator) => {
        operator.classList.remove("operator--active");
      });
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
    // if (e.target.innerText == "%")
    //   display.innerText.split("").splice(0, 0, "0.0").join("");
    if (haveOpr === false) haveOpr = true;
    else return;

    data.num1 = temp;
    temp = "";
    data.opr = e.target.innerText;
    operator.classList.add("operator--active");
  });
});

clear.addEventListener("click", () => {
  data = {
    num1: "",
    num2: "",
    opr: "",
  };
  haveDot = false;
  haveOpr = false;
  temp = "";
  display.innerText = "0";
});

del.addEventListener("click", () => {
  temp = display.innerText;
  temp = temp.substring(0, temp.length - 1);
  display.innerText = temp !== "" ? temp : "0";
});

equal.addEventListener("click", () => {
  data.num2 = temp;
  console.log(JSON.stringify(data));
  if (data.num2 === "0" && data.opr === "÷") display.innerHTML = "ERROR";
  else {
    fetch("http://127.0.0.1:8080/submit", {
      method: "POST",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Success:", data);
        display.innerText = data;
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }
});

// keyboard
window.addEventListener("keydown", (e) => {
  if (
    e.key === "0" ||
    e.key === "1" ||
    e.key === "2" ||
    e.key === "3" ||
    e.key === "4" ||
    e.key === "5" ||
    e.key === "6" ||
    e.key === "7" ||
    e.key === "8" ||
    e.key === "9" ||
    e.key === "."
  ) {
    clickNumber(e.key);
  } else if (e.key === "+" || e.key === "-" || e.key === "%") {
    clickOperator(e.key);
  } else if (e.key === "/") {
    clickOperator("÷");
  } else if (e.key === "*") {
    clickOperator("×");
  } else if (e.key == "Enter" || e.key === "=") {
    clickEqual();
  } else if (e.key === "Backspace") {
    clickDel();
  }
});

function clickNumber(key) {
  numbers.forEach((number) => {
    if (number.innerText === key) {
      number.click();
    }
  });
}
function clickOperator(key) {
  operators.forEach((operator) => {
    if (operator.innerText === key) {
      operator.click();
    }
  });
}
function clickEqual() {
  equal.click();
}
function clickDel() {
  del.click();
}
