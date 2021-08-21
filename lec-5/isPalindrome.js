const isPalindrome = (str) => {
  const inputStr = str.replace(/[^a-zA-Z0-9]/g, "").toLowerCase();
  const reverseStr = inputStr.split("").reverse().join("");

  return inputStr === reverseStr;
};

console.log(isPalindrome("race car"));
