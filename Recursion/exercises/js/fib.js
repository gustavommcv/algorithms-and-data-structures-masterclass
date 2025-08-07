export default function fib(input) {
  if (input === 0) {
    return 0;
  }

  let arr = [0, 1];

  let helper = () => {
    if (input === arr.length - 1) {
      return;
    }

    arr.push(arr[arr.length - 1] + arr[arr.length - 2]);
    helper();
  };

  helper();

  return arr[arr.length - 1];
}
