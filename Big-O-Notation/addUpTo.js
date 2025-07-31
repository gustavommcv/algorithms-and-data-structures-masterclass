function addUpTo(number) {
  total = 0;

  for (let index = 1; index <= number; index++) {
    total += index;
  }

  return total;
}

function addUpToOptimized(number) {
  return (number * (number + 1)) / 2;
}

var t1 = performance.now();
addUpToOptimized(100000000);
var t2 = performance.now();
console.log("time elapsed: ", (t2 - t1) / 1000);
