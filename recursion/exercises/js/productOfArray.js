export default function productOfArray(arr) {
  if (arr.length < 1) {
    return 1;
  }

  let lastElement = arr.pop()
  return lastElement * productOfArray(arr);
}
