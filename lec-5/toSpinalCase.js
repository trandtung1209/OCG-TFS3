const toSpinalCase = (str) => {
  let spinalStr = str
    .replace(/(?!^)([A-Z])/g, " $1")
    .replace(/([^a-zA-Z])+/g, "-")
    .toLowerCase();

  return spinalStr;
};

console.log(toSpinalCase("my_Name_is Quan"));
