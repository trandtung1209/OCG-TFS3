const uniqueUnion = (...arr) => {
  let newArr = [].concat(...arr);
  uniArr = [...new Set(newArr)];

  return uniArr;
};

uniqueUnion([1, 2, 3, 4], [2, 3, 4, 5]);
