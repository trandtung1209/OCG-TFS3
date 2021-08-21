const drop = (arr, func) => {
  const index = arr.findIndex(func);
  return index === -1 ? [] : arr.slice(index);
};

console.log(
  drop([1, 2, 3, 4], function (n) {
    return n > 5;
  })
);

console.log(
  drop([0, 1, 0, 1], function (n) {
    return n === 1;
  })
);
