const seekAndDestroy = (arr, ...args) => {
  let newArr = arr.filter((e) => !args.includes(e));

  return newArr;
};

console.log(seekAndDestroy([1, 2, 3, 5, 1, 2, 3], 2, 3));
